package measure


type ChartData struct {

    Dim  map[int32]string `json:"dim,omitempty"`

    Value  map[int32]string `json:"value,omitempty"`

    IsZeroSpec  *bool `json:"is_zero_spec,omitempty"`

}

type OAPIChart struct {

    ChartName  *string `json:"chart_name,omitempty"`

    ChartID  *string `json:"chart_id,omitempty"`

}

type OAPIChartData struct {

    Name  *string `json:"name,omitempty"`

    ChartID  *string `json:"chart_id,omitempty"`

    ChartDataList  [][]ChartData `json:"chart_data_list,omitempty"`

    DimTitles  []string `json:"dim_titles,omitempty"`

    QuotaTitles  []string `json:"quota_titles,omitempty"`

}

type OAPIChartPage struct {

    Total  *int32 `json:"total,omitempty"`

    PageNum  *int32 `json:"page_num,omitempty"`

    PageSize  *int32 `json:"page_size,omitempty"`

}
