package view

import (
    "fmt"
   "github.com/larksuite/project-oapi-sdk-golang/core"
    
)


type CreateFixViewReq struct {
	apiReq *core.APIReq
}
type CreateFixViewReqBody struct {

    WorkItemIDList  []int64 `json:"work_item_id_list,omitempty"`

    Name  *string `json:"name,omitempty"`

    CooperationMode  *int64 `json:"cooperation_mode,omitempty"`

    CooperationUserKeys  []string `json:"cooperation_user_keys,omitempty"`

    CooperationTeamIDs  []int64 `json:"cooperation_team_ids,omitempty"`

    CooperationTeams  []Team `json:"cooperation_teams,omitempty"`

}

type CreateFixViewResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       *FixView         `json:"data"`
	
	Pagination       *Pagination         `json:"pagination"`
	
}

type CreateFixViewReqBuilder struct {
	apiReq *core.APIReq
}

func NewCreateFixViewReqBuilder() *CreateFixViewReqBuilder {
	builder := &CreateFixViewReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &CreateFixViewReqBody{},
	}
	return builder
}

func (builder *CreateFixViewReqBuilder) ProjectKey(projectKey string) *CreateFixViewReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}


func (builder *CreateFixViewReqBuilder) WorkItemIDList(workItemIDList []int64) *CreateFixViewReqBuilder {
	builder.apiReq.Body.(*CreateFixViewReqBody).WorkItemIDList = workItemIDList
	return builder
}

func (builder *CreateFixViewReqBuilder) Name(name string) *CreateFixViewReqBuilder {
	builder.apiReq.Body.(*CreateFixViewReqBody).Name = &name
	return builder
}


func (builder *CreateFixViewReqBuilder) WorkItemTypeKey(workItemTypeKey string) *CreateFixViewReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(workItemTypeKey))
	return builder
}


func (builder *CreateFixViewReqBuilder) CooperationMode(cooperationMode *int64) *CreateFixViewReqBuilder {
	builder.apiReq.Body.(*CreateFixViewReqBody).CooperationMode = cooperationMode
	return builder
}

func (builder *CreateFixViewReqBuilder) CooperationUserKeys(cooperationUserKeys []string) *CreateFixViewReqBuilder {
	builder.apiReq.Body.(*CreateFixViewReqBody).CooperationUserKeys = cooperationUserKeys
	return builder
}

func (builder *CreateFixViewReqBuilder) CooperationTeamIDs(cooperationTeamIDs []int64) *CreateFixViewReqBuilder {
	builder.apiReq.Body.(*CreateFixViewReqBody).CooperationTeamIDs = cooperationTeamIDs
	return builder
}

func (builder *CreateFixViewReqBuilder) CooperationTeams(cooperationTeams []Team) *CreateFixViewReqBuilder {
	builder.apiReq.Body.(*CreateFixViewReqBody).CooperationTeams = cooperationTeams
	return builder
}
func (builder *CreateFixViewReqBuilder) Build() *CreateFixViewReq {
	req := &CreateFixViewReq{}
	req.apiReq = builder.apiReq
	return req
}

type DeleteFixViewReq struct {
	apiReq *core.APIReq
}

type DeleteFixViewResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
}

type DeleteFixViewReqBuilder struct {
	apiReq *core.APIReq
}

func NewDeleteFixViewReqBuilder() *DeleteFixViewReqBuilder {
	builder := &DeleteFixViewReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}

func (builder *DeleteFixViewReqBuilder) ProjectKey(projectKey string) *DeleteFixViewReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}


func (builder *DeleteFixViewReqBuilder) ViewID(viewID string) *DeleteFixViewReqBuilder {
	builder.apiReq.PathParams.Set("view_id", fmt.Sprint(viewID))
	return builder
}

func (builder *DeleteFixViewReqBuilder) Build() *DeleteFixViewReq {
	req := &DeleteFixViewReq{}
	req.apiReq = builder.apiReq
	return req
}

type QueryWorkItemDetailsByViewIDReq struct {
	apiReq *core.APIReq
}
type QueryWorkItemDetailsByViewIDReqBody struct {

    PageSize  *int64 `json:"page_size,omitempty"`

    PageNum  *int64 `json:"page_num,omitempty"`

    Expand  *Expand `json:"expand,omitempty"`

}

type QueryWorkItemDetailsByViewIDResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       []WorkItemInfo         `json:"data"`
	
	Pagination       *Pagination         `json:"pagination"`
	
}

type QueryWorkItemDetailsByViewIDReqBuilder struct {
	apiReq *core.APIReq
}

func NewQueryWorkItemDetailsByViewIDReqBuilder() *QueryWorkItemDetailsByViewIDReqBuilder {
	builder := &QueryWorkItemDetailsByViewIDReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &QueryWorkItemDetailsByViewIDReqBody{},
	}
	return builder
}

func (builder *QueryWorkItemDetailsByViewIDReqBuilder) ProjectKey(projectKey string) *QueryWorkItemDetailsByViewIDReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}


func (builder *QueryWorkItemDetailsByViewIDReqBuilder) ViewID(viewID string) *QueryWorkItemDetailsByViewIDReqBuilder {
	builder.apiReq.PathParams.Set("view_id", fmt.Sprint(viewID))
	return builder
}


func (builder *QueryWorkItemDetailsByViewIDReqBuilder) PageSize(pageSize *int64) *QueryWorkItemDetailsByViewIDReqBuilder {
	builder.apiReq.Body.(*QueryWorkItemDetailsByViewIDReqBody).PageSize = pageSize
	return builder
}

func (builder *QueryWorkItemDetailsByViewIDReqBuilder) PageNum(pageNum *int64) *QueryWorkItemDetailsByViewIDReqBuilder {
	builder.apiReq.Body.(*QueryWorkItemDetailsByViewIDReqBody).PageNum = pageNum
	return builder
}

func (builder *QueryWorkItemDetailsByViewIDReqBuilder) Expand(expand *Expand) *QueryWorkItemDetailsByViewIDReqBuilder {
	builder.apiReq.Body.(*QueryWorkItemDetailsByViewIDReqBody).Expand = expand
	return builder
}
func (builder *QueryWorkItemDetailsByViewIDReqBuilder) Build() *QueryWorkItemDetailsByViewIDReq {
	req := &QueryWorkItemDetailsByViewIDReq{}
	req.apiReq = builder.apiReq
	return req
}

type UpdateFixViewReq struct {
	apiReq *core.APIReq
}
type UpdateFixViewReqBody struct {

    AddWorkItemIDs  []int64 `json:"add_work_item_ids,omitempty"`

    RemoveWorkItemIDs  []int64 `json:"remove_work_item_ids,omitempty"`

    CooperationMode  *int64 `json:"cooperation_mode,omitempty"`

    CooperationUserKeys  []string `json:"cooperation_user_keys,omitempty"`

    CooperationTeamIDs  []int64 `json:"cooperation_team_ids,omitempty"`

    CooperationTeams  []Team `json:"cooperation_teams,omitempty"`

}

type UpdateFixViewResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       *FixView         `json:"data"`
	
	Pagination       *Pagination         `json:"pagination"`
	
}

type UpdateFixViewReqBuilder struct {
	apiReq *core.APIReq
}

func NewUpdateFixViewReqBuilder() *UpdateFixViewReqBuilder {
	builder := &UpdateFixViewReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &UpdateFixViewReqBody{},
	}
	return builder
}

func (builder *UpdateFixViewReqBuilder) ProjectKey(projectKey string) *UpdateFixViewReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}


func (builder *UpdateFixViewReqBuilder) ViewID(viewID string) *UpdateFixViewReqBuilder {
	builder.apiReq.PathParams.Set("view_id", fmt.Sprint(viewID))
	return builder
}


func (builder *UpdateFixViewReqBuilder) AddWorkItemIDs(addWorkItemIDs []int64) *UpdateFixViewReqBuilder {
	builder.apiReq.Body.(*UpdateFixViewReqBody).AddWorkItemIDs = addWorkItemIDs
	return builder
}

func (builder *UpdateFixViewReqBuilder) RemoveWorkItemIDs(removeWorkItemIDs []int64) *UpdateFixViewReqBuilder {
	builder.apiReq.Body.(*UpdateFixViewReqBody).RemoveWorkItemIDs = removeWorkItemIDs
	return builder
}

func (builder *UpdateFixViewReqBuilder) WorkItemTypeKey(workItemTypeKey string) *UpdateFixViewReqBuilder {
	builder.apiReq.PathParams.Set("work_item_type_key", fmt.Sprint(workItemTypeKey))
	return builder
}


func (builder *UpdateFixViewReqBuilder) CooperationMode(cooperationMode *int64) *UpdateFixViewReqBuilder {
	builder.apiReq.Body.(*UpdateFixViewReqBody).CooperationMode = cooperationMode
	return builder
}

func (builder *UpdateFixViewReqBuilder) CooperationUserKeys(cooperationUserKeys []string) *UpdateFixViewReqBuilder {
	builder.apiReq.Body.(*UpdateFixViewReqBody).CooperationUserKeys = cooperationUserKeys
	return builder
}

func (builder *UpdateFixViewReqBuilder) CooperationTeamIDs(cooperationTeamIDs []int64) *UpdateFixViewReqBuilder {
	builder.apiReq.Body.(*UpdateFixViewReqBody).CooperationTeamIDs = cooperationTeamIDs
	return builder
}

func (builder *UpdateFixViewReqBuilder) CooperationTeams(cooperationTeams []Team) *UpdateFixViewReqBuilder {
	builder.apiReq.Body.(*UpdateFixViewReqBody).CooperationTeams = cooperationTeams
	return builder
}
func (builder *UpdateFixViewReqBuilder) Build() *UpdateFixViewReq {
	req := &UpdateFixViewReq{}
	req.apiReq = builder.apiReq
	return req
}

type ViewListReq struct {
	apiReq *core.APIReq
}
type ViewListReqBody struct {

    WorkItemTypeKey  *string `json:"work_item_type_key,omitempty"`

    ViewIDs  []string `json:"view_ids,omitempty"`

    CreatedBy  *string `json:"created_by,omitempty"`

    CreatedAt  *TimeInterval `json:"created_at,omitempty"`

    PageSize  *int64 `json:"page_size,omitempty"`

    PageNum  *int64 `json:"page_num,omitempty"`

    ViewName  *string `json:"view_name,omitempty"`

}

type ViewListResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       []ViewConf         `json:"data"`
	
	Pagination       *Pagination         `json:"pagination"`
	
}

type ViewListReqBuilder struct {
	apiReq *core.APIReq
}

func NewViewListReqBuilder() *ViewListReqBuilder {
	builder := &ViewListReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &ViewListReqBody{},
	}
	return builder
}

func (builder *ViewListReqBuilder) ProjectKey(projectKey string) *ViewListReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}


func (builder *ViewListReqBuilder) WorkItemTypeKey(workItemTypeKey string) *ViewListReqBuilder {
	builder.apiReq.Body.(*ViewListReqBody).WorkItemTypeKey = &workItemTypeKey
	return builder
}


func (builder *ViewListReqBuilder) ViewIDs(viewIDs []string) *ViewListReqBuilder {
	builder.apiReq.Body.(*ViewListReqBody).ViewIDs = viewIDs
	return builder
}

func (builder *ViewListReqBuilder) CreatedBy(createdBy string) *ViewListReqBuilder {
	builder.apiReq.Body.(*ViewListReqBody).CreatedBy = &createdBy
	return builder
}


func (builder *ViewListReqBuilder) CreatedAt(createdAt *TimeInterval) *ViewListReqBuilder {
	builder.apiReq.Body.(*ViewListReqBody).CreatedAt = createdAt
	return builder
}

func (builder *ViewListReqBuilder) PageSize(pageSize *int64) *ViewListReqBuilder {
	builder.apiReq.Body.(*ViewListReqBody).PageSize = pageSize
	return builder
}

func (builder *ViewListReqBuilder) PageNum(pageNum *int64) *ViewListReqBuilder {
	builder.apiReq.Body.(*ViewListReqBody).PageNum = pageNum
	return builder
}

func (builder *ViewListReqBuilder) ViewName(viewName string) *ViewListReqBuilder {
	builder.apiReq.Body.(*ViewListReqBody).ViewName = &viewName
	return builder
}

func (builder *ViewListReqBuilder) Build() *ViewListReq {
	req := &ViewListReq{}
	req.apiReq = builder.apiReq
	return req
}

type WorkItemListReq struct {
	apiReq *core.APIReq
}

type WorkItemListResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       *FixView         `json:"data"`
	
	Pagination       *Pagination         `json:"pagination"`
	
}

type WorkItemListReqBuilder struct {
	apiReq *core.APIReq
}

func NewWorkItemListReqBuilder() *WorkItemListReqBuilder {
	builder := &WorkItemListReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}

func (builder *WorkItemListReqBuilder) ProjectKey(projectKey string) *WorkItemListReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}


func (builder *WorkItemListReqBuilder) ViewID(viewID string) *WorkItemListReqBuilder {
	builder.apiReq.PathParams.Set("view_id", fmt.Sprint(viewID))
	return builder
}


func (builder *WorkItemListReqBuilder) PageSize(pageSize *int64) *WorkItemListReqBuilder {
	builder.apiReq.QueryParams.Set("page_size", fmt.Sprint(*pageSize))
	return builder
}

func (builder *WorkItemListReqBuilder) PageNum(pageNum *int64) *WorkItemListReqBuilder {
	builder.apiReq.QueryParams.Set("page_num", fmt.Sprint(*pageNum))
	return builder
}
func (builder *WorkItemListReqBuilder) Build() *WorkItemListReq {
	req := &WorkItemListReq{}
	req.apiReq = builder.apiReq
	return req
}

