package measure

import (
	"context"
	"fmt"

	"github.com/larksuite/project-oapi-sdk-golang/core"
)


const APIPath_GetChartData = "/open_api/:project_key/measure/:chart_id"


func NewService(config *core.Config) *MeasureService {
	a := &MeasureService{config: config}
	return a
}

type MeasureService struct {
	config *core.Config
}


/*
 * @name: 拉取图表信息
 * @desc: 可以拉取普通图表或者魔法公式图表
 */
func (a *MeasureService) GetChartData(ctx context.Context, req *GetChartDataReq, options ...core.RequestOptionFunc) (*GetChartDataResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_GetChartData
	apiReq.HttpMethod = "GET"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[GetChartData] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &GetChartDataResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[GetChartData] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

