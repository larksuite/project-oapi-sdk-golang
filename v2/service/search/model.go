package search


type SearchGroup struct {

    SearchParams  []SearchParam `json:"search_params,omitempty"`

    Conjunction  *string `json:"conjunction,omitempty"`

    SearchGroups  []SearchGroup `json:"search_groups,omitempty"`

}

type SearchParam struct {

    ParamKey  *string `json:"param_key,omitempty"`

    Value  interface{} `json:"value,omitempty"`

    Operator  *string `json:"operator,omitempty"`

    PreOperator  *string `json:"pre_operator,omitempty"`

    ValueSearchGroups  *SearchGroup `json:"value_search_groups,omitempty"`

}

type Team struct {

    ID  *string `json:"id,omitempty"`

    NotCascade  *bool `json:"not_cascade,omitempty"`

    NotContainTeamManager  *bool `json:"not_contain_team_manager,omitempty"`

}
