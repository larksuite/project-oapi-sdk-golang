package user

import (
	"context"
	"fmt"

	"github.com/larksuite/project-oapi-sdk-golang/core"
)


const APIPath_OAPIActiveAccount = "/open_api/account/active"

const APIPath_OAPICreateAccount = "/open_api/account/create"

const APIPath_OAPICreateDepartment = "/open_api/department/create"

const APIPath_OAPICreateUserGroup = "/open_api/:project_key/user_group"

const APIPath_OAPIDeleteDepartment = "/open_api/department/delete"

const APIPath_OAPIPageUserGroupMembers = "/open_api/:project_key/user_groups/members/page"

const APIPath_OAPIPatchUserGroupMembers = "/open_api/:project_key/user_group/members"

const APIPath_OAPIQueryDepartment = "/open_api/department/query"

const APIPath_OAPIQueryTenant = "/open_api/tenant/query"

const APIPath_OAPIQueryUsersForTenantPlugin = "/open_api/user/tp/query"

const APIPath_OAPISearchUserByWord = "/open_api/user/search"

const APIPath_OAPIStopAccount = "/open_api/account/stop"

const APIPath_OAPIUpdateAccount = "/open_api/account/update"

const APIPath_OAPIUpdateDepartment = "/open_api/department/update"

const APIPath_QueryUserDetail = "/open_api/user/query"


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
func (a *UserService) OAPIActiveAccount(ctx context.Context, req *OAPIActiveAccountReq, options ...core.RequestOptionFunc) (*OAPIActiveAccountResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPIActiveAccount
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIActiveAccount] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPIActiveAccountResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIActiveAccount] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPICreateAccount
 * @desc: OAPI新增账号
 */
func (a *UserService) OAPICreateAccount(ctx context.Context, req *OAPICreateAccountReq, options ...core.RequestOptionFunc) (*OAPICreateAccountResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPICreateAccount
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPICreateAccount] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPICreateAccountResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPICreateAccount] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPICreateDepartment
 * @desc: OAPI新增部门
 */
func (a *UserService) OAPICreateDepartment(ctx context.Context, req *OAPICreateDepartmentReq, options ...core.RequestOptionFunc) (*OAPICreateDepartmentResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPICreateDepartment
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPICreateDepartment] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPICreateDepartmentResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPICreateDepartment] fail to unmarshal response body, error: %v", err.Error()))
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
func (a *UserService) OAPICreateUserGroup(ctx context.Context, req *OAPICreateUserGroupReq, options ...core.RequestOptionFunc) (*OAPICreateUserGroupResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPICreateUserGroup
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPICreateUserGroup] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPICreateUserGroupResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPICreateUserGroup] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIDeleteDepartment
 * @desc: OAPI删除部门
 */
func (a *UserService) OAPIDeleteDepartment(ctx context.Context, req *OAPIDeleteDepartmentReq, options ...core.RequestOptionFunc) (*OAPIDeleteDepartmentResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPIDeleteDepartment
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIDeleteDepartment] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPIDeleteDepartmentResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIDeleteDepartment] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: 获取指定用户组成员列表
 * @desc: 允许获取预置用户组成员列表；只返回用户，不返回部门
 * @err : https://lark-devops.bytedance.net/page/tools/error-code/list
 */
func (a *UserService) OAPIPageUserGroupMembers(ctx context.Context, req *OAPIPageUserGroupMembersReq, options ...core.RequestOptionFunc) (*OAPIPageUserGroupMembersResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPIPageUserGroupMembers
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIPageUserGroupMembers] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPIPageUserGroupMembersResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIPageUserGroupMembers] fail to unmarshal response body, error: %v", err.Error()))
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
func (a *UserService) OAPIPatchUserGroupMembers(ctx context.Context, req *OAPIPatchUserGroupMembersReq, options ...core.RequestOptionFunc) (*OAPIPatchUserGroupMembersResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPIPatchUserGroupMembers
	apiReq.HttpMethod = "PATCH"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIPatchUserGroupMembers] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPIPatchUserGroupMembersResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIPatchUserGroupMembers] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIQueryDepartment
 * @desc: OAPI查询部门
 */
func (a *UserService) OAPIQueryDepartment(ctx context.Context, req *OAPIQueryDepartmentReq, options ...core.RequestOptionFunc) (*OAPIQueryDepartmentResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPIQueryDepartment
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIQueryDepartment] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPIQueryDepartmentResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIQueryDepartment] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIQueryTenant
 * @desc: OAPI查询租户
 */
func (a *UserService) OAPIQueryTenant(ctx context.Context, req *OAPIQueryTenantReq, options ...core.RequestOptionFunc) (*OAPIQueryTenantResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPIQueryTenant
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIQueryTenant] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPIQueryTenantResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIQueryTenant] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIQueryUsersForTenantPlugin
 * @desc: 获取用户详情信息(租户插件)
 */
func (a *UserService) OAPIQueryUsersForTenantPlugin(ctx context.Context, req *OAPIQueryUsersForTenantPluginReq, options ...core.RequestOptionFunc) (*OAPIQueryUsersForTenantPluginResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPIQueryUsersForTenantPlugin
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIQueryUsersForTenantPlugin] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPIQueryUsersForTenantPluginResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIQueryUsersForTenantPlugin] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPISearchUserByWord
 * @desc: 模糊查询指定空间的用户列表
 */
func (a *UserService) OAPISearchUserByWord(ctx context.Context, req *OAPISearchUserByWordReq, options ...core.RequestOptionFunc) (*OAPISearchUserByWordResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPISearchUserByWord
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPISearchUserByWord] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPISearchUserByWordResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPISearchUserByWord] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIStopAccount
 * @desc: OAPI停用账号
 */
func (a *UserService) OAPIStopAccount(ctx context.Context, req *OAPIStopAccountReq, options ...core.RequestOptionFunc) (*OAPIStopAccountResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPIStopAccount
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIStopAccount] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPIStopAccountResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIStopAccount] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIUpdateAccount
 * @desc: OAPI更新账号信息
 */
func (a *UserService) OAPIUpdateAccount(ctx context.Context, req *OAPIUpdateAccountReq, options ...core.RequestOptionFunc) (*OAPIUpdateAccountResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPIUpdateAccount
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIUpdateAccount] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPIUpdateAccountResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIUpdateAccount] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, err
}

/*
 * @name: OAPIUpdateDepartment
 * @desc: OAPI更新部门
 */
func (a *UserService) OAPIUpdateDepartment(ctx context.Context, req *OAPIUpdateDepartmentReq, options ...core.RequestOptionFunc) (*OAPIUpdateDepartmentResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = APIPath_OAPIUpdateDepartment
	apiReq.HttpMethod = "POST"
	apiResp, err := core.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIUpdateDepartment] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	// 反序列响应结果
	resp := &OAPIUpdateDepartmentResp{APIResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		a.config.Logger.Error(ctx, fmt.Sprintf("[OAPIUpdateDepartment] fail to unmarshal response body, error: %v", err.Error()))
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

