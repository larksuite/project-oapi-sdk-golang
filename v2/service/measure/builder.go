package measure

import (
    "fmt"
   "github.com/larksuite/project-oapi-sdk-golang/core"
    
)


type OAPIGetChartDataReq struct {
	apiReq *core.APIReq
}

type OAPIGetChartDataResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       *OAPIChartData         `json:"data"`
	
}

type OAPIGetChartDataReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPIGetChartDataReqBuilder() *OAPIGetChartDataReqBuilder {
	builder := &OAPIGetChartDataReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}

func (builder *OAPIGetChartDataReqBuilder) ProjectKey(projectKey string) *OAPIGetChartDataReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}


func (builder *OAPIGetChartDataReqBuilder) ChartID(chartID string) *OAPIGetChartDataReqBuilder {
	builder.apiReq.PathParams.Set("chart_id", fmt.Sprint(chartID))
	return builder
}

func (builder *OAPIGetChartDataReqBuilder) Build() *OAPIGetChartDataReq {
	req := &OAPIGetChartDataReq{}
	req.apiReq = builder.apiReq
	return req
}

