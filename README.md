# 飞书项目开放接口SDK

旨在让开发者便捷的调用飞书项目OPEN API

## 目录

<!-- toc -->

- [安装](#安装)
- [API Client](#api-client)
    - [创建API Client](#创建api-client)
    - [配置API Client](#配置api-client)

- [API调用](#api调用)
    - [基本用法](#基本用法)
    - [设置请求选项](#设置请求选项)
- [加入答疑群](#加入答疑群)

<!-- tocstop -->

## 安装

```go
go get -u code.byted.org/bits/project-oapi-sdk-golang
```

## API Client

开发者在调用 API 前，需要先创建一个 API Client，然后才可以基于 API Client 发起 API 调用，Client对象线程安全，全局初始化一次即可。

### 创建API Client

```go
var client = project.NewClient("PluginID", "PluginSecret") //默认插件身份凭证
```

### 配置API Client

创建 API Client 时，可对 API Client 进行一定的配置，比如我们可以在创建 API Client 时设置日志级别、设置 http 请求超时时间等等：

```go
var client = project.NewClient("PluginID", "PluginSecret",
project.WithLogLevel(projectcore.LogLevelDebug),
project.WithReqTimeout(3*time.Second),
project.WithEnableTokenCache(true),
project.WithHttpClient(http.DefaultClient))
```

每个配置选项的具体含义，如下表格：

<table>
  <thead align=left>
    <tr>
      <th>
        配置选项
      </th>
      <th>
        配置方式
      </th>
       <th>
        描述
      </th>
    </tr>
  </thead>
  <tbody align=left valign=top>
    <tr>
      <th>
        <code>LogLevel</code>
      </th>
      <td>
        <code>project.WithLogLevel(logLevel projectcore.LogLevel)</code>
      </td>
      <td>
设置 API Client 的日志输出级别(默认为 Info 级别)，枚举值如下：

- LogLevelDebug
- LogLevelInfo
- LogLevelWarn
- LogLevelError

</td>
</tr>

<tr>
      <th>
        <code>Logger</code>
      </th>
      <td>
        <code>project.WithLogger(logger projectcore.Logger)</code>
      </td>
      <td>
设置API Client的日志器，默认日志输出到标准输出。

开发者可通过实现下面的 Logger 接口，来设置自定义的日志器:

```go
type Logger interface {
Debug(context.Context, ...interface{})
Info(context.Context, ...interface{})
Warn(context.Context, ...interface{})
Error(context.Context, ...interface{})
}
```

</td>
</tr>

<tr>
      <th>
        <code>LogReqAtDebug</code>
      </th>
      <td>
        <code>project.WithLogReqAtDebug(printReqRespLog bool)</code>
      </td>
      <td>
设置是否开启 Http 请求参数和响应参数的日志打印开关；开启后，在 debug 模式下会打印 http 请求和响应的 headers,body 等信息。

在排查问题时，开启该选项，有利于问题的排查。

</td>
</tr>


<tr>
      <th>
        <code>BaseUrl</code>
      </th>
      <td>
        <code>project.WithOpenBaseUrl(baseUrl string)</code>
      </td>
      <td>
设置飞书项目域名，默认为https://project.feishu.cn

</td>
</tr>

<tr>
      <th>
        <code>AccessTokenType</code>
      </th>
      <td>
        <code>project.WithAccessTokenType(accessTokenType projectcore.AccessTokenType)</code>
      </td>
      <td>
设置 plugin_token 类型，枚举值如下：

- AccessTokenTypePlugin（插件身份凭证）
- AccessTokenTypeVirtualPlugin（虚拟plugin_token）
- AccessTokenTypeUserPlugin（用户身份凭证）

</td>
</tr>

<tr>
      <th>
        <code>TokenCache</code>
      </th>
      <td>
        <code>project.WithTokenCache(cache projectcore.Cache)</code>
      </td>
      <td>
设置 token 缓存器，用来缓存 token, 默认实现为内存。

如开发者想要定制 token 缓存器，需实现下面 Cache 接口:

```go
type Cache interface {
Set(ctx context.Context, key string, value string, expireTime time.Duration) error
Get(ctx context.Context, key string) (string, error)
}
```

</td>
</tr>


<tr>
      <th>
        <code>EnableTokenCache</code>
      </th>
      <td>
        <code>project.WithEnableTokenCache(enableTokenCache bool)</code>
      </td>
      <td>
设置是否开启 Token 的自动获取与缓存。

默认开启，如需要关闭可传递 false。
</td>
</tr>

<tr>
      <th>
        <code>ReqTimeout</code>
      </th>
      <td>
        <code>project.WithReqTimeout(time time.Duration)</code>
      </td>
      <td>
设置 SDK 内置的 Http Client 的请求超时时间，默认为0代表永不超时。
</td>
</tr>

<tr>
      <th>
        <code>HttpClient</code>
      </th>
      <td>
        <code>project.WithHttpClient(httpClient projectcore.HttpClient)</code>
      </td>
      <td>
设置 HttpClient，用于替换 SDK 提供的默认实现。

开发者可通过实现下面的 HttpClient 接口来设置自定义的 HttpClient:

```go
type HttpClient interface {
Do(*http.Request) (*http.Response, error)
}

```

</td>
</tr>

  </tbody>
</table>

## API调用

创建完毕 API Client，我们可以使用 ``Client.业务域.方法名称`` 来定位具体的 API 方法，然后对具体的 API 发起调用。

飞书项目开放平台开放的所有 API
列表，可点击[这里查看](https://bytedance.feishu.cn/docs/doccntEfMPoh8Qv3hDCshfRLEuY#9kPTAz)

### 基本用法

如下示例我们通过 client 调用空间业务的 ListProjectWorkItemType
方法，获取空间下的工作项类型列表（注：请求头中的X-USER-KEY和X-IDEM-UUID参数可以通过ReqBuilder中的AccessUser和UUID方法进行设置）：

``` go
import (
	"context"
	"fmt"

	sdk "code.byted.org/bits/project-oapi-sdk-golang"
	sdkcore "code.byted.org/bits/project-oapi-sdk-golang/core"
	"code.byted.org/bits/project-oapi-sdk-golang/service/project"
	"net/http"
)


func main() {
	// 创建 client
	client := sdk.NewClient("PluginID", "PluginSecret")

	// 发起请求
	resp, err := client.Project.ListProjectWorkItemType(context.Background(), project.NewListProjectWorkItemTypeReqBuilder().
		AccessUser("user_key").
		ProjectKey("project_key").
		Build(),
	)

	//处理错误
	if err != nil {
           // 处理err
           return
	}

	// 服务端错误处理
	if !resp.Success() {
           fmt.Println(resp.Code, resp.Msg, resp.RequestId())
	   return 
	}

	// 业务数据处理
	fmt.Println(sdkcore.Prettify(resp.Data))
}
```

更多 API 调用示例：[./sample/demo.go](./sample/demo.go)

### 设置请求选项

开发者在每次发起 API 调用时，可以设置请求级别的一些参数，比如传递 UserPluginAccessToken ,自定义 Headers 等：

```go
import (
    "context"
    "fmt"

    sdk "code.byted.org/bits/project-oapi-sdk-golang"
    sdkcore "code.byted.org/bits/project-oapi-sdk-golang/core"
    "code.byted.org/bits/project-oapi-sdk-golang/service/project"
)

func main() {
    // 创建 client
    client := sdk.NewClient("PluginID", "PluginSecret", sdk.WithAccessTokenType(sdkcore.AccessTokenTypeUserPlugin))
    header := make(http.Header)
    header.Add("k1", "v1")
    // 发起请求
    resp, err := client.Project.ListProjectWorkItemType(context.Background(), project.NewListProjectWorkItemTypeReqBuilder().
                                AccessUser("user_key").
                                ProjectKey("project_key").
                                Build(),
                sdkcore.WithUserPluginAccessToken("user_plugin_token"), //设置用户身份凭证
                sdkcore.WithHeaders(header),                            //设置head
                )

    //处理错误
    if err != nil {
        // 处理err
        return
    }

    // 服务端错误处理
    if !resp.Success() {
        fmt.Println(resp.Code, resp.Msg, resp.RequestId())
        return
    }

    // 业务数据处理
    fmt.Println(sdkcore.Prettify(resp.Data))
}

```

如下表格，展示了所有请求级别可设置的选项：

<table>
  <thead align=left>
    <tr>
      <th>
        配置选项
      </th>
      <th>
        配置方式
      </th>
       <th>
        描述
      </th>
    </tr>
  </thead>
  <tbody align=left valign=top>
    <tr>
      <th>
        <code>Header</code>
      </th>
      <td>
        <code>projectcore.WithHeaders(header http.Header)</code>
      </td>
      <td>
设置自定义请求头，开发者可在发起请求时，这些请求头会被透传到飞书项目开放平台服务端。

</td>
</tr>

<tr>
      <th>
        <code>PluginAccessToken</code>
      </th>
      <td>
        <code>projectcore.WithPluginAccessToken(pluginAccessToken string)</code>
      </td>
      <td>
设置plugin_token，当开发者自己维护token 时（即创建Client时EnableTokenCache设置为了false），需通过该选项传递 plugin_token。

</td>
</tr>

<tr>
      <th>
        <code>VirtualPluginAccessToken</code>
      </th>
      <td>
        <code>projectcore.WithVirtualPluginAccessToken(virtualPluginAccessToken string)</code>
      </td>
      <td>
设置虚拟plugin_token，当开发者自己维护token 时（即创建Client时EnableTokenCache设置为了false），需通过该选项传递 virtual_plugin_token。

</td>
</tr>

<tr>
      <th>
        <code>UserPluginAccessToken</code>
      </th>
      <td>
        <code>projectcore.WithUserPluginAccessToken(userPluginAccessToken string)</code>
      </td>
      <td>
设置user_plugin_token, 当开发者使用用户身份凭证时（即创建Client时AccessTokenType设置为了AccessTokenTypeUserPlugin），需通过该选项传递 user_plugin_token。
</td>
</tr>


  </tbody>
</table>

## 错误自查

[查看错误码](https://bytedance.feishu.cn/docs/doccn3CyRRA52nL9HiYR2v9XC8K)



