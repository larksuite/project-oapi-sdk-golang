package measure


type ChartData struct {

    Dim  map[int32]string `json:"dim,omitempty"`

    Value  map[int32]string `json:"value,omitempty"`

}

type OAPIChartData struct {

    Name  *string `json:"name,omitempty"`

    ChartID  *string `json:"chart_id,omitempty"`

    ChartDataList  [][]ChartData `json:"chart_data_list,omitempty"`

    DimTitles  []string `json:"dim_titles,omitempty"`

    QuotaTitles  []string `json:"quota_titles,omitempty"`

}
