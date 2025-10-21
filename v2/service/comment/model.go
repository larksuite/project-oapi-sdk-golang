package comment

type CommentForOpenAPI struct {
	ID *int64 `json:"id,omitempty"`

	WorkItemID *int64 `json:"work_item_id,omitempty"`

	WorkItemTypeKey *string `json:"work_item_type_key,omitempty"`

	CreatedAt *int64 `json:"created_at,omitempty"`

	Operator *string `json:"operator,omitempty"`

	Content *string `json:"content,omitempty"`

	DocRichText *DocTypeRichText `json:"doc_rich_text,omitempty"`
}

type DocTypeRichText struct {
	Doc *string `json:"doc,omitempty"`

	DocText *string `json:"doc_text,omitempty"`

	DocHTML *string `json:"doc_html,omitempty"`

	DocImg []string `json:"doc_img,omitempty"`

	IsEmpty *bool `json:"is_empty,omitempty"`
}

type FileCommentForOpenAPI struct {
	ID *int64 `json:"id,omitempty"`

	WorkItemID *int64 `json:"work_item_id,omitempty"`

	WorkItemTypeKey *string `json:"work_item_type_key,omitempty"`

	CreatedAt *int64 `json:"created_at,omitempty"`

	Operator *string `json:"operator,omitempty"`

	CommentFile *FileDetail `json:"comment_file,omitempty"`
}

type FileDetail struct {
	FileToken *string `json:"file_token,omitempty"`

	FileName *string `json:"file_name,omitempty"`

	FileSize *int64 `json:"file_size,omitempty"`

	FileType *string `json:"file_type,omitempty"`

	FileUrl *string `json:"file_url,omitempty"`
}

type Pagination struct {
	PageNum *int64 `json:"page_num,omitempty"`

	PageSize *int64 `json:"page_size,omitempty"`

	Total *int64 `json:"total,omitempty"`
}
