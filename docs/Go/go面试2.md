## 三色算法

在Go中这个算法的官方名称是叫做三色标记清除算法（tricolor mark-and-sweep algorithm）。它可以和程序一起并发工作并且使用写屏障（write barrier）。这就意味着，当Go程序员运行起来，go调度器去负责应用程序的调度，而垃圾回收器会像调度器处理常规应用程序一样，去使用多个goroutines去进行工作。

白色集合:要被清理的
黑色集合:不被清理的
灰色集合:要去扫描的

垃圾回收开始时,全部对象标记为白色,然后垃圾回收器遍历所有根对象(程序能直接访问到的对象,包括全局变量和栈里的东西)并标记为灰色.

然后垃圾回收器开始遍历灰色对象,先将遍历到的灰色对象变成黑色,然后看这个对象有没有指针指向白色集合的对象,然后将指向的白色集合的对象放进灰色集合里.

如果一个灰色对象指针指向了另一个灰色对象,那么会将另一个灰色对象变成黑色 .

垃圾回收器会遍历所有灰色对象,直至没有灰色对象,然后回收白色对象

!> 如果垃圾回收过程中，一个灰色对象在某些情况变为不可达状态，它在那次垃圾回收中就不会被回收了，但是不是说下次也不会回收！尽管这不是最佳情况，但也没有那么糟

Go允许你通过在你的Go代码里放一个runtime.GC()的声明来手动去开启一次垃圾回收。但是，要记住一点，runtime.GC()会阻塞调用器，并且它可能会阻塞整个程序，尤其是如果你想运行一个非常繁忙的而且有很多对象的go程序。这种情况发生主要是因为你不能在其他任何事都在频繁变化的时候去处理垃圾回收，因为这种情况不会给垃圾回收器机会去清楚地确定白色、黑色和灰色集合里的成员。这种垃圾回收状态也被称作是垃圾回收安全点(garbage collection safe-point)。

!> Go垃圾回收器是并发的，更多描述见Mastering Go

## Unsafe code

Unsafe code是一种绕过go类型安全和内存安全检查的Go代码。大多数情况，unsafe code是和指针相关的。但是要记住使用unsafe code有可能会损害你的程序，所以，如果你不完全确定是否需要用到unsafe code就不要使用它。

```go
package main
import (
    "fmt"
    "unsafe"
)
func main() {
    var value int64 = 5
    var p1 = &value
    var p2 = (*int32)(unsafe.Pointer(p1))
```

这里使用了unsafe.Pointer()方法，这个方法能让你创造一个int32的p2指针去指向一个int64的value变量，而这个变量是使用p1指针去访问的，注意这种做法是有风险的。

下面的代码,pointer变量指向array[0]的地址，array[0]是整型数组的第一个元素。接下来指向整数值的pointer变量会传入unsafe.Pointer()方法，然后传入uintptr。最后结果存到了memAddr里

```go
  array := [...]int{0, 1, -2, 3, 4}
  pointer := &array[0]
  fmt.Print(*pointer, " ")
  memAddr := uintptr(unsafe.Pointer(pointer)) + unsafe.Sizeof(array[0])
```

任何go指针都可以转化为unsafe.Pointer指针。

## uintptr 和 unsafe.Pointer的区别

- unsafe.Pointer只是单纯的通用指针类型，用于转换不同类型指针，它不可以参与指针运算；而uintptr是用于指针运算的，GC 不把 uintptr 当指针，也就是说 uintptr 无法持有对象， - uintptr 类型的目标会被回收；
- 看源码注释说明，uintptr是一个整数类型，其大小足以容纳任何指针。指针可以被转换为uintptr，反之不行。
- unsafe.Pointer 可以和 普通指针 进行相互转换；
- unsafe.Pointer 可以和 uintptr 进行相互转换。
- 通常，将uintptr转为Pointer是非法的。uintptr是整数，而不是引用。将Pointer转换为uintptr会创建一个去除指针语义的整数值。即使uintptr拥有某个对象的地址，垃圾回收器也不会在对象移动时更新该uintptr的值，而uintptr也不会阻止所指对象被回收。
- 将uintptr转换为Pointer的唯一合法模式是将一个指针p转换为uintptr，加上一个偏移量，再将其转换回Pointer，uintptr在转回Pointer前，uintptr不能存储在变量中，此模式最常见的用法是访问结构体中字段或数组中的元素。

## defer关键字

defer关键字的作用是当外围函数返回之后才执行被推迟的函数。

defer函数在外围函数返回之后，以后进先出(LIFO)的原则执行。简单点说，在一个外围函数中有3个defer函数：f1()最先出现，然后f2()，最后f3()，当外围函数执行返回之后，f3()最先被执行，接着是f2()，最后是f1()。

第一段代码：

```go
package main
import (
    "fmt"
)
func d1() {
    for i := 3; i > 0; i-- {
        defer fmt.Print(i, " ")
    }
}
```

答案很明显了，输出是1 2 3

第二段代码：

```go
func d2() {
    for i := 3; i > 0; i-- {
        defer func() {
            fmt.Print(i, " ")
        }()
    }
    fmt.Println()
}
```

这里defer执行的是匿名函数，但是没有参数，所有输出的是循环结束后的i值 0 0 0

第三段代码：

```go
func d3() {
    for i := 3; i > 0; i-- {
        defer func(n int) {
            fmt.Print(n, " ")
        }(i)
    }
}
```

这里匿名函数的参数传入了被延迟时 i 的值，所以输出是 1 2 3

都遵循 LIFO 顺序

## Panic 和 Recover

panic()是一个内置的Go函数，它终止Go程序的当前流程并开始panicking！ 另一方面，recover()函数也是一个内置的Go函数，允许你收回那些使用了panic()函数的goroutine的控制权。

```go
package main
import (
    "fmt"
)
func a() {
    fmt.Println("Inside a()")
    defer func() {
        if c := recover(); c != nil {
            fmt.Println("Recover inside a()!")
           }
    }()
    fmt.Println("About to call b()")
    b()
    fmt.Println("b() exited!")
    fmt.Println("Exiting a()")
}
// b()
func b() {
    fmt.Println("Inside b()")
    panic("Panic in b()!")
    fmt.Println("Exiting b()")
}

func main() {
    a()
    fmt.Println("main() ended!")
}
```
```sh
$ go run panicRecover.go
Inside a()
About to call b()
Inside b()
Recover inside a()!
main() ended!
```

可以看到，a没有正常结束，，因为它的最后两个语句没有执行：

```go
fmt.Println("b() exited!")
fmt.Println("Exiting a()")
```
然而，好消息是 panicRecover.go 根据我们的意愿结束而没有 panicking，因为 defer 中使用的匿名函数控制了局面！ 另请注意，函数b对函数a一无所知。 但是，函数a包含处理b函数的panic 情况的 Go 代码！

##


