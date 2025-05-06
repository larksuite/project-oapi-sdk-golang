package user

import (
	"context"
	"fmt"

	"github.com/larksuite/project-oapi-sdk-golang/core"
)


const APIPath_ActiveAccount = "/open_api/account/active"

const APIPath_CreateAccount = "/open_api/account/create"

const APIPath_CreateDepartment = "/open_api/department/create"

const APIPath_CreateUserGroup = "/open_api/:project_key/user_group"

const APIPath_DeleteDepartment = "/open_api/department/delete"

const APIPath_ListProjectTeam = "/open_api/:project_key/teams/all"

const APIPath_PatchUserGroupMembers = "/open_api/:project_key/user_group/members"

const APIPath_QueryDepartment = "/open_api/department/query"

const APIPath_QueryTenant = "/open_api/tenant/query"

const APIPath_QueryUserDetail = "/open_api/user/query"

const APIPath_QueryUsersForTenantPlugin = "/open_api/user/tp/query"

const APIPath_SearchUserByWord = "/open_api/user/search"

const APIPath_StopAccount = "/open_api/account/stop"

const APIPath_UpdateAccount = "/open_api/account/update"

const APIPath_UpdateDepartment = "/open_api/department/update"

const APIPath_UserGroupMembersPage = "/open_api/:project_key/user_groups/members/page"


func NewService(config *core.Config) *UserService {
	a := &UserService{config: config}
	return a
}

type UserService struct {
	config *core.Config
}


/*
 * @name: OAPIActiveAccount
 * @desc: OAPI启用账号
 */
func (a *UserService) ActiveAccount(ctx context.Context, req *ActiveAccountReq, options ...core.RequestOptionFunc) (*ActiveAccountResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_ActiveAccount
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[ActiveAccount] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &ActiveAccountResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[ActiveAccount] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPICreateAccount
 * @desc: OAPI新增账号
 */
func (a *UserService) CreateAccount(ctx context.Context, req *CreateAccountReq, options ...core.RequestOptionFunc) (*CreateAccountResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_CreateAccount
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[CreateAccount] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateAccountResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[CreateAccount] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPICreateDepartment
 * @desc: OAPI新增部门
 */
func (a *UserService) CreateDepartment(ctx context.Context, req *CreateDepartmentReq, options ...core.RequestOptionFunc) (*CreateDepartmentResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_CreateDepartment
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[CreateDepartment] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateDepartmentResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[CreateDepartment] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * 用户组接口
@name: 创建自定义用户组
 * @desc: 只能创建自定义用户组，不能创建预置用户组，如空间管理员用户组&&空间成员用户组；只能添加用户，不能添加部门
 * @err : https://lark-devops.bytedance.net/page/tools/error-code/list
 */
func (a *UserService) CreateUserGroup(ctx context.Context, req *CreateUserGroupReq, options ...core.RequestOptionFunc) (*CreateUserGroupResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_CreateUserGroup
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[CreateUserGroup] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateUserGroupResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[CreateUserGroup] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIDeleteDepartment
 * @desc: OAPI删除部门
 */
func (a *UserService) DeleteDepartment(ctx context.Context, req *DeleteDepartmentReq, options ...core.RequestOptionFunc) (*DeleteDepartmentResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_DeleteDepartment
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[DeleteDepartment] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &DeleteDepartmentResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[DeleteDepartment] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * open_api下沉接口
@name: 获取空间下团队人员
 * @desc: open_api下沉
 */
func (a *UserService) ListProjectTeam(ctx context.Context, req *ListProjectTeamReq, options ...core.RequestOptionFunc) (*ListProjectTeamResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_ListProjectTeam
	apiReq.HttpMethod = "GET"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[ListProjectTeam] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &ListProjectTeamResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[ListProjectTeam] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * open.scope="project", open.func="ListProjectTeam"
@name: 更新用户组成员
 * @desc: 允许同时新增/删除用户组成员，允许操作预置用户组；只能操作用户，不能操作部门
 * @err : https://lark-devops.bytedance.net/page/tools/error-code/list
 */
func (a *UserService) PatchUserGroupMembers(ctx context.Context, req *PatchUserGroupMembersReq, options ...core.RequestOptionFunc) (*PatchUserGroupMembersResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_PatchUserGroupMembers
	apiReq.HttpMethod = "PATCH"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[PatchUserGroupMembers] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &PatchUserGroupMembersResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[PatchUserGroupMembers] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIQueryDepartment
 * @desc: OAPI查询部门
 */
func (a *UserService) QueryDepartment(ctx context.Context, req *QueryDepartmentReq, options ...core.RequestOptionFunc) (*QueryDepartmentResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_QueryDepartment
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[QueryDepartment] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &QueryDepartmentResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[QueryDepartment] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIQueryTenant
 * @desc: OAPI查询租户
 */
func (a *UserService) QueryTenant(ctx context.Context, req *QueryTenantReq, options ...core.RequestOptionFunc) (*QueryTenantResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_QueryTenant
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[QueryTenant] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &QueryTenantResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[QueryTenant] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIQueryUsers
 * @desc: 获取用户详情信息
 */
func (a *UserService) QueryUserDetail(ctx context.Context, req *QueryUserDetailReq, options ...core.RequestOptionFunc) (*QueryUserDetailResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_QueryUserDetail
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[QueryUserDetail] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &QueryUserDetailResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[QueryUserDetail] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIQueryUsersForTenantPlugin
 * @desc: 获取用户详情信息(租户插件)
 */
func (a *UserService) QueryUsersForTenantPlugin(ctx context.Context, req *QueryUsersForTenantPluginReq, options ...core.RequestOptionFunc) (*QueryUsersForTenantPluginResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_QueryUsersForTenantPlugin
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[QueryUsersForTenantPlugin] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &QueryUsersForTenantPluginResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[QueryUsersForTenantPlugin] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPISearchUserByWord
 * @desc: 模糊查询指定空间的用户列表
 */
func (a *UserService) SearchUserByWord(ctx context.Context, req *SearchUserByWordReq, options ...core.RequestOptionFunc) (*SearchUserByWordResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_SearchUserByWord
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[SearchUserByWord] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &SearchUserByWordResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[SearchUserByWord] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIStopAccount
 * @desc: OAPI停用账号
 */
func (a *UserService) StopAccount(ctx context.Context, req *StopAccountReq, options ...core.RequestOptionFunc) (*StopAccountResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_StopAccount
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[StopAccount] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &StopAccountResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[StopAccount] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIUpdateAccount
 * @desc: OAPI更新账号信息
 */
func (a *UserService) UpdateAccount(ctx context.Context, req *UpdateAccountReq, options ...core.RequestOptionFunc) (*UpdateAccountResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_UpdateAccount
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[UpdateAccount] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &UpdateAccountResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[UpdateAccount] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIUpdateDepartment
 * @desc: OAPI更新部门
 */
func (a *UserService) UpdateDepartment(ctx context.Context, req *UpdateDepartmentReq, options ...core.RequestOptionFunc) (*UpdateDepartmentResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_UpdateDepartment
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[UpdateDepartment] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &UpdateDepartmentResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[UpdateDepartment] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: 获取指定用户组成员列表
 * @desc: 允许获取预置用户组成员列表；只返回用户，不返回部门
 * @err : https://lark-devops.bytedance.net/page/tools/error-code/list
 */
func (a *UserService) UserGroupMembersPage(ctx context.Context, req *UserGroupMembersPageReq, options ...core.RequestOptionFunc) (*UserGroupMembersPageResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_UserGroupMembersPage
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[UserGroupMembersPage] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &UserGroupMembersPageResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[UserGroupMembersPage] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

