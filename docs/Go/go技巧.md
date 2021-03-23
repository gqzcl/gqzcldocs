## go 技巧参考资料
- [go 实效编程](https://go-zh.org/doc/effective_go.html)
- [go by example](https://gobyexample-cn.github.io/)



判断变量类型

#### 方法一

```go
package main

import (
 "fmt"
)

func main() {
	v1 := "abc"
	for i, c := range v1 {
		//%T 必须是Printf
		fmt.Printf("type:%T ", c)
		fmt.Printf("type:%T\n", v1[i])
	}
}
```

output:

>type:int32 type:uint8
type:int32 type:uint8
type:int32 type:uint8


#### 方法二

```go
package main

import (
 "fmt"
 "reflect"
)

func main() {

    v1 := "abc"        
    for i, c := range v1 {
    	//reflect
		fmt.Println("type:", reflect.TypeOf(c) ,reflect.TypeOf(v1[i]))
	}
}
```

output:


>type: int32 uint8
type: int32 uint8
type: int32 uint8

## string 和 int 类型互转

- string 转成 int：

> int, err := strconv.Atoi(string)

- string 转成 int64：

> int64, err := strconv.ParseInt(string, 10, 64)

- int 转成 string：

> string := strconv.Itoa(int)

- int64 转成 string：

> string := strconv.FormatInt(int64,10)

