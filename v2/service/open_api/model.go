package open_api


type PluginInfo struct {

    PluginKey  *string `json:"plugin_Key,omitempty"`

    PluginVersion  *string `json:"plugin_version,omitempty"`

    PointType  *string `json:"point_type,omitempty"`

    PointKey  *string `json:"point_key,omitempty"`

}

type UserInfo struct {

    UserKey  *string `json:"user_key,omitempty"`

    Email  *string `json:"email,omitempty"`

    NameCn  *string `json:"name_cn,omitempty"`

    NameEn  *string `json:"name_en,omitempty"`

}

type WorkItemData struct {

    WorkItemID  *int64 `json:"work_item_id,omitempty"`

    Result  *bool `json:"result,omitempty"`

    ErrMsg  *string `json:"err_msg,omitempty"`

    WorkItemData  interface{} `json:"work_item_data,omitempty"`

    Editable  *bool `json:"editable,omitempty"`

}

type WorkItemInfo struct {

    WorkItemID  *int64 `json:"work_item_id,omitempty"`

    WorkItemName  *string `json:"work_item_name,omitempty"`

    WorkItemTypeKey  *string `json:"work_item_type_key,omitempty"`

}
