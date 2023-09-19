package project_test

import (
	"code.byted.org/bits/project-oapi-sdk-golang/core"
	sdkcore "code.byted.org/bits/project-oapi-sdk-golang/core"
	"code.byted.org/bits/project-oapi-sdk-golang/service/project"

	"context"
	"encoding/json"
	"testing"
)

func TestProjectService_GetProjectDetail(t *testing.T) {
	type fields struct {
		config *core.Config
	}
	type args struct {
		ctx     context.Context
		req     *project.GetProjectDetailReq
		options []core.RequestOptionFunc
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *project.GetProjectDetailResp
		wantErr bool
	}{
		{
			name: "获取空间详情", // 这个是case的名称
			args: args{
				ctx:     context.Background(),
				req:     project.NewGetProjectDetailReqBuilder().ProjectKeys([]string{"63b3f8ee3a3c74bbbd88d546"}).Build(),
				options: []core.RequestOptionFunc{sdkcore.WithUserKey("7116154537953476628")},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := client.Project.GetProjectDetail(tt.args.ctx, tt.args.req, tt.args.options...)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetProjectDetail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			marshal, _ := json.Marshal(got)
			t.Logf("res: %s", marshal)
		})
	}
}

func TestProjectService_ListProject(t *testing.T) {
	type fields struct {
		config *core.Config
	}
	type args struct {
		ctx     context.Context
		req     *project.ListProjectReq
		options []core.RequestOptionFunc
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "获取空间列表", // 这个是case的名称
			args: args{
				ctx:     context.Background(),
				req:     project.NewListProjectReqBuilder().UserKey("7116154537953476628").Build(),
				options: []core.RequestOptionFunc{sdkcore.WithUserKey("7116154537953476628"), sdkcore.WithAccessToken("123")},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := client.Project.ListProject(tt.args.ctx, tt.args.req, tt.args.options...)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListProject() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			marshal, _ := json.Marshal(got)
			t.Logf("res: %s", marshal)
		})
	}
}

func TestProjectService_ListProjectBusiness(t *testing.T) {
	type fields struct {
		config *core.Config
	}
	type args struct {
		ctx     context.Context
		req     *project.ListProjectBusinessReq
		options []core.RequestOptionFunc
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *project.ListProjectBusinessResp
		wantErr bool
	}{
		{
			name: "获取空间下业务线详情", // 这个是case的名称
			args: args{
				ctx:     context.Background(),
				req:     project.NewListProjectBusinessReqBuilder().Build(),
				options: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := client.Project.ListProjectBusiness(tt.args.ctx, tt.args.req, tt.args.options...)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListProjectBusiness() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			marshal, _ := json.Marshal(got)
			t.Logf("res: %s", marshal)
		})
	}
}

func TestProjectService_ListProjectTeam(t *testing.T) {
	type fields struct {
		config *core.Config
	}
	type args struct {
		ctx     context.Context
		req     *project.ListProjectTeamReq
		options []core.RequestOptionFunc
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *project.ListProjectTeamResp
		wantErr bool
	}{
		{
			name: "获取空间下团队人员", // 这个是case的名称
			args: args{
				ctx:     context.Background(),
				req:     project.NewListProjectTeamReqBuilder().Build(),
				options: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := client.Project.ListProjectTeam(tt.args.ctx, tt.args.req, tt.args.options...)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListProjectTeam() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			marshal, _ := json.Marshal(got)
			t.Logf("res: %s", marshal)
		})
	}
}

func TestProjectService_ListProjectWorkItemType(t *testing.T) {
	type fields struct {
		config *core.Config
	}
	type args struct {
		ctx     context.Context
		req     *project.ListProjectWorkItemTypeReq
		options []core.RequestOptionFunc
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *project.ListProjectWorkItemTypeResp
		wantErr bool
	}{
		{
			name: "获取空间下工作项类型", // 这个是case的名称
			args: args{
				ctx:     context.Background(),
				req:     project.NewListProjectWorkItemTypeReqBuilder().Build(),
				options: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := client.Project.ListProjectWorkItemType(tt.args.ctx, tt.args.req, tt.args.options...)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListProjectWorkItemType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			marshal, _ := json.Marshal(got)
			t.Logf("res: %s", marshal)
		})
	}
}
