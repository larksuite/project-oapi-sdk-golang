package project_test

import (
	sdk "code.byted.org/bits/project-oapi-sdk-golang"
	"code.byted.org/bits/project-oapi-sdk-golang/core"
	"os"
	"testing"
)

var client *sdk.Client

func TestMain(t *testing.M) {

	client = sdk.NewClient("MII_635675D9C2830014", "CE168E6DA69974C87FDEE1B488FE87F1",
		sdk.WithLogReqAtDebug(true),
		sdk.WithLogLevel(core.LogLevelDebug),
		sdk.WithAccessTokenType(core.AccessTokenTypePlugin),
		sdk.WithOpenBaseUrl("https://meego.feishu-boe.cn"),
	)

	os.Exit(t.Run())
}
