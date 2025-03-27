package main

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/larksuite/project-oapi-sdk-golang/service/workitem"

	meegoSDK "github.com/larksuite/project-oapi-sdk-golang"
	sdkcore "github.com/larksuite/project-oapi-sdk-golang/core"
)

func main() {
	//plugin_id := "MII_661E3D92B64A4004"
	//plugin_secret := "40F8C24AD263B9B12ECF55AF49C5AF45"
	//project_id := "662e151cd7fd734ae1874213" //copydemo
	//operater_id := "7356795280408297476"

	//baseUrl := "https://meego.larkoffice.com"

	plugin_id := "MII_673FF7824A400023"
	plugin_secret := "796111A38554F88C703827AB61BBE4B2"
	project_id := "6645776a0752c9b1fb2b9ef8" //--
	operater_id := "7426311932258287635"

	baseUrl := "https://meego.feishu-boe.cn"

	option := func(cfg *sdkcore.Config) {
		//cfg.AccessTokenType = sdkcore.AccessTokenTypeVirtualPlugin
		cfg.BaseUrl = baseUrl
	}
	client := meegoSDK.NewClient(plugin_id, plugin_secret, option)

	// header := make(http.Header)
	// header.Add("x-tt-env", "boe_feature_jiangxin_dev_new")
	// header.Add("x-use-ppe", "1")

	// resp, err := client.WorkItem.Filter(context.Background(), workitem.NewFilterReqBuilder().
	// 	ProjectKey(project_id).
	// 	WorkItemTypeKeys([]string{"story"}).
	// 	// WorkItemIDs([]int64{4676574614}).
	// 	// Expand(&workitem.Expand{
	// 	// 	NeedMultiText: true,
	// 	// }).
	// 	SearchID("Rs5WNERHR").
	// 	Build(),
	// 	sdkcore.WithUserKey(operater_id),
	// )

	//isAutoFalse := false
	//isAutoTrue := true
	// resp, err := client.WorkItem.NodeUpdate(context.Background(),
	// 	workitem.NewNodeUpdateReqBuilder().
	// 		WorkItemTypeKey("story").
	// 		WorkItemID(5370997663).
	// 		ProjectKey(project_id).
	// 		NodeID("doing").
	// 		Schedules([]*workitem.Schedule{
	// 			// &workitem.Schedule{
	// 			// 	Points: 2,
	// 			// 	// EstimateStartDate: 1723132800000,
	// 			// 	// EstimateEndDate:   1723564799999,
	// 			// 	IsAuto: &isAutoTrue,
	// 			// 	Owners: []string{
	// 			// 		"7356795280408297476",
	// 			// 	},
	// 			// },
	// 			&workitem.Schedule{
	// 				Points: 1,
	// 				// EstimateStartDate: 1723132800000,
	// 				// EstimateEndDate:   1723564799999,
	// 				// IsAuto: &isAutoTrue,
	// 				Owners: []string{
	// 					"7311891981507100700",
	// 				},
	// 			},
	// 		}).
	// 		Build(),
	// 	sdkcore.WithUserKey(operater_id))

	//points := 10.0
	//resp, err := client.Task.UpdateSubTask(context.Background(),
	//	task.NewUpdateSubTaskReqBuilder().
	//		WorkItemTypeKey("story").
	//		WorkItemID(5370997663).
	//		ProjectKey(project_id).
	//		NodeID("doing").
	//		TaskID(5371641336).
	//		Schedule(
	//			&workitem.Schedule{
	//				Points: &points,
	//				//EstimateStartDate: 1723132800000,
	//				//EstimateEndDate:   1723564799999,
	//				IsAuto: &isAutoFalse,
	//				Owners: []string{
	//					"7356795280408297476",
	//				},
	//			},
	//		).
	//		Build(),
	//	sdkcore.WithUserKey(operater_id))

	searchGroup := workitem.SearchGroup{
		SearchParams: []*workitem.SearchParam{
			&workitem.SearchParam{
				ParamKey: "created_at",
				Value:    1654064482000,
				Operator: ">",
			},
		},
		Conjunction: "AND",
	}
	resp, err := client.WorkItem.UniversalSearch(context.Background(), workitem.NewUniversalSearchReqBuilder().DataSources([]workitem.DataSource{
		workitem.DataSource{
			ProjectKey:       &project_id,
			WorkItemTypeKeys: thrift.StringPtr("story"),
		},
	}).SearchGroup(&searchGroup).Build(),
		sdkcore.WithUserKey(operater_id))

	//处理错误
	if err != nil {
		fmt.Printf("err:%v \n", err)
		// 处理err
		return
	}

	// 服务端错误处理
	if !resp.Success() {
		fmt.Println(resp.RequestId(), resp.CodeError)
		return
	}

	// 业务数据处理
	fmt.Println(sdkcore.Prettify(resp))
}
