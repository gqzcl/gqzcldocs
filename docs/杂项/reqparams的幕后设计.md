前面已经讲了 req.query 和 req.body，这篇我们来聊聊 req.params。这次也一样，我们在 request.js 里面也没有找到相关赋值。

先看看 params 到底用在哪里？

当我们访问 /test/1 的时候，可以从 req.params 获取：

首先呢，我们没有直接发现有类似 app.get 和 app.post，但是我们发现了：

这个 methods 是什么呢？

一个第三方的工具包

地址：https://www.npmjs.com/package/methods

我们打印了一下 methods：

这里面都是 http 的方法名，而且都是小写。

熟悉 nodejs 核心包 http 的同学应该都会记得 http.METHODS，我们也打印一下：

返回的都是大写，所以如果要用的化，需要转换一下。

接着我们看 app[method] 内部，我们之前 app.get 传入了 2 个参数：

path 和 fn，所以跳过前面的：调用了lazyrouter

这个我们稍微提一下，之前 app.use 也会调用它，然后来 new Router，到了后面的 app.get 其实已经注入了 this._router，所以跳过，后面我们介绍 app.use 会重点讲。

所以，其实我们慢慢地引到了 router/index.js 里面。

注意：这里的router.handle很重要

先看route方法，先实例化了一下：newRoute

Route 是什么呢？

原来相对目录下有一个 route.js 文件

我们看一下结构：

然后：newLayer

Layer 是什么呢？

原来是相对目录有一个 layer.js：

从结果来看，keys和regexp是设计的关键：

1、id 被取出来，放到 keys 里面

2、出现了以 test 相关的正则表达式，放到 regexp 里面

看一下 Layer 的结构：

然后就是赋值了：handle、name、params、path、regexp

我们重点看一下 regexp 的逻辑：

用到了第三方工具包path-to-regexp把 path 参数传进去

那当有 get   /test/1 的请求过来之后，在 router/index.js

我们看一下 matchLayer 函数：

调用了 layer.match，用上面的正则去匹配 path

匹配结果：



- 发表于: 2019-09-29
- 原文链接：https://kuaibao.qq.com/s/20190929A0E1MR00?refer=cp_1026