### 保持对Go的enthusiasm :bulb:

引言：
关于取名，为什么取这么一个俗气的名字呢，因为我是51自学网（手）著（动）名（滑）学（稽）员

enthusiasm definition: a thing that arouses feelings of intense and eager enjoyment.

:bulb::seedling::bike::cn:
目标: 夯实基础，学习更多的Go语言特性

1. 关注gopheracdemy ardanlabs 以及GCTT（Go国内翻译小组挑选的文章）

2. 学习golang wiki里面的东西

3. 每周关注golang weekly news

> Reference:
> 1. [Golang Wiki](https://github.com/golang/go/wiki) 
> 2. [GCTT翻译小组](https://studygolang.com/gctt)
> 3. [gopher academy](https://blog.gopheracademy.com/)
> 4. [Golang Weekly News](https://golangweekly.com/issues)
> 5. [golang-nut google group](https://groups.google.com/forum/#!forum/golang-nuts)
> 6. [Go language specification](https://golang.org/ref/spec)
> 7. [Go recommend blogs](https://github.com/golang/go/wiki/Blogs)
> 8. [Effective Go](https://golang.org/doc/effective_go.html)


学习分为几大块:


1 Go语言并发模型  | 2 GC垃圾回收 |  3 Goroutine调度 | 4 Go semantics | 5 Go编码常犯错误 |  6 Go sdk 源码探究和学习 
:-----------: | :-----------: | :-----------: | :-----------: | :-----------: | :-----------: 

  

- [ ] 不勾选
- [x] 勾选



#### Done

- [x] 1.[函数是第一公民](https://studygolang.com/articles/12789) FirstClassFunction.go tag: semantics
有下列知识点 1. 头等函数作为变量的几种用法 2 高阶函数概念 3 闭包(closure)的变量捕获 外围变量 4 高阶函数在web业务路由Handler功能 5 range和闭包编码常犯的trick 6 Go是值传递，类型分为引用类型(切片 map channel)和值类型(数组)

- [x] 2 [Go的字符串](https://studygolang.com/articles/12261) 1 字符串不可变性质 2 字符串 用 byte char 特别注意Go采用UTF8编码，遇到特殊字符，会有字形和字符串单个字符不一样，这时，需用rune(可以表示任何占据2个字节及以上的单字符)来做处理 3 字符串长度用法辨析ge
   [Go txt normalize](https://blog.golang.org/normalization) [Go string](https://blog.golang.org/strings) 
   其实这2篇文章可能还不如上面这边文章讲的通彻,其实UTF8 unicode ascii 为什么乱码，字形，代码点(unicode code point)搞清楚就可以了解Go采用Utf8 以及rune char byte string类型的关系了
- [x] 3 [Go 指针](https://studygolang.com/articles/12262) 1 指针的表示和解引用(取指针指向的对象数据) 2 指针类型的zero value是零值 ，但是唯独 struct type 没有零值 所以 struct type  不能判断是否nil 参考这2篇文章 Ref: A:https://golang.org/ref/spec#Comparison_operators B: Go各种数据类型的zero value 默认值 https://golang.org/ref/spec#The_zero_value

- [x] 4 [Go方法](https://studygolang.com/articles/12264)  1 为什么要有方法，因为Go不是纯粹的面向对象系统 2 基于类型的方法是一个实现面向OOP的好方法 2 值接收器 ∈ 指针接收器，也就是值接收器的方法，可以用指针接收器。如果是值接收器，那么不能代替，而是Go自动加一个&指针运算符 3 什么时候用值接收器OR指针接收器 参考Go FAQ 章节:Should I define methods on values or pointers?  一般实践原则:用值接收器，除非你想改遍指针的数据  4  结构体成员-->匿名字段 的方法可以直接调用 5 非结构体的方法定义需要和非结构体在同一个包中.
[Blog MethodSet](https://github.com/golang/go/wiki/MethodSets) 这篇文章讲述的内容和上面的差不多,

- [x] 5 [Go 接口①](https://studygolang.com/articles/12266) [Go接口②](https://studygolang.com/articles/12325)
1.接口内部表示(type,value) 2. 接口类型断言和switch选择 3. 实现多个接口 4 接口组合嵌套 reader writer

- [x] 6  [组合](https://studygolang.com/articles/12680) [多态](https://studygolang.com/articles/12681)
Golang的组合和多态玩法,没有继承特性  
 
- [x] 7 [defer](https://studygolang.com/articles/12719) 1 defer用法 入栈, defer 实参的值在函数执行入栈时传入,而非方法正常体结束玩时候在调用(也就是说不在出栈时候调用) 2 入栈遵循Last In First Out 3 defer在sync包 waitgroup的使用
[官方Blog Defer panic recover](Defer, Panic, and Recover
https://blog.golang.org/defer-panic-and-recover) defer的3原则
> 1. A deferred function's arguments are evaluated when the defer statement is evaluated.
> 2. Deferred function calls are executed in Last In First Out order after the surrounding function returns.
> 3. Deferred functions may read and assign to the returning function's named return values.

http://bit.ly/2XyYkzY 

- [x] 8 [panic recover 实践](https://studygolang.com/articles/12785)  总结:函数/方法 发生panic后,不会再执行正常流程,执行完毕所有defer出栈之后,程序控制会转移给调用方,直到当前协程的所有函数退出,程序打印处panic 堆栈信息.可以用recover恢复同一个协程的panic,但是要注意此时的panic信息,要用debug.PrintStack方法打印.

- [x] 9 [Go反射](https://studygolang.com/articles/13178) 什么是反射  反射能获取变量什么信息 类型 类名 struct成员变量各种属性 ,Go反射常用于 Go sql语句组装

- [x] 10 [error 判断3种方法和自定义error来优雅处理err](https://studygolang.com/articles/12784,https://studygolang.com/articles/12724) https://github.com/golang/go/wiki/LearnErrorHandling

- [x] 11 [文件读取](https://studygolang.com/articles/14669) 读取整个文件ioutil 分块读取bus read 逐行读取  buf scanner [文件写入](https://studygolang.com/articles/19475) 并发写

- [x] 12  [Go FAQ](https://golang.org/doc/faq) 读后感: new make区别 一个allocate mem，一个还要继续 initialize ,struct type 不能和nil 做compare 操作，array是值类型。但是slice map channel都是引用类型，不过要强调一点是go都是按值传递 you see is what you want, 另外就是method receiver, where value reciever function, the same is pointer receiver，opposite not work! 需要说明的，如果value receiver 用在pointer receiver上，Go 会自动给她加上&。其他就是Go 语言设计思想了 goroutine 取代thread csp 模型, mutex sync包.取消泛型,同样需要GC。
Go memory model 另外叙述。

- [x] 13 [Err are values]( https://blog.golang.org/errors-are-values) Rob Pike 首先提出观点err are values,并展示一般的处理方法,区别于一般的try-catch-finally,再通过一个再tokyo会上的故事江苏如何优雅处理err的一个实践例子--->errWriter。这个最佳实践也被用在sdk bufio writer archive等包中。

- [x] 14 [Go Slices: usage and internals](https://blog.golang.org/go-slices-usage-and-internals) slice内部meta属性维护指向array关系，slice按需增长,不用担心index out of range。另外，有个possible gotcha 就是 因为slice 指向array. 如果array 过大(通常发生在读取大文件时候)， 这时候可以使用append函数。主要是丢弃array,或者说array-->slice ,slice之间用append合并。

- [x] 15 [JSON and Go](https://blog.golang.org/json-and-go) 讲述 json规范 marshal和unmarshal 以及go如何将复杂的json denote struct type. 
encode原则
>Only data structures that can be represented as valid JSON will be encoded:
>JSON objects only support strings as keys; to encode a Go map type it must be of the form map[string]T (where T is any Go type supported by the json package).
>Channel, complex, and function types cannot be encoded.
>Cyclic data structures are not supported; they will cause Marshal to go into an infinite loop.
>Pointers will be encoded as the values they point to (or 'null' if the pointer is nil).
decode原则
> -->tagName--->fieldName--->case insentive filedName

- [x] 16 [constant](https://blog.golang.org/constants) constant is untyped value with default type which is refered by syntax. constant give  freedom ,对于Numeric 可以不用损失精度 比如PI. 另外叙述了如何表示最大值问题.取反操作(而非类型转换)

- [x] 17 [Gob Go自有序列化协议](https://blog.golang.org/gobs-of-data) 自描述  不需要额外维护字段信息 节省传输空间 ,proto buffer3个misfeature 1 只支持struct类型 不支持primitive type arrary 2 额外维护字段默认值 3字段分optional 和must不够灵活
- [x] 18 [Iota玩法](https://blog.learngoprogramming.com/golang-const-type-enums-iota-bc4befd096d3) [iota spec](https://golang.org/ref/spec#Iota)
iota的基本套路 1 create new Type ---> 2 Declare related constants for new Type  玩法 1 递增  2 递减 3 幂级扩大 4 重置 , iota+1 initialize

- [x] 19 [Go Map玩法](https://blog.golang.org/go-maps-in-action),map常用玩法 , key type 必须可以compare , 然后, 读写锁实现并发安全,不保证遍历顺序.
 

- [x] [defer 配合匿名函数的闭包的几个坑 注意理解defer的3个规则](https://studygolang.com/articles/12061,https://studygolang.com/articles/12136,https://studygolang.com/articles/12319)
[Defer spece](https://golang.org/ref/spec#Defer_statements)
defer function  are executed after any result parameters are set by that return statement but before the function returns to its caller
还有这句话 
Each time a "defer" statement executes, the function value and parameters to the call are evaluated as usual and saved anew but the actual function is not invoked.
[关于defer执行时机的提问](http://bit.ly/2XAXhzM)
名词解释:Function literals is A function literal represents an anonymous function.
- [x] [SignalHandling](https://github.com/golang/go/wiki/SignalHandling)
主要利用chan 阻塞,直到signal 入列.最后调用defer函数,完成资源清理工作.

- [x] [Go单例模式](https://www.ardanlabs.com/blog/2013/07/singleton-design-pattern-in-go.html)
这篇主要讲如何实现goroutine safe的资源load,维护一个私有的变量map,读取只会一次,之后不会改变。没看懂，单例模式和java的单例模式完全概念不同。这篇文章讲的主要是单列模式管理资源如何设计。
- [x] [Effective_go cocurrency chapter](https://golang.org/doc/effective_go.html#concurrency)
原则 share memory by communicate , not communicate by shared memory 
- [x] [Go statements, Channel types, Send statements](https://golang.org/ref/spec#Channel_types)
1. channel的 len cap 和关闭 以及 channel 在何时blocked
2. nil channel , full channel, unbuffered channel will be blocked.
3. Channel type,len dynamic change,cap,close,

channel 有无缓冲. channels of channel ,将每一次请求的数据,数据处理方法和一个存放数据的channel 抽象成一个请求 入列即可.

利用goroutine也可以写出parallelize 程序.

A leaky buffer : 通过一个buffered channel 在client 和 server端进行通信.
 
#### Doing
- [ ] [map 底层原理实现需要看下](https://www.jianshu.com/p/aa0d4808cbb8 https://tiancaiamao.gitbooks.io/go-internals/content/zh/02.3.html)
- [ ] Go语言机制 https://studygolang.com/subject/74
- [ ] 8  并发编程 https://github.com/golang/go/wiki/LearnConcurrency
- [ ] 9 Go读写文件 
- [ ] 10 Go select 关键字 https://studygolang.com/articles/12522
- [ ] 11 Go channel https://studygolang.com/articles/12402
- [ ] 12 Go mutex 和 WaitGroup 用法
- [ ] 13 Go mutex 和 WaitGroup 源码分析
- [ ] 14 GO Ratelimiter https://github.com/golang/go/wiki/RateLimiting
- [ ] 15 go code review https://github.com/golang/go/wiki/CodeReviewComments
- [ ] 16 为什么不支持generic http://bit.ly/2wG5zu8
- [ ] 优秀 article https://github.com/golang/go/wiki/Articles
https://github.com/golang/go/wiki/LearnServerProgramming 服务端编程  如何写中间件


ToDo
[Testable Examples in Go](https://blog.golang.org/examples)


 Go泛型话题讨论 http://bit.ly/2wG5zu8

Miscellaneous
- [x] 1 [GO标准库中文文档](https://studygolang.com/static/pkgdoc/,https://golang.org/pkg/)，这个只能看看库的目录和大致功能，其实没啥用处，直接看英文稳定即可。

- [x] 2 [Go swagger 教程](https://gocn.vip/article/610) Go的swagger 其实代码耦合的挺严重的不好看。
- [x] 3 [interfaceSlice need explicit type conversion](https://github.com/golang/go/wiki/InterfaceSlice，https://golang.org/doc/faq#convert_slice_of_interface) 
- [x] 4 [Go data race检测](https://golang.org/doc/articles/race_detector.html,https://brantou.github.io/2017/05/23/go-race-detector/) 用build tag 排除某些包的检测 参考这个[链接戳我](https://golang.org/pkg/go/build/#hdr-Build_Constraints)
    go build run test -race 用build constraint 去除某些包的data race检测 .rac 会成倍性能损耗的.
- [x] 5 [从其他语言迁移到Go的Gopher需要维持的信仰Or原则](https://blog.rubylearning.com/best-practices-for-a-new-go-developer-8660384302fc)
 >Understand the power of interfaces, they are one of Go’s great gifts, potentially more important than channels or goroutines.
 >If you are coming from another language, be it a dynamic language like Python or Ruby, or a compiled language like Java or C#, leave your OO baggage at the door. Go is an object-oriented language, but it is not a class-based language and does not support inheritance.
 
 >By removing inheritance from the language, the opportunity to practice the mantra of composition over inheritance is made manifest, and fighting it will only lead to frustration.
 
 >If you’re waiting for generic types and functions to be added to the language, my advice is to stop holding your breath and learn to love the language we have today.
 
 >With the 5th point release done, and the 6th on the way in a few months, the possibilities of a new language feature or syntax tweak are remote. The focus from here on out is tools, reliability, and performance.
- [x] [Godoc约定](https://blog.golang.org/godoc-documenting-go-code) Godoc约定 depreciate bug author 第一行展示
- [x] [组织你的GO代码](https://blog.golang.org/organizing-go-code)
- [x] [Go debug](https://blog.golang.org/debugging-what-you-deploy) Goland断掉调试背后就是DELVE
- [x] [Arrays, slices (and strings): The mechanics of 'append'
](https://blog.golang.org/slices) 这篇和DONE列表里面讲的其实差不多,相当于前面几篇汇总 包括copy append make 内建函数 string 额外用法
- [x] [Go data structure](https://research.swtch.com/godata)
- [x] [Go package 管理路程](https://blog.golang.org/versioning-proposal) 可以作为茶歇读物
- [x] [Go syntax](https://blog.golang.org/gos-declaration-syntax) go语法表达为什么和C分风格反过来,易读阿!举C例子证明不易读
- [x] [error-handling-and-go](https://blog.golang.org/error-handling-and-go) 这篇文章降的一般,就是error封装 ,simplize and reduce error handle
- [x] [What happens with closures running as goroutines?](https://golang.org/doc/faq#closures_and_goroutines) 闭包的传参和值引用问题,这个问题好多次了,不想讲了
- [x] [Why doesn't my program run faster with more CPUs?](https://golang.org/doc/faq#parallel_slow) switch context at the cost of more cpu
--->[concurrent is not parallelism](https://blog.golang.org/concurrency-is-not-parallelism)
- [x] [How can I control the number of CPUs? ](https://golang.org/doc/faq#number_cpus) use func GOMAXPROCS(n int) int

//todo
[go memory model](https://golang.org/ref/mem)
https://golang.org/pkg/sync/atomic/
[Rethinking Classical Concurrency Patterns](http://bit.ly/2XDt3vT)

https://blog.golang.org/share-memory-by-communicating

https://golang.org/doc/codewalk/sharemem/

https://gobyexample.com/stateful-goroutines

https://inconshreveable.com/07-08-2014/principles-of-designing-go-apis-with-channels/

https://www.youtube.com/watch?v=3EW1hZ8DVyw

https://blog.golang.org/pipelines

https://blog.golang.org/concurrency-is-not-parallelism

GopherCon 2015: Richard Fliam - A Practical Guide to Preventing Deadlocks and Leaks in Go

https://blog.golang.org/go-concurrency-patterns-timing-out-and

https://blog.golang.org/share-memory-by-communicating

https://talks.golang.org/2012/waza.slide#1

https://talks.golang.org/2013/advconc.slide#1

https://golang.org/ref/spec#Select_statements


https://golang.org/pkg/sync/

https://barrgroup.com/Embedded-Systems/How-To/RTOS-Mutex-Semaphore

https://blog.golang.org/context

https://talks.golang.org/2012/waza.slide#30

https://blog.golang.org/go-at-io-frequently-asked-questions

计算机世界里的定律
https://github.com/nusr/hacker-laws-zh


Debugging what you deploy in Go 1.12
https://blog.golang.org/debugging-what-you-deploy

go interview
https://github.com/goquiz/goquiz.github.io

深入解析go https://tiancaiamao.gitbooks.io/go-internals/content/zh/


https://hackernoon.com/dancing-with-go-s-mutexes-92407ae927bf
https://medium.com/@matryer/stopping-goroutines-golang-1bf28799c1cb
https://medium.com/@cep21/what-accept-interfaces-return-structs-means-in-go-2fe879e25ee8
https://medium.com/@trevor4e/learning-gos-concurrency-through-illustrations-8c4aff603b3
https://www.sohamkamani.com/blog/2018/02/18/golang-data-race-and-how-to-fix-it/
https://guzalexander.com/2013/12/06/golang-channels-tutorial.html
https://medium.com/@cep21/gos-append-is-not-always-thread-safe-a3034db7975
http://go-database-sql.org/overview.html
https://goquiz.github.io/#gc
https://drive.google.com/file/d/1u0be-HVTaT_gz4vvSgkUwvCCrNSTeP17/view?usp=sharing
https://github.com/emluque/golang-internals-resources
//go学习资料 重点推荐
https://github.com/golang/go/wiki/Learn

// go 语言设计思想
https://talks.golang.org/2012/splash.article

https://github.com/avelino/awesome-go 

https://github.com/golang/go/wiki/LearnServerProgramming

https://github.com/golang/go/wiki/GoTalks
https://blog.golang.org/two-recent-go-talks
https://blog.golang.org/two-recent-go-articles
PS: 有些需要编码加深记忆， 有些看看总结。
Go modle教程 https://blog.golang.org/using-go-modules

Go 2.0设计 https://blog.golang.org/toward-go2

https://github.com/golang/go/wiki/LearnErrorHandling


https://blog.learngoprogramming.com/golang-defer-simplified-77d3b2b817ff

go database sql http://go-database-sql.org/

好blog: 


https://blog.golang.org/two-recent-go-talks
https://www.ardanlabs.com/blog/2017/06/for-range-semantics.html

https://www.ardanlabs.com/blog/2014/10/error-handling-in-go-part-i.html

https://www.ardanlabs.com/blog/2014/11/error-handling-in-go-part-ii.html

https://www.ardanlabs.com/blog/2014/12/using-pointers-in-go.html

https://www.ardanlabs.com/blog/2013/09/iterating-over-slices-in-go.html


https://www.ardanlabs.com/blog/2013/06/understanding-defer-panic-and-recover.html

:whale2: :whale2: :whale2: 

接下来计划 学完learn concurrency部分 总结
学Garbage Collector机制 总结
学Gorotutine Scheduler 总结
