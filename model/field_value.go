package model

type Schedule struct {
	StartTime int64 `json:"start_time"`

	EndTime int64 `json:"end_time"`
}

type Radio struct {
	Label int64 `json:"label"`

	Value int64 `json:"value"`
}

type Select struct {
	Label int64 `json:"label"`

	Value int64 `json:"value"`
}

type MUltiSelect []*Select

type TreeSelect struct {
	Label int64 `json:"label"`

	Value int64 `json:"value"`

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
	UUID string `json:"uuid"`
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
