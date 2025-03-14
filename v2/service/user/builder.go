package user

import (
    "fmt"
   "github.com/larksuite/project-oapi-sdk-golang/core"
    
)


type OAPIActiveAccountReq struct {
	apiReq *core.APIReq
}
type OAPIActiveAccountReqBody struct {

    UserMeegoKey  *string `json:"user_meego_key,omitempty"`

}

type OAPIActiveAccountResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
}

type OAPIActiveAccountReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPIActiveAccountReqBuilder() *OAPIActiveAccountReqBuilder {
	builder := &OAPIActiveAccountReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &OAPIActiveAccountReqBody{},
	}
	return builder
}
func (builder *OAPIActiveAccountReqBuilder) UserMeegoKey(userMeegoKey *string) *OAPIActiveAccountReqBuilder {
	builder.apiReq.Body.(*OAPIActiveAccountReqBody).UserMeegoKey = userMeegoKey
	return builder
}
func (builder *OAPIActiveAccountReqBuilder) Build() *OAPIActiveAccountReq {
	req := &OAPIActiveAccountReq{}
	req.apiReq = builder.apiReq
	return req
}

type OAPICreateAccountReq struct {
	apiReq *core.APIReq
}
type OAPICreateAccountReqBody struct {

    OutUserID  *string `json:"out_user_id,omitempty"`

    Name  map[string]string `json:"name,omitempty"`

    LoginPlatformType  *string `json:"login_platform_type,omitempty"`

    AvatarUrl  *string `json:"avatar_url,omitempty"`

    DepartmentMeegoKeys  []string `json:"department_meego_keys,omitempty"`

}

type OAPICreateAccountResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	UserMeegoKey       *string         `json:"user_meego_key"`
	
}

type OAPICreateAccountReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPICreateAccountReqBuilder() *OAPICreateAccountReqBuilder {
	builder := &OAPICreateAccountReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &OAPICreateAccountReqBody{},
	}
	return builder
}
func (builder *OAPICreateAccountReqBuilder) OutUserID(outUserID *string) *OAPICreateAccountReqBuilder {
	builder.apiReq.Body.(*OAPICreateAccountReqBody).OutUserID = outUserID
	return builder
}
func (builder *OAPICreateAccountReqBuilder) Name(name map[string]string) *OAPICreateAccountReqBuilder {
	builder.apiReq.Body.(*OAPICreateAccountReqBody).Name = name
	return builder
}
func (builder *OAPICreateAccountReqBuilder) LoginPlatformType(loginPlatformType *string) *OAPICreateAccountReqBuilder {
	builder.apiReq.Body.(*OAPICreateAccountReqBody).LoginPlatformType = loginPlatformType
	return builder
}
func (builder *OAPICreateAccountReqBuilder) AvatarUrl(avatarUrl *string) *OAPICreateAccountReqBuilder {
	builder.apiReq.Body.(*OAPICreateAccountReqBody).AvatarUrl = avatarUrl
	return builder
}
func (builder *OAPICreateAccountReqBuilder) DepartmentMeegoKeys(departmentMeegoKeys []string) *OAPICreateAccountReqBuilder {
	builder.apiReq.Body.(*OAPICreateAccountReqBody).DepartmentMeegoKeys = departmentMeegoKeys
	return builder
}
func (builder *OAPICreateAccountReqBuilder) Build() *OAPICreateAccountReq {
	req := &OAPICreateAccountReq{}
	req.apiReq = builder.apiReq
	return req
}

type OAPICreateDepartmentReq struct {
	apiReq *core.APIReq
}
type OAPICreateDepartmentReqBody struct {

    Name  map[string]string `json:"name,omitempty"`

    ParentDepartmentMeegoKey  *string `json:"parent_department_meego_key,omitempty"`

}

type OAPICreateDepartmentResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	DepartmentMeegoKey       *string         `json:"department_meego_key"`
	
}

type OAPICreateDepartmentReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPICreateDepartmentReqBuilder() *OAPICreateDepartmentReqBuilder {
	builder := &OAPICreateDepartmentReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &OAPICreateDepartmentReqBody{},
	}
	return builder
}
func (builder *OAPICreateDepartmentReqBuilder) Name(name map[string]string) *OAPICreateDepartmentReqBuilder {
	builder.apiReq.Body.(*OAPICreateDepartmentReqBody).Name = name
	return builder
}
func (builder *OAPICreateDepartmentReqBuilder) ParentDepartmentMeegoKey(parentDepartmentMeegoKey *string) *OAPICreateDepartmentReqBuilder {
	builder.apiReq.Body.(*OAPICreateDepartmentReqBody).ParentDepartmentMeegoKey = parentDepartmentMeegoKey
	return builder
}
func (builder *OAPICreateDepartmentReqBuilder) Build() *OAPICreateDepartmentReq {
	req := &OAPICreateDepartmentReq{}
	req.apiReq = builder.apiReq
	return req
}

type OAPICreateUserGroupReq struct {
	apiReq *core.APIReq
}
type OAPICreateUserGroupReqBody struct {

    Name  *string `json:"name,omitempty"`

    Users  []string `json:"users,omitempty"`

}

type OAPICreateUserGroupResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	ID       *string         `json:"id"`
	
}

type OAPICreateUserGroupReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPICreateUserGroupReqBuilder() *OAPICreateUserGroupReqBuilder {
	builder := &OAPICreateUserGroupReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &OAPICreateUserGroupReqBody{},
	}
	return builder
}
func (builder *OAPICreateUserGroupReqBuilder) ProjectKey(projectKey *string) *OAPICreateUserGroupReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(*projectKey))
	return builder
}
func (builder *OAPICreateUserGroupReqBuilder) Name(name *string) *OAPICreateUserGroupReqBuilder {
	builder.apiReq.Body.(*OAPICreateUserGroupReqBody).Name = name
	return builder
}
func (builder *OAPICreateUserGroupReqBuilder) Users(users []string) *OAPICreateUserGroupReqBuilder {
	builder.apiReq.Body.(*OAPICreateUserGroupReqBody).Users = users
	return builder
}
func (builder *OAPICreateUserGroupReqBuilder) Build() *OAPICreateUserGroupReq {
	req := &OAPICreateUserGroupReq{}
	req.apiReq = builder.apiReq
	return req
}

type OAPIDeleteDepartmentReq struct {
	apiReq *core.APIReq
}
type OAPIDeleteDepartmentReqBody struct {

    DepartmentMeegoKey  *string `json:"department_meego_key,omitempty"`

}

type OAPIDeleteDepartmentResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
}

type OAPIDeleteDepartmentReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPIDeleteDepartmentReqBuilder() *OAPIDeleteDepartmentReqBuilder {
	builder := &OAPIDeleteDepartmentReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &OAPIDeleteDepartmentReqBody{},
	}
	return builder
}
func (builder *OAPIDeleteDepartmentReqBuilder) DepartmentMeegoKey(departmentMeegoKey *string) *OAPIDeleteDepartmentReqBuilder {
	builder.apiReq.Body.(*OAPIDeleteDepartmentReqBody).DepartmentMeegoKey = departmentMeegoKey
	return builder
}
func (builder *OAPIDeleteDepartmentReqBuilder) Build() *OAPIDeleteDepartmentReq {
	req := &OAPIDeleteDepartmentReq{}
	req.apiReq = builder.apiReq
	return req
}

type OAPIPageUserGroupMembersReq struct {
	apiReq *core.APIReq
}
type OAPIPageUserGroupMembersReqBody struct {

    UserGroupType  *string `json:"user_group_type,omitempty"`

    UserGroupIDs  []string `json:"user_group_ids,omitempty"`

    PageNum  *int64 `json:"page_num,omitempty"`

    PageSize  *int64 `json:"page_size,omitempty"`

}

type OAPIPageUserGroupMembersResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	List       []UserGroupDetail         `json:"list"`
	
	Pagination       *Pagination         `json:"pagination"`
	
}

type OAPIPageUserGroupMembersReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPIPageUserGroupMembersReqBuilder() *OAPIPageUserGroupMembersReqBuilder {
	builder := &OAPIPageUserGroupMembersReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &OAPIPageUserGroupMembersReqBody{},
	}
	return builder
}
func (builder *OAPIPageUserGroupMembersReqBuilder) ProjectKey(projectKey *string) *OAPIPageUserGroupMembersReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(*projectKey))
	return builder
}
func (builder *OAPIPageUserGroupMembersReqBuilder) UserGroupType(userGroupType *string) *OAPIPageUserGroupMembersReqBuilder {
	builder.apiReq.Body.(*OAPIPageUserGroupMembersReqBody).UserGroupType = userGroupType
	return builder
}
func (builder *OAPIPageUserGroupMembersReqBuilder) UserGroupIDs(userGroupIDs []string) *OAPIPageUserGroupMembersReqBuilder {
	builder.apiReq.Body.(*OAPIPageUserGroupMembersReqBody).UserGroupIDs = userGroupIDs
	return builder
}
func (builder *OAPIPageUserGroupMembersReqBuilder) PageNum(pageNum *int64) *OAPIPageUserGroupMembersReqBuilder {
	builder.apiReq.Body.(*OAPIPageUserGroupMembersReqBody).PageNum = pageNum
	return builder
}
func (builder *OAPIPageUserGroupMembersReqBuilder) PageSize(pageSize *int64) *OAPIPageUserGroupMembersReqBuilder {
	builder.apiReq.Body.(*OAPIPageUserGroupMembersReqBody).PageSize = pageSize
	return builder
}
func (builder *OAPIPageUserGroupMembersReqBuilder) Build() *OAPIPageUserGroupMembersReq {
	req := &OAPIPageUserGroupMembersReq{}
	req.apiReq = builder.apiReq
	return req
}

type OAPIPatchUserGroupMembersReq struct {
	apiReq *core.APIReq
}
type OAPIPatchUserGroupMembersReqBody struct {

    UserGroupType  *string `json:"user_group_type,omitempty"`

    UserGroupID  *string `json:"user_group_id,omitempty"`

    AddUsers  []string `json:"add_users,omitempty"`

    DeleteUsers  []string `json:"delete_users,omitempty"`

    ReplaceUsers  []string `json:"replace_users,omitempty"`

}

type OAPIPatchUserGroupMembersResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
}

type OAPIPatchUserGroupMembersReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPIPatchUserGroupMembersReqBuilder() *OAPIPatchUserGroupMembersReqBuilder {
	builder := &OAPIPatchUserGroupMembersReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &OAPIPatchUserGroupMembersReqBody{},
	}
	return builder
}
func (builder *OAPIPatchUserGroupMembersReqBuilder) ProjectKey(projectKey *string) *OAPIPatchUserGroupMembersReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(*projectKey))
	return builder
}
func (builder *OAPIPatchUserGroupMembersReqBuilder) UserGroupType(userGroupType *string) *OAPIPatchUserGroupMembersReqBuilder {
	builder.apiReq.Body.(*OAPIPatchUserGroupMembersReqBody).UserGroupType = userGroupType
	return builder
}
func (builder *OAPIPatchUserGroupMembersReqBuilder) UserGroupID(userGroupID *string) *OAPIPatchUserGroupMembersReqBuilder {
	builder.apiReq.Body.(*OAPIPatchUserGroupMembersReqBody).UserGroupID = userGroupID
	return builder
}
func (builder *OAPIPatchUserGroupMembersReqBuilder) AddUsers(addUsers []string) *OAPIPatchUserGroupMembersReqBuilder {
	builder.apiReq.Body.(*OAPIPatchUserGroupMembersReqBody).AddUsers = addUsers
	return builder
}
func (builder *OAPIPatchUserGroupMembersReqBuilder) DeleteUsers(deleteUsers []string) *OAPIPatchUserGroupMembersReqBuilder {
	builder.apiReq.Body.(*OAPIPatchUserGroupMembersReqBody).DeleteUsers = deleteUsers
	return builder
}
func (builder *OAPIPatchUserGroupMembersReqBuilder) ReplaceUsers(replaceUsers []string) *OAPIPatchUserGroupMembersReqBuilder {
	builder.apiReq.Body.(*OAPIPatchUserGroupMembersReqBody).ReplaceUsers = replaceUsers
	return builder
}
func (builder *OAPIPatchUserGroupMembersReqBuilder) Build() *OAPIPatchUserGroupMembersReq {
	req := &OAPIPatchUserGroupMembersReq{}
	req.apiReq = builder.apiReq
	return req
}

type OAPIQueryDepartmentReq struct {
	apiReq *core.APIReq
}
type OAPIQueryDepartmentReqBody struct {

    Scopes  []string `json:"scopes,omitempty"`

    DepartmentMeegoKey  *string `json:"department_meego_key,omitempty"`

}

type OAPIQueryDepartmentResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	DepartmentMeegoKey       *string         `json:"department_meego_key"`
	
	Name       map[string]string         `json:"name"`
	
	ExternalIDs       map[string]string         `json:"external_ids"`
	
	OrgLevel       *string         `json:"org_level"`
	
	Parent       *DepartmentBasicInfo         `json:"parent"`
	
	AllChildren       []DepartmentBasicInfo         `json:"all_children"`
	
	DirectEmployees       []AccountInfo         `json:"direct_employees"`
	
	AllEmployees       []AccountInfo         `json:"all_employees"`
	
}

type OAPIQueryDepartmentReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPIQueryDepartmentReqBuilder() *OAPIQueryDepartmentReqBuilder {
	builder := &OAPIQueryDepartmentReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &OAPIQueryDepartmentReqBody{},
	}
	return builder
}
func (builder *OAPIQueryDepartmentReqBuilder) Scopes(scopes []string) *OAPIQueryDepartmentReqBuilder {
	builder.apiReq.Body.(*OAPIQueryDepartmentReqBody).Scopes = scopes
	return builder
}
func (builder *OAPIQueryDepartmentReqBuilder) DepartmentMeegoKey(departmentMeegoKey *string) *OAPIQueryDepartmentReqBuilder {
	builder.apiReq.Body.(*OAPIQueryDepartmentReqBody).DepartmentMeegoKey = departmentMeegoKey
	return builder
}
func (builder *OAPIQueryDepartmentReqBuilder) Build() *OAPIQueryDepartmentReq {
	req := &OAPIQueryDepartmentReq{}
	req.apiReq = builder.apiReq
	return req
}

type OAPIQueryTenantReq struct {
	apiReq *core.APIReq
}
type OAPIQueryTenantReqBody struct {

    Scopes  []string `json:"scopes,omitempty"`

}

type OAPIQueryTenantResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	TenantMeegoKey       *string         `json:"tenant_meego_key"`
	
	Name       *string         `json:"name"`
	
	ExternalIDs       map[string]string         `json:"external_ids"`
	
	Icon       *string         `json:"icon"`
	
	Status       *string         `json:"status"`
	
	PlatformType       *string         `json:"platform_type"`
	
	Departments       []DepartmentInfo         `json:"departments"`
	
	Accounts       []AccountInfo         `json:"accounts"`
	
}

type OAPIQueryTenantReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPIQueryTenantReqBuilder() *OAPIQueryTenantReqBuilder {
	builder := &OAPIQueryTenantReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &OAPIQueryTenantReqBody{},
	}
	return builder
}
func (builder *OAPIQueryTenantReqBuilder) Scopes(scopes []string) *OAPIQueryTenantReqBuilder {
	builder.apiReq.Body.(*OAPIQueryTenantReqBody).Scopes = scopes
	return builder
}
func (builder *OAPIQueryTenantReqBuilder) Build() *OAPIQueryTenantReq {
	req := &OAPIQueryTenantReq{}
	req.apiReq = builder.apiReq
	return req
}

type OAPIQueryUsersForTenantPluginReq struct {
	apiReq *core.APIReq
}
type OAPIQueryUsersForTenantPluginReqBody struct {

    UserKeys  []string `json:"user_keys,omitempty"`

    Emails  []string `json:"emails,omitempty"`

    OutUsers  []OutUserInfo `json:"out_users,omitempty"`

}

type OAPIQueryUsersForTenantPluginResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       []UserBasicInfo         `json:"data"`
	
}

type OAPIQueryUsersForTenantPluginReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPIQueryUsersForTenantPluginReqBuilder() *OAPIQueryUsersForTenantPluginReqBuilder {
	builder := &OAPIQueryUsersForTenantPluginReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &OAPIQueryUsersForTenantPluginReqBody{},
	}
	return builder
}
func (builder *OAPIQueryUsersForTenantPluginReqBuilder) UserKeys(userKeys []string) *OAPIQueryUsersForTenantPluginReqBuilder {
	builder.apiReq.Body.(*OAPIQueryUsersForTenantPluginReqBody).UserKeys = userKeys
	return builder
}
func (builder *OAPIQueryUsersForTenantPluginReqBuilder) Emails(emails []string) *OAPIQueryUsersForTenantPluginReqBuilder {
	builder.apiReq.Body.(*OAPIQueryUsersForTenantPluginReqBody).Emails = emails
	return builder
}
func (builder *OAPIQueryUsersForTenantPluginReqBuilder) OutUsers(outUsers []OutUserInfo) *OAPIQueryUsersForTenantPluginReqBuilder {
	builder.apiReq.Body.(*OAPIQueryUsersForTenantPluginReqBody).OutUsers = outUsers
	return builder
}
func (builder *OAPIQueryUsersForTenantPluginReqBuilder) Build() *OAPIQueryUsersForTenantPluginReq {
	req := &OAPIQueryUsersForTenantPluginReq{}
	req.apiReq = builder.apiReq
	return req
}

type OAPISearchUserByWordReq struct {
	apiReq *core.APIReq
}
type OAPISearchUserByWordReqBody struct {

    Query  *string `json:"query,omitempty"`

    ProjectKey  *string `json:"project_key,omitempty"`

}

type OAPISearchUserByWordResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       []UserBasicInfo         `json:"data"`
	
}

type OAPISearchUserByWordReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPISearchUserByWordReqBuilder() *OAPISearchUserByWordReqBuilder {
	builder := &OAPISearchUserByWordReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &OAPISearchUserByWordReqBody{},
	}
	return builder
}
func (builder *OAPISearchUserByWordReqBuilder) Query(query *string) *OAPISearchUserByWordReqBuilder {
	builder.apiReq.Body.(*OAPISearchUserByWordReqBody).Query = query
	return builder
}
func (builder *OAPISearchUserByWordReqBuilder) ProjectKey(projectKey *string) *OAPISearchUserByWordReqBuilder {
	builder.apiReq.Body.(*OAPISearchUserByWordReqBody).ProjectKey = projectKey
	return builder
}
func (builder *OAPISearchUserByWordReqBuilder) Build() *OAPISearchUserByWordReq {
	req := &OAPISearchUserByWordReq{}
	req.apiReq = builder.apiReq
	return req
}

type OAPIStopAccountReq struct {
	apiReq *core.APIReq
}
type OAPIStopAccountReqBody struct {

    UserMeegoKey  *string `json:"user_meego_key,omitempty"`

}

type OAPIStopAccountResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
}

type OAPIStopAccountReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPIStopAccountReqBuilder() *OAPIStopAccountReqBuilder {
	builder := &OAPIStopAccountReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &OAPIStopAccountReqBody{},
	}
	return builder
}
func (builder *OAPIStopAccountReqBuilder) UserMeegoKey(userMeegoKey *string) *OAPIStopAccountReqBuilder {
	builder.apiReq.Body.(*OAPIStopAccountReqBody).UserMeegoKey = userMeegoKey
	return builder
}
func (builder *OAPIStopAccountReqBuilder) Build() *OAPIStopAccountReq {
	req := &OAPIStopAccountReq{}
	req.apiReq = builder.apiReq
	return req
}

type OAPIUpdateAccountReq struct {
	apiReq *core.APIReq
}
type OAPIUpdateAccountReqBody struct {

    UserMeegoKey  *string `json:"user_meego_key,omitempty"`

    Name  map[string]string `json:"name,omitempty"`

    AvatarUrl  *string `json:"avatar_url,omitempty"`

    DepartmentMeegoKeys  []string `json:"department_meego_keys,omitempty"`

}

type OAPIUpdateAccountResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
}

type OAPIUpdateAccountReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPIUpdateAccountReqBuilder() *OAPIUpdateAccountReqBuilder {
	builder := &OAPIUpdateAccountReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &OAPIUpdateAccountReqBody{},
	}
	return builder
}
func (builder *OAPIUpdateAccountReqBuilder) UserMeegoKey(userMeegoKey *string) *OAPIUpdateAccountReqBuilder {
	builder.apiReq.Body.(*OAPIUpdateAccountReqBody).UserMeegoKey = userMeegoKey
	return builder
}
func (builder *OAPIUpdateAccountReqBuilder) Name(name map[string]string) *OAPIUpdateAccountReqBuilder {
	builder.apiReq.Body.(*OAPIUpdateAccountReqBody).Name = name
	return builder
}
func (builder *OAPIUpdateAccountReqBuilder) AvatarUrl(avatarUrl *string) *OAPIUpdateAccountReqBuilder {
	builder.apiReq.Body.(*OAPIUpdateAccountReqBody).AvatarUrl = avatarUrl
	return builder
}
func (builder *OAPIUpdateAccountReqBuilder) DepartmentMeegoKeys(departmentMeegoKeys []string) *OAPIUpdateAccountReqBuilder {
	builder.apiReq.Body.(*OAPIUpdateAccountReqBody).DepartmentMeegoKeys = departmentMeegoKeys
	return builder
}
func (builder *OAPIUpdateAccountReqBuilder) Build() *OAPIUpdateAccountReq {
	req := &OAPIUpdateAccountReq{}
	req.apiReq = builder.apiReq
	return req
}

type OAPIUpdateDepartmentReq struct {
	apiReq *core.APIReq
}
type OAPIUpdateDepartmentReqBody struct {

    DepartmentMeegoKey  *string `json:"department_meego_key,omitempty"`

    Name  map[string]string `json:"name,omitempty"`

    ParentDepartmentMeegoKey  *string `json:"parent_department_meego_key,omitempty"`

}

type OAPIUpdateDepartmentResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
}

type OAPIUpdateDepartmentReqBuilder struct {
	apiReq *core.APIReq
}

func NewOAPIUpdateDepartmentReqBuilder() *OAPIUpdateDepartmentReqBuilder {
	builder := &OAPIUpdateDepartmentReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &OAPIUpdateDepartmentReqBody{},
	}
	return builder
}
func (builder *OAPIUpdateDepartmentReqBuilder) DepartmentMeegoKey(departmentMeegoKey *string) *OAPIUpdateDepartmentReqBuilder {
	builder.apiReq.Body.(*OAPIUpdateDepartmentReqBody).DepartmentMeegoKey = departmentMeegoKey
	return builder
}
func (builder *OAPIUpdateDepartmentReqBuilder) Name(name map[string]string) *OAPIUpdateDepartmentReqBuilder {
	builder.apiReq.Body.(*OAPIUpdateDepartmentReqBody).Name = name
	return builder
}
func (builder *OAPIUpdateDepartmentReqBuilder) ParentDepartmentMeegoKey(parentDepartmentMeegoKey *string) *OAPIUpdateDepartmentReqBuilder {
	builder.apiReq.Body.(*OAPIUpdateDepartmentReqBody).ParentDepartmentMeegoKey = parentDepartmentMeegoKey
	return builder
}
func (builder *OAPIUpdateDepartmentReqBuilder) Build() *OAPIUpdateDepartmentReq {
	req := &OAPIUpdateDepartmentReq{}
	req.apiReq = builder.apiReq
	return req
}

type QueryUserDetailReq struct {
	apiReq *core.APIReq
}
type QueryUserDetailReqBody struct {

    UserKeys  []string `json:"user_keys,omitempty"`

    OutIDs  []string `json:"out_ids,omitempty"`

    Emails  []string `json:"emails,omitempty"`

    TenantKey  *string `json:"tenant_key,omitempty"`

}

type QueryUserDetailResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       []UserBasicInfo         `json:"data"`
	
}

type QueryUserDetailReqBuilder struct {
	apiReq *core.APIReq
}

func NewQueryUserDetailReqBuilder() *QueryUserDetailReqBuilder {
	builder := &QueryUserDetailReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &QueryUserDetailReqBody{},
	}
	return builder
}
func (builder *QueryUserDetailReqBuilder) UserKeys(userKeys []string) *QueryUserDetailReqBuilder {
	builder.apiReq.Body.(*QueryUserDetailReqBody).UserKeys = userKeys
	return builder
}
func (builder *QueryUserDetailReqBuilder) OutIDs(outIDs []string) *QueryUserDetailReqBuilder {
	builder.apiReq.Body.(*QueryUserDetailReqBody).OutIDs = outIDs
	return builder
}
func (builder *QueryUserDetailReqBuilder) Emails(emails []string) *QueryUserDetailReqBuilder {
	builder.apiReq.Body.(*QueryUserDetailReqBody).Emails = emails
	return builder
}
func (builder *QueryUserDetailReqBuilder) TenantKey(tenantKey *string) *QueryUserDetailReqBuilder {
	builder.apiReq.Body.(*QueryUserDetailReqBody).TenantKey = tenantKey
	return builder
}
func (builder *QueryUserDetailReqBuilder) Build() *QueryUserDetailReq {
	req := &QueryUserDetailReq{}
	req.apiReq = builder.apiReq
	return req
}

