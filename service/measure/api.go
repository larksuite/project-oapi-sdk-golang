/*
 * Copyright (c) 2023 Lark Technologies Pte. Ltd.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package measure

import (
	"context"
	"fmt"

	"github.com/larksuite/project-oapi-sdk-golang/core"
)

const APIPathQueryChartData = "/open_api/:project_key/measure/:chart_id"

func NewService(config *core.Config) *MeasureService {
	a := &MeasureService{config: config}
	return a
}

type MeasureService struct {
	config *core.Config
}

// 度量查询接口
func (a *MeasureService) QueryChartData(ctx context.Context, req *QueryChartDataReq, options ...core.RequestOptionFunc) (*QueryChartDataResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = APIPathQueryChartData
	apiReq.HttpMethod = "GET"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[QueryChartData] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &QueryChartDataResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[QueryChartData] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}
