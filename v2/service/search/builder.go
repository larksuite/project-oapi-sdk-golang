package search

import (
	"github.com/larksuite/project-oapi-sdk-golang/core"
)

type OAPICreateConditionViewReq struct {
	apiReq *core.APIReq
}
type OAPICreateConditionViewReqBody struct {
	ProjectKey *string `json:"project_key,omitempty"`

	WorkItemTypeKey *string `json:"work_item_type_key,omitempty"`

	SearchGroup *SearchGroup `json:"search_group,omitempty"`

	CooperationMode *int64 `json:"cooperation_mode,omitempty"`

	CooperationUserKeys []string `json:"cooperation_user_keys,omitempty"`

	CooperationTeamIDs []int64 `json:"cooperation_team_ids,omitempty"`

	Name *string `json:"name,omitempty"`

	CooperationTeams []Team `json:"cooperation_teams,omitempty"`
}

type OAPICreateConditionViewResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	ViewID *string `json:"view_id"`
}

type OAPICreateConditionViewReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPICreateConditionViewReqBuilder() *OAPICreateConditionViewReqBuilder {
	builder := &OAPICreateConditionViewReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &OAPICreateConditionViewReqBody{},
	}
	return builder
}
func (builder *OAPICreateConditionViewReqBuilder) ProjectKey(projectKey *string) *OAPICreateConditionViewReqBuilder {
	builder.apiReq.Body.(*OAPICreateConditionViewReqBody).ProjectKey = projectKey
	return builder
}
func (builder *OAPICreateConditionViewReqBuilder) WorkItemTypeKey(workItemTypeKey *string) *OAPICreateConditionViewReqBuilder {
	builder.apiReq.Body.(*OAPICreateConditionViewReqBody).WorkItemTypeKey = workItemTypeKey
	return builder
}
func (builder *OAPICreateConditionViewReqBuilder) SearchGroup(searchGroup *SearchGroup) *OAPICreateConditionViewReqBuilder {
	builder.apiReq.Body.(*OAPICreateConditionViewReqBody).SearchGroup = searchGroup
	return builder
}
func (builder *OAPICreateConditionViewReqBuilder) CooperationMode(cooperationMode *int64) *OAPICreateConditionViewReqBuilder {
	builder.apiReq.Body.(*OAPICreateConditionViewReqBody).CooperationMode = cooperationMode
	return builder
}
func (builder *OAPICreateConditionViewReqBuilder) CooperationUserKeys(cooperationUserKeys []string) *OAPICreateConditionViewReqBuilder {
	builder.apiReq.Body.(*OAPICreateConditionViewReqBody).CooperationUserKeys = cooperationUserKeys
	return builder
}
func (builder *OAPICreateConditionViewReqBuilder) CooperationTeamIDs(cooperationTeamIDs []int64) *OAPICreateConditionViewReqBuilder {
	builder.apiReq.Body.(*OAPICreateConditionViewReqBody).CooperationTeamIDs = cooperationTeamIDs
	return builder
}
func (builder *OAPICreateConditionViewReqBuilder) Name(name *string) *OAPICreateConditionViewReqBuilder {
	builder.apiReq.Body.(*OAPICreateConditionViewReqBody).Name = name
	return builder
}
func (builder *OAPICreateConditionViewReqBuilder) CooperationTeams(cooperationTeams []Team) *OAPICreateConditionViewReqBuilder {
	builder.apiReq.Body.(*OAPICreateConditionViewReqBody).CooperationTeams = cooperationTeams
	return builder
}
func (builder *OAPICreateConditionViewReqBuilder) Build() *OAPICreateConditionViewReq {
	req := &OAPICreateConditionViewReq{}
	req.apiReq = builder.apiReq
	return req
}

type OAPIUpdateConditionViewReq struct {
	apiReq *core.APIReq
}
type OAPIUpdateConditionViewReqBody struct {
	ProjectKey *string `json:"project_key,omitempty"`

	WorkItemTypeKey *string `json:"work_item_type_key,omitempty"`

	SearchGroup *SearchGroup `json:"search_group,omitempty"`

	CooperationMode *int64 `json:"cooperation_mode,omitempty"`

	CooperationUserKeys []string `json:"cooperation_user_keys,omitempty"`

	CooperationTeamIDs []int64 `json:"cooperation_team_ids,omitempty"`

	Name *string `json:"name,omitempty"`

	ViewID *string `json:"view_id,omitempty"`

	CooperationTeams []Team `json:"cooperation_teams,omitempty"`
}

type OAPIUpdateConditionViewResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
}

type OAPIUpdateConditionViewReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPIUpdateConditionViewReqBuilder() *OAPIUpdateConditionViewReqBuilder {
	builder := &OAPIUpdateConditionViewReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &OAPIUpdateConditionViewReqBody{},
	}
	return builder
}
func (builder *OAPIUpdateConditionViewReqBuilder) ProjectKey(projectKey *string) *OAPIUpdateConditionViewReqBuilder {
	builder.apiReq.Body.(*OAPIUpdateConditionViewReqBody).ProjectKey = projectKey
	return builder
}
func (builder *OAPIUpdateConditionViewReqBuilder) WorkItemTypeKey(workItemTypeKey *string) *OAPIUpdateConditionViewReqBuilder {
	builder.apiReq.Body.(*OAPIUpdateConditionViewReqBody).WorkItemTypeKey = workItemTypeKey
	return builder
}
func (builder *OAPIUpdateConditionViewReqBuilder) SearchGroup(searchGroup *SearchGroup) *OAPIUpdateConditionViewReqBuilder {
	builder.apiReq.Body.(*OAPIUpdateConditionViewReqBody).SearchGroup = searchGroup
	return builder
}
func (builder *OAPIUpdateConditionViewReqBuilder) CooperationMode(cooperationMode *int64) *OAPIUpdateConditionViewReqBuilder {
	builder.apiReq.Body.(*OAPIUpdateConditionViewReqBody).CooperationMode = cooperationMode
	return builder
}
func (builder *OAPIUpdateConditionViewReqBuilder) CooperationUserKeys(cooperationUserKeys []string) *OAPIUpdateConditionViewReqBuilder {
	builder.apiReq.Body.(*OAPIUpdateConditionViewReqBody).CooperationUserKeys = cooperationUserKeys
	return builder
}
func (builder *OAPIUpdateConditionViewReqBuilder) CooperationTeamIDs(cooperationTeamIDs []int64) *OAPIUpdateConditionViewReqBuilder {
	builder.apiReq.Body.(*OAPIUpdateConditionViewReqBody).CooperationTeamIDs = cooperationTeamIDs
	return builder
}
func (builder *OAPIUpdateConditionViewReqBuilder) Name(name *string) *OAPIUpdateConditionViewReqBuilder {
	builder.apiReq.Body.(*OAPIUpdateConditionViewReqBody).Name = name
	return builder
}
func (builder *OAPIUpdateConditionViewReqBuilder) ViewID(viewID *string) *OAPIUpdateConditionViewReqBuilder {
	builder.apiReq.Body.(*OAPIUpdateConditionViewReqBody).ViewID = viewID
	return builder
}
func (builder *OAPIUpdateConditionViewReqBuilder) CooperationTeams(cooperationTeams []Team) *OAPIUpdateConditionViewReqBuilder {
	builder.apiReq.Body.(*OAPIUpdateConditionViewReqBody).CooperationTeams = cooperationTeams
	return builder
}
func (builder *OAPIUpdateConditionViewReqBuilder) Build() *OAPIUpdateConditionViewReq {
	req := &OAPIUpdateConditionViewReq{}
	req.apiReq = builder.apiReq
	return req
}
