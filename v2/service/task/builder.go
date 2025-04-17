package task

import (
	"github.com/larksuite/project-oapi-sdk-golang/core"
)

type SearchSubtaskReq struct {
	apiReq *core.APIReq
}
type SearchSubtaskReqBody struct {
	PageSize *int64 `json:"page_size,omitempty"`

	PageNum *int64 `json:"page_num,omitempty"`

	Name *string `json:"name,omitempty"`

	UserKeys []string `json:"user_keys,omitempty"`

	Status *int32 `json:"status,omitempty"`

	CreatedAt *TimeInterval `json:"created_at,omitempty"`

	SimpleNames []string `json:"simple_names,omitempty"`
}

type SearchSubtaskResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Pagination *Pagination `json:"pagination"`

	Data []SubDetail `json:"data"`
}

type SearchSubtaskReqBuilder struct {
	apiReq *core.APIReq
}

func NewSearchSubtaskReqBuilder() *SearchSubtaskReqBuilder {
	builder := &SearchSubtaskReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &SearchSubtaskReqBody{},
	}
	return builder
}

func (builder *SearchSubtaskReqBuilder) PageSize(pageSize int64) *SearchSubtaskReqBuilder {
	builder.apiReq.Body.(*SearchSubtaskReqBody).PageSize = &pageSize
	return builder
}

func (builder *SearchSubtaskReqBuilder) PageNum(pageNum int64) *SearchSubtaskReqBuilder {
	builder.apiReq.Body.(*SearchSubtaskReqBody).PageNum = &pageNum
	return builder
}

func (builder *SearchSubtaskReqBuilder) Name(name string) *SearchSubtaskReqBuilder {
	builder.apiReq.Body.(*SearchSubtaskReqBody).Name = &name
	return builder
}

func (builder *SearchSubtaskReqBuilder) UserKeys(userKeys []string) *SearchSubtaskReqBuilder {
	builder.apiReq.Body.(*SearchSubtaskReqBody).UserKeys = userKeys
	return builder
}

func (builder *SearchSubtaskReqBuilder) Status(status int32) *SearchSubtaskReqBuilder {
	builder.apiReq.Body.(*SearchSubtaskReqBody).Status = &status
	return builder
}

func (builder *SearchSubtaskReqBuilder) CreatedAt(createdAt *TimeInterval) *SearchSubtaskReqBuilder {
	builder.apiReq.Body.(*SearchSubtaskReqBody).CreatedAt = createdAt
	return builder
}

func (builder *SearchSubtaskReqBuilder) SimpleNames(simpleNames []string) *SearchSubtaskReqBuilder {
	builder.apiReq.Body.(*SearchSubtaskReqBody).SimpleNames = simpleNames
	return builder
}
func (builder *SearchSubtaskReqBuilder) Build() *SearchSubtaskReq {
	req := &SearchSubtaskReq{}
	req.apiReq = builder.apiReq
	return req
}
