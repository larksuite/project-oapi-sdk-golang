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

// 上传文件（表单）返回结果
type UploadFileByFormResult struct {
	FileToken string `json:"file_token"`
}

// 分片初始化返回结果
type UploadPreProcessResult struct {
	IsMultipart bool           `json:"is_multipart"`
	Multipart   *MultipartInfo `json:"multipart"`
	ResourceID  string         `json:"resource_id"`
}

// 分片信息
type MultipartInfo struct {
	IsNew     bool    `json:"is_new"`
	PartCount int64   `json:"part_count"`
	PartSize  int64   `json:"part_size"`
	Need      []int64 `json:"need"`
}

// 完成分片上传返回结果
type UploadPartFinishResult struct {
	FileToken string `json:"file_token"`
}
