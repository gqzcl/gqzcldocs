## Home

**Iris** is a free, open-source [Go](https://golang.org/) web framework, created by [Gerasimos Maropoulos](https://twitter.com/MakisMaropoulos) and intended for the development of modern web applications.

Iris 是一个免费的开源 Go web 框架，由 Gerasimos Maropoulos 创建，用于开发现代 web 应用程序。

**Iris** helps back-end developers to quickly create extremely fast web applications in Go with minimal effort.

Iris 帮助后端开发人员在 Go 中快速创建速度极快的 web 应用程序。

**Iris** is the only one Go module with **first-class** support to develop web applications following the model-view-controller [MVC](https://en.wikipedia.org/wiki/Model–view–controller) architectural pattern.

Iris 是唯一一个使用模型-视图-控制器 MVC 体系结构模式开发 web 应用程序的一流 Go 模块。

**Iris** can be used as a web port for [gRPC](https://grpc.io/). REST API for gRPC services.

Iris 可用作 gRPC.restapi 的 gRPC 服务的 web 端口。

The source code of Iris is hosted on [GitHub ](https://github.com/kataras/iris)and licensed under the terms of [BSD 3-clause License](https://opensource.org/licenses/BSD-3-Clause), like the [Go project](https://github.com/golang/go) itself.

Iris 的源代码托管在 GitHub 上，并根据 BSD 3-clause License 的条款进行许可，就像 Go 项目本身一样。

> Documentation refers to the [v12.1.8 stable release](https://github.com/kataras/iris/tree/v12.1.8). Want to learn about the upcoming release? Head over to the [HISTORY.md#next](https://github.com/kataras/iris/blob/master/HISTORY.md#next) instead.
>
> 文档指的是 v12.1.8稳定版本。

## Getting Started

Iris is a cross-platform software.

Iris 是一个跨平台的软件。

The only requirement is the [Go Programming Language](https://golang.org/dl/), version 1.14 and above.

唯一的要求是 Go 编程语言，版本1.14及以上。

```
$ go env -w GO111MODULE=on
```

## Install 安装

```
$ go get github.com/kataras/iris/v12@v12.1.8
```

Or edit your project's `go.mod` file.

或者编辑你项目的 go.mod 文件。

```
module your_project_name

go 1.14

require (
    github.com/kataras/iris/v12 v12.1.8
)
```

> ```
> $ go build
> ```

## Troubleshooting 故障排除

If you get a network error during installation please make sure you set a valid [GOPROXY environment variable](https://github.com/golang/go/wiki/Modules#are-there-always-on-module-repositories-and-enterprise-proxies).

如果你在安装过程中出现网络错误，请确保你设置了一个有效的 GOPROXY 环境变量。

```
go env -w GOPROXY=https://goproxy.cn,https://gocenter.io,https://goproxy.io,direct
```

## Quick start

Create an empty file, let's assume its name is `example.go`, then open it and copy-paste the below code.

创建一个空文件，假设它的名称是 example.go，然后打开它并复制粘贴下面的代码。

```
package main

import "github.com/kataras/iris/v12"

func main() {
    app := iris.Default()
    app.Use(myMiddleware)

    app.Handle("GET", "/ping", func(ctx iris.Context) {
        ctx.JSON(iris.Map{"message": "pong"})
    })

    // Listens and serves incoming http requests
    // on http://localhost:8080.
    app.Listen(":8080")
}

func myMiddleware(ctx iris.Context) {
    ctx.Application().Logger().Infof("Runs before %s", ctx.Path())
    ctx.Next()
}
```

Start a terminal session and execute the following.

启动终端会话并执行以下操作。

```
# run example.go and visit http://localhost:8080/ping on browser
$ go run example.go
```

## Show me more! 

Let's take a small overview of how easy is to get up and running.

让我们来简单概述一下起床和跑步是多么容易。

```
package main

import "github.com/kataras/iris/v12"

func main() {
    app := iris.New()
    // Load all templates from the "./views" folder
    // where extension is ".html" and parse them
    // using the standard `html/template` package.
    app.RegisterView(iris.HTML("./views", ".html"))

    // Method:    GET
    // Resource:  http://localhost:8080
    app.Get("/", func(ctx iris.Context) {
        // Bind: {{.message}} with "Hello world!"
        ctx.ViewData("message", "Hello world!")
        // Render template file: ./views/hello.html
        ctx.View("hello.html")
    })

    // Method:    GET
    // Resource:  http://localhost:8080/user/42
    //
    // Need to use a custom regexp instead?
    // Easy;
    // Just mark the parameter's type to 'string'
    // which accepts anything and make use of
    // its `regexp` macro function, i.e:
    // app.Get("/user/{id:string regexp(^[0-9]+$)}")
    app.Get("/user/{id:uint64}", func(ctx iris.Context) {
        userID, _ := ctx.Params().GetUint64("id")
        ctx.Writef("User ID: %d", userID)
    })

    // Start the server using a network address.
    app.Listen(":8080")
}
<!-- file: ./views/hello.html -->
<html>
<head>
    <title>Hello Page</title>
</head>
<body>
    <h1>{{.message}}</h1>
</body>
</html>
```

> Wanna re-start your app automatically when source code changes happens? Install the [iris-cli](https://github.com/kataras/iris-cli) tool and execute `iris-cli run` instead of `go run main.go`.
>
> 想在源代码发生变化时自动重新启动你的应用程序吗？安装 iris-cli 工具并执行 iris-cli run 而不是 go run main.go。

At the next section we will learn more about [Routing](https://github.com/kataras/iris/wiki/Routing).

在下一节中，我们将了解更多关于路由的内容。

# Host

## Automatic Public Address with TLS 使用 TLS 的自动公共地址

Wouldn't be great to test your web application server in a more "real-world environment" like a public, remote, address instead of localhost?

在一个更“真实的环境”中(比如一个公共的、远程的、地址而不是本地主机)测试 web 应用服务器不是很好吗？

There are plenty of third-party tools offering such a feature, but in my opinion, the [ngrok](https://github.com/inconshreveable/ngrok) one is the best among them. It's popular and tested for years, like Iris.

有很多第三方工具提供这样的功能，但在我看来，ngrok 工具是其中最好的。它很受欢迎，并且经过了多年的测试，就像爱丽丝一样。

Iris offers ngrok integration. This feature is simple yet very powerful. It really helps when you want to quickly show your development progress to your colleagues or the project leader at a remote conference.

Iris 提供 ngrok 集成。这个功能简单但是非常强大。当您希望在远程会议上向您的同事或项目负责人快速显示您的开发进度时，这种方法非常有用。

Follow the steps below to, temporarily, convert your local Iris web server to a public one.

按照下面的步骤，暂时将您的本地 Iris web 服务器转换为公共服务器。

1. Go head and 去吧[download ngrok 下载 ngrok](https://ngrok.io/), add it to your $PATH environment variable, ，添加到您的 $PATH 环境变量,
2. Simply pass the 简单地传递`WithTunneling` configurator in your 配置器在你的`app.Run`,
3. You are ready to 你已经准备好了[GO 开始](https://www.facebook.com/iris.framework/photos/a.2420499271295384/3261189020559734/?type=3&theater)!

![tunneling_screenshot](https://user-images.githubusercontent.com/22900943/81442996-42731800-917d-11ea-90da-7d6475a6b365.png)

- `ctx.Application().ConfigurationReadOnly().GetVHost()` returns the public domain value. Rarely useful but it's there for you. Most of the times you use relative url paths instead of absolute(or you should to). 返回公共域值。很少有用，但它就在那里等着你。大多数时候，你使用的是相对 url 路径而不是绝对路径(或者你应该这样做)
- It doesn't matter if ngrok is already running or not, Iris framework is smart enough to use ngrok's Ngrok 是否已经在运行并不重要，Iris 框架非常聪明，可以使用 ngrok’s[web API](https://ngrok.com/docs) to create a tunnel. 来建造一条隧道

Full `Tunneling` configuration:

全隧道配置:

```
app.Listen(":8080", iris.WithConfiguration(
	iris.Configuration{
		Tunneling: iris.TunnelingConfiguration{
			AuthToken:    "my-ngrok-auth-client-token",
			Bin:          "/bin/path/for/ngrok",
			Region:       "eu",
			WebInterface: "127.0.0.1:4040",
			Tunnels: []iris.Tunnel{
				{
					Name: "MyApp",
					Addr: ":8080",
				},
			},
		},
}))
```

## Configuration配置

At the [previous](https://github.com/kataras/iris/wiki/Host) section we've learnt about the first input argument of the `app.Run`, here we will take a look of what the second one is.

在前面的部分，我们已经了解了应用程序的第一个输入参数。跑，这里我们来看看第二个是什么。

Let's start from basics. The `iris.New` function returns an `iris.Application`. This Application value can be configured through its `Configure(...iris.Configurator)` and `Run` methods.

让我们从基础开始。虹膜。新函数返回一个虹膜。申请。这个应用程序值可以通过它的 Configure (... iris)进行配置。配置器)和运行方法。

The second optional, varadiac argument of `app.Run/Listen` method accepts one or more `iris.Configurator`. An `iris.Configurator` is just a type of: `func(app *iris.Application)`. Custom `iris.Configurator` can be passed to modify yours `*iris.Application` as well.

应用程序的第二个可选的 varadiac 参数。Run/Listen 方法接受一个或多个虹膜。配置器。虹膜。Configurator 只是一种类型: func (app * iris)。申请)。自定义虹膜。可以通过配置程序修改你的 * 虹膜。也可以申请。

There are built-in`iris.Configurator`s for each of the core [Configuration](https://godoc.org/github.com/kataras/iris#Configuration)'s fields, such as `iris.WithoutStartupLog`, `iris.WithCharset("UTF-8")`, `iris.WithOptimizations` and `iris.WithConfiguration(iris.Configuration{...})` functions.

有嵌入式系统。每个核心配置字段的配置器，比如 iris。没有启动日志，爱丽丝。用 charset (“ UTF-8”)、 iris。优化和虹膜。用配置(iris。配置{ ... })函数。

Each "module" like the iris view engine, websockets, sessions and each middleware has its own configurations and options.

每个“模块”，如虹膜视图引擎，websockets，会话和每个中间件都有自己的配置和选项。

## Using the 使用[Configuration 配置](https://godoc.org/github.com/kataras/iris#Configuration)

All of the `iris.Configuration` fields are defaulted to the most common use cases. If you want to make use of a custom `iris.Configurator` at any point call the `app.Configure(accepts iris.Configurator)` method to pass your configurator(s) there.

所有的虹膜。配置字段默认为最常见的用例。如果你想使用自定义虹膜。配置器在任何时候都可以调用应用程序。配置(接受虹膜。方法来传递你的配置器到那里。

```
config := iris.WithConfiguration(iris.Configuration {
  DisableStartupLog: true,
  Optimizations: true,
  Charset: "UTF-8",
})

app.Listen(":8080", config)
```

List of all available settings:

所有可用设置列表:

```
// Tunnel is the Tunnels field of the TunnelingConfiguration structure.
type Tunnel struct {
	// Name is the only one required field,
	// it is used to create and close tunnels, e.g. "MyApp".
	// If this field is not empty then ngrok tunnels will be created
	// when the iris app is up and running.
	Name string `json:"name" yaml:"Name" toml:"Name"`
	// Addr is basically optionally as it will be set through
	// Iris built-in Runners, however, if `iris.Raw` is used
	// then this field should be set of form 'hostname:port'
	// because framework cannot be aware
	// of the address you used to run the server on this custom runner.
	Addr string `json:"addr,omitempty" yaml:"Addr" toml:"Addr"`
}

// TunnelingConfiguration contains configuration
// for the optional tunneling through ngrok feature.
// Note that the ngrok should be already installed at the host machine.
type TunnelingConfiguration struct {
	// AuthToken field is optionally and can be used
	// to authenticate the ngrok access.
	// ngrok authtoken <YOUR_AUTHTOKEN>
	AuthToken string `json:"authToken,omitempty" yaml:"AuthToken" toml:"AuthToken"`

	// No...
	// Config is optionally and can be used
	// to load ngrok configuration from file system path.
	//
	// If you don't specify a location for a configuration file,
	// ngrok tries to read one from the default location $HOME/.ngrok2/ngrok.yml.
	// The configuration file is optional; no error is emitted if that path does not exist.
	// Config string `json:"config,omitempty" yaml:"Config" toml:"Config"`

	// Bin is the system binary path of the ngrok executable file.
	// If it's empty then the framework will try to find it through system env variables.
	Bin string `json:"bin,omitempty" yaml:"Bin" toml:"Bin"`

	// WebUIAddr is the web interface address of an already-running ngrok instance.
	// Iris will try to fetch the default web interface address(http://127.0.0.1:4040)
	// to determinate if a ngrok instance is running before try to start it manually.
	// However if a custom web interface address is used,
	// this field must be set e.g. http://127.0.0.1:5050.
	WebInterface string `json:"webInterface,omitempty" yaml:"WebInterface" toml:"WebInterface"`

	// Region is optionally, can be used to set the region which defaults to "us".
	// Available values are:
	// "us" for United States
	// "eu" for Europe
	// "ap" for Asia/Pacific
	// "au" for Australia
	// "sa" for South America
	// "jp" forJapan
	// "in" for India
	Region string `json:"region,omitempty" yaml:"Region" toml:"Region"`

	// Tunnels the collection of the tunnels.
	// One tunnel per Iris Host per Application, usually you only need one.
	Tunnels []Tunnel `json:"tunnels" yaml:"Tunnels" toml:"Tunnels"`
}

// Configuration holds the necessary settings for an Iris Application instance.
// All fields are optionally, the default values will work for a common web application.
//
// A Configuration value can be passed through `WithConfiguration` Configurator.
// Usage:
// conf := iris.Configuration{ ... }
// app := iris.New()
// app.Configure(iris.WithConfiguration(conf)) OR
// app.Run/Listen(..., iris.WithConfiguration(conf)).
type Configuration struct {
	// LogLevel is the log level the application should use to output messages.
	// Logger, by default, is mostly used on Build state but it is also possible
	// that debug error messages could be thrown when the app is running, e.g.
	// when malformed data structures try to be sent on Client (i.e Context.JSON/JSONP/XML...).
	//
	// Defaults to "info". Possible values are:
	// * "disable"
	// * "fatal"
	// * "error"
	// * "warn"
	// * "info"
	// * "debug"
	LogLevel string `json:"logLevel" yaml:"LogLevel" toml:"LogLevel" env:"LOG_LEVEL"`

	// Tunneling can be optionally set to enable ngrok http(s) tunneling for this Iris app instance.
	// See the `WithTunneling` Configurator too.
	Tunneling TunnelingConfiguration `json:"tunneling,omitempty" yaml:"Tunneling" toml:"Tunneling"`

	// IgnoreServerErrors will cause to ignore the matched "errors"
	// from the main application's `Run` function.
	// This is a slice of string, not a slice of error
	// users can register these errors using yaml or toml configuration file
	// like the rest of the configuration fields.
	//
	// See `WithoutServerError(...)` function too.
	//
	// Example: https://github.com/kataras/iris/tree/master/_examples/http-server/listen-addr/omit-server-errors
	//
	// Defaults to an empty slice.
	IgnoreServerErrors []string `json:"ignoreServerErrors,omitempty" yaml:"IgnoreServerErrors" toml:"IgnoreServerErrors"`

	// DisableStartupLog if set to true then it turns off the write banner on server startup.
	//
	// Defaults to false.
	DisableStartupLog bool `json:"disableStartupLog,omitempty" yaml:"DisableStartupLog" toml:"DisableStartupLog"`
	// DisableInterruptHandler if set to true then it disables the automatic graceful server shutdown
	// when control/cmd+C pressed.
	// Turn this to true if you're planning to handle this by your own via a custom host.Task.
	//
	// Defaults to false.
	DisableInterruptHandler bool `json:"disableInterruptHandler,omitempty" yaml:"DisableInterruptHandler" toml:"DisableInterruptHandler"`

	// DisablePathCorrection disables the correcting
	// and redirecting or executing directly the handler of
	// the requested path to the registered path
	// for example, if /home/ path is requested but no handler for this Route found,
	// then the Router checks if /home handler exists, if yes,
	// (permanent)redirects the client to the correct path /home.
	//
	// See `DisablePathCorrectionRedirection` to enable direct handler execution instead of redirection.
	//
	// Defaults to false.
	DisablePathCorrection bool `json:"disablePathCorrection,omitempty" yaml:"DisablePathCorrection" toml:"DisablePathCorrection"`
	// DisablePathCorrectionRedirection works whenever configuration.DisablePathCorrection is set to false
	// and if DisablePathCorrectionRedirection set to true then it will fire the handler of the matching route without
	// the trailing slash ("/") instead of send a redirection status.
	//
	// Defaults to false.
	DisablePathCorrectionRedirection bool `json:"disablePathCorrectionRedirection,omitempty" yaml:"DisablePathCorrectionRedirection" toml:"DisablePathCorrectionRedirection"`
	// EnablePathIntelligence if set to true,
	// the router will redirect HTTP "GET" not found pages to the most closest one path(if any). For example
	// you register a route at "/contact" path -
	// a client tries to reach it by "/cont", the path will be automatic fixed
	// and the client will be redirected to the "/contact" path
	// instead of getting a 404 not found response back.
	//
	// Defaults to false.
	EnablePathIntelligence bool `json:"enablePathIntelligence,omitempty" yaml:"EnablePathIntelligence" toml:"EnablePathIntelligence"`
	// EnablePathEscape when is true then its escapes the path and the named parameters (if any).
	// When do you need to Disable(false) it:
	// accepts parameters with slash '/'
	// Request: http://localhost:8080/details/Project%2FDelta
	// ctx.Param("project") returns the raw named parameter: Project%2FDelta
	// which you can escape it manually with net/url:
	// projectName, _ := url.QueryUnescape(c.Param("project").
	//
	// Defaults to false.
	EnablePathEscape bool `json:"enablePathEscape,omitempty" yaml:"EnablePathEscape" toml:"EnablePathEscape"`
	// ForceLowercaseRouting if enabled, converts all registered routes paths to lowercase
	// and it does lowercase the request path too for matching.
	//
	// Defaults to false.
	ForceLowercaseRouting bool `json:"forceLowercaseRouting,omitempty" yaml:"ForceLowercaseRouting" toml:"ForceLowercaseRouting"`
	// FireMethodNotAllowed if it's true router checks for StatusMethodNotAllowed(405) and
	//  fires the 405 error instead of 404
	// Defaults to false.
	FireMethodNotAllowed bool `json:"fireMethodNotAllowed,omitempty" yaml:"FireMethodNotAllowed" toml:"FireMethodNotAllowed"`
	// DisableAutoFireStatusCode if true then it turns off the http error status code
	// handler automatic execution on error code from a `Context.StatusCode` call.
	// By-default a custom http error handler will be fired when "Context.StatusCode(errorCode)" called.
	//
	// Defaults to false.
	DisableAutoFireStatusCode bool `json:"disableAutoFireStatusCode,omitempty" yaml:"DisableAutoFireStatusCode" toml:"DisableAutoFireStatusCode"`
	// ResetOnFireErrorCode if true then any previously response body or headers through
	// response recorder or gzip writer will be ignored and the router
	// will fire the registered (or default) HTTP error handler instead.
	// See `core/router/handler#FireErrorCode` and `Context.EndRequest` for more details.
	//
	// Read more at: https://github.com/kataras/iris/issues/1531
	//
	// Defaults to false.
	ResetOnFireErrorCode bool `json:"resetOnFireErrorCode,omitempty" yaml:"ResetOnFireErrorCode" toml:"ResetOnFireErrorCode"`

	// EnableOptimization when this field is true
	// then the application tries to optimize for the best performance where is possible.
	//
	// Defaults to false.
	EnableOptimizations bool `json:"enableOptimizations,omitempty" yaml:"EnableOptimizations" toml:"EnableOptimizations"`
	// DisableBodyConsumptionOnUnmarshal manages the reading behavior of the context's body readers/binders.
	// If set to true then it
	// disables the body consumption by the `context.UnmarshalBody/ReadJSON/ReadXML`.
	//
	// By-default io.ReadAll` is used to read the body from the `context.Request.Body which is an `io.ReadCloser`,
	// if this field set to true then a new buffer will be created to read from and the request body.
	// The body will not be changed and existing data before the
	// context.UnmarshalBody/ReadJSON/ReadXML will be not consumed.
	DisableBodyConsumptionOnUnmarshal bool `json:"disableBodyConsumptionOnUnmarshal,omitempty" yaml:"DisableBodyConsumptionOnUnmarshal" toml:"DisableBodyConsumptionOnUnmarshal"`
	// FireEmptyFormError returns if set to tue true then the `context.ReadBody/ReadForm`
	// will return an `iris.ErrEmptyForm` on empty request form data.
	FireEmptyFormError bool `json:"fireEmptyFormError,omitempty" yaml:"FireEmptyFormError" yaml:"FireEmptyFormError"`

	// TimeFormat time format for any kind of datetime parsing
	// Defaults to  "Mon, 02 Jan 2006 15:04:05 GMT".
	TimeFormat string `json:"timeFormat,omitempty" yaml:"TimeFormat" toml:"TimeFormat"`

	// Charset character encoding for various rendering
	// used for templates and the rest of the responses
	// Defaults to "utf-8".
	Charset string `json:"charset,omitempty" yaml:"Charset" toml:"Charset"`

	// PostMaxMemory sets the maximum post data size
	// that a client can send to the server, this differs
	// from the overral request body size which can be modified
	// by the `context#SetMaxRequestBodySize` or `iris#LimitRequestBodySize`.
	//
	// Defaults to 32MB or 32 << 20 if you prefer.
	PostMaxMemory int64 `json:"postMaxMemory" yaml:"PostMaxMemory" toml:"PostMaxMemory"`
	//  +----------------------------------------------------+
	//  | Context's keys for values used on various featuers |
	//  +----------------------------------------------------+

	// Context values' keys for various features.
	//
	// LocaleContextKey is used by i18n to get the current request's locale, which contains a translate function too.
	//
	// Defaults to "iris.locale".
	LocaleContextKey string `json:"localeContextKey,omitempty" yaml:"LocaleContextKey" toml:"LocaleContextKey"`
	// LanguageContextKey is the context key which a language can be modified by a middleware.
	// It has the highest priority over the rest and if it is empty then it is ignored,
	// if it set to a static string of "default" or to the default language's code
	// then the rest of the language extractors will not be called at all and
	// the default language will be set instead.
	//
	// Use with `Context.SetLanguage("el-GR")`.
	//
	// See `i18n.ExtractFunc` for a more organised way of the same feature.
	// Defaults to "iris.locale.language".
	LanguageContextKey string `json:"languageContextKey,omitempty" yaml:"LanguageContextKey" toml:"LanguageContextKey"`
	// VersionContextKey is the context key which an API Version can be modified
	// via a middleware through `SetVersion` method, e.g. `ctx.SetVersion("1.0, 1.1")`.
	// Defaults to "iris.api.version".
	VersionContextKey string `json:"versionContextKey" yaml:"VersionContextKey" toml:"VersionContextKey"`
	// GetViewLayoutContextKey is the key of the context's user values' key
	// which is being used to set the template
	// layout from a middleware or the main handler.
	// Overrides the parent's or the configuration's.
	//
	// Defaults to "iris.ViewLayout"
	ViewLayoutContextKey string `json:"viewLayoutContextKey,omitempty" yaml:"ViewLayoutContextKey" toml:"ViewLayoutContextKey"`
	// GetViewDataContextKey is the key of the context's user values' key
	// which is being used to set the template
	// binding data from a middleware or the main handler.
	//
	// Defaults to "iris.viewData"
	ViewDataContextKey string `json:"viewDataContextKey,omitempty" yaml:"ViewDataContextKey" toml:"ViewDataContextKey"`
	// RemoteAddrHeaders are the allowed request headers names
	// that can be valid to parse the client's IP based on.
	// By-default no "X-" header is consired safe to be used for retrieving the
	// client's IP address, because those headers can manually change by
	// the client. But sometimes are useful e.g., when behind a proxy
	// you want to enable the "X-Forwarded-For" or when cloudflare
	// you want to enable the "CF-Connecting-IP", inneed you
	// can allow the `ctx.RemoteAddr()` to use any header
	// that the client may sent.
	//
	// Defaults to an empty map but an example usage is:
	// RemoteAddrHeaders {
	//	"X-Real-Ip":             true,
	//  "X-Forwarded-For":       true,
	// 	"CF-Connecting-IP": 	 true,
	//	}
	//
	// Look `context.RemoteAddr()` for more.
	RemoteAddrHeaders map[string]bool `json:"remoteAddrHeaders,omitempty" yaml:"RemoteAddrHeaders" toml:"RemoteAddrHeaders"`
	// RemoteAddrPrivateSubnets defines the private sub-networks.
	// They are used to be compared against
	// IP Addresses fetched through `RemoteAddrHeaders` or `Context.Request.RemoteAddr`.
	// For details please navigate through: https://github.com/kataras/iris/issues/1453
	// Defaults to:
	// {
	// 	Start: net.ParseIP("10.0.0.0"),
	// 	End:   net.ParseIP("10.255.255.255"),
	// },
	// {
	// 	Start: net.ParseIP("100.64.0.0"),
	// 	End:   net.ParseIP("100.127.255.255"),
	// },
	// {
	// 	Start: net.ParseIP("172.16.0.0"),
	// 	End:   net.ParseIP("172.31.255.255"),
	// },
	// {
	// 	Start: net.ParseIP("192.0.0.0"),
	// 	End:   net.ParseIP("192.0.0.255"),
	// },
	// {
	// 	Start: net.ParseIP("192.168.0.0"),
	// 	End:   net.ParseIP("192.168.255.255"),
	// },
	// {
	// 	Start: net.ParseIP("198.18.0.0"),
	// 	End:   net.ParseIP("198.19.255.255"),
	// }
	//
	// Look `Context.RemoteAddr()` for more.
	RemoteAddrPrivateSubnets []netutil.IPRange `json:"remoteAddrPrivateSubnets" yaml:"RemoteAddrPrivateSubnets" toml:"RemoteAddrPrivateSubnets"`
	// SSLProxyHeaders defines the set of header key values
	// that would indicate a valid https Request (look `Context.IsSSL()`).
	// Example: `map[string]string{"X-Forwarded-Proto": "https"}`.
	//
	// Defaults to empty map.
	SSLProxyHeaders map[string]string `json:"sslProxyHeaders" yaml:"SSLProxyHeaders" toml:"SSLProxyHeaders"`
	// HostProxyHeaders defines the set of headers that may hold a proxied hostname value for the clients.
	// Look `Context.Host()` for more.
	// Defaults to empty map.
	HostProxyHeaders map[string]bool `json:"hostProxyHeaders" yaml:"HostProxyHeaders" toml:"HostProxyHeaders"`
	// Other are the custom, dynamic options, can be empty.
	// This field used only by you to set any app's options you want.
	//
	// Defaults to empty map.
	Other map[string]interface{} `json:"other,omitempty" yaml:"Other" toml:"Other"`
}
```

### Load from 载入[YAML 男名男子名](https://yaml.org/)

Using the `iris.YAML("path")`.

使用 iris.YAML (“ path”)。

File: **iris.yml**

文件: iris.yml

```
FireMethodNotAllowed: true
DisableBodyConsumptionOnUnmarshal: true
TimeFormat: Mon, 01 Jan 2006 15:04:05 GMT
Charset: UTF-8
```

File: **main.go**

文件: main.go

```
config := iris.WithConfiguration(iris.YAML("./iris.yml"))
app.Listen(":8080", config)
```

### Load from 载入[TOML 汤姆林](https://github.com/toml-lang/toml)

Using the `iris.TOML("path")`.

使用 iris.TOML (“ path”)。

File: **iris.tml**

文件: iris.tml

```
FireMethodNotAllowed = true
DisableBodyConsumptionOnUnmarshal = false
TimeFormat = "Mon, 01 Jan 2006 15:04:05 GMT"
Charset = "UTF-8"

[Other]
    ServerName = "my fancy iris server"
    ServerOwner = "admin@example.com"
```

File: **main.go**

文件: main.go

```
config := iris.WithConfiguration(iris.TOML("./iris.tml"))
app.Listen(":8080", config)
```

## Using the functional way 使用函数式方法

As we already mention, you can pass any number of `iris.Configurator` in the `app.Run/Listen`’s second argument. Iris provides an option for each of its `iris.Configuration`’s fields.

正如我们已经提到的，你可以通过任何数量的虹膜。应用程序中的配置器。Run/Listen 的第二个参数。Iris 为它的每个虹膜提供了一个选项。配置的字段。

```
app.Listen(":8080", iris.WithoutInterruptHandler,
    iris.WithoutBodyConsumptionOnUnmarshal,
    iris.WithoutAutoFireStatusCode,
    iris.WithLowercaseRouting,
    iris.WithPathIntelligence,
    iris.WithOptimizations,
    iris.WithTimeFormat("Mon, 01 Jan 2006 15:04:05 GMT"),
)
```

Good when you want to change some of the configuration's field. Prefix: "With" or "Without", code editors will help you navigate through all configuration options without even a glitch to the documentation.

当您想要更改一些配置的字段时，这是很好的。前缀: “带”或“不带”，代码编辑器将帮助您浏览所有配置选项，甚至没有任何故障的文档。

List of functional options:

功能选项清单:

```
var (
    WithGlobalConfiguration Configurator
    WithoutStartupLog, WithoutBanner Configurator
    WithoutInterruptHandler Configurator
    WithoutPathCorrection Configurator
    WithPathIntelligence Configurator
    WithoutPathCorrectionRedirection Configurator
    WithoutBodyConsumptionOnUnmarshal Configurator
    WithEmptyFormError Configurator
    WithPathEscape Configurator
    WithLowercaseRouting Configurator
    WithOptimizations Configurator
    WithFireMethodNotAllowed Configurator
    WithoutAutoFireStatusCode Configurator
    WithResetOnFireErrorCode Configurator
    WithTunneling Configurator
)

func WithLogLevel(level string) Configurator
func WithoutServerError(errors ...error) Configurator
func WithTimeFormat(timeformat string) Configurator
func WithCharset(charset string) Configurator
func WithPostMaxMemory(limit int64) Configurator
func WithRemoteAddrHeader(header ...string) Configurator
func WithoutRemoteAddrHeader(headerName string) Configurator
func WithRemoteAddrPrivateSubnet(startIP, endIP string) Configurator
func WithSSLProxyHeader(headerKey, headerValue string) Configurator
func WithHostProxyHeader(headers ...string) Configurator
func WithOtherValue(key string, val interface{}) Configurator
func WithSitemap(startURL string) Configurator
```

## Custom values 自定义值

The `iris.Configuration` contains a field named `Other map[string]interface{}` which accepts any custom `key:value` option, therefore you can use that field to pass specific values that your app expects based on the custom requirements.

虹膜。配置包含一个名为 Other map [ string ] interface {}的字段，该字段接受任何自定义 key: value 选项，因此您可以使用该字段传递应用程序基于自定义需求预期的特定值。

```
app.Listen(":8080", 
    iris.WithOtherValue("ServerName", "my amazing iris server"),
    iris.WithOtherValue("ServerOwner", "admin@example.com"),
)
```

You can access those fields via `app.ConfigurationReadOnly`.

你可以通过 app.ConfigurationReadOnly 访问这些字段。

## Access Configuration from Context 来自上下文的访问配置

```
func (ctx iris.Context) {
    cfg := app.Application().ConfigurationReadOnly().GetOther()
    srvName := cfg["MyServerName"]
    srvOwner := cfg["ServerOwner"]
}
```

