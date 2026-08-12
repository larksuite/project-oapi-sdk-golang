package comment


type BizObject struct {

    Type  *string `json:"type,omitempty"`

    ParentID  *string `json:"parent_id,omitempty"`

    WorkItemTypeKey  *string `json:"work_item_type_key,omitempty"`

    WorkItemID  *string `json:"work_item_id,omitempty"`

    FieldKey  *string `json:"field_key,omitempty"`

}

type Comment struct {

    CommentID  *string `json:"comment_id,omitempty"`

    ProjectKey  *string `json:"project_key,omitempty"`

    Object  *BizObject `json:"object,omitempty"`

    SequenceID  *string `json:"sequence_id,omitempty"`

    Status  *string `json:"status,omitempty"`

    Content  *Content `json:"content,omitempty"`

    Creator  *string `json:"creator,omitempty"`

    HasMoreChildren  *bool `json:"has_more_children,omitempty"`

}

type CommentForOpenAPI struct {

    ID  *int64 `json:"id,omitempty"`

    WorkItemID  *int64 `json:"work_item_id,omitempty"`

    WorkItemTypeKey  *string `json:"work_item_type_key,omitempty"`

    CreatedAt  *int64 `json:"created_at,omitempty"`

    Operator  *string `json:"operator,omitempty"`

    Content  *string `json:"content,omitempty"`

    DocRichText  *DocTypeRichText `json:"doc_rich_text,omitempty"`

}

type Content struct {

    Type  *string `json:"type,omitempty"`

    RichText  *RichText `json:"rich_text,omitempty"`

    FileToken  *string `json:"file_token,omitempty"`

    UpdateTime  *int64 `json:"update_time,omitempty"`

    CreateTime  *int64 `json:"create_time,omitempty"`

}

type DocTypeRichText struct {

    Doc  *string `json:"doc,omitempty"`

    DocText  *string `json:"doc_text,omitempty"`

    DocHTML  *string `json:"doc_html,omitempty"`

    DocImg  []string `json:"doc_img,omitempty"`

    IsEmpty  *bool `json:"is_empty,omitempty"`

}

type MutateCreateContent struct {

    Type  *string `json:"type,omitempty"`

    RichText  interface{} `json:"rich_text,omitempty"`

    FileToken  *string `json:"file_token,omitempty"`

    Text  *string `json:"text,omitempty"`

    Markdown  *string `json:"markdown,omitempty"`

}

type MutateUpdateContent struct {

    Type  *string `json:"type,omitempty"`

    RichText  interface{} `json:"rich_text,omitempty"`

    Text  *string `json:"text,omitempty"`

    Markdown  *string `json:"markdown,omitempty"`

}

type Pagination struct {

    PageNum  *int64 `json:"page_num,omitempty"`

    PageSize  *int64 `json:"page_size,omitempty"`

    Total  *int64 `json:"total,omitempty"`

}

type Paginator struct {

    Direction  *string `json:"direction,omitempty"`

    Cursor  *string `json:"cursor,omitempty"`

    PageSize  *int32 `json:"page_size,omitempty"`

}

type RichText struct {

    Doc  *string `json:"doc,omitempty"`

    DocText  *string `json:"doc_text,omitempty"`

    DocHTML  *string `json:"doc_html,omitempty"`

    DocImg  []string `json:"doc_img,omitempty"`

    IsEmpty  *bool `json:"is_empty,omitempty"`

    MarkdownContent  *RichTextMarkDown `json:"markdown_content,omitempty"`

}

type RichTextMarkDown struct {

    Markdown  *string `json:"markdown,omitempty"`

    MentionUserKeys  []UserKeys `json:"mention_user_keys,omitempty"`

}

type UserKeys struct {

    ID  *string `json:"id,omitempty"`

    UserKey  *string `json:"user_key,omitempty"`

}
