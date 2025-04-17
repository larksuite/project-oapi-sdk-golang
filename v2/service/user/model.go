package user


type AccountInfo struct {

    UserMeegoKey  *string `json:"user_meego_key,omitempty"`

    TenantMeegoKey  *string `json:"tenant_meego_key,omitempty"`

    Name  map[string]string `json:"name,omitempty"`

    ExternalIDs  map[string]string `json:"external_ids,omitempty"`

    LoginPlatformType  *string `json:"login_platform_type,omitempty"`

    Status  *string `json:"status,omitempty"`

    AvatarUrl  *string `json:"avatar_url,omitempty"`

}

type DepartmentBasicInfo struct {

    DepartmentMeegoKey  *string `json:"department_meego_key,omitempty"`

    Name  map[string]string `json:"name,omitempty"`

    ExternalIDs  map[string]string `json:"external_ids,omitempty"`

    MemberCount  *int64 `json:"memberCount,omitempty"`

    OrgLevel  *string `json:"org_level,omitempty"`

}

type DepartmentInfo struct {

    DepartmentMeegoKey  *string `json:"department_meego_key,omitempty"`

    Name  map[string]string `json:"name,omitempty"`

    ExternalIDs  map[string]string `json:"external_ids,omitempty"`

    OrgLevel  *string `json:"org_level,omitempty"`

    Parent  *DepartmentBasicInfo `json:"parent,omitempty"`

    AllChildren  []DepartmentBasicInfo `json:"all_children,omitempty"`

    DirectEmployees  []AccountInfo `json:"direct_employees,omitempty"`

    AllEmployees  []AccountInfo `json:"all_employees,omitempty"`

}

type OutUserInfo struct {

    OutUserID  *string `json:"out_user_id,omitempty"`

    LoginPlatformType  *string `json:"login_platform_type,omitempty"`

}

type Pagination struct {

    PageNum  *int64 `json:"page_num,omitempty"`

    PageSize  *int64 `json:"page_size,omitempty"`

    HasMore  *bool `json:"has_more,omitempty"`

}

type Team struct {

    TeamID  *int64 `json:"team_id,omitempty"`

    TeamName  *string `json:"team_name,omitempty"`

    UserKeys  []string `json:"user_keys,omitempty"`

    Administrators  []string `json:"administrators,omitempty"`

}

type UserBasicInfo struct {

    UserID  *int64 `json:"user_id,omitempty"`

    UserKey  *string `json:"user_key,omitempty"`

    Username  *string `json:"username,omitempty"`

    Email  *string `json:"email,omitempty"`

    AvatarUrl  *string `json:"avatar_url,omitempty"`

    NameCn  *string `json:"name_cn,omitempty"`

    NameEn  *string `json:"name_en,omitempty"`

    OutID  *string `json:"out_id,omitempty"`

    Status  *string `json:"status,omitempty"`

    OutUser  *OutUserInfo `json:"out_user,omitempty"`

    Departments  []DepartmentBasicInfo `json:"departments,omitempty"`

    Name  map[string]string `json:"name,omitempty"`

}

type UserGroupDetail struct {

    ID  *string `json:"id,omitempty"`

    Name  *string `json:"name,omitempty"`

    UserCount  *int64 `json:"user_count,omitempty"`

    UserMembers  []string `json:"user_members,omitempty"`

}
