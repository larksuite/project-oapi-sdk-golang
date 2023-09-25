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

type ChartData struct {
	Dim map[int32]string `json:"dim"`

	Value map[int32]string `json:"value"`
}

type MeasureData struct {
	Name string `json:"name"`

	ChartID string `json:"chart_id"`

	ChartDataList [][]*ChartData `json:"chart_data_list"`

	DimTitles []string `json:"dim_titles"`

	QuotaTitles []string `json:"quota_titles"`
}
