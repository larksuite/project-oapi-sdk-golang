package project

import (
	"github.com/larksuite/project-oapi-sdk-golang/core"
)

type BatchQueryProjectInfoReq struct {
	apiReq *core.APIReq
}
type BatchQueryProjectInfoReqBody struct {
	UserKey       *string  `json:"user_key,omitempty"`
	SimpleNames   []string `json:"simple_names,omitempty"`
	TenantGroupID *int64   `json:"tenant_group_id,omitempty"`
}
type BatchQueryProjectInfoResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data map[string]Project `json:"data"`
}

type BatchQueryProjectInfoReqBuilder struct {
	apiReq *core.APIReq
}

func NewBatchQueryProjectInfoReqBuilder() *BatchQueryProjectInfoReqBuilder {
	builder := &BatchQueryProjectInfoReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &BatchQueryProjectInfoReqBody{},
	}
	return builder
}

func (builder *BatchQueryProjectInfoReqBuilder) UserKey(userKey string) *BatchQueryProjectInfoReqBuilder {
	builder.apiReq.Body.(*BatchQueryProjectInfoReqBody).UserKey = &userKey
	return builder
}

func (builder *BatchQueryProjectInfoReqBuilder) SimpleNames(simpleNames []string) *BatchQueryProjectInfoReqBuilder {
	builder.apiReq.Body.(*BatchQueryProjectInfoReqBody).SimpleNames = simpleNames
	return builder
}

func (builder *BatchQueryProjectInfoReqBuilder) TenantGroupID(tenantGroupID int64) *BatchQueryProjectInfoReqBuilder {
	builder.apiReq.Body.(*BatchQueryProjectInfoReqBody).TenantGroupID = &tenantGroupID
	return builder
}

func (builder *BatchQueryProjectInfoReqBuilder) Build() *BatchQueryProjectInfoReq {
	req := &BatchQueryProjectInfoReq{}
	req.apiReq = builder.apiReq
	return req
}

type QueryProjectsReq struct {
	apiReq *core.APIReq
}
type QueryProjectsReqBody struct {
	UserKey       *string  `json:"user_key,omitempty"`
	TenantGroupID *int64   `json:"tenant_group_id,omitempty"`
	AssetKey      *string  `json:"asset_key,omitempty"`
	Order         []string `json:"order,omitempty"`
}
type QueryProjectsResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data []string `json:"data"`
}

type QueryProjectsReqBuilder struct {
	apiReq *core.APIReq
}

func NewQueryProjectsReqBuilder() *QueryProjectsReqBuilder {
	builder := &QueryProjectsReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &QueryProjectsReqBody{},
	}
	return builder
}

func (builder *QueryProjectsReqBuilder) UserKey(userKey string) *QueryProjectsReqBuilder {
	builder.apiReq.Body.(*QueryProjectsReqBody).UserKey = &userKey
	return builder
}

func (builder *QueryProjectsReqBuilder) TenantGroupID(tenantGroupID int64) *QueryProjectsReqBuilder {
	builder.apiReq.Body.(*QueryProjectsReqBody).TenantGroupID = &tenantGroupID
	return builder
}

func (builder *QueryProjectsReqBuilder) AssetKey(assetKey string) *QueryProjectsReqBuilder {
	builder.apiReq.Body.(*QueryProjectsReqBody).AssetKey = &assetKey
	return builder
}

func (builder *QueryProjectsReqBuilder) Order(order []string) *QueryProjectsReqBuilder {
	builder.apiReq.Body.(*QueryProjectsReqBody).Order = order
	return builder
}
func (builder *QueryProjectsReqBuilder) Build() *QueryProjectsReq {
	req := &QueryProjectsReq{}
	req.apiReq = builder.apiReq
	return req
}
