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
	"fmt"

	"github.com/larksuite/project-oapi-sdk-golang/core"
)

type QueryChartDataReq struct {
	apiReq *core.APIReq
}

type QueryChartDataResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data *MeasureData `json:"data"`
}

type QueryChartDataReqBuilder struct {
	apiReq *core.APIReq
}

func NewQueryChartDataReqBuilder() *QueryChartDataReqBuilder {
	builder := &QueryChartDataReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}
func (builder *QueryChartDataReqBuilder) ProjectKey(projectKey string) *QueryChartDataReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}
func (builder *QueryChartDataReqBuilder) ChartID(chartID string) *QueryChartDataReqBuilder {
	builder.apiReq.PathParams.Set("chart_id", fmt.Sprint(chartID))
	return builder
}
func (builder *QueryChartDataReqBuilder) Build() *QueryChartDataReq {
	req := &QueryChartDataReq{}
	req.apiReq = builder.apiReq
	return req
}
