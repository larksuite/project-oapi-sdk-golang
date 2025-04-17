package project


type Project struct {

    ProjectKey  *string `json:"project_key,omitempty"`

    Name  *string `json:"name,omitempty"`

    SimpleName  *string `json:"simple_name,omitempty"`

    Administrators  []string `json:"administrators,omitempty"`

}
