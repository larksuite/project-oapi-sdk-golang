package workitem


type ActualTimeInfo struct {

    ActualStartTime  *int64 `json:"actual_start_time,omitempty"`

    ActualFinishTime  *int64 `json:"actual_finish_time,omitempty"`

}

type BotJoinChatInfo struct {

    ChatID  *string `json:"chat_id,omitempty"`

    SuccessMembers  []string `json:"success_members,omitempty"`

    FailedMembers  []string `json:"failed_members,omitempty"`

}

type Business struct {

    ID  *string `json:"id,omitempty"`

    Name  *string `json:"name,omitempty"`

    Project  *string `json:"project,omitempty"`

    Labels  []string `json:"labels,omitempty"`

    RoleOwners  map[string]WorkItem_work_item_RoleOwner `json:"role_owners,omitempty"`

    Watchers  []string `json:"watchers,omitempty"`

    Order  *float64 `json:"order,omitempty"`

    SuperMasters  []string `json:"super_masters,omitempty"`

    Parent  *string `json:"parent,omitempty"`

    Disabled  *bool `json:"disabled,omitempty"`

    LevelID  *int64 `json:"level_id,omitempty"`

    Children  []Business `json:"children,omitempty"`

    TemplateType  *string `json:"template_type,omitempty"`

}

type Checker struct {

    CheckedTime  *int64 `json:"checked_time,omitempty"`

    Owner  *string `json:"owner,omitempty"`

    Status  *int32 `json:"status,omitempty"`

}

type ChildOrder struct {

    UUID  *string `json:"uuid,omitempty"`

    OrderIndex  *int64 `json:"order_index,omitempty"`

}

type Common_RoleOwner struct {

    Role  *string `json:"role,omitempty"`

    RoleKey  *string `json:"role_key,omitempty"`

    Owners  []UserDetail `json:"owners,omitempty"`

}

type CompInfo struct {

    ID  *string `json:"ID,omitempty"`

    Name  *string `json:"name,omitempty"`

    WorkItemTypeKey  *string `json:"WorkItemTypeKey,omitempty"`

    ProjectKey  *string `json:"ProjectKey,omitempty"`

    CreatedBy  *string `json:"CreatedBy,omitempty"`

    CreatedAt  *int64 `json:"CreatedAt,omitempty"`

    SearchHit  []string `json:"SearchHit,omitempty"`

    ViewScopeKey  *string `json:"ViewScopeKey,omitempty"`

    ProjectKeys  []string `json:"ProjectKeys,omitempty"`

}

type Condition struct {

    FieldItem  *FieldItem `json:"field_item,omitempty"`

    Operator  *string `json:"operator,omitempty"`

    OriginalValue  *string `json:"original_value,omitempty"`

    Formula  *string `json:"formula,omitempty"`

    Version  *string `json:"version,omitempty"`

    PreOperator  *string `json:"pre_operator,omitempty"`

    ValueGroup  *Filter `json:"value_group,omitempty"`

    Source  *FieldItemSource `json:"source,omitempty"`

}

type ConfirmForm struct {

    Action  *int64 `json:"action,omitempty"`

    StateKey  *string `json:"state_key,omitempty"`

}

type Connection struct {

    SourceStateKey  *string `json:"source_state_key,omitempty"`

    TargetStateKey  *string `json:"target_state_key,omitempty"`

    TransitionID  *int64 `json:"transition_id,omitempty"`

}

type CreateWorkItemRelationData struct {

    RelationID  *string `json:"relation_id,omitempty"`

}

type CreateWorkingHourRecord struct {

    ResourceType  *string `json:"resource_type,omitempty"`

    ResourceID  *string `json:"resource_id,omitempty"`

    WorkTime  *string `json:"work_time,omitempty"`

    WorkDescription  *string `json:"work_description,omitempty"`

}

type DataSource struct {

    ProjectKey  *string `json:"project_key,omitempty"`

    WorkItemTypeKeys  *string `json:"work_item_type_keys,omitempty"`

}

type DefaultValue struct {

    DefaultAppear  *int32 `json:"default_appear,omitempty"`

    Value  interface{} `json:"value,omitempty"`

}

type DeliveryRelatedInfo struct {

    RootWorkItem  *DeliveryRelatedInfoItem `json:"root_work_item,omitempty"`

    SourceWorkItem  *DeliveryRelatedInfoItem `json:"source_work_item,omitempty"`

}

type DeliveryRelatedInfoItem struct {

    WorkItemID  *int64 `json:"work_item_id,omitempty"`

    ProjectKey  *string `json:"project_key,omitempty"`

    WorkItemTypeKey  *string `json:"work_item_type_key,omitempty"`

    Name  *string `json:"name,omitempty"`

}

type Department struct {

    ID  *string `json:"id,omitempty"`

    Name  *string `json:"name,omitempty"`

    EnName  *string `json:"en_name,omitempty"`

}

type DependencyInfo struct {

    WorkItemID  *int64 `json:"work_item_id,omitempty"`

    StateKey  *string `json:"state_key,omitempty"`

    Name  *string `json:"name,omitempty"`

    Type  *string `json:"type,omitempty"`

}

type DraftViewSubWorkItemConf struct {

    WorkItemTypeKey  *string `json:"work_item_type_key,omitempty"`

    WorkItemTypeName  *string `json:"work_item_type_name,omitempty"`

    Relations  []DraftViewSubWorkItemConfRelation `json:"relations,omitempty"`

    EnableModelResourceLib  *bool `json:"enable_model_resource_lib,omitempty"`

}

type DraftViewSubWorkItemConfRelation struct {

    RelationKey  *string `json:"relation_key,omitempty"`

    RelationName  *string `json:"relation_name,omitempty"`

    RelationConfUUID  *string `json:"relation_conf_uuid,omitempty"`

}

type EditablePersonnelScope struct {

    EditablePersonnelRangeType  *string `json:"editable_personnel_range_type,omitempty"`

    EditableRoles  []string `json:"editable_roles,omitempty"`

}

type Expand struct {

    NeedWorkflow  *bool `json:"need_workflow,omitempty"`

    RelationFieldsDetail  *bool `json:"relation_fields_detail,omitempty"`

    NeedMultiText  *bool `json:"need_multi_text,omitempty"`

    NeedUserDetail  *bool `json:"need_user_detail,omitempty"`

    NeedSubTaskParent  *bool `json:"need_sub_task_parent,omitempty"`

}

type FieldConf struct {

    IsRequired  *int32 `json:"is_required,omitempty"`

    IsVisibility  *int32 `json:"is_visibility,omitempty"`

    RoleAssign  []RoleAssign `json:"role_assign,omitempty"`

    FieldName  *string `json:"field_name,omitempty"`

    FieldKey  *string `json:"field_key,omitempty"`

    FieldAlias  *string `json:"field_alias,omitempty"`

    FieldTypeKey  *string `json:"field_type_key,omitempty"`

    DefaultValue  *DefaultValue `json:"default_value,omitempty"`

    Options  []OptionConf `json:"options,omitempty"`

    CompoundFields  []FieldConf `json:"compound_fields,omitempty"`

    IsValidity  *int32 `json:"is_validity,omitempty"`

    Label  *string `json:"label,omitempty"`

    WorkItemRelation  *WorkItemRelation `json:"work_item_relation,omitempty"`

    FieldUUID  *string `json:"field_uuid,omitempty"`

    FreeAdd  *bool `json:"free_add,omitempty"`

    FieldTips  *string `json:"field_tips,omitempty"`

    SubTypeLevelMode  *string `json:"sub_type_level_mode,omitempty"`

    SubTypeLevelClass  *int64 `json:"sub_type_level_class,omitempty"`

    EditablePersonnelScope  *EditablePersonnelScope `json:"editable_personnel_scope,omitempty"`

}

type FieldDeliverableItem struct {

    FieldInfo  *WorkItem_work_item_FieldValuePair `json:"field_info,omitempty"`

    Placeholder  *string `json:"placeholder,omitempty"`

    Remark  *string `json:"remark,omitempty"`

    Status  *int64 `json:"status,omitempty"`

}

type FieldDeliveryData struct {

    ProjectKey  *string `json:"project_key,omitempty"`

    WorkItemTypeKey  *string `json:"work_item_type_key,omitempty"`

    FieldKey  *string `json:"field_key,omitempty"`

}

type FieldDetail struct {

    StoryID  *int64 `json:"story_id,omitempty"`

    WorkItemTypeKey  *string `json:"work_item_type_key,omitempty"`

    ProjectKey  *string `json:"project_key,omitempty"`

}

type FieldItem struct {

    Class  *string `json:"class,omitempty"`

    Strategy  *int32 `json:"strategy,omitempty"`

    Key  *string `json:"key,omitempty"`

    Type  *string `json:"type,omitempty"`

    ProjectKey  *string `json:"project_key,omitempty"`

    WorkItemTypeKey  *string `json:"work_item_type_key,omitempty"`

    Params  *string `json:"params,omitempty"`

    Extras  map[string]string `json:"extras,omitempty"`

    ParentFieldKey  *string `json:"parent_field_key,omitempty"`

    ParentFieldType  *string `json:"parent_field_type,omitempty"`

    EntityScope  *string `json:"entity_scope,omitempty"`

}

type FieldItemSource struct {

    Usage  *int32 `json:"usage,omitempty"`

    UniqueID  *string `json:"unique_id,omitempty"`

}

type FieldValuePair struct {

    FieldKey  *string `json:"field_key,omitempty"`

    FieldValue  interface{} `json:"field_value,omitempty"`

    TargetState  *TargetState `json:"target_state,omitempty"`

    FieldTypeKey  *string `json:"field_type_key,omitempty"`

    FieldAlias  *string `json:"field_alias,omitempty"`

    HelpDescription  *string `json:"help_description,omitempty"`

}

type Filter struct {

    Conjunction  *string `json:"conjunction,omitempty"`

    Conditions  []Condition `json:"conditions,omitempty"`

    Groups  []Filter `json:"groups,omitempty"`

}

type FinishedInfo struct {

    SummaryMode  *string `json:"summary_mode,omitempty"`

    Conclusion  *NodeFinishedConclusion `json:"conclusion,omitempty"`

    Opinion  *NodeFinishedOpinion `json:"opinion,omitempty"`

}

type Group struct {

    FieldItem  *FieldItem `json:"field_item,omitempty"`

    NeedEmpty  *bool `json:"needEmpty,omitempty"`

    NeedsConfig  *bool `json:"needsConfig,omitempty"`

}

type InstanceDeliverableItem struct {

    Name  *string `json:"name,omitempty"`

    WorkItemID  *int64 `json:"work_item_id,omitempty"`

    Deletable  *bool `json:"deletable,omitempty"`

    MustComplete  *bool `json:"must_complete,omitempty"`

    StateKey  *string `json:"state_key,omitempty"`

    StateName  *string `json:"state_name,omitempty"`

    Owners  []string `json:"owners,omitempty"`

    Remark  *string `json:"remark,omitempty"`

}

type InstanceDeliveryData struct {

    ProjectKey  *string `json:"project_key,omitempty"`

    WorkItemTypeKey  *string `json:"work_item_type_key,omitempty"`

    ResourceID  *int64 `json:"resource_id,omitempty"`

}

type Leader struct {

    Email  *string `json:"email,omitempty"`

    EmployeeNumber  *string `json:"employee_number,omitempty"`

    ID  *string `json:"id,omitempty"`

    Name  *string `json:"name,omitempty"`

    UName  *string `json:"u_name,omitempty"`

}

type LinkedResourceInfo struct {

    ResourceID  *int64 `json:"resource_id,omitempty"`

      *int64 `json:",omitempty"`

}

type ManHourRecord struct {

    ID  *int64 `json:"id,omitempty"`

    RelatedWorkItemID  *int64 `json:"related_work_item_id,omitempty"`

    RelatedWorkItemTypeKey  *string `json:"related_work_item_type_key,omitempty"`

    RelatedWorkItemName  *string `json:"related_work_item_name,omitempty"`

    ResourceType  *string `json:"resource_type,omitempty"`

    ResourceID  *string `json:"resource_id,omitempty"`

    WorkDescription  *string `json:"work_description,omitempty"`

    WorkTime  *float64 `json:"work_time,omitempty"`

    WorkUserKey  *string `json:"work_user_key,omitempty"`

    CreatedAt  *int64 `json:"created_at,omitempty"`

    UpdatedAt  *int64 `json:"updated_at,omitempty"`

    ResourceName  *string `json:"resource_name,omitempty"`

    WorkDate  *int64 `json:"work_date,omitempty"`

}

type MultiSignal struct {

    Status  *string `json:"status,omitempty"`

    Detail  []MultiSignalDetail `json:"detail,omitempty"`

}

type MultiSignalDetail struct {

    ID  *string `json:"id,omitempty"`

    Title  *string `json:"title,omitempty"`

    Status  *string `json:"status,omitempty"`

    ViewLink  *string `json:"view_link,omitempty"`

    QueryLink  *QueryLink `json:"query_link,omitempty"`

}

type MultiText struct {

    FieldKey  *string `json:"field_key,omitempty"`

    FieldValue  *MultiTextDetail `json:"field_value,omitempty"`

}

type MultiTextDetail struct {

    Doc  *string `json:"doc,omitempty"`

    DocText  *string `json:"doc_text,omitempty"`

    IsEmpty  *bool `json:"is_empty,omitempty"`

    NotifyUserList  []string `json:"notify_user_list,omitempty"`

    NotifyUserType  *string `json:"notify_user_type,omitempty"`

    DocHTML  *string `json:"doc_html,omitempty"`

}

type NodeBasicInfo struct {

    ID  *string `json:"id,omitempty"`

    Name  *string `json:"name,omitempty"`

    Owners  []string `json:"owners,omitempty"`

    Milestone  *bool `json:"milestone,omitempty"`

}

type NodeConf struct {

    StateKey  *string `json:"state_key,omitempty"`

    NodeName  *string `json:"node_name,omitempty"`

    NodeTags  []string `json:"node_tags,omitempty"`

    NodeType  *string `json:"node_type,omitempty"`

    IsVisibility  *int32 `json:"is_visibility,omitempty"`

    NeedSchedule  *bool `json:"need_schedule,omitempty"`

    Owner  *OwnerConf `json:"owner,omitempty"`

    WbsStatusMap  *WbsStatusMap `json:"wbs_status_map,omitempty"`

    NodeSubProcess  *SubProcessConf `json:"node_sub_process,omitempty"`

    WbsNodeMap  *WbsNodeMap `json:"wbs_node_map,omitempty"`

}

type NodeElement struct {

    ElementKey  *string `json:"element_key,omitempty"`

    Name  *string `json:"name,omitempty"`

}

type NodeField struct {

    FieldKey  *string `json:"field_key,omitempty"`

    FieldValue  interface{} `json:"field_value,omitempty"`

    FieldAlias  *string `json:"field_alias,omitempty"`

    FieldTypeKey  *string `json:"field_type_key,omitempty"`

}

type NodeFinishedConclusion struct {

    FinishedConclusionResult  *NodeFinishedConclusionOption `json:"finished_conclusion_result,omitempty"`

    OwnersFinishedConclusionResult  []NodeOwnerFinishedConclusion `json:"owners_finished_conclusion_result,omitempty"`

}

type NodeFinishedConclusionOption struct {

    Key  *string `json:"key,omitempty"`

    Label  *string `json:"label,omitempty"`

    OriginalLabel  *string `json:"original_label,omitempty"`

}

type NodeFinishedOpinion struct {

    FinishedOpinionResult  *string `json:"finished_opinion_result,omitempty"`

    OwnersFinishedOpinionResult  []NodeOwnerFinishedOpinion `json:"owners_finished_opinion_result,omitempty"`

}

type NodeOwnerFinishedConclusion struct {

    Owner  *Owner `json:"owner,omitempty"`

    FinishedConclusionResult  *NodeFinishedConclusionOption `json:"finished_conclusion_result,omitempty"`

}

type NodeOwnerFinishedOpinion struct {

    Owner  *Owner `json:"owner,omitempty"`

    FinishedOpinionResult  *string `json:"finished_opinion_result,omitempty"`

}

type NodeRequiredItemRes struct {

    FormItems  []RequiredFormItem `json:"form_items,omitempty"`

    Tasks  []RequiredTask `json:"tasks,omitempty"`

    Deliverables  []RequiredDeliverable `json:"deliverables,omitempty"`

    NodeFields  []RequiredField `json:"node_fields,omitempty"`

}

type NodeTask struct {

    ID  *string `json:"id,omitempty"`

    StateKey  *string `json:"state_key,omitempty"`

    SubTasks  []WorkItem_work_item_SubTask `json:"sub_tasks,omitempty"`

    NodeName  *string `json:"node_name,omitempty"`

    TemplateID  *int64 `json:"template_id,omitempty"`

    Version  *int64 `json:"version,omitempty"`

}

type NodeWBSRoleOwners struct {

    Path  *string `json:"path,omitempty"`

    WorkItemID  *int64 `json:"work_item_id,omitempty"`

    Name  *string `json:"name,omitempty"`

    RoleOwners  []WorkItem_work_item_RoleOwner `json:"role_owners,omitempty"`

    WorkItemTypeKey  *string `json:"work_item_type_key,omitempty"`

}

type NodesConnections struct {

    WorkflowNodes  []WorkflowNode `json:"workflow_nodes,omitempty"`

    Connections  []Connection `json:"connections,omitempty"`

    StateFlowNodes  []StateFlowNode `json:"state_flow_nodes,omitempty"`

    UserDetails  []UserDetail `json:"user_details,omitempty"`

}

type NumberConfig struct {

    ScalingRatio  *string `json:"scaling_ratio,omitempty"`

    DisplayDigits  *int64 `json:"display_digits,omitempty"`

    SymbolSetting  *SymbolSetting `json:"symbol_setting,omitempty"`

    Thousandth  *bool `json:"thousandth,omitempty"`

}

type OAPIBatchQueryConclusionOptionItem struct {

    NodeID  *string `json:"node_id,omitempty"`

    FinishedConclusionOption  []OAPIFinishedConclusionResultItem `json:"finished_conclusion_option,omitempty"`

    FinishedOwnersConclusionOption  []OAPIFinishedConclusionResultItem `json:"finished_owners_conclusion_option,omitempty"`

    FinishedOverallConclusionOption  []OAPIFinishedConclusionResultItem `json:"finished_overall_conclusion_option,omitempty"`

}

type OAPIBatchQueryDeliverable struct {

    DeliverableUUID  *string `json:"deliverable_uuid,omitempty"`

    DeliverableType  *string `json:"deliverable_type,omitempty"`

    DeliverableInfo  *OAPIBatchQueryDeliverableInfo `json:"deliverable_info,omitempty"`

}

type OAPIBatchQueryDeliverableInfo struct {

    WorkItemID  *int64 `json:"work_item_id,omitempty"`

    Name  *string `json:"name,omitempty"`

    InstanceLinkedVirtualResourceWorkItem  *int64 `json:"instance_linked_virtual_resource_workitem,omitempty"`

    TemplateResources  *bool `json:"template_resources,omitempty"`

    TemplateType  *string `json:"template_type,omitempty"`

    Deleted  *bool `json:"deleted,omitempty"`

    DeliveryRelatedInfo  *DeliveryRelatedInfo `json:"delivery_related_info,omitempty"`

}

type OAPICreateWorkItemInfo struct {

    WorkItemID  *int64 `json:"work_item_id,omitempty"`

}

type OAPIFinishedConclusionInfo struct {

    FinishedConclusionResult  *OAPIFinishedConclusionResultItem `json:"finished_conclusion_result,omitempty"`

    OwnersFinishedConclusionResult  []OAPIFinishedConclusionOwnersResultItem `json:"owners_finished_conclusion_result,omitempty"`

}

type OAPIFinishedConclusionOwnersResultItem struct {

    Owner  *string `json:"owner,omitempty"`

    OwnersFinishedConclusionResult  *OAPIFinishedConclusionResultItem `json:"owners_finished_conclusion_result,omitempty"`

}

type OAPIFinishedConclusionResultItem struct {

    Key  *string `json:"key,omitempty"`

    Label  *string `json:"label,omitempty"`

    OriginLabel  *string `json:"origin_label,omitempty"`

}

type OAPIFinishedInfoItem struct {

    NodeID  *string `json:"node_id,omitempty"`

    SummaryMode  *string `json:"summary_mode,omitempty"`

    Opinion  *OAPIFinishedOpinionInfo `json:"opinion,omitempty"`

    Conclusion  *OAPIFinishedConclusionInfo `json:"conclusion,omitempty"`

}

type OAPIFinishedOpinionInfo struct {

    FinishedOpinionResult  *string `json:"finished_opinion_result,omitempty"`

    OwnersFinishedOpinionResult  []OAPIFinishedOpinionOwnersResultItem `json:"owners_finished_opinion_result,omitempty"`

}

type OAPIFinishedOpinionOwnersResultItem struct {

    Owner  *string `json:"owner,omitempty"`

    FinishedOpinionResult  *string `json:"finished_opinion_result,omitempty"`

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

    Extra  map[string][]string `json:"extra,omitempty"`

}

type OAPIResourceCreateInstanceResponseData struct {

    WorkItemID  *int64 `json:"work_item_id,omitempty"`

    IgnoreCreateInfo  *OAPIResourceCreateInstanceResponseDataIgnoreCreateInfo `json:"ignore_create_info,omitempty"`

}

type OAPIResourceCreateInstanceResponseDataIgnoreCreateInfo struct {

    FieldKeys  []string `json:"field_keys,omitempty"`

    RoleIDs  []string `json:"role_ids,omitempty"`

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

type Operation struct {

    UUID  *string `json:"uuid,omitempty"`

    OperationType  *string `json:"operation_type,omitempty"`

    OperationTarget  *string `json:"operation_target,omitempty"`

    Value  *OperationValue `json:"value,omitempty"`

    WorkItemID  *int64 `json:"work_item_id,omitempty"`

}

type OperationCreate struct {

    ElementKey  *string `json:"element_key,omitempty"`

    Name  *string `json:"name,omitempty"`

    ID  *int64 `json:"id,omitempty"`

    SeqOrderInfo  *SeqOrderInfo `json:"seq_order_info,omitempty"`

    WbsStatus  *string `json:"wbs_status,omitempty"`

}

type OperationValue struct {

    Name  *string `json:"name,omitempty"`

    RoleOwners  []WorkItem_work_item_RoleOwner `json:"role_owners,omitempty"`

    OperationCreate  *OperationCreate `json:"operation_create,omitempty"`

    ChildOrders  []ChildOrder `json:"child_orders,omitempty"`

    DeletedUuids  []string `json:"deleted_uuids,omitempty"`

    DeleteLinkageUuids  *bool `json:"delete_linkage_uuids,omitempty"`

    Reason  *string `json:"reason,omitempty"`

    Schedule  *WorkItem_work_item_Schedule `json:"schedule,omitempty"`

    Schedules  []WorkItem_work_item_Schedule `json:"schedules,omitempty"`

    WbsWorkItem  *WBSWorkItem `json:"wbs_work_item,omitempty"`

    Connections  []WorkItem_work_item_Connection `json:"connections,omitempty"`

    RestoreUuids  []string `json:"restore_uuids,omitempty"`

    WbsStatus  *string `json:"wbs_status,omitempty"`

    SeqOrderInfo  *SeqOrderInfo `json:"seq_order_info,omitempty"`

    SubInstanceCreate  *SubInstanceCreate `json:"sub_instance_create,omitempty"`

    DismantleMode  *int64 `json:"dismantle_mode,omitempty"`

    UnionDeliveries  []UnionDelivery `json:"union_deliveries,omitempty"`

}

type Option struct {

    Label  *string `json:"label,omitempty"`

    Value  *string `json:"value,omitempty"`

    Children  []Option `json:"children,omitempty"`

    WorkItemTypeKey  *string `json:"work_item_type_key,omitempty"`

    Order  *int64 `json:"order,omitempty"`

    Color  *string `json:"color,omitempty"`

    IsVisibility  *int64 `json:"is_visibility,omitempty"`

    IsDisabled  *int64 `json:"is_disabled,omitempty"`

}

type OptionConf struct {

    Label  *string `json:"label,omitempty"`

    Value  *string `json:"value,omitempty"`

    IsDisabled  *int32 `json:"is_disabled,omitempty"`

    IsVisibility  *int32 `json:"is_visibility,omitempty"`

    Children  []OptionConf `json:"children,omitempty"`

}

type OrderInfo struct {

    ValueType  *string `json:"value_type,omitempty"`

    Vaule  *string `json:"vaule,omitempty"`

    Path  *string `json:"path,omitempty"`

}

type Owner struct {

    ID  *string `json:"id,omitempty"`

    Username  *string `json:"username,omitempty"`

    EmployeeId  *string `json:"employeeId,omitempty"`

    Email  *string `json:"email,omitempty"`

    Nickname  *string `json:"nickname,omitempty"`

    Avatar  *string `json:"avatar,omitempty"`

    AvatarUrl  *string `json:"avatar_url,omitempty"`

    EnName  *string `json:"en_name,omitempty"`

    Name  *string `json:"name,omitempty"`

    LarkOpenId  *string `json:"larkOpenId,omitempty"`

    Department  *Department `json:"department,omitempty"`

    Leader  *Leader `json:"leader,omitempty"`

}

type OwnerConf struct {

    OwnerUsageMode  *string `json:"owner_usage_mode,omitempty"`

    OwnerRoles  []string `json:"owner_roles,omitempty"`

    UserKeys  []string `json:"user_keys,omitempty"`

}

type Pagination struct {

    PageNum  *int64 `json:"page_num,omitempty"`

    PageSize  *int64 `json:"page_size,omitempty"`

    Total  *int64 `json:"total,omitempty"`

}

type ParentWorkItem struct {

    RelationUUID  *string `json:"relation_uuid,omitempty"`

    ParentWorkItemID  *int64 `json:"parent_work_item_id,omitempty"`

    ParentNodeKey  *string `json:"parent_node_key,omitempty"`

}

type ProjectRelationRule struct {

    RemoteProjectKey  *string `json:"remote_project_key,omitempty"`

    RemoteProjectName  *string `json:"remote_project_name,omitempty"`

    Rules  []RelationRule `json:"rules,omitempty"`

}

type Query struct {

    Filter  *Filter `json:"filter,omitempty"`

    Sorts  []Sort `json:"sorts,omitempty"`

    Groups  []Group `json:"groups,omitempty"`

    QueryType  *int32 `json:"query_type,omitempty"`

    Pagination  *Search_apicommon_Pagination `json:"pagination,omitempty"`

    Projection  map[string]int64 `json:"projection,omitempty"`

    InstantQuery  *bool `json:"instant_query,omitempty"`

    FindAborted  *bool `json:"find_aborted,omitempty"`

    FindDeleted  *bool `json:"find_deleted,omitempty"`

    NeedCheckDirtyFields  *bool `json:"need_check_dirty_fields,omitempty"`

    ProjectKeys  []string `json:"project_keys,omitempty"`

    WorkItemTypeKeys  []string `json:"work_item_type_keys,omitempty"`

    AllowTruncate  *bool `json:"allow_truncate,omitempty"`

}

type QueryLink struct {

    Url  *string `json:"url,omitempty"`

    Method  *string `json:"method,omitempty"`

    Headers  interface{} `json:"headers,omitempty"`

    Body  interface{} `json:"body,omitempty"`

    Params  interface{} `json:"params,omitempty"`

}

type QueryStoryRelationData struct {

    WorkItemProjectMap  map[int64]string `json:"work_item_project_map,omitempty"`

    Value  map[int64][]int64 `json:"value,omitempty"`

}

type RelatedFieldExtraDisplayInfo struct {

    ProjectKey  *string `json:"project_key,omitempty"`

    WorkItemTypeKey  *string `json:"work_item_type_key,omitempty"`

    DisplayFieldKeys  []string `json:"display_field_keys,omitempty"`

    DisplayRoleKeys  []string `json:"display_role_keys,omitempty"`

    DisplayControls  []string `json:"display_controls,omitempty"`

}

type RelationBindInstance struct {

    ProjectKey  *string `json:"project_key,omitempty"`

    WorkItemTypeKey  *string `json:"work_item_type_key,omitempty"`

    WorkItemID  *int64 `json:"work_item_id,omitempty"`

    ChatGroupMerge  *int64 `json:"chat_group_merge,omitempty"`

}

type RelationCaseDetailForUpdate struct {

    RelationUUID  *string `json:"relation_uuid,omitempty"`

    ParentWorkItemID  *int64 `json:"parent_work_item_id,omitempty"`

    NodeKey  *string `json:"node_key,omitempty"`

    RelationConfUUID  *string `json:"relation_conf_uuid,omitempty"`

}

type RelationDetail struct {

    WorkItemTypeKey  *string `json:"work_item_type_key,omitempty"`

    WorkItemTypeName  *string `json:"work_item_type_name,omitempty"`

    ProjectKey  *string `json:"project_key,omitempty"`

    ProjectName  *string `json:"project_name,omitempty"`

}

type RelationFieldDetail struct {

    FieldKey  *string `json:"field_key,omitempty"`

    Detail  []FieldDetail `json:"detail,omitempty"`

}

type RelationInstance struct {

    RelationWorkItemID  *int64 `json:"relation_work_item_id,omitempty"`

    RelationWorkItemName  *string `json:"relation_work_item_name,omitempty"`

    RelationWorkItemTypeName  *string `json:"relation_work_item_type_name,omitempty"`

    RelationWorkItemTypeKey  *string `json:"relation_work_item_type_key,omitempty"`

    ProjectRelationRuleID  *string `json:"project_relation_rule_id,omitempty"`

    ProjectRelationRuleName  *string `json:"project_relation_rule_name,omitempty"`

    RelationProjectKey  *string `json:"relation_project_key,omitempty"`

    RelationProjectName  *string `json:"relation_project_name,omitempty"`

}

type RelationRule struct {

    ID  *string `json:"id,omitempty"`

    Name  *string `json:"name,omitempty"`

    Disabled  *int64 `json:"disabled,omitempty"`

    WorkItemRelationID  *string `json:"work_item_relation_id,omitempty"`

    WorkItemRelationName  *string `json:"work_item_relation_name,omitempty"`

    CurrentWorkItemTypeKey  *string `json:"current_work_item_type_key,omitempty"`

    CurrentWorkItemTypeName  *string `json:"current_work_item_type_name,omitempty"`

    RemoteWorkItemTypeKey  *string `json:"remote_work_item_type_key,omitempty"`

    RemoteWorkItemTypeName  *string `json:"remote_work_item_type_name,omitempty"`

    ChatGroupMerge  *int64 `json:"chat_group_merge,omitempty"`

}

type RequiredDeliverable struct {

    DeliverableID  *int64 `json:"deliverable_id,omitempty"`

    Finished  *bool `json:"finished,omitempty"`

}

type RequiredField struct {

    FieldKey  *string `json:"field_key,omitempty"`

    FieldTypeKey  *string `json:"field_type_key,omitempty"`

    Finished  *bool `json:"finished,omitempty"`

    NotFinishedOwner  []string `json:"not_finished_owner,omitempty"`

    SubField  []RequiredField `json:"sub_field,omitempty"`

}

type RequiredFormItem struct {

    Class  *string `json:"class,omitempty"`

    Key  *string `json:"key,omitempty"`

    Finished  *bool `json:"finished,omitempty"`

    FieldTypeKey  *string `json:"field_type_key,omitempty"`

    NotFinishedOwner  []string `json:"not_finished_owner,omitempty"`

    SubField  []RequiredField `json:"sub_field,omitempty"`

    StateInfo  *RequiredStateInfo `json:"state_info,omitempty"`

}

type RequiredStateInfo struct {

    StateKey  *string `json:"state_key,omitempty"`

    NodeFields  []RequiredField `json:"node_fields,omitempty"`

}

type RequiredTask struct {

    TaskID  *int64 `json:"task_id,omitempty"`

    Finished  *bool `json:"finished,omitempty"`

}

type ResourceLibSetting struct {

    EnableFields  []SimpleField `json:"enable_fields,omitempty"`

    EnableRoles  []SimpleRoleConf `json:"enable_roles,omitempty"`

}

type ResourceWorkItemInfo struct {

    WorkItemID  *int64 `json:"work_item_id,omitempty"`

    Name  *string `json:"name,omitempty"`

    WorkItemTypeKey  *string `json:"work_item_type_key,omitempty"`

    ProjectKey  *string `json:"project_key,omitempty"`

    TemplateType  *string `json:"template_type,omitempty"`

    CreatedBy  *UserDetail `json:"created_by,omitempty"`

    UpdatedBy  *UserDetail `json:"updated_by,omitempty"`

    CreatedAt  *int64 `json:"created_at,omitempty"`

    UpdatedAt  *int64 `json:"updated_at,omitempty"`

    Owner  *UserDetail `json:"owner,omitempty"`

    RoleOwners  []Common_RoleOwner `json:"role_owners,omitempty"`

    TemplateVersion  *int64 `json:"template_version,omitempty"`

    Template  *Template `json:"template,omitempty"`

}

type RoleAssign struct {

    Role  *string `json:"role,omitempty"`

    Name  *string `json:"name,omitempty"`

    DefaultAppear  *int32 `json:"default_appear,omitempty"`

    Deletable  *int32 `json:"deletable,omitempty"`

    MemberAssign  *int32 `json:"member_assign,omitempty"`

    Members  []string `json:"members,omitempty"`

}

type RoleConfCreate struct {

    ID  *string `json:"id,omitempty"`

    Name  *string `json:"name,omitempty"`

    IsOwner  *bool `json:"is_owner,omitempty"`

    AutoEnterGroup  *bool `json:"auto_enter_group,omitempty"`

    MemberAssignMode  *int32 `json:"member_assign_mode,omitempty"`

    Members  []string `json:"members,omitempty"`

    IsMemberMulti  *bool `json:"is_member_multi,omitempty"`

    RoleAlias  *string `json:"role_alias,omitempty"`

    LockScope  []string `json:"lock_scope,omitempty"`

}

type RoleConfDetail struct {

    ID  *string `json:"id,omitempty"`

    Name  *string `json:"name,omitempty"`

    IsOwner  *bool `json:"is_owner,omitempty"`

    RoleAppearMode  *int64 `json:"role_appear_mode,omitempty"`

    Deletable  *bool `json:"deletable,omitempty"`

    AutoEnterGroup  *bool `json:"auto_enter_group,omitempty"`

    MemberAssignMode  *int64 `json:"member_assign_mode,omitempty"`

    Members  []string `json:"members,omitempty"`

    Key  *string `json:"key,omitempty"`

    AllowDelete  *bool `json:"allow_delete,omitempty"`

    AuthorizationRoleKeys  []string `json:"authorization_role_keys,omitempty"`

    IsMemberMulti  *bool `json:"is_member_multi,omitempty"`

    IsRequired  *int32 `json:"is_required,omitempty"`

    RoleAlias  *string `json:"role_alias,omitempty"`

    LockAppID  *string `json:"lock_app_id,omitempty"`

    LockScope  []string `json:"lock_scope,omitempty"`

}

type RoleConfUpdate struct {

    Name  *string `json:"name,omitempty"`

    IsOwner  *bool `json:"is_owner,omitempty"`

    AutoEnterGroup  *bool `json:"auto_enter_group,omitempty"`

    MemberAssignMode  *int32 `json:"member_assign_mode,omitempty"`

    Members  []string `json:"members,omitempty"`

    IsMemberMulti  *bool `json:"is_member_multi,omitempty"`

    RoleAlias  *string `json:"role_alias,omitempty"`

    LockScope  []string `json:"lock_scope,omitempty"`

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

type ScheduleConstraintRule struct {

    SubTask  *bool `json:"sub_task,omitempty"`

    Node  *bool `json:"node,omitempty"`

    SubProcessNode  *bool `json:"sub_process_node,omitempty"`

    WbsSubInstanceType  map[string]bool `json:"wbs_sub_instance_type,omitempty"`

}

type ScheduleReferenceValue struct {

    WorkItemID  *int64 `json:"work_item_id,omitempty"`

    StateKey  *string `json:"state_key,omitempty"`

    Name  *string `json:"name,omitempty"`

    Type  *string `json:"type,omitempty"`

}

type SearchGroup struct {

    SearchParams  []SearchParam `json:"search_params,omitempty"`

    Conjunction  *string `json:"conjunction,omitempty"`

    SearchGroups  []SearchGroup `json:"search_groups,omitempty"`

}

type SearchParam struct {

    ParamKey  *string `json:"param_key,omitempty"`

    Value  interface{} `json:"value,omitempty"`

    Operator  *string `json:"operator,omitempty"`

    PreOperator  *string `json:"pre_operator,omitempty"`

    ValueSearchGroups  *SearchGroup `json:"value_search_groups,omitempty"`

}

type SearchUser struct {

    UserKeys  []string `json:"user_keys,omitempty"`

    FieldKey  *string `json:"field_key,omitempty"`

    Role  *string `json:"role,omitempty"`

}

type Search_SearchGroup struct {

    SearchParams  []Search_SearchParam `json:"search_params,omitempty"`

    Conjunction  *string `json:"conjunction,omitempty"`

    SearchGroups  []Search_SearchGroup `json:"search_groups,omitempty"`

}

type Search_SearchParam struct {

    ParamKey  *string `json:"param_key,omitempty"`

    Value  interface{} `json:"value,omitempty"`

    Operator  *string `json:"operator,omitempty"`

    PreOperator  *string `json:"pre_operator,omitempty"`

    ValueSearchGroups  *Search_SearchGroup `json:"value_search_groups,omitempty"`

}

type Search_apicommon_Pagination struct {

    PageNum  *int64 `json:"page_num,omitempty"`

    PageSize  *int64 `json:"page_size,omitempty"`

    SearchAfter  *string `json:"search_after,omitempty"`

    Pit  *string `json:"pit,omitempty"`

}

type Search_concisesearch_Pagination struct {

    PageSize  *int64 `json:"page_size,omitempty"`

    SearchAfter  *string `json:"search_after,omitempty"`

    Pit  *string `json:"pit,omitempty"`

    Total  *int64 `json:"total,omitempty"`

}

type Search_concisesearch_Sort struct {

    FieldKey  *string `json:"field_key,omitempty"`

    FieldType  *string `json:"field_type,omitempty"`

    WorkItemTypeKey  *string `json:"work_item_type_key,omitempty"`

    Order  *string `json:"order,omitempty"`

    Params  map[string]string `json:"params,omitempty"`

}

type SeqOrderInfo struct {

    Pre  []OrderInfo `json:"pre,omitempty"`

    Post  []OrderInfo `json:"post,omitempty"`

    Parent  *OrderInfo `json:"parent,omitempty"`

}

type SimpleField struct {

    FieldKey  *string `json:"field_key,omitempty"`

    FieldTypeKey  *string `json:"field_type_key,omitempty"`

    Options  []Option `json:"options,omitempty"`

    CompoundFields  []SimpleField `json:"compound_fields,omitempty"`

    FieldAlias  *string `json:"field_alias,omitempty"`

    FieldName  *string `json:"field_name,omitempty"`

    IsCustomField  *bool `json:"is_custom_field,omitempty"`

    IsObsoleted  *bool `json:"is_obsoleted,omitempty"`

    WorkItemScopes  []string `json:"work_item_scopes,omitempty"`

    ValueGenerateMode  *string `json:"value_generate_mode,omitempty"`

    RelationID  *string `json:"relation_id,omitempty"`

    EditablePersonnelScope  *EditablePersonnelScope `json:"editable_personnel_scope,omitempty"`

    NumberConfig  *NumberConfig `json:"number_config,omitempty"`

}

type SimpleRoleConf struct {

    ID  *string `json:"id,omitempty"`

    Name  *string `json:"name,omitempty"`

    Key  *string `json:"key,omitempty"`

}

type Sort struct {

    FieldItem  *FieldItem `json:"field_item,omitempty"`

    Order  *string `json:"order,omitempty"`

}

type StateFlowConfInfo struct {

    StateKey  *string `json:"state_key,omitempty"`

    Name  *string `json:"name,omitempty"`

    StateType  *int64 `json:"state_type,omitempty"`

    AuthorizedRoles  []string `json:"authorized_roles,omitempty"`

    ConfirmFormList  []ConfirmForm `json:"confirm_form_list,omitempty"`

    Action  *int64 `json:"action,omitempty"`

}

type StateFlowNode struct {

    ID  *string `json:"id,omitempty"`

    Name  *string `json:"name,omitempty"`

    RoleOwners  []RoleOwner `json:"role_owners,omitempty"`

    Status  *int32 `json:"status,omitempty"`

    ActualBeginTime  *string `json:"actual_begin_time,omitempty"`

    ActualFinishTime  *string `json:"actual_finish_time,omitempty"`

    Fields  []FieldValuePair `json:"fields,omitempty"`

}

type StateTime struct {

    StateKey  *string `json:"state_key,omitempty"`

    StartTime  *int64 `json:"start_time,omitempty"`

    EndTime  *int64 `json:"end_time,omitempty"`

    Name  *string `json:"name,omitempty"`

}

type StatusConf struct {

    StatusKey  *string `json:"status_key,omitempty"`

    StatusName  *string `json:"status_name,omitempty"`

    StatusOrderIndex  *int32 `json:"status_order_index,omitempty"`

}

type StoryRelationEntity struct {

    LocalWorkItemID  *int64 `json:"local_work_item_id,omitempty"`

    RemoteWorkItemID  *int64 `json:"remote_work_item_id,omitempty"`

    RelationID  *string `json:"relation_id,omitempty"`

    MergeGroup  *bool `json:"merge_group,omitempty"`

}

type SubDetail struct {

    WorkItemID  *int64 `json:"work_item_id,omitempty"`

    WorkItemName  *string `json:"work_item_name,omitempty"`

    NodeID  *string `json:"node_id,omitempty"`

    SubTask  *SubTask `json:"sub_task,omitempty"`

}

type SubInstanceCreate struct {

    ProjectKey  *string `json:"project_key,omitempty"`

    WorkItemTypeKey  *string `json:"work_item_type_key,omitempty"`

    FieldValues  []WorkItem_work_item_FieldValuePair `json:"field_values,omitempty"`

    RelationCase  *RelationCaseDetailForUpdate `json:"relation_case,omitempty"`

    TemplateID  *int64 `json:"template_id,omitempty"`

    WorkItemID  *int64 `json:"work_item_id,omitempty"`

    LinkedResources  []LinkedResourceInfo `json:"linked_resources,omitempty"`

}

type SubProcessConf struct {

    TemplateKey  *string `json:"template_key,omitempty"`

    TemplateName  *string `json:"template_name,omitempty"`

    TemplateID  *string `json:"template_id,omitempty"`

    WorkItemTypeKey  *string `json:"work_item_type_key,omitempty"`

    FlowMode  *string `json:"flow_mode,omitempty"`

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

type SubTaskParentInfo struct {

    WorkItemID  *int64 `json:"work_item_id,omitempty"`

    WorkItemName  *string `json:"work_item_name,omitempty"`

    WorkItemTypeKey  *string `json:"work_item_type_key,omitempty"`

    NodeID  *string `json:"node_id,omitempty"`

}

type SubWorkItemConfInfo struct {

    Name  *string `json:"name,omitempty"`

    RelationUUID  *string `json:"relation_uuid,omitempty"`

    RelationName  *string `json:"relation_name,omitempty"`

    WorkItemTypeKey  *string `json:"work_item_type_key,omitempty"`

    WorkItemTypeName  *string `json:"work_item_type_name,omitempty"`

}

type SymbolSetting struct {

    Display  *string `json:"display,omitempty"`

    Value  *string `json:"value,omitempty"`

    Layout  *string `json:"layout,omitempty"`

}

type TargetState struct {

    StateKey  *string `json:"state_key,omitempty"`

    TransitionID  *int64 `json:"transition_id,omitempty"`

}

type TaskConf struct {

    Action  *int64 `json:"action,omitempty"`

    Name  *string `json:"name,omitempty"`

    ID  *string `json:"id,omitempty"`

    DeliverableFieldID  *string `json:"deliverable_field_id,omitempty"`

    PassMode  *int64 `json:"pass_mode,omitempty"`

    NodePassRequiredMode  *int64 `json:"node_pass_required_mode,omitempty"`

}

type TaskConfInfo struct {

    Name  *string `json:"name,omitempty"`

    OwnerUsageMode  *int64 `json:"owner_usage_mode,omitempty"`

    OwnerRoles  []string `json:"owner_roles,omitempty"`

    Owners  []string `json:"owners,omitempty"`

}

type TaskElement struct {

    ElementKey  *string `json:"element_key,omitempty"`

    Name  *string `json:"name,omitempty"`

}

type TeamDataScope struct {

    TeamID  *string `json:"team_id,omitempty"`

    Cascade  *bool `json:"cascade,omitempty"`

}

type TeamOption struct {

    TeamDataScopes  []TeamDataScope `json:"team_data_scopes,omitempty"`

    TeamMode  *string `json:"team_mode,omitempty"`

}

type Template struct {

    ID  *int64 `json:"id,omitempty"`

    StateFlowKey  *string `json:"state_flow_key,omitempty"`

    WorkFlowKey  *string `json:"work_flow_key,omitempty"`

}

type TemplateConf struct {

    TemplateID  *string `json:"template_id,omitempty"`

    TemplateName  *string `json:"template_name,omitempty"`

    IsDisabled  *int32 `json:"is_disabled,omitempty"`

    Version  *int64 `json:"version,omitempty"`

    UniqueKey  *string `json:"unique_key,omitempty"`

    TemplateKey  *string `json:"template_key,omitempty"`

}

type TemplateDetail struct {

    WorkflowConfs  []WorkflowConfInfo `json:"workflow_confs,omitempty"`

    StateFlowConfs  []StateFlowConfInfo `json:"state_flow_confs,omitempty"`

    Connections  []WorkItem_work_item_Connection `json:"connections,omitempty"`

    TemplateID  *int64 `json:"template_id,omitempty"`

    TemplateName  *string `json:"template_name,omitempty"`

    Version  *int64 `json:"version,omitempty"`

    IsDisabled  *int64 `json:"is_disabled,omitempty"`

    WorkItemTypeKey  *string `json:"work_item_type_key,omitempty"`

}

type TimeInterval struct {

    Start  *int64 `json:"start,omitempty"`

    End  *int64 `json:"end,omitempty"`

}

type UnionDeliverable struct {

    FieldDeliverables  []FieldDeliverableItem `json:"field_deliverables,omitempty"`

    InstanceDeliverables  []InstanceDeliverableItem `json:"instance_deliverables,omitempty"`

}

type UnionDelivery struct {

    Type  *string `json:"type,omitempty"`

    FieldDelivery  *FieldDeliveryData `json:"field_delivery,omitempty"`

    InstanceDelivery  *InstanceDeliveryData `json:"instance_delivery,omitempty"`

}

type UpdateWorkingHourRecord struct {

    ID  *int64 `json:"id,omitempty"`

    WorkTime  *string `json:"work_time,omitempty"`

    WorkDescription  *string `json:"work_description,omitempty"`

}

type UserDetail struct {

    UserKey  *string `json:"user_key,omitempty"`

    Username  *string `json:"username,omitempty"`

    Email  *string `json:"email,omitempty"`

    NameCn  *string `json:"name_cn,omitempty"`

    NameEn  *string `json:"name_en,omitempty"`

}

type WBSInfo struct {

    TemplateKey  *string `json:"template_key,omitempty"`

    RelatedSubWorkItems  []WBSWorkItem `json:"related_sub_work_items,omitempty"`

    TemplateVersion  *int64 `json:"template_version,omitempty"`

    TemplateName  *string `json:"template_name,omitempty"`

    TemplateID  *string `json:"template_id,omitempty"`

    RelatedParentWorkItem  *WBSParentWorkItem `json:"related_parent_work_item,omitempty"`

    UserDetails  []WorkItem_work_item_UserDetail `json:"user_details,omitempty"`

    Connections  []WorkItem_work_item_Connection `json:"connections,omitempty"`

    RelationChainPath  *WBSRelationChainPath `json:"relation_chain_path,omitempty"`

    RelationChainEntity  *WBSRelationChainEntity `json:"relation_chain_entity,omitempty"`

}

type WBSParentWorkItem struct {

    IsTop  *bool `json:"is_top,omitempty"`

    WorkItemID  *int64 `json:"work_item_id,omitempty"`

    WorkItemTypeKey  *string `json:"work_item_type_key,omitempty"`

    Name  *string `json:"name,omitempty"`

    TemplateID  *string `json:"template_id,omitempty"`

    TemplateKey  *string `json:"template_key,omitempty"`

    TemplateName  *string `json:"template_name,omitempty"`

    TemplateVersion  *int64 `json:"template_version,omitempty"`

    RelationNodeID  *string `json:"relation_node_id,omitempty"`

    RelationNodeName  *string `json:"relation_node_name,omitempty"`

    RelationNodeTags  []string `json:"relation_node_tags,omitempty"`

    RelationNodeUUID  *string `json:"relation_node_uuid,omitempty"`

}

type WBSRelationChainEntity struct {

    WbsRelationChainEntity  []WBSRelationChainEntityItem `json:"wbs_relation_chain_entity,omitempty"`

    WorkItemID  *int64 `json:"work_item_id,omitempty"`

}

type WBSRelationChainEntityItem struct {

    ProjectKey  *string `json:"project_key,omitempty"`

    WorkItemTypeKey  *string `json:"work_item_type_key,omitempty"`

    WorkItemID  *int64 `json:"work_item_id,omitempty"`

    Type  *string `json:"type,omitempty"`

    WorkItemName  *string `json:"work_item_name,omitempty"`

    StateKey  *string `json:"state_key,omitempty"`

    NodeName  *string `json:"node_name,omitempty"`

    SubTaskID  *int64 `json:"sub_task_id,omitempty"`

    SubTaskName  *string `json:"sub_task_name,omitempty"`

    Level  *int32 `json:"level,omitempty"`

}

type WBSRelationChainPath struct {

    WbsRelationChainPath  []WBSRelationChainPathItem `json:"wbs_relation_chain_path,omitempty"`

    WorkItemID  *int64 `json:"work_item_id,omitempty"`

}

type WBSRelationChainPathItem struct {

    ProjectKey  *string `json:"project_key,omitempty"`

    Level  *int32 `json:"level,omitempty"`

    Type  *string `json:"type,omitempty"`

    WorkItemTypeKey  *string `json:"work_item_type_key,omitempty"`

    Name  *string `json:"name,omitempty"`

    NodeName  *string `json:"node_name,omitempty"`

}

type WBSWorkItem struct {

    NodeUUID  *string `json:"node_uuid,omitempty"`

    WorkItemID  *int64 `json:"work_item_id,omitempty"`

    Type  *string `json:"type,omitempty"`

    WbsStatus  *string `json:"wbs_status,omitempty"`

    WbsStatusMap  map[string]string `json:"wbs_status_map,omitempty"`

    SubWorkItem  []WBSWorkItem `json:"sub_work_item,omitempty"`

    Name  *string `json:"name,omitempty"`

    Deliverable  []WorkItem_work_item_FieldValuePair `json:"deliverable,omitempty"`

    Schedule  *WorkItem_work_item_Schedule `json:"schedule,omitempty"`

    Schedules  []WorkItem_work_item_Schedule `json:"schedules,omitempty"`

    Points  *float64 `json:"points,omitempty"`

    RoleOwners  []WorkItem_work_item_RoleOwner `json:"role_owners,omitempty"`

    WorkItemTypeKey  *string `json:"work_item_type_key,omitempty"`

    Milestone  *bool `json:"milestone,omitempty"`

    Connections  []WorkItem_work_item_Connection `json:"connections,omitempty"`

    StateKey  *string `json:"state_key,omitempty"`

    StartTime  *int64 `json:"start_time,omitempty"`

    FinishTime  *int64 `json:"finish_time,omitempty"`

    OwnerConf  *OwnerConf `json:"owner_conf,omitempty"`

    IsDraftCreate  *bool `json:"is_draft_create,omitempty"`

    NeedActualWorkTime  *bool `json:"need_actual_work_time,omitempty"`

    NodeWbsRoleOwners  []NodeWBSRoleOwners `json:"node_wbs_role_owners,omitempty"`

    TaskType  *int64 `json:"task_type,omitempty"`

    UUIDInConf  *string `json:"uuid_in_conf,omitempty"`

    UnionDeliverable  *UnionDeliverable `json:"union_deliverable,omitempty"`

    SeqOrderInfo  *SeqOrderInfo `json:"seq_order_info,omitempty"`

    DismantleMode  *int64 `json:"dismantle_mode,omitempty"`

    Status  *int64 `json:"status,omitempty"`

    Dependencies  []DependencyInfo `json:"dependencies,omitempty"`

    RelativeScheduleV2  []ScheduleReferenceValue `json:"relative_schedule_v2,omitempty"`

    IsScheduleAggItem  *bool `json:"is_schedule_agg_item,omitempty"`

    FieldValues  []WorkItem_work_item_FieldValuePair `json:"field_values,omitempty"`

}

type WbsDraft struct {

    TemplateUUID  *string `json:"template_uuid,omitempty"`

    TemplateVersion  *int64 `json:"template_version,omitempty"`

    WorkItemTypeKey  *string `json:"work_item_type_key,omitempty"`

    RelatedSubWorkItems  []WBSWorkItem `json:"related_sub_work_items,omitempty"`

    Connections  []WorkItem_work_item_Connection `json:"connections,omitempty"`

    DraftID  *string `json:"draft_id,omitempty"`

    DeleteUuids  []string `json:"delete_uuids,omitempty"`

    RoleOwners  map[int64][]WorkItem_work_item_RoleOwner `json:"role_owners,omitempty"`

    UserDetails  []WorkItem_work_item_UserDetail `json:"user_details,omitempty"`

    WorkItemID  *int64 `json:"work_item_id,omitempty"`

}

type WbsNodeMap struct {

    StateKey  *string `json:"state_key,omitempty"`

    StatusName  *string `json:"status_name,omitempty"`

}

type WbsStatusMap struct {

    StatusKey  *string `json:"status_key,omitempty"`

    StatusName  *string `json:"status_name,omitempty"`

}

type WbsTemplate struct {

    TemplateKey  *string `json:"template_key,omitempty"`

    TemplateName  *string `json:"template_name,omitempty"`

    TemplateID  *string `json:"template_id,omitempty"`

    IsDisabled  *bool `json:"is_disabled,omitempty"`

    Version  *int64 `json:"version,omitempty"`

    WorkflowConf  *WorkflowConf `json:"workflow_conf,omitempty"`

}

type WorkItemInfo struct {

    ID  *int64 `json:"id,omitempty"`

    Name  *string `json:"name,omitempty"`

    WorkItemTypeKey  *string `json:"work_item_type_key,omitempty"`

    ProjectKey  *string `json:"project_key,omitempty"`

    TemplateType  *string `json:"template_type,omitempty"`

    Pattern  *string `json:"pattern,omitempty"`

    SubStage  *string `json:"sub_stage,omitempty"`

    CurrentNodes  []NodeBasicInfo `json:"current_nodes,omitempty"`

    CreatedBy  *string `json:"created_by,omitempty"`

    UpdatedBy  *string `json:"updated_by,omitempty"`

    CreatedAt  *int64 `json:"created_at,omitempty"`

    UpdatedAt  *int64 `json:"updated_at,omitempty"`

    Fields  []FieldValuePair `json:"fields,omitempty"`

    WorkItemStatus  *WorkItemStatus `json:"work_item_status,omitempty"`

    DeletedBy  *string `json:"deleted_by,omitempty"`

    DeletedAt  *int64 `json:"deleted_at,omitempty"`

    SimpleName  *string `json:"simple_name,omitempty"`

    TemplateID  *int64 `json:"template_id,omitempty"`

    WorkflowInfos  *NodesConnections `json:"workflow_infos,omitempty"`

    StateTimes  []StateTime `json:"state_times,omitempty"`

    MultiTexts  []MultiText `json:"multi_texts,omitempty"`

    RelationFieldsDetail  []RelationFieldDetail `json:"relation_fields_detail,omitempty"`

    UserDetails  []UserDetail `json:"user_details,omitempty"`

    SubTaskParentInfo  *SubTaskParentInfo `json:"sub_task_parent_info,omitempty"`

}

type WorkItemKeyType struct {

    TypeKey  *string `json:"type_key,omitempty"`

    Name  *string `json:"name,omitempty"`

    IsDisable  *int32 `json:"is_disable,omitempty"`

    APIName  *string `json:"api_name,omitempty"`

    EnableModelResourceLib  *bool `json:"enable_model_resource_lib,omitempty"`

}

type WorkItemRelation struct {

    ID  *string `json:"id,omitempty"`

    Name  *string `json:"name,omitempty"`

    Disabled  *bool `json:"disabled,omitempty"`

    WorkItemTypeKey  *string `json:"work_item_type_key,omitempty"`

    WorkItemTypeName  *string `json:"work_item_type_name,omitempty"`

    RelationDetails  []RelationDetail `json:"relation_details,omitempty"`

    RelationType  *int64 `json:"relation_type,omitempty"`

}

type WorkItemStatus struct {

    StateKey  *string `json:"state_key,omitempty"`

    IsArchivedState  *bool `json:"is_archived_state,omitempty"`

    IsInitState  *bool `json:"is_init_state,omitempty"`

    UpdatedAt  *int64 `json:"updated_at,omitempty"`

    UpdatedBy  *string `json:"updated_by,omitempty"`

    History  []WorkItemStatus `json:"history,omitempty"`

}

type WorkItemTypeInfo struct {

    TypeKey  *string `json:"type_key,omitempty"`

    Name  *string `json:"name,omitempty"`

    FlowMode  *string `json:"flow_mode,omitempty"`

    APIName  *string `json:"api_name,omitempty"`

    Description  *string `json:"description,omitempty"`

    IsDisabled  *bool `json:"is_disabled,omitempty"`

    IsPinned  *bool `json:"is_pinned,omitempty"`

    EnableSchedule  *bool `json:"enable_schedule,omitempty"`

    ScheduleFieldKey  *string `json:"schedule_field_key,omitempty"`

    ScheduleFieldName  *string `json:"schedule_field_name,omitempty"`

    EstimatePointFieldKey  *string `json:"estimate_point_field_key,omitempty"`

    EstimatePointFieldName  *string `json:"estimate_point_field_name,omitempty"`

    ActualWorkTimeFieldKey  *string `json:"actual_work_time_field_key,omitempty"`

    ActualWorkTimeFieldName  *string `json:"actual_work_time_field_name,omitempty"`

    BelongRoles  []SimpleRoleConf `json:"belong_roles,omitempty"`

    ActualWorkTimeSwitch  *bool `json:"actual_work_time_switch,omitempty"`

    EnableModelResourceLib  *bool `json:"enable_model_resource_lib,omitempty"`

    ResourceLibSetting  *ResourceLibSetting `json:"resource_lib_setting,omitempty"`

}

type WorkItem_common_Pagination struct {

    PageNum  *int64 `json:"page_num,omitempty"`

    PageSize  *int64 `json:"page_size,omitempty"`

    Total  *int64 `json:"total,omitempty"`

}

type WorkItem_work_item_Checker struct {

    CheckedTime  *int64 `json:"checked_time,omitempty"`

    Owner  *string `json:"owner,omitempty"`

    Status  *int32 `json:"status,omitempty"`

}

type WorkItem_work_item_Connection struct {

    SourceStateKey  *string `json:"source_state_key,omitempty"`

    TargetStateKey  *string `json:"target_state_key,omitempty"`

    TransitionID  *int64 `json:"transition_id,omitempty"`

    Fields  []FieldConf `json:"fields,omitempty"`

}

type WorkItem_work_item_Expand struct {

    NeedWorkflow  *bool `json:"need_workflow,omitempty"`

    RelationFieldsDetail  *bool `json:"relation_fields_detail,omitempty"`

    NeedMultiText  *bool `json:"need_multi_text,omitempty"`

    NeedUserDetail  *bool `json:"need_user_detail,omitempty"`

    NeedSubTaskParent  *bool `json:"need_sub_task_parent,omitempty"`

    NeedUnionDeliverable  *bool `json:"need_union_deliverable,omitempty"`

    NeedWbsRelationChainEntity  *bool `json:"need_wbs_relation_chain_entity,omitempty"`

    NeedWbsRelationChainPath  *bool `json:"need_wbs_relation_chain_path,omitempty"`

    NeedGroupUUIDForCompound  *bool `json:"need_group_uuid_for_compound,omitempty"`

    NeedParentWorkItem  *bool `json:"need_parent_workitem,omitempty"`

}

type WorkItem_work_item_FieldDetail struct {

    StoryID  *int64 `json:"story_id,omitempty"`

    WorkItemTypeKey  *string `json:"work_item_type_key,omitempty"`

    ProjectKey  *string `json:"project_key,omitempty"`

}

type WorkItem_work_item_FieldValuePair struct {

    FieldKey  *string `json:"field_key,omitempty"`

    FieldValue  interface{} `json:"field_value,omitempty"`

    TargetState  *WorkItem_work_item_TargetState `json:"target_state,omitempty"`

    FieldTypeKey  *string `json:"field_type_key,omitempty"`

    FieldAlias  *string `json:"field_alias,omitempty"`

    HelpDescription  *string `json:"help_description,omitempty"`

    UpdateMode  *int32 `json:"update_mode,omitempty"`

}

type WorkItem_work_item_MultiText struct {

    FieldKey  *string `json:"field_key,omitempty"`

    FieldValue  *WorkItem_work_item_MultiTextDetail `json:"field_value,omitempty"`

}

type WorkItem_work_item_MultiTextDetail struct {

    Doc  *string `json:"doc,omitempty"`

    DocText  *string `json:"doc_text,omitempty"`

    IsEmpty  *bool `json:"is_empty,omitempty"`

    NotifyUserList  []string `json:"notify_user_list,omitempty"`

    NotifyUserType  *string `json:"notify_user_type,omitempty"`

    DocHTML  *string `json:"doc_html,omitempty"`

}

type WorkItem_work_item_NodeBasicInfo struct {

    ID  *string `json:"id,omitempty"`

    Name  *string `json:"name,omitempty"`

    Owners  []string `json:"owners,omitempty"`

    Milestone  *bool `json:"milestone,omitempty"`

}

type WorkItem_work_item_NodesConnections struct {

    WorkflowNodes  []WorkItem_work_item_WorkflowNode `json:"workflow_nodes,omitempty"`

    Connections  []WorkItem_work_item_Connection `json:"connections,omitempty"`

    StateFlowNodes  []WorkItem_work_item_StateFlowNode `json:"state_flow_nodes,omitempty"`

    UserDetails  []WorkItem_work_item_UserDetail `json:"user_details,omitempty"`

    TemplateID  *int64 `json:"template_id,omitempty"`

    Version  *int64 `json:"version,omitempty"`

}

type WorkItem_work_item_RelationFieldDetail struct {

    FieldKey  *string `json:"field_key,omitempty"`

    Detail  []WorkItem_work_item_FieldDetail `json:"detail,omitempty"`

}

type WorkItem_work_item_ResourceWorkItemInfo struct {

    ID  *int64 `json:"id,omitempty"`

    Name  *string `json:"name,omitempty"`

    WorkItemTypeKey  *string `json:"work_item_type_key,omitempty"`

    ProjectKey  *string `json:"project_key,omitempty"`

    TemplateType  *string `json:"template_type,omitempty"`

    CreatedBy  *string `json:"created_by,omitempty"`

    UpdatedBy  *string `json:"updated_by,omitempty"`

    CreatedAt  *int64 `json:"created_at,omitempty"`

    UpdatedAt  *int64 `json:"updated_at,omitempty"`

    Fields  []WorkItem_work_item_FieldValuePair `json:"fields,omitempty"`

    SimpleName  *string `json:"simple_name,omitempty"`

    TemplateID  *int64 `json:"template_id,omitempty"`

    MultiTexts  []WorkItem_work_item_MultiText `json:"multi_texts,omitempty"`

    RelationFieldsDetail  []WorkItem_work_item_RelationFieldDetail `json:"relation_fields_detail,omitempty"`

    UserDetails  []WorkItem_work_item_UserDetail `json:"user_details,omitempty"`

    ParentWorkItem  []ParentWorkItem `json:"parent_work_item,omitempty"`

}

type WorkItem_work_item_RoleOwner struct {

    Role  *string `json:"role,omitempty"`

    Name  *string `json:"name,omitempty"`

    Owners  []string `json:"owners,omitempty"`

    Exist  *bool `json:"exist,omitempty"`

    IsMemberMulti  *bool `json:"is_member_multi,omitempty"`

}

type WorkItem_work_item_Schedule struct {

    Points  *float64 `json:"points,omitempty"`

    EstimateStartDate  *int64 `json:"estimate_start_date,omitempty"`

    EstimateEndDate  *int64 `json:"estimate_end_date,omitempty"`

    Owners  []string `json:"owners,omitempty"`

    ActualWorkTime  *float64 `json:"actual_work_time,omitempty"`

    IsAuto  *bool `json:"is_auto,omitempty"`

    PlannedConstructionPeriod  *int32 `json:"planned_construction_period,omitempty"`

}

type WorkItem_work_item_StateFlowNode struct {

    ID  *string `json:"id,omitempty"`

    Name  *string `json:"name,omitempty"`

    RoleOwners  []WorkItem_work_item_RoleOwner `json:"role_owners,omitempty"`

    Status  *int32 `json:"status,omitempty"`

    ActualBeginTime  *string `json:"actual_begin_time,omitempty"`

    ActualFinishTime  *string `json:"actual_finish_time,omitempty"`

    Fields  []WorkItem_work_item_FieldValuePair `json:"fields,omitempty"`

}

type WorkItem_work_item_StateTime struct {

    StateKey  *string `json:"state_key,omitempty"`

    StartTime  *int64 `json:"start_time,omitempty"`

    EndTime  *int64 `json:"end_time,omitempty"`

    Name  *string `json:"name,omitempty"`

}

type WorkItem_work_item_SubTask struct {

    ID  *string `json:"id,omitempty"`

    Name  *string `json:"name,omitempty"`

    Schedules  []WorkItem_work_item_Schedule `json:"schedules,omitempty"`

    Order  *float64 `json:"order,omitempty"`

    Details  *string `json:"details,omitempty"`

    Passed  *bool `json:"passed,omitempty"`

    Owners  []string `json:"owners,omitempty"`

    Note  *string `json:"note,omitempty"`

    ActualBeginTime  *string `json:"actual_begin_time,omitempty"`

    ActualFinishTime  *string `json:"actual_finish_time,omitempty"`

    Assignee  []string `json:"assignee,omitempty"`

    RoleAssignee  []WorkItem_work_item_RoleOwner `json:"role_assignee,omitempty"`

    Deliverable  []WorkItem_work_item_FieldValuePair `json:"deliverable,omitempty"`

    OwnerRoles  []string `json:"owner_roles,omitempty"`

    OwnerUsageMode  *int64 `json:"owner_usage_mode,omitempty"`

    Fields  []WorkItem_work_item_FieldValuePair `json:"fields,omitempty"`

}

type WorkItem_work_item_SubTaskParentInfo struct {

    WorkItemID  *int64 `json:"work_item_id,omitempty"`

    WorkItemName  *string `json:"work_item_name,omitempty"`

    WorkItemTypeKey  *string `json:"work_item_type_key,omitempty"`

    NodeID  *string `json:"node_id,omitempty"`

}

type WorkItem_work_item_TargetState struct {

    StateKey  *string `json:"state_key,omitempty"`

    TransitionID  *int64 `json:"transition_id,omitempty"`

}

type WorkItem_work_item_UserDetail struct {

    UserKey  *string `json:"user_key,omitempty"`

    Username  *string `json:"username,omitempty"`

    Email  *string `json:"email,omitempty"`

    NameCn  *string `json:"name_cn,omitempty"`

    NameEn  *string `json:"name_en,omitempty"`

}

type WorkItem_work_item_WorkItemInfo struct {

    ID  *int64 `json:"id,omitempty"`

    Name  *string `json:"name,omitempty"`

    WorkItemTypeKey  *string `json:"work_item_type_key,omitempty"`

    ProjectKey  *string `json:"project_key,omitempty"`

    TemplateType  *string `json:"template_type,omitempty"`

    Pattern  *string `json:"pattern,omitempty"`

    SubStage  *string `json:"sub_stage,omitempty"`

    CurrentNodes  []WorkItem_work_item_NodeBasicInfo `json:"current_nodes,omitempty"`

    CreatedBy  *string `json:"created_by,omitempty"`

    UpdatedBy  *string `json:"updated_by,omitempty"`

    CreatedAt  *int64 `json:"created_at,omitempty"`

    UpdatedAt  *int64 `json:"updated_at,omitempty"`

    Fields  []WorkItem_work_item_FieldValuePair `json:"fields,omitempty"`

    WorkItemStatus  *WorkItem_work_item_WorkItemStatus `json:"work_item_status,omitempty"`

    DeletedBy  *string `json:"deleted_by,omitempty"`

    DeletedAt  *int64 `json:"deleted_at,omitempty"`

    SimpleName  *string `json:"simple_name,omitempty"`

    TemplateID  *int64 `json:"template_id,omitempty"`

    WorkflowInfos  *WorkItem_work_item_NodesConnections `json:"workflow_infos,omitempty"`

    StateTimes  []WorkItem_work_item_StateTime `json:"state_times,omitempty"`

    MultiTexts  []WorkItem_work_item_MultiText `json:"multi_texts,omitempty"`

    RelationFieldsDetail  []WorkItem_work_item_RelationFieldDetail `json:"relation_fields_detail,omitempty"`

    UserDetails  []WorkItem_work_item_UserDetail `json:"user_details,omitempty"`

    SubTaskParentInfo  *WorkItem_work_item_SubTaskParentInfo `json:"sub_task_parent_info,omitempty"`

    CompoundFieldExtra  []WorkItem_work_item_FieldValuePair `json:"compound_field_extra,omitempty"`

    ParentWorkItem  []ParentWorkItem `json:"parent_work_item,omitempty"`

}

type WorkItem_work_item_WorkItemStatus struct {

    StateKey  *string `json:"state_key,omitempty"`

    IsArchivedState  *bool `json:"is_archived_state,omitempty"`

    IsInitState  *bool `json:"is_init_state,omitempty"`

    UpdatedAt  *int64 `json:"updated_at,omitempty"`

    UpdatedBy  *string `json:"updated_by,omitempty"`

    History  []WorkItem_work_item_WorkItemStatus `json:"history,omitempty"`

}

type WorkItem_work_item_WorkflowNode struct {

    ID  *string `json:"id,omitempty"`

    StateKey  *string `json:"state_key,omitempty"`

    Name  *string `json:"name,omitempty"`

    Status  *int32 `json:"status,omitempty"`

    Fields  []WorkItem_work_item_FieldValuePair `json:"fields,omitempty"`

    Owners  []string `json:"owners,omitempty"`

    NodeSchedule  *WorkItem_work_item_Schedule `json:"node_schedule,omitempty"`

    Schedules  []WorkItem_work_item_Schedule `json:"schedules,omitempty"`

    SubTasks  []WorkItem_work_item_SubTask `json:"sub_tasks,omitempty"`

    ActualBeginTime  *string `json:"actual_begin_time,omitempty"`

    ActualFinishTime  *string `json:"actual_finish_time,omitempty"`

    RoleAssignee  []WorkItem_work_item_RoleOwner `json:"role_assignee,omitempty"`

    DifferentSchedule  *bool `json:"different_schedule,omitempty"`

    SubStatus  []WorkItem_work_item_Checker `json:"sub_status,omitempty"`

    Milestone  *bool `json:"milestone,omitempty"`

    Participants  []string `json:"participants,omitempty"`

    OwnerRoles  []string `json:"owner_roles,omitempty"`

    OwnerUsageMode  *int64 `json:"owner_usage_mode,omitempty"`

    NodeFields  []NodeField `json:"node_fields,omitempty"`

    FinishedInfos  *FinishedInfo `json:"finished_infos,omitempty"`

    Checker  []WorkItem_work_item_Checker `json:"checker,omitempty"`

}

type WorkflowConf struct {

    StatusInfos  []StatusConf `json:"status_infos,omitempty"`

    NodeInfos  []NodeConf `json:"node_infos,omitempty"`

    Connections  []WorkItem_work_item_Connection `json:"connections,omitempty"`

}

type WorkflowConfInfo struct {

    StateKey  *string `json:"state_key,omitempty"`

    Name  *string `json:"name,omitempty"`

    Tags  []string `json:"tags,omitempty"`

    OwnerUsageMode  *int64 `json:"owner_usage_mode,omitempty"`

    OwnerRoles  []string `json:"owner_roles,omitempty"`

    Owners  []string `json:"owners,omitempty"`

    NeedSchedule  *bool `json:"need_schedule,omitempty"`

    DifferentSchedule  *bool `json:"different_schedule,omitempty"`

    VisibilityUsageMode  *int64 `json:"visibility_usage_mode,omitempty"`

    Deletable  *bool `json:"deletable,omitempty"`

    DeletableOperationRole  []string `json:"deletable_operation_role,omitempty"`

    PassMode  *int64 `json:"pass_mode,omitempty"`

    IsLimitNode  *bool `json:"is_limit_node,omitempty"`

    DoneOperationRole  []string `json:"done_operation_role,omitempty"`

    DoneSchedule  *bool `json:"done_schedule,omitempty"`

    DoneAllocateOwner  *bool `json:"done_allocate_owner,omitempty"`

    Action  *int64 `json:"action,omitempty"`

    PreNodeStateKey  []string `json:"pre_node_state_key,omitempty"`

    CompletionTips  *string `json:"completion_tips,omitempty"`

    TaskConfs  []TaskConf `json:"task_confs,omitempty"`

    BelongStatus  *string `json:"belong_status,omitempty"`

    DoneActualPoint  *bool `json:"done_actual_point,omitempty"`

    IsMilestone  *bool `json:"is_milestone,omitempty"`

    DoneFinishTime  *bool `json:"done_finish_time,omitempty"`

    Fields  []FieldConf `json:"fields,omitempty"`

    SubWorkItems  []SubWorkItemConfInfo `json:"sub_work_items,omitempty"`

    SubTasks  []TaskConfInfo `json:"sub_tasks,omitempty"`

}

type WorkflowNode struct {

    ID  *string `json:"id,omitempty"`

    StateKey  *string `json:"state_key,omitempty"`

    Name  *string `json:"name,omitempty"`

    Status  *int32 `json:"status,omitempty"`

    Fields  []FieldValuePair `json:"fields,omitempty"`

    Owners  []string `json:"owners,omitempty"`

    NodeSchedule  *Schedule `json:"node_schedule,omitempty"`

    Schedules  []Schedule `json:"schedules,omitempty"`

    SubTasks  []SubTask `json:"sub_tasks,omitempty"`

    ActualBeginTime  *string `json:"actual_begin_time,omitempty"`

    ActualFinishTime  *string `json:"actual_finish_time,omitempty"`

    RoleAssignee  []RoleOwner `json:"role_assignee,omitempty"`

    DifferentSchedule  *bool `json:"different_schedule,omitempty"`

    SubStatus  []Checker `json:"sub_status,omitempty"`

    Milestone  *bool `json:"milestone,omitempty"`

    Participants  []string `json:"participants,omitempty"`

}
