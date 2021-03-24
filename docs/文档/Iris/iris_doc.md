## 安装

Installation 

Iris is a cross-platform software.
Iris 是一个跨平台的软件。

The only requirement is the Go Programming Language, version 1.14 and above.
唯一的要求是 Go 编程语言，版本1.14及以上。

```bash
$ mkdir myapp
$ cd myapp
$ go mod init myapp
$ go get github.com/kataras/iris/v12@master
```

Import it in your code:
在你的代码中导入:

```import "github.com/kataras/iris/v12"```

#### 故障排除

Troubleshooting 

If you get a network error during installation please make sure you set a valid GOPROXY environment variable.
如果你在安装过程中出现网络错误，请确保你设置了一个有效的 GOPROXY 环境变量。

```go env -w GOPROXY=https://goproxy.cn,https://gocenter.io,https://goproxy.io,direct```

Perform a clean of your go modules cache if none of the above worked:
如果上面的模块都不起作用，那么清理一下你的 go 模块缓存:

```go clean --modcache```

## 快速启动

Quick start 

```
# assume the following codes in main.go file
$ cat main.go
```

```go
package main

import "github.com/kataras/iris/v12"

func main() {
    app := iris.New()

    booksAPI := app.Party("/books")
    {
        booksAPI.Use(iris.Compression)
    
        // GET: http://localhost:8080/books
        booksAPI.Get("/", list)
        // POST: http://localhost:8080/books
        booksAPI.Post("/", create)
    }
    
    app.Listen(":8080")
}

// Book example.
type Book struct {
    Title string `json:"title"`
}

func list(ctx iris.Context) {
    books := []Book{
        {"Mastering Concurrency in Go"},
        {"Go Design Patterns"},
        {"Black Hat Go"},
    }

    ctx.JSON(books)
    // TIP: negotiate the response between server's prioritizes
    // and client's requirements, instead of ctx.JSON:
    // ctx.Negotiation().JSON().MsgPack().Protobuf()
    // ctx.Negotiate(books)
}

func create(ctx iris.Context) {
    var b Book
    err := ctx.ReadJSON(&b)
    // TIP: use ctx.ReadBody(&b) to bind
    // any type of incoming data instead.
    if err != nil {
        ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
            Title("Book creation failure").DetailErr(err))
        // TIP: use ctx.StopWithError(code, err) when only
        // plain text responses are expected on errors.
        return
    }

    println("Received Book: " + b.Title)
    
    ctx.StatusCode(iris.StatusCreated)
}
```

MVC equivalent:
相当于 MVC:

```go
import "github.com/kataras/iris/v12/mvc"
```

```go
m := mvc.New(booksAPI)
m.Handle(new(BookController))
```

```go
type BookController struct {
    /* dependencies */
}

// GET: http://localhost:8080/books
func (c *BookController) Get() []Book {
    return []Book{
        {"Mastering Concurrency in Go"},
        {"Go Design Patterns"},
        {"Black Hat Go"},
    }
}

// POST: http://localhost:8080/books
func (c *BookController) Post(b Book) int {
    println("Received Book: " + b.Title)

    return iris.StatusCreated
}
```

Run your Iris web server:
运行你的 Iris 网络服务器:

```
$ go run main.go
> Now listening on: http://localhost:8080
> Application started. Press CTRL+C to shut down.
```

List Books:
书单:

```bash
$ curl --header 'Accept-Encoding:gzip' http://localhost:8080/books

[
  {
    "title": "Mastering Concurrency in Go"
  },
  {
    "title": "Go Design Patterns"
  },
  {
    "title": "Black Hat Go"
  }
]
```

Create a new Book:
创建一本新书:

```bash
$ curl -i -X POST \
--header 'Content-Encoding:gzip' \
--header 'Content-Type:application/json' \
--data "{\"title\":\"Writing An Interpreter In Go\"}" \
http://localhost:8080/books

> HTTP/1.1 201 Created
```

That's how an error response looks like:
这就是错误响应的样子:

```bash
$ curl -X POST --data "{\"title\" \"not valid one\"}" \
http://localhost:8080/books

> HTTP/1.1 400 Bad Request

{
  "status": 400,
  "title": "Book creation failure"
  "detail": "invalid character '\"' after object key",
}
```

run [in the browser](https://replit.com/@kataras/Iris-Hello-World#main.go)

Benchmarks 
基准

Iris uses a custom version of muxie.
Iris 使用自定义版本的 muxie。

See all benchmarks
参见所有基准测试

📖 Fires 200000 requests with a dynamic parameter of int, sends JSON as request body and receives JSON as response.
触发200000个动态参数为 int 的请求，发送 JSON 作为请求主体，接收 JSON 作为响应。
								延迟			吞吐量

| Name    | Language   | Reqs/sec | Latency  | Throughput | Time To Complete |
| ------- | ---------- | -------- | -------- | ---------- | ---------------- |
| Iris    | Go         | 150430   | 826.05us | 41.25MB    | 1.33s            |
| Chi     | Go         | 146274   | 0.85ms   | 39.32MB    | 1.37s            |
| Gin     | Go         | 141664   | 0.88ms   | 38.74MB    | 1.41s            |
| Echo    | Go         | 138915   | 0.90ms   | 38.15MB    | 1.44s            |
| Kestrel | C#         | 136935   | 0.91ms   | 39.79MB    | 1.47s            |
| Martini | Go         | 128590   | 0.97ms   | 34.57MB    | 1.56s            |
| Buffalo | Go         | 58954    | 2.12ms   | 16.18MB    | 3.40s            |
| Koa     | Javascript | 50948    | 2.61ms   | 14.15MB    | 4.19s            |
| Express | Javascript | 38451    | 3.24ms   | 13.77MB    | 5.21s            |


## API 示例

API Examples 


You can find a number of ready-to-run examples at Iris examples repository.
您可以在 Iris 示例存储库中找到许多可以直接运行的示例。

## 使用 GET，POST，PUT，PATCH，DELETE 和 OPTIONS

Using GET, POST, PUT, PATCH, DELETE and OPTIONS


```go
func main() {
    // Creates an iris application with default middleware:
    // Default with "debug" Logger Level.
    // Localization enabled on "./locales" directory
    // and HTML templates on "./views" or "./templates" directory.
    // It runs with the AccessLog on "./access.log",
    // Recovery (crash-free) and Request ID middleware already attached.
    app := iris.Default()

    app.Get("/someGet", getting)
    app.Post("/somePost", posting)
    app.Put("/somePut", putting)
    app.Delete("/someDelete", deleting)
    app.Patch("/somePatch", patching)
    app.Header("/someHead", head)
    app.Options("/someOptions", options)
    
    app.Listen(":8080")
}
```

## 路径中的参数

Parameters in path 

```go
func main() {
    app := iris.Default()

    // This handler will match /user/john but will not match /user/ or /user
    app.Get("/user/{name}", func(ctx iris.Context) {
        name := ctx.Params().Get("name")
        ctx.Writef("Hello %s", name)
    })
    
    // However, this one will match /user/john/ and also /user/john/send
    // If no other routers match /user/john, it will redirect to /user/john/
    app.Get("/user/{name}/{action:path}", func(ctx iris.Context) {
        name := ctx.Params().Get("name")
        action := ctx.Params().Get("action")
        message := name + " is " + action
        ctx.WriteString(message)
    })
    
    // For each matched request Context will hold the route definition
    app.Post("/user/{name:string}/{action:path}", func(ctx iris.Context) {
        ctx.GetCurrentRoute().Tmpl().Src == "/user/{name:string}/{action:path}" // true
    })
    
    app.Listen(":8080")
}
```

Builtin available parameter types:
内置的可用参数类型:

| Param Type    | Go Type  | Validation 验证                                              | Retrieve Helper 检索助手 |
| ------------- | -------- | ------------------------------------------------------------ | ------------------------ |
| :string       | string   | anything (single path segment)任何内容(单路径段)             | Params().Get             |
| :uuid         | string   | uuidv4 or v1 (single path segment) Uuuidv4或 v1(单路径段)    | Params().Get             |
| :int          | int      | -9223372036854775808 to 9223372036854775807 (x64) or -2147483648 to 2147483647 (x32), depends on the host arch | Params().GetInt          |
| :int8         | int8     | -128 to 127                                                  | Params().GetInt8         |
| :int16        | int16    | -32768 to 32767                                              | Params().GetInt16        |
| :int32        | int32    | -2147483648 to 2147483647                                    | Params().GetInt32        |
| :int64        | int64 64 | -9223372036854775808 to 9223372036854775807                  | Params().GetInt64        |
| :uint         | uint     | 0 to 18446744073709551615 (x64) or 0 to 4294967295 (x32), depends on the host arch | Params().GetUint         |
| :uint8        | uint8    | 0 to 255                                                     | Params().GetUint8        |
| :uint16       | uint16   | 0 to 65535                                                   | Params().GetUint16       |
| :uint32       | uint32   | 0 to 4294967295                                              | Params().GetUint32       |
| :uint64       | uint64   | 0 to 18446744073709551615                                    | Params().GetUint64       |
| :bool         | bool     | "1" or "t" or "T" or "TRUE" or "true" or "True" or "0" or "f" or "F" or "FALSE" or "false" or "False" | Params().GetBool         |
| :alphabetical | string   | lowercase or uppercase letters                               | Params().Get             |
| :file         | string   | lowercase or uppercase letters, numbers, underscore (_), dash (-), point (.) and no spaces or other special characters that are not valid for filenames | Params().Get             |
| :path         | string   | anything, can be separated by slashes (path segments) but should be the last part of the route path | Params().Get             |


More examples can be found at: [_examples/routing](https://github.com/kataras/iris/tree/master/_examples/routing).
更多的例子可以在: _ examples/routing 中找到。

## 查询字符串参数

Querystring parameters 


```go
func main() {
    app := iris.Default()

    // Query string parameters are parsed using the existing underlying request object.
    // The request responds to a url matching:  /welcome?firstname=Jane&lastname=Doe
    app.Get("/welcome", func(ctx iris.Context) {
        firstname := ctx.URLParamDefault("firstname", "Guest")
        lastname := ctx.URLParam("lastname") // shortcut for ctx.Request().URL.Query().Get("lastname")
    
        ctx.Writef("Hello %s %s", firstname, lastname)
    })
    app.Listen(":8080")
}
```

## 多部分/字符编码形式

Multipart/Urlencoded Form 


```go
func main() {
    app := iris.Default()

    app.Post("/form_post", func(ctx iris.Context) {
        message := ctx.PostValue("message")
        nick := ctx.PostValueDefault("nick", "anonymous")
    
        ctx.JSON(iris.Map{
            "status":  "posted",
            "message": message,
            "nick":    nick,
        })
    })
    app.Listen(":8080")
}
```

## 另一个例子: query + post form

Another example: query + post form


```bash
POST /post?id=1234&page=1 HTTP/1.1
Content-Type: application/x-www-form-urlencoded

name=kataras&message=this_is_great
```

```go
func main() {
    app := iris.Default()

    app.Post("/post", func(ctx iris.Context) {
        id, err := ctx.URLParamInt("id", 0)
        if err != nil {
            ctx.StopWithError(iris.StatusBadRequest, err)
            return
        }
    
        page := ctx.URLParamIntDefault("page", 0)
        name := ctx.PostValue("name")
        message := ctx.PostValue("message")
    
        ctx.Writef("id: %d; page: %d; name: %s; message: %s", id, page, name, message)
    })
    app.Listen(":8080")
}
```

```id: 1234; page: 1; name: kataras; message: this_is_great```

## 查询和张贴表单参数

Query and post form parameters 

```bash
POST /post?id=a&id=b&id=c&name=john&name=doe&name=kataras
Content-Type: application/x-www-form-urlencoded
```

```go
func main() {
    app := iris.Default()

    app.Post("/post", func(ctx iris.Context) {
    
        ids := ctx.URLParamSlice("id")
        names, err := ctx.PostValues("name")
        if err != nil {
            ctx.StopWithError(iris.StatusBadRequest, err)
            return
        }
    
        ctx.Writef("ids: %v; names: %v", ids, names)
    })
    app.Listen(":8080")
}
```

``` ids: [a b c], names: [john doe kataras]```


## 上传文件

Upload files 

### 单文件

Single file 

```go
const maxSize = 8 * iris.MB

func main() {
    app := iris.Default()

    app.Post("/upload", func(ctx iris.Context) {
        // Set a lower memory limit for multipart forms (default is 32 MiB)
        ctx.SetMaxRequestBodySize(maxSize)
        // OR
        // app.Use(iris.LimitRequestBodySize(maxSize))
        // OR
        // OR iris.WithPostMaxMemory(maxSize)
    
        // single file
        file, fileHeader, err:= ctx.FormFile("file")
        if err != nil {
            ctx.StopWithError(iris.StatusBadRequest, err)
            return
        }
    
        // Upload the file to specific destination.
        dest := filepath.Join("./uploads", fileHeader.Filename)
        ctx.SaveFormFile(file, dest)
    
        ctx.Writef("File: %s uploaded!", fileHeader.Filename)
    })
    
    app.Listen(":8080")
}
```

How to curl:
如何curl:

```bash
curl -X POST http://localhost:8080/upload \
  -F "file=@/Users/kataras/test.zip" \
  -H "Content-Type: multipart/form-data"
```

### 多文件

Multiple files

See the detail [example code](https://github.com/kataras/iris/tree/master/_examples/file-server/upload-files).
请参阅详细的示例代码。

```go
func main() {
    app := iris.Default()
    app.Post("/upload", func(ctx iris.Context) {
        files, n, err := ctx.UploadFormFiles("./uploads")
        if err != nil {
            ctx.StopWithStatus(iris.StatusInternalServerError)
            return
        }

        ctx.Writef("%d files of %d total size uploaded!", len(files), n))
    })
    
    app.Listen(":8080", iris.WithPostMaxMemory(8 * iris.MB))
}
```

How to curl:

```bash
curl -X POST http://localhost:8080/upload \
  -F "upload[]=@/Users/kataras/test1.zip" \
  -F "upload[]=@/Users/kataras/test2.zip" \
  -H "Content-Type: multipart/form-data"
```

## 路由分组

Grouping routes 

```go
func main() {
    app := iris.Default()

    // Simple group: v1
    v1 := app.Party("/v1")
    {
        v1.Post("/login", loginEndpoint)
        v1.Post("/submit", submitEndpoint)
        v1.Post("/read", readEndpoint)
    }
    
    // Simple group: v2
    v2 := app.Party("/v2")
    {
        v2.Post("/login", loginEndpoint)
        v2.Post("/submit", submitEndpoint)
        v2.Post("/read", readEndpoint)
    }
    
    app.Listen(":8080")
}
```

## 缺省情况下没有中间件的空白 Iris

Blank Iris without middleware by default 


Use
使用

```go
app := iris.New()
```

instead of
而不是

```go
// Default with "debug" Logger Level.
// Localization enabled on "./locales" directory
// and HTML templates on "./views" or "./templates" directory.
// It runs with the AccessLog on "./access.log",
// Recovery and Request ID middleware already attached.
app := iris.Default()
```

## 使用中间件

Using middleware

```go
package main

import (
    "github.com/kataras/iris/v12"

    "github.com/kataras/iris/v12/middleware/recover"
)

func main() {
    // Creates an iris application without any middleware by default
    app := iris.New()

    // Global middleware using `UseRouter`.
    //
    // Recovery middleware recovers from any panics and writes a 500 if there was one.
    app.UseRouter(recover.New())
    
    // Per route middleware, you can add as many as you desire.
    app.Get("/benchmark", MyBenchLogger(), benchEndpoint)
    
    // Authorization group
    // authorized := app.Party("/", AuthRequired())
    // exactly the same as:
    authorized := app.Party("/")
    // per group middleware! in this case we use the custom created
    // AuthRequired() middleware just in the "authorized" group.
    authorized.Use(AuthRequired())
    {
        authorized.Post("/login", loginEndpoint)
        authorized.Post("/submit", submitEndpoint)
        authorized.Post("/read", readEndpoint)
    
        // nested group
        testing := authorized.Party("testing")
        testing.Get("/analytics", analyticsEndpoint)
    }
    
    // Listen and serve on 0.0.0.0:8080
    app.Listen(":8080")
}
```

## 应用程序文件记录器

Application File Logger 

```go
func main() {
    app := iris.Default()
    // Logging to a file.
    // Colors are automatically disabled when writing to a file.
    f, _ := os.Create("iris.log")
    app.Logger().SetOutput(f)

    // Use the following code if you need to write the logs
    // to file and console at the same time.
    // app.Logger().AddOutput(os.Stdout)
    
    app.Get("/ping", func(ctx iris.Context) {
        ctx.WriteString("pong")
    })

   app.Listen(":8080")
}
```

### 控制日志输出着色

Controlling Log output coloring 

By default, logs output on console should be colorized depending on the detected TTY.
默认情况下，控制台上的日志输出应该根据检测到的 TTY 着色。

Customize level title, text, color and styling at general.
自定义级别标题，文本，颜色和样式在 general。

Import ```golog``` and ```pio```:
导入 golog 和 pio:

```go
import (
    "github.com/kataras/golog"
    "github.com/kataras/pio"
    // [...]
)
```

Get a level to customize e.g. ```DebugLevel```:
获得一个自定义级别，例如 DebugLevel:

```go
level := golog.Levels[golog.DebugLevel]
```

You have full control over his text, title and style:
你可以完全控制他的文字、标题和风格:

```go
// The Name of the Level
// that named (lowercased) will be used
// to convert a string level on `SetLevel`
// to the correct Level type.
Name string
// AlternativeNames are the names that can be referred to this specific log level.
// i.e Name = "warn"
// AlternativeNames = []string{"warning"}, it's an optional field,
// therefore we keep Name as a simple string and created this new field.
AlternativeNames []string
// Tha Title is the prefix of the log level.
// See `ColorCode` and `Style` too.
// Both `ColorCode` and `Style` should be respected across writers.
Title string
// ColorCode a color for the `Title`.
ColorCode int
// Style one or more rich options for the `Title`.
Style []pio.RichOption
```

Example Code:

```go
level := golog.Levels[golog.DebugLevel]
level.Name = "debug" // default
level.Title = "[DBUG]" // default
level.ColorCode = pio.Yellow // default
```

To change the output format:
更改输出格式:

```go
app.Logger().SetFormat("json", "    ")
```

To register a custom Formatter:
注册自定义格式化程序:

```go
app.Logger().RegisterFormatter(new(myFormatter))
```

The [golog.Formatter interface](https://github.com/kataras/golog/blob/master/formatter.go) looks like this:
格式化程序接口如下所示:

```go
// Formatter is responsible to print a log to the logger's writer.
type Formatter interface {
    // The name of the formatter.
    String() string
    // Set any options and return a clone,
    // generic. See `Logger.SetFormat`.
    Options(opts ...interface{}) Formatter
    // Writes the "log" to "dest" logger.
    Format(dest io.Writer, log *Log) bool
}
```

To change the output and the format per level:
更改每个级别的输出和格式:

```
app.Logger().SetLevelOutput("error", os.Stderr)
app.Logger().SetLevelFormat("json")
```

## 请求记录

Request Logging

The application logger we've seen above it's used to log application-releated information and errors. At the other hand, the Access Logger, we see below, is used to log the incoming HTTP requests and responses.
我们在上面看到的应用程序日志记录器用于记录应用程序发布的信息和错误。另一方面，我们在下面看到的 Access Logger 用于记录传入的 HTTP 请求和响应。

```go
package main

import (
    "os"

    "github.com/kataras/iris/v12"
    "github.com/kataras/iris/v12/middleware/accesslog"
)

// Read the example and its comments carefully.
func makeAccessLog() *accesslog.AccessLog {
    // Initialize a new access log middleware.
    ac := accesslog.File("./access.log")
    // Remove this line to disable logging to console:
    ac.AddOutput(os.Stdout)

    // The default configuration:
    ac.Delim = '|'
    ac.TimeFormat = "2006-01-02 15:04:05"
    ac.Async = false
    ac.IP = true
    ac.BytesReceivedBody = true
    ac.BytesSentBody = true
    ac.BytesReceived = false
    ac.BytesSent = false
    ac.BodyMinify = true
    ac.RequestBody = true
    ac.ResponseBody = false
    ac.KeepMultiLineError = true
    ac.PanicLog = accesslog.LogHandler
    
    // Default line format if formatter is missing:
    // Time|Latency|Code|Method|Path|IP|Path Params Query Fields|Bytes Received|Bytes Sent|Request|Response|
    //
    // Set Custom Formatter:
    ac.SetFormatter(&accesslog.JSON{
        Indent:    "  ",
        HumanTime: true,
    })
    // ac.SetFormatter(&accesslog.CSV{})
    // ac.SetFormatter(&accesslog.Template{Text: "{{.Code}}"})
    
    return ac
}

func main() {
    ac := makeAccessLog()
    defer ac.Close() // Close the underline file.

    app := iris.New()
    // Register the middleware (UseRouter to catch http errors too).
    app.UseRouter(ac.Handler)
    
    app.Get("/", indexHandler)
    
    app.Listen(":8080")
}

func indexHandler(ctx iris.Context) {
    ctx.WriteString("OK")
}
```

Read more examples at:[ _examples/logging/request-logger]https://github.com/kataras/iris/tree/master/_examples/logging/request-logger.


## 模型绑定和验证

Model binding and validation 

To bind a request body into a type, use model binding. We currently support binding of ```JSON```、 ```JSONProtobuf```、 ```Protobuf```、 ```MsgPack```、 ```XML```、 ```YAML``` and standard form values (foo=bar&boo=baz).
若要将请求体绑定到类型，请使用模型绑定。我们目前支持 ```JSON```、 ```JSONProtobuf```、 ```Protobuf```、 ```MsgPack```、 ```XML```、 ```YAML``` 和标准表单值(foo = bar & boo = baz)的绑定。

```go
ReadJSON(outPtr interface{}) error
ReadJSONProtobuf(ptr proto.Message, opts ...ProtoUnmarshalOptions) error
ReadProtobuf(ptr proto.Message) error
ReadMsgPack(ptr interface{}) error
ReadXML(outPtr interface{}) error
ReadYAML(outPtr interface{}) error
ReadForm(formObject interface{}) error
ReadQuery(ptr interface{}) error
```

When using the ```ReadBody```, Iris tries to infer the binder depending on the Content-Type header. If you are sure what you are binding, you can use the specific ```ReadXXX``` methods, e.g. ```ReadJSON``` or ```ReadProtobuf``` and e.t.c.
当使用 ```ReadBody``` 时，Iris 尝试根据 Content-Type 标题推断活页夹。如果您确定要绑定的内容，可以使用特定的 ```ReadXXX``` 方法，例如 ```ReadJSON``` 或 ```ReadProtobuf``` 和 e.t.c。

```go
ReadBody(ptr interface{}) error
```

Iris, wisely, not features a builtin data validation. However, it does allow you to attach a validator which will automatically called on methods like ```ReadJSON```, ```ReadXML```.... In this example we will learn how to use the go-playground/validator/v10 for request body validation.
Iris 明智地不具备内建数据验证功能。但是，它允许您附加一个验证器，该验证器将自动调用 ```ReadJSON```、 ```ReadXML```... 等方法。在这个示例中，我们将学习如何使用 go-playground/validator/v10进行请求体验证。

Note that you need to set the corresponding binding tag on all fields you want to bind. For example, when binding from JSON, set ```json:"fieldname"```.
请注意，您需要在所有要绑定的字段上设置相应的绑定标记。例如，在从 JSON 绑定时，设置 ```JSON: “ fieldname”```。

You can also specify that specific fields are required. If a field is decorated with ```binding:"required"``` and has a empty value when binding, an error will be returned.

您还可以指定需要特定的字段。如果字段使用绑定修饰: ```binding:"required"```，并且在绑定时具有空值，则将返回错误。

```go
package main

import (
    "fmt"

    "github.com/kataras/iris/v12"
    "github.com/go-playground/validator/v10"
)

func main() {
    app := iris.New()
    app.Validator = validator.New()

    userRouter := app.Party("/user")
    {
        userRouter.Get("/validation-errors", resolveErrorsDocumentation)
        userRouter.Post("/", postUser)
    }
    app.Listen(":8080")
}

// User contains user information.
type User struct {
    FirstName      string     `json:"fname" validate:"required"`
    LastName       string     `json:"lname" validate:"required"`
    Age            uint8      `json:"age" validate:"gte=0,lte=130"`
    Email          string     `json:"email" validate:"required,email"`
    FavouriteColor string     `json:"favColor" validate:"hexcolor|rgb|rgba"`
    Addresses      []*Address `json:"addresses" validate:"required,dive,required"`
}

// Address houses a users address information.
type Address struct {
    Street string `json:"street" validate:"required"`
    City   string `json:"city" validate:"required"`
    Planet string `json:"planet" validate:"required"`
    Phone  string `json:"phone" validate:"required"`
}

type validationError struct {
    ActualTag string `json:"tag"`
    Namespace string `json:"namespace"`
    Kind      string `json:"kind"`
    Type      string `json:"type"`
    Value     string `json:"value"`
    Param     string `json:"param"`
}

func wrapValidationErrors(errs validator.ValidationErrors) []validationError {
    validationErrors := make([]validationError, 0, len(errs))
    for _, validationErr := range errs {
        validationErrors = append(validationErrors, validationError{
            ActualTag: validationErr.ActualTag(),
            Namespace: validationErr.Namespace(),
            Kind:      validationErr.Kind().String(),
            Type:      validationErr.Type().String(),
            Value:     fmt.Sprintf("%v", validationErr.Value()),
            Param:     validationErr.Param(),
        })
    }

    return validationErrors
}

func postUser(ctx iris.Context) {
    var user User
    err := ctx.ReadJSON(&user)
    if err != nil {
        // Handle the error, below you will find the right way to do that...

        if errs, ok := err.(validator.ValidationErrors); ok {
            // Wrap the errors with JSON format, the underline library returns the errors as interface.
            validationErrors := wrapValidationErrors(errs)
    
            // Fire an application/json+problem response and stop the handlers chain.
            ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
                Title("Validation error").
                Detail("One or more fields failed to be validated").
                Type("/user/validation-errors").
                Key("errors", validationErrors))
    
            return
        }
    
        // It's probably an internal JSON error, let's dont give more info here.
        ctx.StopWithStatus(iris.StatusInternalServerError)
        return
    }
    
    ctx.JSON(iris.Map{"message": "OK"})
}

func resolveErrorsDocumentation(ctx iris.Context) {
    ctx.WriteString("A page that should document to web developers or users of the API on how to resolve the validation errors")
}
```

#### Sample request


```go
{
    "fname": "",
    "lname": "",
    "age": 45,
    "email": "mail@example.com",
    "favColor": "#000",
    "addresses": [{
        "street": "Eavesdown Docks",
        "planet": "Persphone",
        "phone": "none",
        "city": "Unknown"
    }]
}
```

#### Sample response

```go
{
    "title": "Validation error",
    "detail": "One or more fields failed to be validated",
    "type": "http://localhost:8080/user/validation-errors",
    "status": 400,
    "fields": [
    {
        "tag": "required",
        "namespace": "User.FirstName",
        "kind": "string",
        "type": "string",
        "value": "",
        "param": ""
    },
    {
        "tag": "required",
        "namespace": "User.LastName",
        "kind": "string",
        "type": "string",
        "value": "",
        "param": ""
    }
    ]
}
```

Learn more about model validation at: https://github.com/go-playground/validator/blob/master/_examples

## 绑定查询字符串

Bind Query String 

The ```ReadQuery``` method only binds the query params and not the post data, use ```ReadForm``` instead to bind post data.

```ReadQuery``` 方法只绑定查询参数而不绑定发布数据，使用 ```ReadForm``` 来绑定发布数据。

```go
package main

import "github.com/kataras/iris/v12"

type Person struct {
    Name    string `url:"name,required"`
    Address string `url:"address"`
}

func main() {
    app := iris.Default()
    app.Any("/", index)
    app.Listen(":8080")
}

func index(ctx iris.Context) {
    var person Person
    if err := ctx.ReadQuery(&person); err!=nil {
        ctx.StopWithError(iris.StatusBadRequest, err)
        return
    }

    ctx.Application().Logger().Infof("Person: %#+v", person)
    ctx.WriteString("Success")
}
```

## 绑定任意

Bind Any 

Bind request body to "ptr" depending on the content-type that client sends the data, e.g. JSON, XML, YAML, MessagePack, Protobuf, Form and URL Query.
根据客户端发送数据的内容类型，将请求体绑定到“ ptr”，例如 JSON、 XML、 YAML、 MessagePack、 Protobuf、 Form 和 URL Query。

```go
package main

import (
    "time"

    "github.com/kataras/iris/v12"
)

type Person struct {
        Name       string    `form:"name" json:"name" url:"name" msgpack:"name"` 
        Address    string    `form:"address" json:"address" url:"address" msgpack:"address"`
        Birthday   time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1" json:"birthday" url:"birthday" msgpack:"birthday"`
        CreateTime time.Time `form:"createTime" time_format:"unixNano" json:"create_time" url:"create_time" msgpack:"createTime"`
        UnixTime   time.Time `form:"unixTime" time_format:"unix" json:"unix_time" url:"unix_time" msgpack:"unixTime"`
}

func main() {
    app := iris.Default()
    app.Any("/", index)
    app.Listen(":8080")
}

func index(ctx iris.Context) {
    var person Person
    if err := ctx.ReadBody(&person); err!=nil {
        ctx.StopWithError(iris.StatusBadRequest, err)
        return
    }

    ctx.Application().Logger().Infof("Person: %#+v", person)
    ctx.WriteString("Success")
}
```

Test it with:
用以下方法进行测试:

```bash
$ curl -X GET "localhost:8085/testing?name=kataras&address=xyz&birthday=1992-03-15&createTime=1562400033000000123&unixTime=1562400033"
```

## 绑定 URL 路径参数

Bind URL Path Parameters 

```go
package main

import "github.com/kataras/iris/v12"

type myParams struct {
    Name string   `param:"name"`
    Age  int      `param:"age"`
    Tail []string `param:"tail"`
}
// All parameters are required, as we already know,
// the router will fire 404 if name or int or tail are missing.

func main() {
    app := iris.Default()
    app.Get("/{name}/{age:int}/{tail:path}", func(ctx iris.Context) {
        var p myParams
        if err := ctx.ReadParams(&p); err != nil {
            ctx.StopWithError(iris.StatusInternalServerError, err)
            return
        }

        ctx.Writef("myParams: %#v", p)
    })
    app.Listen(":8088")
}
```

#### 请求

```bash
$ curl -v http://localhost:8080/kataras/27/iris/web/framework
```

## 绑定 Header 

Bind Header

```go
package main

import "github.com/kataras/iris/v12"


type myHeaders struct {
    RequestID      string `header:"X-Request-Id,required"`
    Authentication string `header:"Authentication,required"`
}

func main() {
    app := iris.Default()
    r.GET("/", func(ctx iris.Context) {
        var hs myHeaders
        if err := ctx.ReadHeaders(&hs); err != nil {
            ctx.StopWithError(iris.StatusInternalServerError, err)
            return
        }

        ctx.JSON(hs)
    })
    
    app.Listen(":8080")
}
```

#### Request

```bash
curl -H "x-request-id:373713f0-6b4b-42ea-ab9f-e2e04bc38e73" -H "authentication: Bearer my-token" \
http://localhost:8080
```

#### Response

```json
{
  "RequestID": "373713f0-6b4b-42ea-ab9f-e2e04bc38e73",
  "Authentication": "Bearer my-token"
}
```

## 绑定 HTML 复选框

Bind HTML checkboxes 

```go
package main

import "github.com/kataras/iris/v12"

func main() {
    app := iris.New()
    app.RegisterView(iris.HTML("./templates", ".html"))

    app.Get("/", showForm)
    app.Post("/", handleForm)
    
    app.Listen(":8080")
}

func showForm(ctx iris.Context) {
    ctx.View("form.html")
}

type formExample struct {
    Colors []string `form:"colors[]"` // or just "colors".
}

func handleForm(ctx iris.Context) {
    var form formExample
    err := ctx.ReadForm(&form)
    if err != nil {
        ctx.StopWithError(iris.StatusBadRequest, err)
        return
    }

    ctx.JSON(iris.Map{"Colors": form.Colors})
}
```

#### templates/form.html

模板/表单. html

```html
<form action="/" method="POST">
    <p>Check one or more colors</p>

    <label for="red">Red</label>
    <!-- name can be "colors" too -->
    <input type="checkbox" name="colors[]" value="red" id="red">
    <label for="green">Green</label>
    <input type="checkbox" name="colors[]" value="green" id="green">
    <label for="blue">Blue</label>
    <input type="checkbox" name="colors[]" value="blue" id="blue">
    <input type="submit">
</form>
```

Response

```json
{
  "Colors": [
    "red",
    "green",
    "blue"
  ]
}
```

## JSON、 JSONP、 XML、 Markdown、 YAML 和 MsgPack 呈现

JSON, JSONP, XML, Markdown, YAML and MsgPack rendering 

Detailed examples can be found [here](https://github.com/kataras/iris/tree/master/_examples/response-writer/write-rest).
详细的例子可以在这里找到。

```go
func main() {
    app := iris.New()

    // iris.Map is an alias of map[string]interface{}
    app.Get("/json", func(ctx iris.Context) {
        ctx.JSON(iris.Map{"message": "hello", "status": iris.StatusOK})
    })
    
    // Use Secure field to prevent json hijacking.
    // It prepends `"while(1),"` to the body when the data is array.
    app.Get("/json_secure", func(ctx iris.Context) {
        response := []string{"val1", "val2", "val3"}
        options := iris.JSON{Indent: "", Secure: true}
        ctx.JSON(response, options)
    
        // Will output: while(1);["val1","val2","val3"]
    })
    
    // Use ASCII field to generate ASCII-only JSON
    // with escaped non-ASCII characters.
    app.Get("/json_ascii", func(ctx iris.Context) {
        response := iris.Map{"lang": "GO-虹膜", "tag": "<br>"}
        options := iris.JSON{Indent: "    ", ASCII: true}
        ctx.JSON(response, options)
    
        /* Will output:
           {
               "lang": "GO-\u8679\u819c",
               "tag": "\u003cbr\u003e"
           }
        */
    })
    
    // Normally, JSON replaces special HTML characters with their unicode entities.
    // If you want to encode such characters literally,
    // you SHOULD set the UnescapeHTML field to true.
    app.Get("/json_raw", func(ctx iris.Context) {
        options := iris.JSON{UnescapeHTML: true}
        ctx.JSON(iris.Map{
            "html": "<b>Hello, world!</b>",
        }, options)
    
        // Will output: {"html":"<b>Hello, world!</b>"}
    })
    
    app.Get("/json_struct", func(ctx iris.Context) {
        // You also can use a struct.
        var msg struct {
            Name    string `json:"user"`
            Message string
            Number  int
        }
        msg.Name = "Mariah"
        msg.Message = "hello"
        msg.Number = 42
        // Note that msg.Name becomes "user" in the JSON.
        // Will output: {"user": "Mariah", "Message": "hello", "Number": 42}
        ctx.JSON(msg)
    })
    
    app.Get("/jsonp", func(ctx iris.Context) {
        ctx.JSONP(iris.Map{"hello": "jsonp"}, iris.JSONP{Callback: "callbackName"})
    })
    
    app.Get("/xml", func(ctx iris.Context) {
        ctx.XML(iris.Map{"message": "hello", "status": iris.StatusOK})
    })
    
    app.Get("/markdown", func(ctx iris.Context) {
        ctx.Markdown([]byte("# Hello Dynamic Markdown -- iris"))
    })
    
    app.Get("/yaml", func(ctx iris.Context) {
        ctx.YAML(iris.Map{"message": "hello", "status": iris.StatusOK})
    })
    
    app.Get("/msgpack", func(ctx iris.Context) {
        u := User{
            Firstname: "John",
            Lastname:  "Doe",
            City:      "Neither FBI knows!!!",
            Age:       25,
        }
    
        ctx.MsgPack(u)
    })
    
    // Render using jsoniter instead of the encoding/json:
    app.Listen(":8080", iris.WithOptimizations)
}
```

### Protobuf

Iris supports native ```protobuf``` with Protobuf and protobuf to JSON encode and decode.
Iris 支持带有 Protobuf 的本地 Protobuf，而 Protobuf 则支持 JSON 编码和解码。

```go
package main

import (
    "app/protos"

    "github.com/kataras/iris/v12"
)

func main() {
    app := iris.New()

    app.Get("/", send)
    app.Get("/json", sendAsJSON)
    app.Post("/read", read)
    app.Post("/read_json", readFromJSON)
    
    app.Listen(":8080")
}

func send(ctx iris.Context) {
    response := &protos.HelloReply{Message: "Hello, World!"}
    ctx.Protobuf(response)
}

func sendAsJSON(ctx iris.Context) {
    response := &protos.HelloReply{Message: "Hello, World!"}
    options := iris.JSON{
        Proto: iris.ProtoMarshalOptions{
            AllowPartial: true,
            Multiline:    true,
            Indent:       "    ",
        },
    }

    ctx.JSON(response, options)
}

func read(ctx iris.Context) {
    var request protos.HelloRequest

    err := ctx.ReadProtobuf(&request)
    if err != nil {
        ctx.StopWithError(iris.StatusBadRequest, err)
        return
    }
    
    ctx.Writef("HelloRequest.Name = %s", request.Name)
}

func readFromJSON(ctx iris.Context) {
    var request protos.HelloRequest

    err := ctx.ReadJSONProtobuf(&request)
    if err != nil {
        ctx.StopWithError(iris.StatusBadRequest, err)
        return
    }
    
    ctx.Writef("HelloRequest.Name = %s", request.Name)
}
```

## 服务静态文件

Serving static files 

```go
func main() {
    app := iris.New()
    app.Favicon("./resources/favicon.ico")
    app.HandleDir("/assets", iris.Dir("./assets"))

    app.Listen(":8080")
}
```

The ```HandleDir``` method accepts a third, optional argument of ```DirOptions```:

```HandleDir```方法接受 ```dirotices``` 的第三个可选参数:

```go
type DirOptions struct {
    // Defaults to "/index.html", if request path is ending with **/*/$IndexName
    // then it redirects to **/*(/) which another handler is handling it,
    // that another handler, called index handler, is auto-registered by the framework
    // if end developer does not managed to handle it by hand.
    IndexName string
    // PushTargets filenames (map's value) to
    // be served without additional client's requests (HTTP/2 Push)
    // when a specific request path (map's key WITHOUT prefix)
    // is requested and it's not a directory (it's an `IndexFile`).
    //
    // Example:
    //     "/": {
    //         "favicon.ico",
    //         "js/main.js",
    //         "css/main.css",
    //     }
    PushTargets map[string][]string
    // PushTargetsRegexp like `PushTargets` but accepts regexp which
    // is compared against all files under a directory (recursively).
    // The `IndexName` should be set.
    //
    // Example:
    // "/": regexp.MustCompile("((.*).js|(.*).css|(.*).ico)$")
    // See `iris.MatchCommonAssets` too.
    PushTargetsRegexp map[string]*regexp.Regexp

    // Cache to enable in-memory cache and pre-compress files.
    Cache DirCacheOptions
    // When files should served under compression.
    Compress bool
    
    // List the files inside the current requested directory if `IndexName` not found.
    ShowList bool
    // If `ShowList` is true then this function will be used instead
    // of the default one to show the list of files of a current requested directory(dir).
    // See `DirListRich` package-level function too.
    DirList DirListFunc
    
    // Files downloaded and saved locally.
    Attachments Attachments
    
    // Optional validator that loops through each requested resource.
    AssetValidator func(ctx *context.Context, name string) bool
}
```

Learn more about [file-server](https://github.com/kataras/iris/tree/master/_examples/file-server).
了解更多关于文件服务器的信息。

## 从上下文中提供数据

Serving data from Context 

```go
SendFile(filename string, destinationName string) error
SendFileWithRate(src, destName string, limit float64, burst int) error
```

Usage


Force-Send a file to the client:
强制-向客户端发送一个文件:

```go
func handler(ctx iris.Context) {
    src := "./files/first.zip"
    ctx.SendFile(src, "client.zip")
}
```

Limit download speed to ~50Kb/s with a burst of 100KB:
将下载速度限制在 ~ 50Kb/s，突发速度为100KB:

```
func handler(ctx iris.Context) {
    src := "./files/big.zip"
    // optionally, keep it empty to resolve the filename based on the "src".
    dest := "" 

    limit := 50.0 * iris.KB
    burst := 100 * iris.KB
    ctx.SendFileWithRate(src, dest, limit, burst)
}
```

```go
ServeContent(content io.ReadSeeker, filename string, modtime time.Time)
ServeContentWithRate(content io.ReadSeeker, filename string, modtime time.Time, limit float64, burst int)

ServeFile(filename string) error
ServeFileWithRate(filename string, limit float64, burst int) error
```

Usage
用法

```go
func handler(ctx iris.Context) {
    ctx.ServeFile("./public/main.js")
}
```

## 模板渲染

Template rendering 

Iris supports 8 template engines out-of-the-box, developers can still use any external golang template engine, as ```Context.ResponseWriter()``` is an ```io.Writer```.
Iris 支持8模板引擎开箱即用，开发人员仍然可以使用任何外部的 golang 模板引擎，作为``Context.ResponseWriter()```是一个 ```io.Writer```。

All template engines share a common API i.e. Parse using embedded assets, Layouts and Party-specific layout, Template Funcs, Partial Render and more.
所有模板引擎共享一个公共 API，即解析使用嵌入 assets ，布局和 Party-specific 的布局，模板函数，Partial Render和更多。

！> render会加载layout文件然后将你render指定的页面替换到layout的$content中，而renderPartial不会加载。

| \#   | Name       | Parser 解析器    |
| ---- | ---------- | ---------------- |
| 1    | HTML       | html/template    |
| 2    | Blocks     | kataras/blocks   |
| 3    | Django     | flosch/pongo2    |
| 4    | Pug        | Joker/jade       |
| 5    | Handlebars | aymerick/raymond |
| 6    | Amber      | eknkc/amber      |
| 7    | Jet        | CloudyKit/jet    |
| 8    | Ace        | yosssi/ace       |


[List of Examples.](https://github.com/kataras/iris/tree/master/_examples/view)

List of Benchmarks.
基准一览表。

A view engine can be registered per-Party. To **register** a view engine use the ```Application/Party.RegisterView(ViewEngine)``` method as shown below.
视图引擎可以按照每一方注册。要**注册**视图引擎，请使用如下所示的 ```Application/Party. RegisterView (ViewEngine)```方法。

Load all templates from the "./views" folder where extension is ".html" and parse them using the standard ```html/template``` package.
从扩展名为 “.html” 的 “./views” 文件夹中加载所有模板，并使用标准的 ```html/template``` 包解析它们。

```go
// [app := iris.New...]
tmpl := iris.HTML("./views", ".html")
app.RegisterView(tmpl)
```

To **render or execute** a view use the ```Context.View``` method inside the main route's handler.
要**呈现或执行**视图，请使用主路由处理程序内的 ```Context. View``` 方法。

```go
ctx.View("hi.html")
```

To **bind** Go values with key-value pattern inside a view through middleware or main handler use the ```Context.ViewData``` method before the ```Context.View``` one.
要通过中间件或主处理程序将 Go的值 与 视图中的键值模式 绑定在一起，请在```Context.View```之前使用 ```Context.ViewData``` 方法。

Bind: ```{{.message}}``` with ```"Hello world!"```.
绑定: {{ . message }与“ Hello world!”。

```go
ctx.ViewData("message", "Hello world!")
```

Root binding:
根绑定:

```go
ctx.View("user-page.html", User{})

// root binding as {{.Name}}
```

To **add a template function** use the ```AddFunc``` method of the preferred view engine.
要添加模板函数，请使用首选视图引擎的 ```AddFunc``` 方法。

```go
//       func name, input arguments, render value
tmpl.AddFunc("greet", func(s string) string {
    return "Greetings " + s + "!"
})
```

To reload on every request call the view engine's Reload method.
要在每个请求上重新加载，请调用视图引擎的 Reload 方法。

```go
tmpl.Reload(true)
```

To use **embedded** templates and not depend on local file system use the [go-bindata](https://github.com/go-bindata/go-bindata) external tool and pass its ```AssetFile()``` generated function to the first input argument of the preferred view engine.
要使用嵌入模板而不依赖于本地文件系统，请使用 [go-bindata](https://github.com/go-bindata/go-bindata) 外部工具并将其 ```AssetFile ()``` 生成的函数传递给首选视图引擎的第一个输入参数。

```go
 tmpl := iris.HTML(AssetFile(), ".html")
```

Example Code:

```go
// file: main.go
package main

import "github.com/kataras/iris/v12"

func main() {
    app := iris.New()

    // Parse all templates from the "./views" folder
    // where extension is ".html" and parse them
    // using the standard `html/template` package.
    tmpl := iris.HTML("./views", ".html")
    // Set custom delimeters.
    tmpl.Delims("{{", "}}")
    // Enable re-build on local template files changes.
    tmpl.Reload(true)
    
    // Default template funcs are:
    //
    // - {{ urlpath "myNamedRoute" "pathParameter_ifNeeded" }}
    // - {{ render "header.html" }}
    // and partial relative path to current page:
    // - {{ render_r "header.html" }} 
    // - {{ yield }}
    // - {{ current }}
    // Register a custom template func:
    tmpl.AddFunc("greet", func(s string) string {
        return "Greetings " + s + "!"
    })
    
    // Register the view engine to the views,
    // this will load the templates.
    app.RegisterView(tmpl)
    
    // Method:    GET
    // Resource:  http://localhost:8080
    app.Get("/", func(ctx iris.Context) {
        // Bind: {{.message}} with "Hello world!"
        ctx.ViewData("message", "Hello world!")
        // Render template file: ./views/hi.html
        ctx.View("hi.html")
    })
    
    app.Listen(":8080")
}
```

```html
<!-- file: ./views/hi.html -->
<html>
<head>
    <title>Hi Page</title>
</head>
<body>
    <h1>{{.message}}</h1>
    <strong>{{greet "to you"}}</strong>
</body>
</html>
```

Open a browser tab at http://localhost:8080.
在 http://localhost:8080 打开浏览器标签。

The rendered result will look like this:
呈现的结果如下:

```html
<html>
<head>
    <title>Hi Page</title>
</head>
<body>
    <h1>Hello world!</h1>
    <strong>Greetings to you!</strong>
</body>
</html>
```

## 多模版 

Multitemplate

Iris allows unlimited number of registered view engines per Application. Besides that, you can register a view engine **per Party or through middleware too!**.
Iris 允许每个应用程序不限数量的注册视图引擎。除此之外，您还可以为每个 Party 或通过中间件注册一个视图引擎！.

```go
// Register a view engine per group of routes.
adminGroup := app.Party("/admin")
adminGroup.RegisterView(iris.Blocks("./views/admin", ".html"))
```

#### Through Middleware 

通过中间件

```go
func middleware(views iris.ViewEngine) iris.Handler {
    return func(ctx iris.Context) {
        ctx.ViewEngine(views)
        ctx.Next()
    }
}
```

Usage

```go
// Register a view engine on-fly for the current chain of handlers.
views := iris.Blocks("./views/on-fly", ".html")
views.Load()

app.Get("/", setViews(views), onFly)
```

## 重定向

Redirects 

Issuing a HTTP redirect is easy. Both internal and external locations are supported. By locations we mean, paths, subdomains, domains and e.t.c.
发出 HTTP 重定向是很容易的。支持内部和外部位置。所谓位置，我们指的是路径、子域、域和 e.t.c。

#### From Handler 

```go
app.Get("/", func(ctx iris.Context) {
    ctx.Redirect("https://golang.org/dl", iris.StatusMovedPermanently)
})
```

Issuing a HTTP redirect from POST.
从 POST 发出 HTTP 重定向。

```go
app.Post("/", func(ctx iris.Context) {
    ctx.Redirect("/login", iris.StatusFound)
})
```

Issuing a local router redirect from a Handler, use ```Application.ServeHTTPC``` or ```Exec()``` like below.
使用如下所示的 ```Application.ServeHTTPC``` 或 ```Exec ()```发出本地路由器重定向。

```go
app.Get("/test", func(ctx iris.Context) {
    r := ctx.Request()
    r.URL.Path = "/test2"

    ctx.Application().ServeHTTPC(ctx)
    // OR
    // ctx.Exec("GET", "/test2")
})

app.Get("/test2", func(ctx iris.Context) {
    ctx.JSON(iris.Map{"hello": "world"})
})
```

#### Globally

全局

Use the syntax we all love.
使用我们都喜欢的语法。

```go
import "github.com/kataras/iris/v12/middleware/rewrite"
```

```go
func main() {
    app := iris.New()
    // [...routes]
    redirects := rewrite.Load("redirects.yml")
    app.WrapRouter(redirects)
    app.Listen(":80")
}
```

The ```"redirects.yml"``` file looks like that:
```“ redirects.yml”```文件如下所示:

```yaml
RedirectMatch:
  # Redirects /seo/* to /*
  - 301 /seo/(.*) /$1

  # Redirects /docs/v12* to /docs
  - 301 /docs/v12(.*) /docs

  # Redirects /old(.*) to /
  - 301 /old(.*) /

  # Redirects http or https://test.* to http or https://newtest.*
  - 301 ^(http|https)://test.(.*) $1://newtest.$2

  # Handles /*.json or .xml as *?format=json or xml,
  # without redirect. See /users route.
  # When Code is 0 then it does not redirect the request,
  # instead it changes the request URL
  # and leaves a route handle the request.
  - 0 /(.*).(json|xml) /$1?format=$2

# Redirects root domain to www.
# Creation of a www subdomain inside the Application is unnecessary,
# all requests are handled by the root Application itself.
PrimarySubdomain: www
```

The full code can be found at the [rewrite middleware example](https://github.com/kataras/iris/tree/master/_examples/routing/rewrite).


## 自定义中间件

Custom Middleware 

```go
func Logger() iris.Handler {
    return func(ctx iris.Context) {
        t := time.Now()

        // Set a shared variable between handlers
        ctx.Values().Set("framework", "iris")
    
        // before request
    
        ctx.Next()
    
        // after request
        latency := time.Since(t)
        log.Print(latency)
    
        // access the status we are sending
        status := ctx.GetStatusCode()
        log.Println(status)
    }
}

func main() {
    app := iris.New()
    app.Use(Logger())

    app.Get("/test", func(ctx iris.Context) {
        // retrieve a value set by the middleware.
        framework := ctx.Values().GetString("framework")
    
        // it would print: "iris"
        log.Println(framework)
    })
    
    app.Listen(":8080")
}
```

## 使用基本认证

Using Basic Authentication 

HTTP Basic Authentication is the simplest technique for enforcing access controls to web resources because it does not require cookies, session identifiers, or login pages; rather, HTTP Basic authentication uses standard fields in the HTTP header.
基本认证是对网络资源实施访问控制的最简单的技术，因为它不需要 cookie、会话标识符或登录页面; 更确切地说，HTTP 基本认证在 HTTP 报头中使用标准字段。

The Basic Authentication middleware is included with the Iris framework, so you do not need to install it separately.
基本身份验证中间件包含在 Iris 框架中，因此不需要单独安装。

**1.** Import the middleware
导入中间件

```go
import "github.com/kataras/iris/v12/middleware/basicauth"
```

**2.** Configure the middleware with its ```Options``` struct:
使用 ```Options``` 配置中间件:

```go
opts := basicauth.Options{
    Allow: basicauth.AllowUsers(map[string]string{
        "username": "password",
    }),
    Realm:        "Authorization Required",
    ErrorHandler: basicauth.DefaultErrorHandler,
    // [...more options]
}
```

**3.** Initialize the middleware:
初始化中间件:

```go
auth := basicauth.New(opts)
```

**3.1** The above steps are the same as the Default function:
以上步骤与 Default 函数相同:

```go
auth := basicauth.Default(map[string]string{
    "username": "password",
})
```

**3.2** Use a custom slice of Users:
使用自定义用户切片:

```
// The struct value MUST contain a Username and Passwords fields
// or GetUsername() string and GetPassword() string methods.
type User struct {
    Username string
    Password string
}

// [...]
auth := basicauth.Default([]User{...})
```

**3.3** Load users from a file optionally, passwords are encrypted using the golang.org/x/crypto/bcrypt package:
可以选择从文件加载用户，密码使用 golang.org/x/crypto/bcrypt 加密包进行加密:

```go
auth := basicauth.Load("users.yml", basicauth.BCRYPT)
```

**3.3.1** The same can be achieved using the Options (recommended):
使用下列方案(建议)亦可达到同样的效果:

```go
opts := basicauth.Options{
    Allow: basicauth.AllowUsersFile("users.yml", basicauth.BCRYPT),
    Realm: basicauth.DefaultRealm,
    // [...more options]
}

auth := basicauth.New(opts)
```

Where the ```users.yml``` may look like that:
其中 ```users.yml``` 可能是这样的:

```yaml
- username: kataras
  password: $2a$10$Irg8k8HWkDlvL0YDBKLCYee6j6zzIFTplJcvZYKA.B8/clHPZn2Ey
  # encrypted of kataras_pass
  role: admin
- username: makis
  password: $2a$10$3GXzp3J5GhHThGisbpvpZuftbmzPivDMo94XPnkTnDe7254x7sJ3O
  # encrypted of makis_pass
  role: member
```

**4.** Register the middleware:
注册中间件:

```go
// Register to all matched routes
// under a Party and its children.
app.Use(auth)

// OR/and register to all http error routes.
app.UseError(auth)

// OR register under a path prefix of a specific Party,
// including all http errors of this path prefix.
app.UseRouter(auth)

// OR register to a specific Route before its main handler.
app.Post("/protected", auth, routeHandler)
```

**5.** Retrieve the username & password:
检索用户名和密码:

```go
func routeHandler(ctx iris.Context) {
    username, password, _ := ctx.Request().BasicAuth()
    // [...]
}
```

**5.1** Retrieve the User value (useful when you register a slice of custom user struct at ```Options.AllowUsers```):
检索 User 值(在 ```Options. AllowUsers``` 中注册一个自定义用户结构片段时非常有用) :

```go
func routeHandler(ctx iris.Context) {
    user := ctx.User().(*iris.SimpleUser)
    // user.Username
    // user.Password
}
```

在 _ examples/auth 阅读更多授权和身份验证示例。
Read more authorization and authentication examples at [_examples/auth](https://github.com/kataras/iris/tree/master/_examples/auth).


## 中间件内部的 goroutine

Goroutines inside a middleware 

When starting new Goroutines inside a middleware or handler, you **SHOULD NOT** use the original context inside it, you have to use a read-only copy.
在中间件或处理程序中启动新的 goroutine 时，不应使用其中的原始上下文，必须使用只读副本。

```go
func main() {
    app := iris.Default()

    app.Get("/long_async", func(ctx iris.Context) {
        // create a clone to be used inside the goroutine
        ctxCopy := ctx.Clone()
        go func() {
            // simulate a long task with time.Sleep(). 5 seconds
            time.Sleep(5 * time.Second)
    
            // note that you are using the copied context "ctxCopy", IMPORTANT
            log.Printf("Done! in path: %s", ctxCopy.Path())
        }()
    })
    
    app.Get("/long_sync", func(ctx iris.Context) {
        // simulate a long task with time.Sleep(). 5 seconds
        time.Sleep(5 * time.Second)
    
        // since we are NOT using a goroutine, we do not have to copy the context
        log.Printf("Done! in path: %s", ctx.Path())
    })
    
    app.Listen(":8080")
}
```

## 自定义 HTTP 配置

Custom HTTP configuration 

在 _ examples/http-server 文件夹中可以找到超过12个关于 http 服务器配置的示例。
More than 12 examples about http server configuration can be found at the [_examples/http-server](https://github.com/kataras/iris/tree/master/_examples/http-server) folder.


Use ```http.ListenAndServe()``` directly, like this:
直接使用 ```http. ListenAndServe ()``` ，像这样:

```go
func main() {
    app := iris.New()
    // [...routes]
    if err := app.Build(); err!=nil{
        panic(err)
    }
    http.ListenAndServe(":8080", app)
}
```

Note that you SHOULD call its ```Build``` method manually to build the application and the router before using it as an ```http.Handler```.
注意，在使用它作为 http 之前，您应该手动调用它的 Build 方法来构建应用程序和路由器。处理程序。

Another example:
另一个例子:

```go
func main() {
    app := iris.New()
    // [...routes]
    app.Build()

    srv := &http.Server{
        Addr:           ":8080",
        Handler:        app,
        ReadTimeout:    10 * time.Second,
        WriteTimeout:   10 * time.Second,
        MaxHeaderBytes: 1 << 20,
    }
    srv.ListenAndServe()
}
```

However, you rarely need an external ```http.Server``` instance with Iris. You can listen using any tcp listener, http server or a custom function via ```Application.Run``` method.
然而，你很少需要外部的 ```http.Server```的 Iris 的实例。您可以使用任何 tcp 监听器，http 服务器或通过```Application.Run``` 方法自定义函数。

```go
app.Run(iris.Listener(l net.Listener)) // listen using a custom net.Listener
app.Run(iris.Server(srv *http.Server)) // listen using a custom http.Server
app.Run(iris.Addr(addr string)) // the app.Listen is a shortcut of this method.
app.Run(iris.TLS(addr string, certFileOrContents, keyFileOrContents string)) // listen TLS.
app.Run(iris.AutoTLS(addr, domain, email string)) // listen using letsencrypt (see below).

// and any custom function that returns an error:
app.Run(iris.Raw(f func() error))
```

## Socket Sharding

套接字切分

This option allows linear scaling server performance on multi-CPU servers. See https://www.nginx.com/blog/socket-sharding-nginx-release-1-9-1/ for details. Enable with ```iris.WithSocketSharding``` configurator.
这个选项允许在多 cpu 服务器上线性调整服务器的性能，详细信息请参阅 https://www.nginx.com/blog/socket-sharding-nginx-release-1-9-1/ ，使用 ```iris.WithSocketSharding``` 启用配置器。

Example Code:

```go
package main

import (
    "time"

    "github.com/kataras/iris/v12"
)

func main() {
    startup := time.Now()

    app := iris.New()
    app.Get("/", func(ctx iris.Context) {
        s := startup.Format(ctx.Application().ConfigurationReadOnly().GetTimeFormat())
        ctx.Writef("This server started at: %s\n", s)
    })
    
    app.Listen(":8080", iris.WithSocketSharding)
    // or app.Run(..., iris.WithSocketSharding)
}
```

## 支持 Let’s Encrypt

Support Let's Encrypt 

Example for 1-line LetsEncrypt HTTPS servers.
单行 LetsEncrypt HTTPS 服务器示例。

```go
package main

import (
    "log"

    "github.com/iris-gonic/autotls"
    "github.com/kataras/iris/v12"
)

func main() {
    app := iris.Default()

    // Ping handler
    app.Get("/ping", func(ctx iris.Context) {
        ctx.WriteString("pong")
    })
    
    app.Run(iris.AutoTLS(":443", "example.com example2.com", "mail@example.com"))
}
```

Example for custom TLS (you can bind an autocert manager too):
自定义 TLS (你也可以绑定一个 autocert 管理器) :

```go
app.Run(
    iris.TLS(":443", "", "", func(su *iris.Supervisor) {
        su.Server.TLSConfig = &tls.Config{
            /* your custom fields */
        },
    }),
)
```

!> All ```iris.Runner``` methods such as: Addr, TLS, AutoTLS, Server, Listener and e.t.c accept a variadic input argument of ```func(*iris.Supervisor)``` to configure the http server instance on build state.

所有的 ```iris.Runner```方法像 Addr、 TLS、 auttls、 Server、 Listener 等都接受可变输入参数的 ```func(*iris.Supervisor)``` 在 build 阶段配置 http server 实例。

## 使用 Iris 运行多个服务

Run multiple service using Iris 

```go
package main

import (
    "log"
    "net/http"
    "time"

    "github.com/kataras/iris/v12"
    "github.com/kataras/iris/v12/middleware/recover"
    
    "golang.org/x/sync/errgroup"
)

var g errgroup.Group

func startApp1() error {
    app := iris.New().SetName("app1")
    app.Use(recover.New())
    app.Get("/", func(ctx iris.Context) {
        app.Get("/", func(ctx iris.Context) {
            ctx.JSON(iris.Map{
                "code":  iris.StatusOK,
                "message": "Welcome server 1",
            })
        })
    })

    app.Build()
   return app.Listen(":8080")
}

func startApp2() error {
    app := iris.New().SetName("app2")
    app.Use(recover.New())
    app.Get("/", func(ctx iris.Context) {
        ctx.JSON(iris.Map{
            "code":  iris.StatusOK,
            "message": "Welcome server 2",
        })
    })

    return app.Listen(":8081")
}

func main() {
    g.Go(startApp1)
    g.Go(startApp2)

    if err := g.Wait(); err != nil {
        log.Fatal(err)
    }
}
```

Manage multiple Iris instances through the ```apps``` package. Read more [here](https://github.com/kataras/iris/tree/master/_examples/http-server/graceful-shutdown).
通过 ```apps``` 包管理多个 Iris 实例。

## 优雅的关闭或重新启动

Graceful shutdown or restart 

There are a few approaches you can use to perform a graceful shutdown or restart. You can make use of third-party packages specifically built for that, or you can use the ```app.Shutdown(context.Context)``` method. Examples can be found [here](https://github.com/kataras/iris/tree/master/_examples/http-server/graceful-shutdown).
有几种方法可以用来优雅地关机或重新启动。您可以使用专门为此构建的第三方软件包，也可以使用该应用程序。关闭(上下文。上下文)方法。例子可以在这里找到。

Register an event on CTRL/CMD+C using ```iris.RegisterOnInterrupt```:
使用 ```iris. RegisterOnInterrupt``` 在 CTRL/CMD + C 上注册事件:

```go
idleConnsClosed := make(chan struct{})
iris.RegisterOnInterrupt(func() {
    timeout := 10 * time.Second
    ctx, cancel := stdContext.WithTimeout(stdContext.Background(), timeout)
    defer cancel()
    // close all hosts.
    app.Shutdown(ctx)
    close(idleConnsClosed)
})

// [...]
app.Listen(":8080", iris.WithoutInterruptHandler, iris.WithoutServerError(iris.ErrServerClosed))
<-idleConnsClosed
```

## 使用模板构建单个二进制文件

Build a single binary with templates 

You can build a server into a single binary containing templates by using [go-bindata](https://github.com/go-bindata/go-bindata) 's ```AssetFile``` generated function.
通过使用[ go-bindata ](https://github.com/go-bindata/go-bindata)的 ```AssetFile``` 生成函数，您可以将服务器构建成包含模板的单个二进制文件。

```sh
$ go get -u github.com/go-bindata/go-bindata/...
$ go-bindata -fs -prefix "templates" ./templates/...
$ go run .
```

Example Code:

```go
func main() {
    app := iris.New()

    tmpl := iris.HTML(AssetFile(), ".html")
    tmpl.Layout("layouts/layout.html")
    tmpl.AddFunc("greet", func(s string) string {
        return "Greetings " + s + "!"
    })
    app.RegisterView(tmpl)
    
    // [...]
}
```

See complete examples at the [_examples/view](https://github.com/kataras/iris/tree/master/_examples/view).
请参阅 [_examples/view](https://github.com/kataras/iris/tree/master/_examples/view) 中的完整示例。

## 尝试将主体绑定到不同的结构中

Try to bind body into different structs 

The normal methods for binding request body consumes ```ctx.Request().Body``` and they cannot be called multiple times, **unless** the ```iris.WithoutBodyConsumptionOnUnmarshal``` configurator is passed to ```app.Run/Listen```.
用于绑定request body的常规方法需要 ```ctx.Request().Body```并且它们不能被多次调用，除非```iris.WithoutBodyConsumptionOnUnmarshal``` 配置器被传递给 ```app.Run/Listen```。

```go
package main

import "github.com/kataras/iris/v12"

func main() {
    app := iris.New()

    app.Post("/", logAllBody, logJSON, logFormValues, func(ctx iris.Context) {
        // body, err := ioutil.ReadAll(ctx.Request().Body) once or
        body, err := ctx.GetBody() // as many times as you need.
        if err != nil {
            ctx.StopWithError(iris.StatusInternalServerError, err)
            return
        }
    
        if len(body) == 0 {
            ctx.WriteString(`The body was empty.`)
        } else {
            ctx.WriteString("OK body is still:\n")
            ctx.Write(body)
        }
    })
    
    app.Listen(":8080", iris.WithoutBodyConsumptionOnUnmarshal)
}

func logAllBody(ctx iris.Context) {
    body, err := ctx.GetBody()
    if err == nil && len(body) > 0 {
        ctx.Application().Logger().Infof("logAllBody: %s", string(body))
    }

    ctx.Next()
}

func logJSON(ctx iris.Context) {
    var p interface{}
    if err := ctx.ReadJSON(&p); err == nil {
        ctx.Application().Logger().Infof("logJSON: %#+v", p)
    }

    ctx.Next()
}

func logFormValues(ctx iris.Context) {
    values := ctx.FormValues()
    if values != nil {
        ctx.Application().Logger().Infof("logFormValues: %v", values)
    }

    ctx.Next()
}
```

You can use the ```ReadBody``` to bind a struct to a request based on the client's content-type. You can also use [Content Negotiation](https://developer.mozilla.org/en-US/docs/Web/HTTP/Content_negotiation). Here's a full example:
可以使用 ReadBody 根据客户端的内容类型将结构绑定到请求。你也可以使用[Content Negotiation](https://developer.mozilla.org/en-US/docs/Web/HTTP/Content_negotiation)。下面是一个完整的例子:

```go
package main

import (
    "github.com/kataras/iris/v12"
)

func main() {
    app := newApp()
    // See main_test.go for usage.
    app.Listen(":8080")
}

func newApp() *iris.Application {
    app := iris.New()
    // To automatically decompress using gzip:
    // app.Use(iris.GzipReader)

    app.Use(setAllowedResponses)
    
    app.Post("/", readBody)
    
    return app
}

type payload struct {
    Message string `json:"message" xml:"message" msgpack:"message" yaml:"Message" url:"message" form:"message"`
}

func readBody(ctx iris.Context) {
    var p payload

    // Bind request body to "p" depending on the content-type that client sends the data,
    // e.g. JSON, XML, YAML, MessagePack, Protobuf, Form and URL Query.
    err := ctx.ReadBody(&p)
    if err != nil {
        ctx.StopWithProblem(iris.StatusBadRequest,
            iris.NewProblem().Title("Parser issue").Detail(err.Error()))
        return
    }
    
    // For the sake of the example, log the received payload.
    ctx.Application().Logger().Infof("Received: %#+v", p)
    
    // Send back the payload depending on the accept content type and accept-encoding of the client,
    // e.g. JSON, XML and so on.
    ctx.Negotiate(p)
}

func setAllowedResponses(ctx iris.Context) {
    // Indicate that the Server can send JSON, XML, YAML and MessagePack for this request.
    ctx.Negotiation().JSON().XML().YAML().MsgPack()
    // Add more, allowed by the server format of responses, mime types here...

    // If client is missing an "Accept: " header then default it to JSON.
    ctx.Negotiation().Accept.JSON()
    
    ctx.Next()
}
```

### HTTP2 服务推送 

HTTP2 server push 

Full example code can be found at[_examples/response-writer/http2push](https://github.com/kataras/iris/tree/master/_examples/response-writer/http2push).
完整的示例代码可以在 _ examples/response-writer/http2push 中找到。

Server push lets the server preemptively "push" website assets to the client without the user having explicitly asked for them. When used with care, we can send what we know the user is going to need for the page they're requesting.
服务器推送可以让服务器先发制人地将 website assets “推送”到客户端，而用户没有明确要求这些资产。当小心使用时，我们可以发送我们知道的用户将要需要并请求的页面。

```go
package main

import (
    "net/http"

    "github.com/kataras/iris/v12"
)

func main() {
    app := iris.New()
    app.Get("/", pushHandler)
    app.Get("/main.js", simpleAssetHandler)

    app.Run(iris.TLS("127.0.0.1:443", "mycert.crt", "mykey.key"))
    // $ openssl req -new -newkey rsa:4096 -x509 -sha256 \
    // -days 365 -nodes -out mycert.crt -keyout mykey.key
}

func pushHandler(ctx iris.Context) {
    // The target must either be an absolute path (like "/path") or an absolute
    // URL that contains a valid host and the same scheme as the parent request.
    // If the target is a path, it will inherit the scheme and host of the
    // parent request.
    target := "/main.js"

    if pusher, ok := ctx.ResponseWriter().Naive().(http.Pusher); ok {
        err := pusher.Push(target, nil)
        if err != nil {
            if err == iris.ErrPushNotSupported {
                ctx.StopWithText(iris.StatusHTTPVersionNotSupported, "HTTP/2 push not supported.")
            } else {
                ctx.StopWithError(iris.StatusInternalServerError, err)
            }
            return
        }
    }
    
    ctx.HTML(`<html><body><script src="%s"></script></body></html>`, target)
}

func simpleAssetHandler(ctx iris.Context) {
    ctx.ServeFile("./public/main.js")
}
```

##设置并得到一个cookie 

Set and get a cookie 

Secure cookies, encoding and decoding, sessions (and sessions scaling), flash messages and more can be found at the [_examples/cookies](https://github.com/kataras/iris/tree/master/_examples/cookies) and [_examples/sessions](https://github.com/kataras/iris/tree/master/_examples/sessions) directories respectfully.
安全 cookie、编码和解码、会话(和会话缩放)、 flash 消息等等都可以在 _ examples/cookies 和 _ examples/session 目录中找到。

```go
import "github.com/kataras/iris/v12"

func main() {
    app := iris.Default()

    app.Get("/cookie", func(ctx iris.Context) {
        value := ctx.GetCookie("my_cookie")
    
        if value == "" {
            value = "NotSet"
            ctx.SetCookieKV("my_cookie", value)
            // Alternatively: ctx.SetCookie(&http.Cookie{...})
            ctx.SetCookie("", "test", 3600, "/", "localhost", false, true)
        }
    
        ctx.Writef("Cookie value: %s \n", cookie)
    })
    
    app.Listen(":8080")
}
```

If you want to set custom the path:
如果你想自定义路径:

```go
ctx.SetCookieKV(name,value,iris.CookiePath("/custom/path/cookie/will/be/stored"))
```

If you want to be visible only to current request path:
如果您希望仅对当前请求路径可见:

```go
ctx.SetCookieKV(name, value, iris.CookieCleanPath /* or iris.CookiePath("") */)
```

More:

- ```iris.CookieAllowReclaim```
- ```iris.CookieAllowSubdomains```
- ```iris.CookieSecure```
- ```iris.CookieHTTPOnly```
- ```iris.CookieSameSite```
- ```iris.CookiePath```
- ```iris.CookieCleanPath```
- ```iris.CookieExpires```
- ```iris.CookieEncoding```

You can add cookie options for the whole request in a middleware too:
你也可以在中间件中为整个请求添加 cookie 选项:

```go
func setCookieOptions(ctx iris.Context) {
    ctx.AddCookieOptions(iris.CookieHTTPOnly(true), iris.CookieExpires(1*time.Hour))
    ctx.Next()
}
```

## JSON Web 令牌

JSON Web Tokens 

JSON Web Token (JWT) is an open standard (RFC 7519) that defines a compact and self-contained way for securely transmitting information between parties as a JSON object. This information can be verified and trusted because it is digitally signed. JWTs can be signed using a secret (with the HMAC algorithm) or a public/private key pair using RSA or ECDSA.
JSON Web 令牌(JWT)是一种开放标准(RFC 7519) ，它定义了一种紧凑和自包含的方式，用于作为 JSON 对象在各方之间安全地传输信息。可以验证和信任此信息，因为它是数字签名的。JWTs 可以使用 secret (使用 HMAC 算法)或使用 RSA 或 ECDSA 的公钥/私钥对进行签名。

When should you use JSON Web Tokens? 
什么时候应该使用 JSON Web 令牌？

Here are some scenarios where JSON Web Tokens are useful:
下面是一些 JSON Web 令牌有用的场景:

Authorization: This is the most common scenario for using JWT. Once the user is logged in, each subsequent request will include the JWT, allowing the user to access routes, services, and resources that are permitted with that token. Single Sign On is a feature that widely uses JWT nowadays, because of its small overhead and its ability to be easily used across different domains.
授权: 这是使用 JWT 最常见的场景。一旦用户登录，每个后续请求都将包含 JWT，允许用户访问该令牌所允许的路由、服务和资源。单点登录(Single Sign On)是目前广泛使用 JWT 的一个特性，因为它的开销很小，而且可以轻松地跨不同的域使用。

Information Exchange: JSON Web Tokens are a good way of securely transmitting information between parties. Because JWTs can be signed—for example, using public/private key pairs—you can be sure the senders are who they say they are. Additionally, as the signature is calculated using the header and the payload, you can also verify that the content hasn't been tampered with.
信息交换: JSON Web 令牌是在各方之间安全地传输信息的好方法。因为可以对 JWTs 进行签名(例如，使用公钥/私钥对) ，所以可以确定发送者就是他们所说的那个人。此外，由于签名是使用标头和有效负载计算的，因此还可以验证内容没有被篡改。

!> Read more about JWT at: https://jwt.io/introduction/

更多关于智威汤逊的信息请点击: https://JWT.io/introduction/

## 在 Iris 使用 JWT  

Using JWT with Iris 

The Iris JWT [middleware](https://github.com/kataras/iris/tree/master/middleware/jwt) was designed with security, performance and simplicity in mind, it protects your tokens from [critical vulnerabilities that you may find in other libraries](https://auth0.com/blog/critical-vulnerabilities-in-json-web-token-libraries/). It is based on [kataras/jwt](https://github.com/kataras/jwt) package.
Iris JWT 中间件在设计时考虑到了安全性、性能和简单性，它保护令牌不受其他库中可能存在的关键漏洞的影响。它基于 kataras/jwt 包。

```go
package main

import (
    "time"

    "github.com/kataras/iris/v12"
    "github.com/kataras/iris/v12/middleware/jwt"
)

var (
    sigKey = []byte("signature_hmac_secret_shared_key")
    encKey = []byte("GCM_AES_256_secret_shared_key_32")
)

type fooClaims struct {
    Foo string `json:"foo"`
}

func main() {
    app := iris.New()

    signer := jwt.NewSigner(jwt.HS256, sigKey, 10*time.Minute)
    // Enable payload encryption with:
    // signer.WithEncryption(encKey, nil)
    app.Get("/", generateToken(signer))
    
    verifier := jwt.NewVerifier(jwt.HS256, sigKey)
    // Enable server-side token block feature (even before its expiration time):
    verifier.WithDefaultBlocklist()
    // Enable payload decryption with:
    // verifier.WithDecryption(encKey, nil)
    verifyMiddleware := verifier.Verify(func() interface{} {
        return new(fooClaims)
    })
    
    protectedAPI := app.Party("/protected")
    // Register the verify middleware to allow access only to authorized clients.
    protectedAPI.Use(verifyMiddleware)
    // ^ or UseRouter(verifyMiddleware) to disallow unauthorized http error handlers too.
    
    protectedAPI.Get("/", protected)
    // Invalidate the token through server-side, even if it's not expired yet.
    protectedAPI.Get("/logout", logout)
    
    // http://localhost:8080
    // http://localhost:8080/protected?token=$token (or Authorization: Bearer $token)
    // http://localhost:8080/protected/logout?token=$token
    // http://localhost:8080/protected?token=$token (401)
    app.Listen(":8080")
}

func generateToken(signer *jwt.Signer) iris.Handler {
    return func(ctx iris.Context) {
        claims := fooClaims{Foo: "bar"}

        token, err := signer.Sign(claims)
        if err != nil {
            ctx.StopWithStatus(iris.StatusInternalServerError)
            return
        }
    
        ctx.Write(token)
    }
}

func protected(ctx iris.Context) {
    // Get the verified and decoded claims.
    claims := jwt.Get(ctx).(*fooClaims)

    // Optionally, get token information if you want to work with them.
    // Just an example on how you can retrieve all the standard claims (set by signer's max age, "exp").
    standardClaims := jwt.GetVerifiedToken(ctx).StandardClaims
    expiresAtString := standardClaims.ExpiresAt().Format(ctx.Application().ConfigurationReadOnly().GetTimeFormat())
    timeLeft := standardClaims.Timeleft()
    
    ctx.Writef("foo=%s\nexpires at: %s\ntime left: %s\n", claims.Foo, expiresAtString, timeLeft)
}

func logout(ctx iris.Context) {
    err := ctx.Logout()
    if err != nil {
        ctx.WriteString(err.Error())
    } else {
        ctx.Writef("token invalidated, a new token is required to access the protected API")
    }
}
```

!> Learn about refresh tokens, blocklist and more at: [_examples/auth/jwt](https://github.com/kataras/iris/tree/master/_examples/auth/jwt).

在: _ examples/auth/jwt 了解关于刷新标记、 blocklist 和更多信息。

## Testing

Iris offers an incredible support for the [httpexpect](https://github.com/gavv/httpexpect), a Testing Framework for web applications. The ```iris/httptest``` subpackage provides helpers for Iris + httpexpect.
Iris 为 httppexpect 提供了令人难以置信的支持，一个 web 应用程序的测试框架。Iris/httptest 子包为 Iris + httppexpect 提供了助手。

if you prefer the Go's standard [net/http/httptest](https://golang.org/pkg/net/http/httptest/) package, you can still use it. Iris as much every http web framework is compatible with any external tool for testing, at the end it's HTTP.
如果你喜欢 Go 的标准 net/http/httptest 包，你仍然可以使用它。尽管每个 HTTP web 框架都可以兼容任何外部测试工具，但最终它还是采用了 HTTP。

### Testing Basic Authentication 

测试基本认证

In our first example we will use the iris/httptest subpackage to test Basic Authentication.
在我们的第一个示例中，我们将使用 iris/httptest 子包来测试基本身份验证。

**1.** The ```main.go``` source file looks like that:

```go
package main

import (
    "github.com/kataras/iris/v12"
    "github.com/kataras/iris/v12/middleware/basicauth"
)

func newApp() *iris.Application {
    app := iris.New()

    opts := basicauth.Options{
        Allow: basicauth.AllowUsers(map[string]string{"myusername": "mypassword"}),
    }
    
    authentication := basicauth.New(opts) // or just: basicauth.Default(map...)
    
    app.Get("/", func(ctx iris.Context) { ctx.Redirect("/admin") })
    
    // to party
    
    needAuth := app.Party("/admin", authentication)
    {
        //http://localhost:8080/admin
        needAuth.Get("/", h)
        // http://localhost:8080/admin/profile
        needAuth.Get("/profile", h)
    
        // http://localhost:8080/admin/settings
        needAuth.Get("/settings", h)
    }
    
    return app
}

func h(ctx iris.Context) {
    // username, password, _ := ctx.Request().BasicAuth()
    // third parameter it will be always true because the middleware
    // makes sure for that, otherwise this handler will not be executed.
    // OR:

    user := ctx.User().(*iris.SimpleUser)
    ctx.Writef("%s %s:%s", ctx.Path(), user.Username, user.Password)
    // ctx.Writef("%s %s:%s", ctx.Path(), username, password)
}

func main() {
    app := newApp()
    app.Listen(":8080")
}
```

**2.** Now, create a ```main_test.go``` file and copy-paste the following.
现在，创建一个 ```main_test.go``` 文件并复制粘贴以下内容。

```go
package main

import (
    "testing"

    "github.com/kataras/iris/v12/httptest"
)

func TestNewApp(t *testing.T) {
    app := newApp()
    e := httptest.New(t, app)

    // redirects to /admin without basic auth
    e.GET("/").Expect().Status(httptest.StatusUnauthorized)
    // without basic auth
    e.GET("/admin").Expect().Status(httptest.StatusUnauthorized)
    
    // with valid basic auth
    e.GET("/admin").WithBasicAuth("myusername", "mypassword").Expect().
        Status(httptest.StatusOK).Body().Equal("/admin myusername:mypassword")
    e.GET("/admin/profile").WithBasicAuth("myusername", "mypassword").Expect().
        Status(httptest.StatusOK).Body().Equal("/admin/profile myusername:mypassword")
    e.GET("/admin/settings").WithBasicAuth("myusername", "mypassword").Expect().
        Status(httptest.StatusOK).Body().Equal("/admin/settings myusername:mypassword")
    
    // with invalid basic auth
    e.GET("/admin/settings").WithBasicAuth("invalidusername", "invalidpassword").
        Expect().Status(httptest.StatusUnauthorized)

}
```

**3.** Open your command line and execute:
打开命令行并执行:

```bash
$ go test -v
```

### Testing Cookies 

测试 Cookies

```go
package main

import (
    "fmt"
    "testing"

    "github.com/kataras/iris/v12/httptest"
)

func TestCookiesBasic(t *testing.T) {
    app := newApp()
    e := httptest.New(t, app, httptest.URL("http://example.com"))

    cookieName, cookieValue := "my_cookie_name", "my_cookie_value"
    
    // Test Set A Cookie.
    t1 := e.GET(fmt.Sprintf("/cookies/%s/%s", cookieName, cookieValue)).
        Expect().Status(httptest.StatusOK)
    // Validate cookie's existence, it should be available now.
    t1.Cookie(cookieName).Value().Equal(cookieValue)
    t1.Body().Contains(cookieValue)
    
    path := fmt.Sprintf("/cookies/%s", cookieName)
    
    // Test Retrieve A Cookie.
    t2 := e.GET(path).Expect().Status(httptest.StatusOK)
    t2.Body().Equal(cookieValue)
    
    // Test Remove A Cookie.
    t3 := e.DELETE(path).Expect().Status(httptest.StatusOK)
    t3.Body().Contains(cookieName)
    
    t4 := e.GET(path).Expect().Status(httptest.StatusOK)
    t4.Cookies().Empty()
    t4.Body().Empty()
}
```

```bash
$ go test -v -run=TestCookiesBasic$
```

Iris web 框架本身使用这个包来测试自己。在 _ examples 资源库目录中，您还可以找到一些有用的测试。欲了解更多信息，请查看并阅读 httpexpect 的文档。
Iris web framework itself uses this package to test itself. In the [_examples repository](https://github.com/kataras/iris/tree/master/_examples) directory you will find some useful tests as well. For more information please take a look and read the [httpexpect's documentation](https://github.com/gavv/httpexpect).


## 本地化

Localization 

Introduction 
引言

Localization features provide a convenient way to retrieve strings in various languages, allowing you to easily support multiple languages within your application. Language strings are stored in files within the ```./locales``` directory. Within this directory there should be a subdirectory for each language supported by the application:
本地化特性为检索各种语言的字符串提供了一种方便的方法，使您能够轻松地在应用程序中支持多种语言。语言字符串存储。```./locales``` 目录。在这个目录中应该有一个应用程序支持的每种语言的子目录:

```bash
│   main.go
└───locales
    ├───el-GR
    │       home.yml
    ├───en-US
    │       home.yml
    └───zh-CN
            home.yml
```

The default language for your application is the first registered language.
应用程序的默认语言是第一个注册的语言。

```go
app := iris.New()

// First parameter: Glob filpath patern,
// Second variadic parameter: Optional language tags,
// the first one is the default/fallback one.
app.I18n.Load("./locales/*/*", "en-US", "el-GR", "zh-CN")
```

Or if you load all languages by:
或者如果你通过以下方式加载所有语言:

```go
app.I18n.Load("./locales/*/*")
// Then set the default language using:
app.I18n.SetDefault("en-US")
```

### 加载嵌入的区域设置

Load embedded locales 

You may want to embed locales with a go-bindata tool within your application executable.
您可能希望在应用程序可执行文件中使用 go-bindata 工具嵌入区域设置。

**1.** install a go-bindata tool, e.g.
安装一个 go-bindata 工具，例如。

```$ go get -u github.com/go-bindata/go-bindata/...```



**2.** embed local files to your application
在应用程序中嵌入本地文件

```$ go-bindata -o locales.go ./locales/...```

**3.** use the ```LoadAssets``` method to initialize and load the languages
使用 ```LoadAssets``` 方法初始化和加载这些语言

^ The ```AssetNames``` and ```Asset``` functions are generated by ```go-bindata```
AssetNames 和 Asset 函数由 go-bindata 生成

```go
ap.I18n.LoadAssets(AssetNames, Asset, "en-US", "el-GR", "zh-CN")
```

## 定义翻译

Defining Translations 

Locale files can be written at YAML(recommended), JSON, TOML or INI form.
语言环境文件可以以 YAML (推荐)、 JSON、 TOML 或 INI 的形式编写。

Each file should contain keys. Keys can have sub-keys(we call them "sections") too.
每个文件都应该包含键。键也可以有子键(我们称之为“部分”)。

Each key's value should be of form ```string``` or ```map``` containing by its translated text (or template) or/and its pluralized key-values.
每个键的值应该是由其翻译文本(或模板)或/及其多元化键值所包含的表单```string```或```map```。

Iris i18n module supports pluralization out-of-the-box, see below.
Iris i18n 模块支持开箱即用的多元化，如下所示。

### Fmt Style

```yaml
hi: "Hi %s!"
```

```go
ctx.Tr("Hi", "John")
// Outputs: Hi John!
```

### Template 

```yaml
hi: "Hi {{.Name}}!"
```

```go
ctx.Tr("Hi", iris.Map{"Name": "John"})
// Outputs: Hi John!
```

### 多元化

Pluralization

Iris i18n supports plural variables. To define a per-locale variable you must define a new section of Vars key.
Iris i18n 支持多个变量。要定义每个区域设置变量，必须定义 Vars 键的新部分。

The acceptable keys for variables are:
可接受的变量键是:

- ```one```
- ```"=x"``` where x is a number 
- ```"<x"```
- ```other```
- ```format```

例子:

```yaml
Vars:
  - Minutes:
      one: "minute"
      other: "minutes"
  - Houses:
      one: "house"
      other: "houses"
```

Then, each message can use this variable, here's how:
然后，每条消息都可以使用这个变量，方法如下:

```yaml
# Using variables in raw string
YouLate: "You are %[1]d ${Minutes} late."
# [x] is the argument position,
# variables always have priority other fmt-style arguments,
# that's why we see [1] for houses and [2] for the string argument.
HouseCount: "%[2]s has %[1]d ${Houses}."
Copy to clipboardErrorCopied
ctx.Tr("YouLate", 1)
// Outputs: You are 1 minute late.
ctx.Tr("YouLate", 10)
// Outputs: You are 10 minutes late.

ctx.Tr("HouseCount", 2, "John")
// Outputs: John has 2 houses.
```

You can select what message will be shown based on a given plural count.
您可以根据给定的复数计数选择要显示的消息。

Except variables, each message can also have its plural form too!
除了变量，每个消息也可以有它的复数形式！

Acceptable keys:
可接受的keys:

- ```zero```
- ```one```
- ```two```
- ```"=x"```
- ```"<x"```
- ```">x"```
- ```other```

Let's create a simple plural-featured message, it can use the Minutes variable we created above too.
让我们创建一个简单的复数形式的消息，它也可以使用我们上面创建的 Minutes 变量。

```yaml
FreeDay:
  "=3": "You have three days and %[2]d ${Minutes} off." # "FreeDay" 3, 15
  one:  "You have a day off." # "FreeDay", 1
  other: "You have %[1]d free days." # "FreeDay", 5
```

```go
ctx.Tr("FreeDay", 3, 15)
// Outputs: You have three days and 15 minutes off.
ctx.Tr("FreeDay", 1)
// Outputs: You have a day off.
ctx.Tr("FreeDay", 5)
// Outputs: You have 5 free days.
```

Let's continue with a bit more advanced example, using template text + functions + plural + variables.
让我们继续使用一个更高级的例子，使用模板文本 + 函数 + 复数 + 变量。

```yaml
Vars:
  - Houses:
      one: "house"
      other: "houses"
  - Gender:
      "=1": "She"
      "=2": "He"

VarTemplatePlural:
  one: "${Gender} is awesome!"
  other: "other (${Gender}) has %[3]d ${Houses}."
  "=5": "{{call .InlineJoin .Names}} are awesome."
```

```go
const (
    female = iota + 1
    male
)

ctx.Tr("VarTemplatePlural", iris.Map{
    "PluralCount": 5,
    "Names":       []string{"John", "Peter"},
    "InlineJoin": func(arr []string) string {
        return strings.Join(arr, ", ")
    },
})
// Outputs: John, Peter are awesome

ctx.Tr("VarTemplatePlural", 1, female)
// Outputs: She is awesome!

ctx.Tr("VarTemplatePlural", 2, female, 5)
// Outputs: other (She) has 5 houses.
```

### Sections 

If the key is not a reserved one (e.g. one, two...) then it acts as a sub section. The sections are separated by dot characters (.).
如果这个密钥不是一个保留的密钥(例如一个、两个...) ，那么它就是一个子部分。这些部分由点字符(.)分隔.

```yaml
Welcome:
  Message: "Welcome {{.Name}}"
```

```go
ctx.Tr("Welcome.Message", iris.Map{"Name": "John"})
// Outputs: Welcome John
```

### 确定当前位置

Determining The Current Locale 


You may use the ```context.GetLocale``` method to determine the current locale or check if the locale is a given value:
你可以利用```context.GetLocale```  方法来确定当前的 locale 或者检查 locale 是否是一个给定的值:

```go
func(ctx iris.Context) {
    locale := ctx.GetLocale()
    // [...]
}
```

The **Locale** interface looks like this.
Locale 接口如下所示。

```go
// Locale is the interface which returns from a `Localizer.GetLocale` metod.
// It serves the transltions based on "key" or format. See `GetMessage`.
type Locale interface {
    // Index returns the current locale index from the languages list.
    Index() int
    // Tag returns the full language Tag attached tothis Locale,
    // it should be uniue across different Locales.
    Tag() *language.Tag
    // Language should return the exact languagecode of this `Locale`
    //that the user provided on `New` function.
    //
    // Same as `Tag().String()` but it's static.
    Language() string
    // GetMessage should return translated text based n the given "key".
    GetMessage(key string, args ...interface{}) string
}
```

### 检索翻译

Retrieving Translation 


Use of ```context.Tr``` method as a shortcut to get a translated text for this request.
使用 ```context.Tr``` 方法作为获取此请求的翻译文本的快捷方式。

```go
func(ctx iris.Context) {
    text := ctx.Tr("hi", "name")
    // [...]
}
```

### 内部视图

Inside Views 

```go
func(ctx iris.Context) {
    ctx.View("index.html", iris.Map{
        "tr": ctx.Tr,
    })
}
```

### [Example](https://github.com/kataras/iris/tree/master/_examples/i18n)

```go
package main

import (
    "github.com/kataras/iris/v12"
)

func newApp() *iris.Application {
    app := iris.New()

    // Configure i18n.
    // First parameter: Glob filpath patern,
    // Second variadic parameter: Optional language tags, the first one is the default/fallback one.
    app.I18n.Load("./locales/*/*.ini", "en-US", "el-GR", "zh-CN")
    // app.I18n.LoadAssets for go-bindata.
    
    // Default values:
    // app.I18n.URLParameter = "lang"
    // app.I18n.Subdomain = true
    //
    // Set to false to disallow path (local) redirects,
    // see https://github.com/kataras/iris/issues/1369.
    // app.I18n.PathRedirect = true
    
    app.Get("/", func(ctx iris.Context) {
        hi := ctx.Tr("hi", "iris")
    
        locale := ctx.GetLocale()
    
        ctx.Writef("From the language %s translated output: %s", locale.Language(), hi)
    })
    
    app.Get("/some-path", func(ctx iris.Context) {
        ctx.Writef("%s", ctx.Tr("hi", "iris"))
    })
    
    app.Get("/other", func(ctx iris.Context) {
        language := ctx.GetLocale().Language()
    
        fromFirstFileValue := ctx.Tr("key1")
        fromSecondFileValue := ctx.Tr("key2")
        ctx.Writef("From the language: %s, translated output:\n%s=%s\n%s=%s",
            language, "key1", fromFirstFileValue,
            "key2", fromSecondFileValue)
    })
    
    // using in inside your views:
    view := iris.HTML("./views", ".html")
    app.RegisterView(view)
    
    app.Get("/templates", func(ctx iris.Context) {
        ctx.View("index.html", iris.Map{
            "tr": ctx.Tr, // word, arguments... {call .tr "hi" "iris"}}
        })
    
        // Note that,
        // Iris automatically adds a "tr" global template function as well,
        // the only difference is the way you call it inside your templates and
        // that it accepts a language code as its first argument.
    })
    //
    
    return app
}

func main() {
    app := newApp()

    // go to http://localhost:8080/el-gr/some-path
    // ^ (by path prefix)
    //
    // or http://el.mydomain.com8080/some-path
    // ^ (by subdomain - test locally with the hosts file)
    //
    // or http://localhost:8080/zh-CN/templates
    // ^ (by path prefix with uppercase)
    //
    // or http://localhost:8080/some-path?lang=el-GR
    // ^ (by url parameter)
    //
    // or http://localhost:8080 (default is en-US)
    // or http://localhost:8080/?lang=zh-CN
    //
    // go to http://localhost:8080/other?lang=el-GR
    // or http://localhost:8080/other (default is en-US)
    // or http://localhost:8080/other?lang=en-US
    //
    // or use cookies to set the language.
    app.Listen(":8080", iris.WithSitemap("http://localhost:8080"))
}
```

## 网站地图

Sitemap 

Sitemap translations are automatically set to each route by path prefix if ```app.I18n.PathRedirect``` is true or by subdomain if ```app.I18n.Subdomain``` is true or by URL query parameter if ```app.I18n.URLParameter``` is not empty.
Sitemap 翻译为每条路由自动设置按**路径前缀**如果 ```app.I18n.PathRedirect```为 true或**按子域**如果 ```app.I18n.Subdomain``` 为 true或通过 **URL 查询参数**如果 ```app.I18n.URLParameter```不为空。

Read more at: https://support.google.com/webmasters/answer/189077?hl=en
更多信息请点击: https://support.google.com/webmasters/answer/189077?hl=en

```bash
GET http://localhost:8080/sitemap.xml
```

```markup
<?xml version="1.0" encoding="utf-8" standalone="yes"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9" xmlns:xhtml="http://www.w3.org/1999/xhtml">
    <url>
        <loc>http://localhost:8080/</loc>
        <xhtml:link rel="alternate" hreflang="en-US" href="http://localhost:8080/"></xhtml:link>
        <xhtml:link rel="alternate" hreflang="el-GR" href="http://localhost:8080/el-GR/"></xhtml:link>
        <xhtml:link rel="alternate" hreflang="zh-CN" href="http://localhost:8080/zh-CN/"></xhtml:link>
    </url>
    <url>
        <loc>http://localhost:8080/some-path</loc>
        <xhtml:link rel="alternate" hreflang="en-US" href="http://localhost:8080/some-path"></xhtml:link>
        <xhtml:link rel="alternate" hreflang="el-GR" href="http://localhost:8080/el-GR/some-path"></xhtml:link>
        <xhtml:link rel="alternate" hreflang="zh-CN" href="http://localhost:8080/zh-CN/some-path"></xhtml:link>
    </url>
    <url>
        <loc>http://localhost:8080/other</loc>
        <xhtml:link rel="alternate" hreflang="en-US" href="http://localhost:8080/other"></xhtml:link>
        <xhtml:link rel="alternate" hreflang="el-GR" href="http://localhost:8080/el-GR/other"></xhtml:link>
        <xhtml:link rel="alternate" hreflang="zh-CN" href="http://localhost:8080/zh-CN/other"></xhtml:link>
    </url>
    <url>
        <loc>http://localhost:8080/templates</loc>
        <xhtml:link rel="alternate" hreflang="en-US" href="http://localhost:8080/templates"></xhtml:link>
        <xhtml:link rel="alternate" hreflang="el-GR" href="http://localhost:8080/el-GR/templates"></xhtml:link>
        <xhtml:link rel="alternate" hreflang="zh-CN" href="http://localhost:8080/zh-CN/templates"></xhtml:link>
    </url>
</urlset>
```

That's all the basics about Iris. This document covers enough for beginners. Want to become an expert and a Certificated Iris Developer, learn about MVC, i18n, dependency-injection, gRPC, lambda functions, websockets, best practises and more? Request the Iris E-Book today and be participated in the development of Iris!
这就是关于 Iris 的所有基本知识。这份文件对初学者来说已经足够了。想成为一个专家和Iris开发者，学习 MVC，i18n，依赖注入，gRPC，lambda 函数，websockets，最佳实践和更多？获取Iris电子书，并参与开发Iris！