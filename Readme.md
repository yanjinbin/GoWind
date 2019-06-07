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

- [x] 5 [Go 接口①](https://studygolang.com/articles/12266) [Go接口②](https://studygolang.com/articles/12325)
1.接口内部表示(type,value) 2. 接口类型断言和switch选择 3. 实现多个接口 4 接口组合嵌套 reader writer

- [x] 6  [组合](https://studygolang.com/articles/12680) [多态](https://studygolang.com/articles/12681)
Golang的组合和多态玩法,没有继承特性  
 
- [x] 7 [defer](https://studygolang.com/articles/12719) 1 defer用法 入栈, defer 实参的值在函数执行入栈时传入,而非方法正常体结束玩时候在调用(也就是说不在出栈时候调用) 2 入栈遵循Last In First Out 3 defer在sync包 waitgroup的使用

- [x] 8 [panic recover 实践](https://studygolang.com/articles/12785)  总结:函数/方法 发生panic后,不会再执行正常流程,执行完毕所有defer出栈之后,程序控制会转移给调用方,直到当前协程的所有函数退出,程序打印处panic 堆栈信息.可以用recover恢复同一个协程的panic,但是要注意此时的panic信息,要用debug.PrintStack方法打印.

- [x] 9 [Go反射](https://studygolang.com/articles/13178) 什么是反射  反射能获取变量什么信息 类型 类名 struct成员变量各种属性 ,Go反射常用于 Go sql语句组装

- [x] 10 [error 判断3种方法和自定义error来优雅处理err](https://studygolang.com/articles/12784,https://studygolang.com/articles/12724) https://github.com/golang/go/wiki/LearnErrorHandling

- [x] 11 [文件读取](https://studygolang.com/articles/14669) 读取整个文件ioutil 分块读取bus read 逐行读取  buf scanner [文件写入](https://studygolang.com/articles/19475) 并发写

- [x] 12  [Go FAQ](https://golang.org/doc/faq) 读后感: new make区别 一个allocate mem，一个还要继续 initialize ,struct type 不能和nil 做compare 操作，array是值类型。但是slice map channel都是引用类型，不过要强调一点是go都是按值传递 you see is what you want, 另外就是method receiver, where value reciever function, the same is pointer receiver，opposite not work! 需要说明的，如果value receiver 用在pointer receiver上，Go 会自动给她加上&。其他就是Go 语言设计思想了 goroutine 取代thread csp 模型, mutex sync包.取消泛型,同样需要GC。
Go memory model 另外叙述。

- [x] 13 [Err are values]( https://blog.golang.org/errors-are-values) Rob Pike 首先提出观点err are values,并展示一般的处理方法,区别于一般的try-catch-finally,再通过一个再tokyo会上的故事江苏如何优雅处理err的一个实践例子--->errWriter。这个最佳实践也被用在sdk bufio writer archive等包中。

#### Doing
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

https://github.com/golang/go/wiki/Articles#concurrency--channels  concurrent 和 channel处理

https://github.com/golang/go/wiki/CommonMistakes

https://github.com/golang/go/wiki/SignalHandling
https://github.com/golang/go/wiki/SimultaneousAssignment
https://github.com/golang/go/wiki/SliceTricks
https://blog.learngoprogramming.com/golang-defer-simplified-77d3b2b817ff
https://blog.rubylearning.com/best-practices-for-a-new-go-developer-8660384302fc

go database sql http://go-database-sql.org/

好blog:https://github.com/golang/go/wiki/Blogs

context https://studygolang.com/articles/13866

ToDo

 Go泛型话题讨论 http://bit.ly/2wG5zu8

Miscellaneous
- [x] 1 [GO标准库中文文档](https://studygolang.com/static/pkgdoc/)，这个只能看看库的目录和大致功能，其实没啥用处，直接看英文稳定即可。

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

//todo
Effective go 学习链接
https://golang.org/doc/effective_go.html


The Go Memory Model
https://golang.org/ref/mem

Advanced Go Concurrency Patterns
https://blog.golang.org/advanced-go-concurrency-patterns

Introducing the Go Race Detector
https://blog.golang.org/race-detector

Share Memory By Communicating
https://blog.golang.org/share-memory-by-communicating
https://golang.org/doc/codewalk/sharemem/

Go Concurrency Patterns: Pipelines and cancellation
https://blog.golang.org/pipelines

Defer, Panic, and Recover
https://blog.golang.org/defer-panic-and-recover

Go Concurrency Patterns: Timing out, moving on
https://blog.golang.org/go-concurrency-patterns-timing-out-and


Getting to Go: The Journey of Go's Garbage Collector
https://blog.golang.org/ismmkeynote


Go Concurrency Patterns: Context
https://blog.golang.org/context

https://blog.golang.org/index

Generating code
https://blog.golang.org/generate

Debugging what you deploy in Go 1.12
https://blog.golang.org/debugging-what-you-deploy

HTTP/2 Server Push
https://blog.golang.org/h2push

Introducing HTTP Tracing
https://blog.golang.org/http-tracing

Go maps in action
https://blog.golang.org/go-maps-in-action

Go Slices: usage and internals
https://blog.golang.org/go-slices-usage-and-internals

JSON and Go
https://blog.golang.org/json-and-go

http://www.itfanr.cc/2017/03/31/learning-golang/

必须的go文章列表
1 go work steal scheduler https://rakyll.org/scheduler/
2 https://golang.org/doc/asm
3 https://studygolang.com/articles/11720#reply2
4 http://tonybai.com/2014/11/15/how-goroutines-work/
5 http://tonybai.com/2017/06/23/an-intro-about-goroutine-scheduler/
6 https://talks.golang.org/2012/waza.slide#6
7 http://tonybai.com/2015/08/25/go-debugging-profiling-optimization/
8 http://tonybai.com/2015/08/27/understanding-go-statements-evaluating-order/
9 http://tonybai.com/2014/11/15/how-goroutines-work/
10 https://gocn.io/article/28
11 http://blog.studygolang.com/2012/12/gdb%E8%B0%83%E8%AF%95go%E7%A8%8B%E5%BA%8F/
12 https://wizardforcel.gitbooks.io/100-gdb-tips/index.html
http://gopl-zh.b0.upaiyun.com/ch5/ch5-10.html

https://segmentfault.com/a/1190000006233292

https://medium.com/@ggiovani/tcp-socket-implementation-on-golang-c38b67c5d8b

----
go wiki 
https://github.com/golang/go/wiki

https://golang.org/ref/spec

//go学习资料 重点推荐
https://github.com/golang/go/wiki/Learn

// go 语言设计思想
https://talks.golang.org/2012/splash.article

// go error handle 
https://github.com/golang/go/wiki/LearnErrorHandling

// common mistakes 
https://github.com/golang/go/wiki/CommonMistakes



----
awsome go

https://github.com/avelino/awesome-go

go talk 
https://github.com/golang/go/wiki/GoTalks


----
advance topic 高级主题

https://github.com/golang/go/wiki/LearnConcurrency



---
https://github.com/golang/go/wiki/MutexOrChannel

---
defer 5个坑

https://studygolang.com/articles/12061

https://deepzz.com/post/how-to-use-defer-in-golang.html

https://my.oschina.net/henrylee2cn/blog/505535

https://xiaozhou.net/something-about-defer-2014-05-25.html#more


学完这个去头条吹比 Go Concurrent
https://github.com/golang/go/wiki/LearnConcurrency

https://github.com/golang/go/wiki/LearnErrorHandling

https://github.com/golang/go/wiki/LearnServerProgramming


https://github.com/golang/go/wiki/CodeTools

https://github.com/golang/go/wiki/Iota

https://github.com/golang/go/wiki/RateLimiting

https://github.com/golang/go/wiki/SliceTricks

https://github.com/golang/go/wiki/SimultaneousAssignment

https://github.com/golang/go/wiki/Timeouts

https://github.com/golang/go/wiki/LockOSThread

https://github.com/golang/go/wiki/MutexOrChannel


https://docs.google.com/document/d/1vrAy9gMpMoS3uaVphB32uVXX4pi-HnNjkMEgyAHX4N4/preview

https://github.com/golang/go/wiki/GoTalks

https://github.com/golang/go/wiki/Articles#concurrency--channels

PS: 有些需要编码加深记忆， 有些看看总结。
:whale2: :whale2: :whale2: 

