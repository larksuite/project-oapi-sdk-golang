package user

import (
    "fmt"
   "github.com/larksuite/project-oapi-sdk-golang/core"
    
)


type ActiveAccountReq struct {
	apiReq *core.APIReq
}
type ActiveAccountReqBody struct {
    UserMeegoKey  *string `json:"user_meego_key,omitempty"`
}
type ActiveAccountResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
}

type ActiveAccountReqBuilder struct {
	apiReq *core.APIReq
}

func NewActiveAccountReqBuilder() *ActiveAccountReqBuilder {
	builder := &ActiveAccountReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &ActiveAccountReqBody{},
	}
	return builder
}

func (builder *ActiveAccountReqBuilder) UserMeegoKey(userMeegoKey string) *ActiveAccountReqBuilder {
	builder.apiReq.Body.(*ActiveAccountReqBody).UserMeegoKey = &userMeegoKey
	return builder
}

func (builder *ActiveAccountReqBuilder) Build() *ActiveAccountReq {
	req := &ActiveAccountReq{}
	req.apiReq = builder.apiReq
	return req
}

type CreateAccountReq struct {
	apiReq *core.APIReq
}
type CreateAccountReqBody struct {
    OutUserID  *string `json:"out_user_id,omitempty"`
    Name  map[string]string `json:"name,omitempty"`
    LoginPlatformType  *string `json:"login_platform_type,omitempty"`
    AvatarUrl  *string `json:"avatar_url,omitempty"`
    DepartmentMeegoKeys  []string `json:"department_meego_keys,omitempty"`
    Email  *string `json:"email,omitempty"`
    Mobile  *string `json:"mobile,omitempty"`
}
type CreateAccountResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data      *CreateAccountRespData        `json:"data,omitempty"`
}

type CreateAccountRespData struct {
	UserMeegoKey       *string         `json:"user_meego_key,omitempty"`
}

type CreateAccountReqBuilder struct {
	apiReq *core.APIReq
}

func NewCreateAccountReqBuilder() *CreateAccountReqBuilder {
	builder := &CreateAccountReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &CreateAccountReqBody{},
	}
	return builder
}

func (builder *CreateAccountReqBuilder) OutUserID(outUserID string) *CreateAccountReqBuilder {
	builder.apiReq.Body.(*CreateAccountReqBody).OutUserID = &outUserID
	return builder
}


func (builder *CreateAccountReqBuilder) Name(name map[string]string) *CreateAccountReqBuilder {
	builder.apiReq.Body.(*CreateAccountReqBody).Name = name
	return builder
}

func (builder *CreateAccountReqBuilder) LoginPlatformType(loginPlatformType string) *CreateAccountReqBuilder {
	builder.apiReq.Body.(*CreateAccountReqBody).LoginPlatformType = &loginPlatformType
	return builder
}


func (builder *CreateAccountReqBuilder) AvatarUrl(avatarUrl string) *CreateAccountReqBuilder {
	builder.apiReq.Body.(*CreateAccountReqBody).AvatarUrl = &avatarUrl
	return builder
}


func (builder *CreateAccountReqBuilder) DepartmentMeegoKeys(departmentMeegoKeys []string) *CreateAccountReqBuilder {
	builder.apiReq.Body.(*CreateAccountReqBody).DepartmentMeegoKeys = departmentMeegoKeys
	return builder
}

func (builder *CreateAccountReqBuilder) Email(email string) *CreateAccountReqBuilder {
	builder.apiReq.Body.(*CreateAccountReqBody).Email = &email
	return builder
}


func (builder *CreateAccountReqBuilder) Mobile(mobile string) *CreateAccountReqBuilder {
	builder.apiReq.Body.(*CreateAccountReqBody).Mobile = &mobile
	return builder
}

func (builder *CreateAccountReqBuilder) Build() *CreateAccountReq {
	req := &CreateAccountReq{}
	req.apiReq = builder.apiReq
	return req
}

type CreateDepartmentReq struct {
	apiReq *core.APIReq
}
type CreateDepartmentReqBody struct {
    Name  map[string]string `json:"name,omitempty"`
    ParentDepartmentMeegoKey  *string `json:"parent_department_meego_key,omitempty"`
}
type CreateDepartmentResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data      *CreateDepartmentRespData        `json:"data,omitempty"`
}

type CreateDepartmentRespData struct {
	DepartmentMeegoKey       *string         `json:"department_meego_key,omitempty"`
}

type CreateDepartmentReqBuilder struct {
	apiReq *core.APIReq
}

func NewCreateDepartmentReqBuilder() *CreateDepartmentReqBuilder {
	builder := &CreateDepartmentReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &CreateDepartmentReqBody{},
	}
	return builder
}

func (builder *CreateDepartmentReqBuilder) Name(name map[string]string) *CreateDepartmentReqBuilder {
	builder.apiReq.Body.(*CreateDepartmentReqBody).Name = name
	return builder
}

func (builder *CreateDepartmentReqBuilder) ParentDepartmentMeegoKey(parentDepartmentMeegoKey string) *CreateDepartmentReqBuilder {
	builder.apiReq.Body.(*CreateDepartmentReqBody).ParentDepartmentMeegoKey = &parentDepartmentMeegoKey
	return builder
}

func (builder *CreateDepartmentReqBuilder) Build() *CreateDepartmentReq {
	req := &CreateDepartmentReq{}
	req.apiReq = builder.apiReq
	return req
}

type CreateUserGroupReq struct {
	apiReq *core.APIReq
}
type CreateUserGroupReqBody struct {
    Name  *string `json:"name,omitempty"`
    Users  []string `json:"users,omitempty"`
}
type CreateUserGroupResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data      *CreateUserGroupRespData        `json:"data,omitempty"`
}

type CreateUserGroupRespData struct {
	ID       *string         `json:"id,omitempty"`
}

type CreateUserGroupReqBuilder struct {
	apiReq *core.APIReq
}

func NewCreateUserGroupReqBuilder() *CreateUserGroupReqBuilder {
	builder := &CreateUserGroupReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &CreateUserGroupReqBody{},
	}
	return builder
}

func (builder *CreateUserGroupReqBuilder) ProjectKey(projectKey string) *CreateUserGroupReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}


func (builder *CreateUserGroupReqBuilder) Name(name string) *CreateUserGroupReqBuilder {
	builder.apiReq.Body.(*CreateUserGroupReqBody).Name = &name
	return builder
}


func (builder *CreateUserGroupReqBuilder) Users(users []string) *CreateUserGroupReqBuilder {
	builder.apiReq.Body.(*CreateUserGroupReqBody).Users = users
	return builder
}
func (builder *CreateUserGroupReqBuilder) Build() *CreateUserGroupReq {
	req := &CreateUserGroupReq{}
	req.apiReq = builder.apiReq
	return req
}

type DeleteDepartmentReq struct {
	apiReq *core.APIReq
}
type DeleteDepartmentReqBody struct {
    DepartmentMeegoKey  *string `json:"department_meego_key,omitempty"`
}
type DeleteDepartmentResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
}

type DeleteDepartmentReqBuilder struct {
	apiReq *core.APIReq
}

func NewDeleteDepartmentReqBuilder() *DeleteDepartmentReqBuilder {
	builder := &DeleteDepartmentReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &DeleteDepartmentReqBody{},
	}
	return builder
}

func (builder *DeleteDepartmentReqBuilder) DepartmentMeegoKey(departmentMeegoKey string) *DeleteDepartmentReqBuilder {
	builder.apiReq.Body.(*DeleteDepartmentReqBody).DepartmentMeegoKey = &departmentMeegoKey
	return builder
}

func (builder *DeleteDepartmentReqBuilder) Build() *DeleteDepartmentReq {
	req := &DeleteDepartmentReq{}
	req.apiReq = builder.apiReq
	return req
}

type ListProjectTeamReq struct {
	apiReq *core.APIReq
}
type ListProjectTeamResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       []Team         `json:"data"`
	
}

type ListProjectTeamReqBuilder struct {
	apiReq *core.APIReq
}

func NewListProjectTeamReqBuilder() *ListProjectTeamReqBuilder {
	builder := &ListProjectTeamReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}

func (builder *ListProjectTeamReqBuilder) ProjectKey(projectKey string) *ListProjectTeamReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}

func (builder *ListProjectTeamReqBuilder) Build() *ListProjectTeamReq {
	req := &ListProjectTeamReq{}
	req.apiReq = builder.apiReq
	return req
}

type PatchUserGroupMembersReq struct {
	apiReq *core.APIReq
}
type PatchUserGroupMembersReqBody struct {
    UserGroupType  *string `json:"user_group_type,omitempty"`
    UserGroupID  *string `json:"user_group_id,omitempty"`
    AddUsers  []string `json:"add_users,omitempty"`
    DeleteUsers  []string `json:"delete_users,omitempty"`
    ReplaceUsers  []string `json:"replace_users,omitempty"`
}
type PatchUserGroupMembersResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
}

type PatchUserGroupMembersReqBuilder struct {
	apiReq *core.APIReq
}

func NewPatchUserGroupMembersReqBuilder() *PatchUserGroupMembersReqBuilder {
	builder := &PatchUserGroupMembersReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &PatchUserGroupMembersReqBody{},
	}
	return builder
}

func (builder *PatchUserGroupMembersReqBuilder) ProjectKey(projectKey string) *PatchUserGroupMembersReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}


func (builder *PatchUserGroupMembersReqBuilder) UserGroupType(userGroupType string) *PatchUserGroupMembersReqBuilder {
	builder.apiReq.Body.(*PatchUserGroupMembersReqBody).UserGroupType = &userGroupType
	return builder
}


func (builder *PatchUserGroupMembersReqBuilder) UserGroupID(userGroupID string) *PatchUserGroupMembersReqBuilder {
	builder.apiReq.Body.(*PatchUserGroupMembersReqBody).UserGroupID = &userGroupID
	return builder
}


func (builder *PatchUserGroupMembersReqBuilder) AddUsers(addUsers []string) *PatchUserGroupMembersReqBuilder {
	builder.apiReq.Body.(*PatchUserGroupMembersReqBody).AddUsers = addUsers
	return builder
}

func (builder *PatchUserGroupMembersReqBuilder) DeleteUsers(deleteUsers []string) *PatchUserGroupMembersReqBuilder {
	builder.apiReq.Body.(*PatchUserGroupMembersReqBody).DeleteUsers = deleteUsers
	return builder
}

func (builder *PatchUserGroupMembersReqBuilder) ReplaceUsers(replaceUsers []string) *PatchUserGroupMembersReqBuilder {
	builder.apiReq.Body.(*PatchUserGroupMembersReqBody).ReplaceUsers = replaceUsers
	return builder
}
func (builder *PatchUserGroupMembersReqBuilder) Build() *PatchUserGroupMembersReq {
	req := &PatchUserGroupMembersReq{}
	req.apiReq = builder.apiReq
	return req
}

type QueryDepartmentReq struct {
	apiReq *core.APIReq
}
type QueryDepartmentReqBody struct {
    Scopes  []string `json:"scopes,omitempty"`
    DepartmentMeegoKey  *string `json:"department_meego_key,omitempty"`
}
type QueryDepartmentResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data      *QueryDepartmentRespData        `json:"data,omitempty"`
}

type QueryDepartmentRespData struct {
	DepartmentMeegoKey       *string         `json:"department_meego_key,omitempty"`
	Name       map[string]string         `json:"name,omitempty"`
	ExternalIDs       map[string]string         `json:"external_ids,omitempty"`
	OrgLevel       *string         `json:"org_level,omitempty"`
	Parent       *DepartmentBasicInfo         `json:"parent,omitempty"`
	AllChildren       []DepartmentBasicInfo         `json:"all_children,omitempty"`
	DirectEmployees       []AccountInfo         `json:"direct_employees,omitempty"`
	AllEmployees       []AccountInfo         `json:"all_employees,omitempty"`
}

type QueryDepartmentReqBuilder struct {
	apiReq *core.APIReq
}

func NewQueryDepartmentReqBuilder() *QueryDepartmentReqBuilder {
	builder := &QueryDepartmentReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &QueryDepartmentReqBody{},
	}
	return builder
}

func (builder *QueryDepartmentReqBuilder) Scopes(scopes []string) *QueryDepartmentReqBuilder {
	builder.apiReq.Body.(*QueryDepartmentReqBody).Scopes = scopes
	return builder
}

func (builder *QueryDepartmentReqBuilder) DepartmentMeegoKey(departmentMeegoKey string) *QueryDepartmentReqBuilder {
	builder.apiReq.Body.(*QueryDepartmentReqBody).DepartmentMeegoKey = &departmentMeegoKey
	return builder
}

func (builder *QueryDepartmentReqBuilder) Build() *QueryDepartmentReq {
	req := &QueryDepartmentReq{}
	req.apiReq = builder.apiReq
	return req
}

type QueryTenantReq struct {
	apiReq *core.APIReq
}
type QueryTenantReqBody struct {
    Scopes  []string `json:"scopes,omitempty"`
}
type QueryTenantResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data      *QueryTenantRespData        `json:"data,omitempty"`
}

type QueryTenantRespData struct {
	TenantMeegoKey       *string         `json:"tenant_meego_key,omitempty"`
	Name       *string         `json:"name,omitempty"`
	ExternalIDs       map[string]string         `json:"external_ids,omitempty"`
	Icon       *string         `json:"icon,omitempty"`
	Status       *string         `json:"status,omitempty"`
	PlatformType       *string         `json:"platform_type,omitempty"`
	Departments       []DepartmentInfo         `json:"departments,omitempty"`
	Accounts       []AccountInfo         `json:"accounts,omitempty"`
}

type QueryTenantReqBuilder struct {
	apiReq *core.APIReq
}

func NewQueryTenantReqBuilder() *QueryTenantReqBuilder {
	builder := &QueryTenantReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &QueryTenantReqBody{},
	}
	return builder
}

func (builder *QueryTenantReqBuilder) Scopes(scopes []string) *QueryTenantReqBuilder {
	builder.apiReq.Body.(*QueryTenantReqBody).Scopes = scopes
	return builder
}
func (builder *QueryTenantReqBuilder) Build() *QueryTenantReq {
	req := &QueryTenantReq{}
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

func (builder *QueryUserDetailReqBuilder) TenantKey(tenantKey string) *QueryUserDetailReqBuilder {
	builder.apiReq.Body.(*QueryUserDetailReqBody).TenantKey = &tenantKey
	return builder
}

func (builder *QueryUserDetailReqBuilder) Build() *QueryUserDetailReq {
	req := &QueryUserDetailReq{}
	req.apiReq = builder.apiReq
	return req
}

type QueryUsersForTenantPluginReq struct {
	apiReq *core.APIReq
}
type QueryUsersForTenantPluginReqBody struct {
    UserKeys  []string `json:"user_keys,omitempty"`
    Emails  []string `json:"emails,omitempty"`
    OutUsers  []OutUserInfo `json:"out_users,omitempty"`
}
type QueryUsersForTenantPluginResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       []UserBasicInfo         `json:"data"`
	
}

type QueryUsersForTenantPluginReqBuilder struct {
	apiReq *core.APIReq
}

func NewQueryUsersForTenantPluginReqBuilder() *QueryUsersForTenantPluginReqBuilder {
	builder := &QueryUsersForTenantPluginReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &QueryUsersForTenantPluginReqBody{},
	}
	return builder
}

func (builder *QueryUsersForTenantPluginReqBuilder) UserKeys(userKeys []string) *QueryUsersForTenantPluginReqBuilder {
	builder.apiReq.Body.(*QueryUsersForTenantPluginReqBody).UserKeys = userKeys
	return builder
}

func (builder *QueryUsersForTenantPluginReqBuilder) Emails(emails []string) *QueryUsersForTenantPluginReqBuilder {
	builder.apiReq.Body.(*QueryUsersForTenantPluginReqBody).Emails = emails
	return builder
}

func (builder *QueryUsersForTenantPluginReqBuilder) OutUsers(outUsers []OutUserInfo) *QueryUsersForTenantPluginReqBuilder {
	builder.apiReq.Body.(*QueryUsersForTenantPluginReqBody).OutUsers = outUsers
	return builder
}
func (builder *QueryUsersForTenantPluginReqBuilder) Build() *QueryUsersForTenantPluginReq {
	req := &QueryUsersForTenantPluginReq{}
	req.apiReq = builder.apiReq
	return req
}

type SearchUserByWordReq struct {
	apiReq *core.APIReq
}
type SearchUserByWordReqBody struct {
    Query  *string `json:"query,omitempty"`
    ProjectKey  *string `json:"project_key,omitempty"`
}
type SearchUserByWordResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data       []UserBasicInfo         `json:"data"`
	
}

type SearchUserByWordReqBuilder struct {
	apiReq *core.APIReq
}

func NewSearchUserByWordReqBuilder() *SearchUserByWordReqBuilder {
	builder := &SearchUserByWordReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &SearchUserByWordReqBody{},
	}
	return builder
}

func (builder *SearchUserByWordReqBuilder) Query(query string) *SearchUserByWordReqBuilder {
	builder.apiReq.Body.(*SearchUserByWordReqBody).Query = &query
	return builder
}


func (builder *SearchUserByWordReqBuilder) ProjectKey(projectKey string) *SearchUserByWordReqBuilder {
	builder.apiReq.Body.(*SearchUserByWordReqBody).ProjectKey = &projectKey
	return builder
}

func (builder *SearchUserByWordReqBuilder) Build() *SearchUserByWordReq {
	req := &SearchUserByWordReq{}
	req.apiReq = builder.apiReq
	return req
}

type StopAccountReq struct {
	apiReq *core.APIReq
}
type StopAccountReqBody struct {
    UserMeegoKey  *string `json:"user_meego_key,omitempty"`
}
type StopAccountResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
}

type StopAccountReqBuilder struct {
	apiReq *core.APIReq
}

func NewStopAccountReqBuilder() *StopAccountReqBuilder {
	builder := &StopAccountReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &StopAccountReqBody{},
	}
	return builder
}

func (builder *StopAccountReqBuilder) UserMeegoKey(userMeegoKey string) *StopAccountReqBuilder {
	builder.apiReq.Body.(*StopAccountReqBody).UserMeegoKey = &userMeegoKey
	return builder
}

func (builder *StopAccountReqBuilder) Build() *StopAccountReq {
	req := &StopAccountReq{}
	req.apiReq = builder.apiReq
	return req
}

type UpdateAccountReq struct {
	apiReq *core.APIReq
}
type UpdateAccountReqBody struct {
    UserMeegoKey  *string `json:"user_meego_key,omitempty"`
    Name  map[string]string `json:"name,omitempty"`
    AvatarUrl  *string `json:"avatar_url,omitempty"`
    DepartmentMeegoKeys  []string `json:"department_meego_keys,omitempty"`
}
type UpdateAccountResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
}

type UpdateAccountReqBuilder struct {
	apiReq *core.APIReq
}

func NewUpdateAccountReqBuilder() *UpdateAccountReqBuilder {
	builder := &UpdateAccountReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &UpdateAccountReqBody{},
	}
	return builder
}

func (builder *UpdateAccountReqBuilder) UserMeegoKey(userMeegoKey string) *UpdateAccountReqBuilder {
	builder.apiReq.Body.(*UpdateAccountReqBody).UserMeegoKey = &userMeegoKey
	return builder
}


func (builder *UpdateAccountReqBuilder) Name(name map[string]string) *UpdateAccountReqBuilder {
	builder.apiReq.Body.(*UpdateAccountReqBody).Name = name
	return builder
}

func (builder *UpdateAccountReqBuilder) AvatarUrl(avatarUrl string) *UpdateAccountReqBuilder {
	builder.apiReq.Body.(*UpdateAccountReqBody).AvatarUrl = &avatarUrl
	return builder
}


func (builder *UpdateAccountReqBuilder) DepartmentMeegoKeys(departmentMeegoKeys []string) *UpdateAccountReqBuilder {
	builder.apiReq.Body.(*UpdateAccountReqBody).DepartmentMeegoKeys = departmentMeegoKeys
	return builder
}
func (builder *UpdateAccountReqBuilder) Build() *UpdateAccountReq {
	req := &UpdateAccountReq{}
	req.apiReq = builder.apiReq
	return req
}

type UpdateDepartmentReq struct {
	apiReq *core.APIReq
}
type UpdateDepartmentReqBody struct {
    DepartmentMeegoKey  *string `json:"department_meego_key,omitempty"`
    Name  map[string]string `json:"name,omitempty"`
    ParentDepartmentMeegoKey  *string `json:"parent_department_meego_key,omitempty"`
}
type UpdateDepartmentResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
}

type UpdateDepartmentReqBuilder struct {
	apiReq *core.APIReq
}

func NewUpdateDepartmentReqBuilder() *UpdateDepartmentReqBuilder {
	builder := &UpdateDepartmentReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &UpdateDepartmentReqBody{},
	}
	return builder
}

func (builder *UpdateDepartmentReqBuilder) DepartmentMeegoKey(departmentMeegoKey string) *UpdateDepartmentReqBuilder {
	builder.apiReq.Body.(*UpdateDepartmentReqBody).DepartmentMeegoKey = &departmentMeegoKey
	return builder
}


func (builder *UpdateDepartmentReqBuilder) Name(name map[string]string) *UpdateDepartmentReqBuilder {
	builder.apiReq.Body.(*UpdateDepartmentReqBody).Name = name
	return builder
}

func (builder *UpdateDepartmentReqBuilder) ParentDepartmentMeegoKey(parentDepartmentMeegoKey string) *UpdateDepartmentReqBuilder {
	builder.apiReq.Body.(*UpdateDepartmentReqBody).ParentDepartmentMeegoKey = &parentDepartmentMeegoKey
	return builder
}

func (builder *UpdateDepartmentReqBuilder) Build() *UpdateDepartmentReq {
	req := &UpdateDepartmentReq{}
	req.apiReq = builder.apiReq
	return req
}

type UserGroupMembersPageReq struct {
	apiReq *core.APIReq
}
type UserGroupMembersPageReqBody struct {
    UserGroupType  *string `json:"user_group_type,omitempty"`
    UserGroupIDs  []string `json:"user_group_ids,omitempty"`
    PageNum  *int64 `json:"page_num,omitempty"`
    PageSize  *int64 `json:"page_size,omitempty"`
}
type UserGroupMembersPageResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data      *UserGroupMembersPageRespData        `json:"data,omitempty"`
}

type UserGroupMembersPageRespData struct {
	List       []UserGroupDetail         `json:"list,omitempty"`
	Pagination       *Pagination         `json:"pagination,omitempty"`
}

type UserGroupMembersPageReqBuilder struct {
	apiReq *core.APIReq
}

func NewUserGroupMembersPageReqBuilder() *UserGroupMembersPageReqBuilder {
	builder := &UserGroupMembersPageReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &UserGroupMembersPageReqBody{},
	}
	return builder
}

func (builder *UserGroupMembersPageReqBuilder) ProjectKey(projectKey string) *UserGroupMembersPageReqBuilder {
	builder.apiReq.PathParams.Set("project_key", fmt.Sprint(projectKey))
	return builder
}


func (builder *UserGroupMembersPageReqBuilder) UserGroupType(userGroupType string) *UserGroupMembersPageReqBuilder {
	builder.apiReq.Body.(*UserGroupMembersPageReqBody).UserGroupType = &userGroupType
	return builder
}


func (builder *UserGroupMembersPageReqBuilder) UserGroupIDs(userGroupIDs []string) *UserGroupMembersPageReqBuilder {
	builder.apiReq.Body.(*UserGroupMembersPageReqBody).UserGroupIDs = userGroupIDs
	return builder
}

func (builder *UserGroupMembersPageReqBuilder) PageNum(pageNum int64) *UserGroupMembersPageReqBuilder {
	builder.apiReq.Body.(*UserGroupMembersPageReqBody).PageNum = &pageNum
	return builder
}


func (builder *UserGroupMembersPageReqBuilder) PageSize(pageSize int64) *UserGroupMembersPageReqBuilder {
	builder.apiReq.Body.(*UserGroupMembersPageReqBody).PageSize = &pageSize
	return builder
}

func (builder *UserGroupMembersPageReqBuilder) Build() *UserGroupMembersPageReq {
	req := &UserGroupMembersPageReq{}
	req.apiReq = builder.apiReq
	return req
}

