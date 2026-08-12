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
	"bytes"
	"io"
	"testing"
)

// TestFormPayloadContentLength 断言 contentLength 预算与 streamBody 实际写入的字节数逐字节相等。
// 覆盖含普通字段 + 文件 part + 需转义的文件名等情形，防止 Content-Length 长度漂移。
func TestFormPayloadContentLength(t *testing.T) {
	cases := []struct {
		name     string
		fields   map[string]string
		fileName string
		mimeType string
		content  []byte
	}{
		{
			name:     "basic",
			fields:   map[string]string{"field_map": `{"a":"b"}`},
			fileName: "bigo.png",
			mimeType: "image/png",
			content:  []byte("payload-content-bytes"),
		},
		{
			name:     "no-field",
			fields:   nil,
			fileName: "unknown-file",
			mimeType: octetStream,
			content:  []byte(""),
		},
		{
			name:     "escape-and-multi-fields",
			fields:   map[string]string{"a": "1", "b": `has "quote" and \back`},
			fileName: `na"me\.bin`,
			mimeType: "application/octet-stream",
			content:  bytes.Repeat([]byte("x"), 1024),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			f := newFormPayload()
			f.mimeType = c.mimeType
			for k, v := range c.fields {
				f.addField(k, v)
			}
			f.setFile(c.fileName, bytes.NewReader(c.content))
			f.setFileSize(int64(len(c.content)))

			_, body, predicted := f.streamBody()
			actual, err := io.Copy(io.Discard, body)
			if err != nil {
				t.Fatalf("copy body: %v", err)
			}
			if err := body.Close(); err != nil {
				t.Fatalf("close body: %v", err)
			}
			if predicted != actual {
				t.Fatalf("content length mismatch: predicted=%d actual=%d", predicted, actual)
			}
		})
	}
}
