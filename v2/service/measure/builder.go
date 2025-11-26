package measure

import (
    "fmt"
   "github.com/larksuite/project-oapi-sdk-golang/core"
    
)


type GetChartDataReq struct {
	apiReq *core.APIReq
}
type GetChartDataResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       *OAPIChartData         `json:"data"`
	
}

type GetChartDataReqBuilder struct {
	apiReq *core.APIReq
}

func NewGetChartDataReqBuilder() *GetChartDataReqBuilder {
	builder := &GetChartDataReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}

func (builder *GetChartDataReqBuilder) ProjectKey(projectKey string) *GetChartDataReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}


func (builder *GetChartDataReqBuilder) ChartID(chartID string) *GetChartDataReqBuilder {
	builder.apiReq.PathParams.Set("chart_id", fmt.Sprint(chartID))
	return builder
}

func (builder *GetChartDataReqBuilder) Build() *GetChartDataReq {
	req := &GetChartDataReq{}
	req.apiReq = builder.apiReq
	return req
}

type GetChartsReq struct {
	apiReq *core.APIReq
}
type GetChartsReqBody struct {
    ProjectKey  *string `json:"project_key,omitempty"`
    ViewID  *string `json:"view_id,omitempty"`
    PageNum  *int32 `json:"page_num,omitempty"`
    PageSize  *int32 `json:"page_size,omitempty"`
}
type GetChartsResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data      *GetChartsRespData        `json:"data,omitempty"`
}

type GetChartsRespData struct {
	ChartList       []OAPIChart         `json:"chart_list,omitempty"`
	ChartPage       *OAPIChartPage         `json:"chart_page,omitempty"`
}

type GetChartsReqBuilder struct {
	apiReq *core.APIReq
}

func NewGetChartsReqBuilder() *GetChartsReqBuilder {
	builder := &GetChartsReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &GetChartsReqBody{},
	}
	return builder
}

func (builder *GetChartsReqBuilder) ProjectKey(projectKey string) *GetChartsReqBuilder {
	builder.apiReq.Body.(*GetChartsReqBody).ProjectKey = &projectKey
	return builder
}


func (builder *GetChartsReqBuilder) ViewID(viewID string) *GetChartsReqBuilder {
	builder.apiReq.Body.(*GetChartsReqBody).ViewID = &viewID
	return builder
}


func (builder *GetChartsReqBuilder) PageNum(pageNum int32) *GetChartsReqBuilder {
	builder.apiReq.Body.(*GetChartsReqBody).PageNum = &pageNum
	return builder
}


func (builder *GetChartsReqBuilder) PageSize(pageSize int32) *GetChartsReqBuilder {
	builder.apiReq.Body.(*GetChartsReqBody).PageSize = &pageSize
	return builder
}

func (builder *GetChartsReqBuilder) Build() *GetChartsReq {
	req := &GetChartsReq{}
	req.apiReq = builder.apiReq
	return req
}

