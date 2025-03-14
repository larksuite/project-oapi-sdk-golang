package project


type CommentForOpenAPI struct {

    ID  *int64 `json:"id,omitempty"`

    WorkItemID  *int64 `json:"work_item_id,omitempty"`

    WorkItemTypeKey  *string `json:"work_item_type_key,omitempty"`

    CreatedAt  *int64 `json:"created_at,omitempty"`

    Operator  *string `json:"operator,omitempty"`

    Content  *string `json:"content,omitempty"`

}

type OAPIOperationRecord struct {

    ProjectKey  *string `json:"project_key,omitempty"`

    OperationType  *string `json:"operation_type,omitempty"`

    OperationTime  *int64 `json:"operation_time,omitempty"`

    WorkItemID  *int64 `json:"work_item_id,omitempty"`

    OperatorType  *string `json:"operator_type,omitempty"`

    Operator  *string `json:"operator,omitempty"`

    OpRecordModule  *string `json:"op_record_module,omitempty"`

    SourceType  *string `json:"source_type,omitempty"`

    Source  *string `json:"source,omitempty"`

    RecordContents  []OAPIRecordContent `json:"record_contents,omitempty"`

    WorkItemTypeKey  *string `json:"work_item_type_key,omitempty"`

}

type OAPIRecordContent struct {

    Object  *OpRecordObject `json:"object,omitempty"`

    ObjectProperty  *string `json:"object_property,omitempty"`

    Old  []string `json:"old,omitempty"`

    New  []string `json:"new,omitempty"`

    Add  []string `json:"add,omitempty"`

    Delete  []string `json:"delete,omitempty"`

    StatusValues  []ObjectStatusValue `json:"status_values,omitempty"`

    BelongObject  []OpRecordObject `json:"belong_object,omitempty"`

}

type ObjectStatusValue struct {

    ObjectType  *string `json:"object_type,omitempty"`

    ObjectProperty  *string `json:"object_property,omitempty"`

    Values  []string `json:"values,omitempty"`

}

type OpRecordObject struct {

    ObjectType  *string `json:"object_type,omitempty"`

    ObjectValue  *string `json:"object_value,omitempty"`

}

type Pagination struct {

    PageNum  *int64 `json:"page_num,omitempty"`

    PageSize  *int64 `json:"page_size,omitempty"`

    Total  *int64 `json:"total,omitempty"`

}

type Project struct {

    ProjectKey  *string `json:"project_key,omitempty"`

    Name  *string `json:"name,omitempty"`

    SimpleName  *string `json:"simple_name,omitempty"`

    Administrators  []string `json:"administrators,omitempty"`

}

type Team struct {

    TeamID  *int64 `json:"team_id,omitempty"`

    TeamName  *string `json:"team_name,omitempty"`

    UserKeys  []string `json:"user_keys,omitempty"`

    Administrators  []string `json:"administrators,omitempty"`

}
