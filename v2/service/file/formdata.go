/*
 * Copyright (c) 2023 Lark Technologies Pte. Ltd.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package file

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/textproto"
	"strings"
)

// formPayload 是 file 包自持的轻量 multipart 表单模型。
// 与 core.FormData 不同，它不把文件内容缓冲成 []byte，而是持有 io.Reader 与已知长度，
// 由 streamBody 借助 io.Pipe + multipart.Writer 边编码边发送，峰值内存为常数级。
type formPayload struct {
	mimeType string            // 文件 part 的 Content-Type，默认 application/octet-stream
	fields   map[string]string // 普通字段（如 field_map）
	fileName string
	file     io.Reader
	fileSize int64 // 必须等于 file 可读字节数，用于预算 Content-Length
}

func newFormPayload() *formPayload {
	return &formPayload{mimeType: octetStream}
}

func (f *formPayload) addField(k, v string) {
	if f.fields == nil {
		f.fields = map[string]string{}
	}
	f.fields[k] = v
}

func (f *formPayload) setFile(name string, r io.Reader) {
	f.fileName = name
	f.file = r
}

func (f *formPayload) setFileSize(size int64) {
	f.fileSize = size
}

// contentLength 按 multipart/form-data 固定格式逐字节预算总长度，供设置 Content-Length。
// 预算须与 streamBody 的实际写入逐字节对齐（见各段注释），boundary 复用同一字符串。
func (f *formPayload) contentLength(boundary string) int64 {
	var total int64
	// 每个字段 part：--boundary\r\n + Content-Disposition: form-data; name="<k>"\r\n\r\n + <v> + \r\n
	for k, v := range f.fields {
		total += int64(len("--" + boundary + "\r\n"))
		total += int64(len(fmt.Sprintf("Content-Disposition: form-data; name=\"%s\"\r\n", escapeQuotes(k))))
		total += int64(len("\r\n"))
		total += int64(len(v))
		total += int64(len("\r\n"))
	}
	// 文件 part：--boundary\r\n + Content-Disposition(含 filename)\r\n + Content-Type\r\n + \r\n + <fileSize> + \r\n
	total += int64(len("--" + boundary + "\r\n"))
	total += int64(len(fmt.Sprintf("Content-Disposition: form-data; name=\"file\"; filename=\"%s\"\r\n", escapeQuotes(f.fileName))))
	total += int64(len(fmt.Sprintf("Content-Type: %s\r\n", f.mimeType)))
	total += int64(len("\r\n"))
	total += f.fileSize
	total += int64(len("\r\n"))
	// 结尾：--boundary--\r\n
	total += int64(len("--" + boundary + "--\r\n"))
	return total
}

// streamBody 以流式方式编码 multipart body，返回 Content-Type、body 读端与预算总长度。
// 调用方（doStreamRequest）负责读取并 Close body；编码错误经 pipe 传播到读端。
func (f *formPayload) streamBody() (contentType string, body io.ReadCloser, length int64) {
	// 先取一个随机 boundary，供预算与实际写入共用，杜绝长度漂移
	boundary := multipart.NewWriter(io.Discard).Boundary()
	length = f.contentLength(boundary)

	pr, pw := io.Pipe()
	go func() {
		mw := multipart.NewWriter(pw)
		_ = mw.SetBoundary(boundary)
		var err error
		defer func() { pw.CloseWithError(err) }()
		for k, v := range f.fields {
			if err = mw.WriteField(k, v); err != nil {
				return
			}
		}
		// 文件 part：复刻 core 的 header 形态（name="file"; filename="<入参 fileName>"）
		h := make(textproto.MIMEHeader)
		h.Set("Content-Disposition",
			fmt.Sprintf(`form-data; name="file"; filename="%s"`, escapeQuotes(f.fileName)))
		h.Set("Content-Type", f.mimeType)
		var part io.Writer
		if part, err = mw.CreatePart(h); err != nil {
			return
		}
		if _, err = io.Copy(part, f.file); err != nil {
			return
		}
		err = mw.Close()
	}()
	return "multipart/form-data; boundary=" + boundary, pr, length
}

// escapeQuotes 与 core 内未导出的同名实现一致；file 包自持以避免改动 core。
var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

func escapeQuotes(s string) string {
	return quoteEscaper.Replace(s)
}
