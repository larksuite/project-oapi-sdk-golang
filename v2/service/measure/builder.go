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

