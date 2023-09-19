package model

type Schedule struct {
	StartTime int64 `json:"start_time"`

	EndTime int64 `json:"end_time"`
}

type Radio struct {
	Label string `json:"label"`

	Value string `json:"value"`
}

type Select struct {
	Label string `json:"label"`

	Value string `json:"value"`
}

type MUltiSelect []*Select

type TreeSelect struct {
	Label string `json:"label"`

	Value string `json:"value"`

	Children *TreeSelect `json:"children"`
}

type MultiTreeSelect []*TreeSelect

type MultiFile struct {
	Name string `json:"name"`

	Type string `json:"type"`

	Size string `json:"size"`

	Uid string `json:"uid"`

	URL string `json:"url"`
}

type RoleOwner struct {
	Owners []string `json:"owners"`

	Role string `json:"role"`
}

type RoleOwners []*RoleOwner

type SubCompoundField struct {
	FieldKey string `json:"field_key"`

	FieldValue string `json:"field_value"`

	FieldTypeKey string `json:"field_type_key"`

	FieldAlias string `json:"field_alias"`
}

type CompoundField []*SubCompoundField

type Attrs struct {
	Src string `json:"src"`
}

type Marks struct {
	Type  string `json:"type"`
	Attrs string `json:"attrs"`
}

type MultiDoc struct {
	Type    string      `json:"type"`
	Attrs   *Attrs      `json:"attrs"`
	Text    string      `json:"text"`
	Marks   *Marks      `json:"marks"`
	Content []*MultiDoc `json:"content"`
}

type MultiDocs []*MultiDoc

type Text string

type Link string

type Date int64

type Number float64

type WorkItemRelatedSelect int64

type WorkItemRelatedMultiSelect int64

type Signal bool

type Bool bool

type User string

type MultiUser []string
