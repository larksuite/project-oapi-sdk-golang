package comment


type CommentForOpenAPI struct {

    ID  *int64 `json:"id,omitempty"`

    WorkItemID  *int64 `json:"work_item_id,omitempty"`

    WorkItemTypeKey  *string `json:"work_item_type_key,omitempty"`

    CreatedAt  *int64 `json:"created_at,omitempty"`

    Operator  *string `json:"operator,omitempty"`

    Content  *string `json:"content,omitempty"`

}

type Pagination struct {

    PageNum  *int64 `json:"page_num,omitempty"`

    PageSize  *int64 `json:"page_size,omitempty"`

    Total  *int64 `json:"total,omitempty"`

}
