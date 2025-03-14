package task


type FieldValuePair struct {

    FieldKey  *string `json:"field_key,omitempty"`

    FieldValue  interface{} `json:"field_value,omitempty"`

    TargetState  *TargetState `json:"target_state,omitempty"`

    FieldTypeKey  *string `json:"field_type_key,omitempty"`

    FieldAlias  *string `json:"field_alias,omitempty"`

    HelpDescription  *string `json:"help_description,omitempty"`

}

type Pagination struct {

    PageNum  *int64 `json:"page_num,omitempty"`

    PageSize  *int64 `json:"page_size,omitempty"`

    Total  *int64 `json:"total,omitempty"`

}

type RoleOwner struct {

    Role  *string `json:"role,omitempty"`

    Name  *string `json:"name,omitempty"`

    Owners  []string `json:"owners,omitempty"`

}

type Schedule struct {

    Points  *float64 `json:"points,omitempty"`

    EstimateStartDate  *int64 `json:"estimate_start_date,omitempty"`

    EstimateEndDate  *int64 `json:"estimate_end_date,omitempty"`

    Owners  []string `json:"owners,omitempty"`

    ActualWorkTime  *float64 `json:"actual_work_time,omitempty"`

}

type SubDetail struct {

    WorkItemID  *int64 `json:"work_item_id,omitempty"`

    WorkItemName  *string `json:"work_item_name,omitempty"`

    NodeID  *string `json:"node_id,omitempty"`

    SubTask  *SubTask `json:"sub_task,omitempty"`

}

type SubTask struct {

    ID  *string `json:"id,omitempty"`

    Name  *string `json:"name,omitempty"`

    Schedules  []Schedule `json:"schedules,omitempty"`

    Order  *float64 `json:"order,omitempty"`

    Details  *string `json:"details,omitempty"`

    Passed  *bool `json:"passed,omitempty"`

    Owners  []string `json:"owners,omitempty"`

    Note  *string `json:"note,omitempty"`

    ActualBeginTime  *string `json:"actual_begin_time,omitempty"`

    ActualFinishTime  *string `json:"actual_finish_time,omitempty"`

    Assignee  []string `json:"assignee,omitempty"`

    RoleAssignee  []RoleOwner `json:"role_assignee,omitempty"`

    Deliverable  []FieldValuePair `json:"deliverable,omitempty"`

    Fields  []FieldValuePair `json:"fields,omitempty"`

}

type TargetState struct {

    StateKey  *string `json:"state_key,omitempty"`

    TransitionID  *int64 `json:"transition_id,omitempty"`

}

type TimeInterval struct {

    Start  *int64 `json:"start,omitempty"`

    End  *int64 `json:"end,omitempty"`

}
