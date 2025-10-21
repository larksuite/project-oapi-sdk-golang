package view

type Checker struct {
	CheckedTime *int64 `json:"checked_time,omitempty"`

	Owner *string `json:"owner,omitempty"`

	Status *int32 `json:"status,omitempty"`
}

type Connection struct {
	SourceStateKey *string `json:"source_state_key,omitempty"`

	TargetStateKey *string `json:"target_state_key,omitempty"`

	TransitionID *int64 `json:"transition_id,omitempty"`
}

type Error struct {
	ErrorCode *int32 `json:"error_code,omitempty"`

	ErrorInfo *string `json:"error_info,omitempty"`
}

type Expand struct {
	NeedWorkflow *bool `json:"need_workflow,omitempty"`

	RelationFieldsDetail *bool `json:"relation_fields_detail,omitempty"`

	NeedMultiText *bool `json:"need_multi_text,omitempty"`

	NeedUserDetail *bool `json:"need_user_detail,omitempty"`

	NeedSubTaskParent *bool `json:"need_sub_task_parent,omitempty"`
}

type FieldDetail struct {
	StoryID *int64 `json:"story_id,omitempty"`

	WorkItemTypeKey *string `json:"work_item_type_key,omitempty"`

	ProjectKey *string `json:"project_key,omitempty"`
}

type FieldValuePair struct {
	FieldKey *string `json:"field_key,omitempty"`

	FieldValue interface{} `json:"field_value,omitempty"`

	TargetState *TargetState `json:"target_state,omitempty"`

	FieldTypeKey *string `json:"field_type_key,omitempty"`

	FieldAlias *string `json:"field_alias,omitempty"`

	HelpDescription *string `json:"help_description,omitempty"`
}

type FixView struct {
	ViewID *string `json:"view_id,omitempty"`

	Name *string `json:"name,omitempty"`

	CreatedBy *string `json:"created_by,omitempty"`

	CreatedAt *int64 `json:"created_at,omitempty"`

	ModifiedBy *string `json:"modified_by,omitempty"`

	WorkItemIDList []int64 `json:"work_item_id_list,omitempty"`

	Editable *bool `json:"editable,omitempty"`
}

type I18nTranslation struct {
	SyncStatus *string `json:"sync_status,omitempty"`

	Translated *bool `json:"translated,omitempty"`

	Language *string `json:"language,omitempty"`

	Translation *string `json:"translation,omitempty"`
}

type MultiText struct {
	FieldKey *string `json:"field_key,omitempty"`

	FieldValue *MultiTextDetail `json:"field_value,omitempty"`
}

type MultiTextDetail struct {
	Doc *string `json:"doc,omitempty"`

	DocText *string `json:"doc_text,omitempty"`

	IsEmpty *bool `json:"is_empty,omitempty"`

	NotifyUserList []string `json:"notify_user_list,omitempty"`

	NotifyUserType *string `json:"notify_user_type,omitempty"`

	DocHTML *string `json:"doc_html,omitempty"`
}

type NodeBasicInfo struct {
	ID *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	Owners []string `json:"owners,omitempty"`

	Milestone *bool `json:"milestone,omitempty"`
}

type NodesConnections struct {
	WorkflowNodes []WorkflowNode `json:"workflow_nodes,omitempty"`

	Connections []Connection `json:"connections,omitempty"`

	StateFlowNodes []StateFlowNode `json:"state_flow_nodes,omitempty"`

	UserDetails []UserDetail `json:"user_details,omitempty"`
}

type Pagination struct {
	PageNum *int64 `json:"page_num,omitempty"`

	PageSize *int64 `json:"page_size,omitempty"`

	Total *int64 `json:"total,omitempty"`
}

type QuickFilter struct {
	QuickFilterID *string `json:"quick_filter_id,omitempty"`

	QuickFilterName *string `json:"quick_filter_name,omitempty"`
}

type RelationFieldDetail struct {
	FieldKey *string `json:"field_key,omitempty"`

	Detail []FieldDetail `json:"detail,omitempty"`
}

type ResourceAuth struct {
	Usable *bool `json:"usable,omitempty"`

	Editable *bool `json:"editable,omitempty"`

	Deletable *bool `json:"deletable,omitempty"`

	Creatable *bool `json:"creatable,omitempty"`

	BatchEditable *bool `json:"batch_editable,omitempty"`
}

type ResourceItem struct {
	Key *string `json:"key,omitempty"`

	Ruuid *string `json:"ruuid,omitempty"`

	Type *string `json:"type,omitempty"`

	Class *string `json:"class,omitempty"`

	ParentKey *string `json:"parent_key,omitempty"`

	CustomKey *string `json:"custom_key,omitempty"`

	Label *string `json:"label,omitempty"`

	DefaultLabel *string `json:"default_label,omitempty"`

	Icon *string `json:"icon,omitempty"`

	Description *string `json:"description,omitempty"`

	Auth *ResourceAuth `json:"auth,omitempty"`

	Params *string `json:"params,omitempty"`

	Available []string `json:"available,omitempty"`

	Unavailable []string `json:"unavailable,omitempty"`

	UsingFeaturePoints []string `json:"using_feature_points,omitempty"`

	ResourceItems []ResourceItem `json:"resource_items,omitempty"`

	ErrorMessage *Error `json:"error_message,omitempty"`

	IsPreset *bool `json:"is_preset,omitempty"`

	BelongsTo *string `json:"belongs_to,omitempty"`

	LabelUniqueKey *string `json:"label_unique_key,omitempty"`

	LabelOrigin *string `json:"label_origin,omitempty"`

	LabelTranslations []I18nTranslation `json:"label_translations,omitempty"`

	ResourceKey *string `json:"resource_key,omitempty"`

	Scope *Scope `json:"scope,omitempty"`

	Status *string `json:"status,omitempty"`

	RelationRuuid *string `json:"relation_ruuid,omitempty"`

	ParentLabel *string `json:"parent_label,omitempty"`

	UiDescription *UIDescription `json:"ui_description,omitempty"`

	GroupLabel *string `json:"group_label,omitempty"`

	IsMultiTarget *bool `json:"is_multi_target,omitempty"`
}

type RoleOwner struct {
	Role *string `json:"role,omitempty"`

	Name *string `json:"name,omitempty"`

	Owners []string `json:"owners,omitempty"`
}

type Schedule struct {
	Points *float64 `json:"points,omitempty"`

	EstimateStartDate *int64 `json:"estimate_start_date,omitempty"`

	EstimateEndDate *int64 `json:"estimate_end_date,omitempty"`

	Owners []string `json:"owners,omitempty"`

	ActualWorkTime *float64 `json:"actual_work_time,omitempty"`
}

type Scope struct {
	ProjectKey *string `json:"project_key,omitempty"`

	WorkItemTypeKey *string `json:"work_item_type_key,omitempty"`
}

type SearchGroup struct {
	SearchParams []SearchParam `json:"search_params,omitempty"`

	Conjunction *string `json:"conjunction,omitempty"`

	SearchGroups []SearchGroup `json:"search_groups,omitempty"`
}

type SearchParam struct {
	ParamKey *string `json:"param_key,omitempty"`

	Value interface{} `json:"value,omitempty"`

	Operator *string `json:"operator,omitempty"`

	PreOperator *string `json:"pre_operator,omitempty"`

	ValueSearchGroups *SearchGroup `json:"value_search_groups,omitempty"`
}

type StateFlowNode struct {
	ID *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	RoleOwners []RoleOwner `json:"role_owners,omitempty"`

	Status *int32 `json:"status,omitempty"`

	ActualBeginTime *string `json:"actual_begin_time,omitempty"`

	ActualFinishTime *string `json:"actual_finish_time,omitempty"`

	Fields []FieldValuePair `json:"fields,omitempty"`
}

type StateTime struct {
	StateKey *string `json:"state_key,omitempty"`

	StartTime *int64 `json:"start_time,omitempty"`

	EndTime *int64 `json:"end_time,omitempty"`

	Name *string `json:"name,omitempty"`
}

type SubTask struct {
	ID *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	Schedules []Schedule `json:"schedules,omitempty"`

	Order *float64 `json:"order,omitempty"`

	Details *string `json:"details,omitempty"`

	Passed *bool `json:"passed,omitempty"`

	Owners []string `json:"owners,omitempty"`

	Note *string `json:"note,omitempty"`

	ActualBeginTime *string `json:"actual_begin_time,omitempty"`

	ActualFinishTime *string `json:"actual_finish_time,omitempty"`

	Assignee []string `json:"assignee,omitempty"`

	RoleAssignee []RoleOwner `json:"role_assignee,omitempty"`

	Deliverable []FieldValuePair `json:"deliverable,omitempty"`

	Fields []FieldValuePair `json:"fields,omitempty"`
}

type SubTaskParentInfo struct {
	WorkItemID *int64 `json:"work_item_id,omitempty"`

	WorkItemName *string `json:"work_item_name,omitempty"`

	WorkItemTypeKey *string `json:"work_item_type_key,omitempty"`

	NodeID *string `json:"node_id,omitempty"`
}

type TargetState struct {
	StateKey *string `json:"state_key,omitempty"`

	TransitionID *int64 `json:"transition_id,omitempty"`
}

type Team struct {
	ID *string `json:"id,omitempty"`

	NotCascade *bool `json:"not_cascade,omitempty"`

	NotContainTeamManager *bool `json:"not_contain_team_manager,omitempty"`
}

type TimeInterval struct {
	Start *int64 `json:"start,omitempty"`

	End *int64 `json:"end,omitempty"`
}

type UIDescription struct {
	UiType *string `json:"ui_type,omitempty"`

	UUID *string `json:"uuid,omitempty"`

	IsMulti *bool `json:"is_multi,omitempty"`

	DateType *int8 `json:"dateType,omitempty"`

	DateFormat *string `json:"dateFormat,omitempty"`

	SubType *int32 `json:"subType,omitempty"`

	SubTypeLevel *int64 `json:"subTypeLevel,omitempty"`
}

type UserDetail struct {
	UserKey *string `json:"user_key,omitempty"`

	Username *string `json:"username,omitempty"`

	Email *string `json:"email,omitempty"`

	NameCn *string `json:"name_cn,omitempty"`

	NameEn *string `json:"name_en,omitempty"`
}

type ViewConf struct {
	ViewID *string `json:"view_id,omitempty"`

	Name *string `json:"name,omitempty"`

	ViewUrl *string `json:"view_url,omitempty"`

	ViewType *int32 `json:"view_type,omitempty"`

	Auth *int32 `json:"auth,omitempty"`

	SystemView *int32 `json:"system_view,omitempty"`

	Collaborators []string `json:"collaborators,omitempty"`

	CreatedAt *int64 `json:"created_at,omitempty"`

	CreatedBy *string `json:"created_by,omitempty"`

	QuickFilters []QuickFilter `json:"quick_filters,omitempty"`
}

type WorkItemInfo struct {
	ID *int64 `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	WorkItemTypeKey *string `json:"work_item_type_key,omitempty"`

	ProjectKey *string `json:"project_key,omitempty"`

	TemplateType *string `json:"template_type,omitempty"`

	Pattern *string `json:"pattern,omitempty"`

	SubStage *string `json:"sub_stage,omitempty"`

	CurrentNodes []NodeBasicInfo `json:"current_nodes,omitempty"`

	CreatedBy *string `json:"created_by,omitempty"`

	UpdatedBy *string `json:"updated_by,omitempty"`

	CreatedAt *int64 `json:"created_at,omitempty"`

	UpdatedAt *int64 `json:"updated_at,omitempty"`

	Fields []FieldValuePair `json:"fields,omitempty"`

	WorkItemStatus *WorkItemStatus `json:"work_item_status,omitempty"`

	DeletedBy *string `json:"deleted_by,omitempty"`

	DeletedAt *int64 `json:"deleted_at,omitempty"`

	SimpleName *string `json:"simple_name,omitempty"`

	TemplateID *int64 `json:"template_id,omitempty"`

	WorkflowInfos *NodesConnections `json:"workflow_infos,omitempty"`

	StateTimes []StateTime `json:"state_times,omitempty"`

	MultiTexts []MultiText `json:"multi_texts,omitempty"`

	RelationFieldsDetail []RelationFieldDetail `json:"relation_fields_detail,omitempty"`

	UserDetails []UserDetail `json:"user_details,omitempty"`

	SubTaskParentInfo *SubTaskParentInfo `json:"sub_task_parent_info,omitempty"`
}

type WorkItemStatus struct {
	StateKey *string `json:"state_key,omitempty"`

	IsArchivedState *bool `json:"is_archived_state,omitempty"`

	IsInitState *bool `json:"is_init_state,omitempty"`

	UpdatedAt *int64 `json:"updated_at,omitempty"`

	UpdatedBy *string `json:"updated_by,omitempty"`

	History []WorkItemStatus `json:"history,omitempty"`
}

type WorkflowNode struct {
	ID *string `json:"id,omitempty"`

	StateKey *string `json:"state_key,omitempty"`

	Name *string `json:"name,omitempty"`

	Status *int32 `json:"status,omitempty"`

	Fields []FieldValuePair `json:"fields,omitempty"`

	Owners []string `json:"owners,omitempty"`

	NodeSchedule *Schedule `json:"node_schedule,omitempty"`

	Schedules []Schedule `json:"schedules,omitempty"`

	SubTasks []SubTask `json:"sub_tasks,omitempty"`

	ActualBeginTime *string `json:"actual_begin_time,omitempty"`

	ActualFinishTime *string `json:"actual_finish_time,omitempty"`

	RoleAssignee []RoleOwner `json:"role_assignee,omitempty"`

	DifferentSchedule *bool `json:"different_schedule,omitempty"`

	SubStatus []Checker `json:"sub_status,omitempty"`

	Milestone *bool `json:"milestone,omitempty"`

	Participants []string `json:"participants,omitempty"`
}
