### 保持对Go的enthusiasm :bulb:

引言：
关于取名，为什么取这么一个俗气的名字呢，因为我是51自学网（手）著（动）名（滑）学（稽）员
enthusiasm definition: a thing that arouses feelings of intense and eager enjoyment.

:bulb::seedling::bike::cn:
目标: 夯实基础，学习更多的Go语言特性

1. 关注gopheracdemy ardanlabs 以及GCTT（Go国内翻译小组挑选的文章）

2. 学习golang wiki里面的东西

3. 每周关注golang weekly news

4. Golang issue里面也有很多东西, 但是有些写的indirect,

> Reference:
> 1. [Golang Wiki](https://github.com/golang/go/wiki) 
> 2. [GCTT翻译小组](https://studygolang.com/gctt)
> 3. [gopher academy](https://blog.gopheracademy.com/)
> 4. [Golang Weekly News](https://golangweekly.com/issues)
> 5. [golang-nut google group](https://groups.google.com/forum/#!forum/golang-nuts)
> 6. [Go language specification](https://golang.org/ref/spec)
> 7. [Go recommend blogs](https://github.com/golang/go/wiki/Blogs)
> 8. [Effective Go](https://golang.org/doc/effective_go.html)
> 9. [awesome go](https://blog.golang.org/ismmkeynote)

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
Golang的组合和多态玩法,没有继承特性 ,组合实现继承,接口实现多态
 
- [x] 7 [defer](https://studygolang.com/articles/12719) 
 1 defer用法 入栈, defer 实参的值在函数执行入栈时传入,而非方法正常体结束玩时候在调用(也就是说不在出栈时候调用)  2 入栈遵循Last In First Out 3 defer在sync包 WaitGroup的使用[官方示例Blog Defer panic recover](https://blog.golang.org/defer-panic-and-recover) 
  defer对于变量的调用顺序3原则
> 1. A deferred function's arguments are evaluated when the defer statement is evaluated.
> 2. Deferred function calls are executed in Last In First Out order after the surrounding function returns.
> 3. Deferred functions may read and assign to the returning function's named return values.

[defer gotcha](http://bit.ly/2XyYkzY ), [defer spec](https://golang.org/ref/spec#Defer_statements)  注意defer的restriction,也就是说 defer len("aaa") 是defer不能work的,defer没用  
[比较tricky case,什么时候捕获](http://bit.ly/2WK1uzL)

- [x] 8 [panic recover 实践](https://studygolang.com/articles/12785)  总结:函数/方法 发生panic后,不会再执行正常流程,执行完毕所有defer出栈之后,程序控制会转移给调用方,直到当前协程的所有函数退出,程序打印处panic 堆栈信息.可以用recover恢复同一个协程的panic,但是要注意此时的panic信息,要用debug.PrintStack方法打印.

- [x] 9 [Go反射](https://studygolang.com/articles/13178) 什么是反射  反射能获取变量什么信息 类型 类名 struct成员变量各种属性 ,Go反射常用于 Go sql语句组装

- [x] 10 error 判断3种方法和自定义error来优雅处理err[①](https://studygolang.com/articles/12784) [②](https://studygolang.com/articles/12724) 
[err继续学习教程](https://github.com/golang/go/wiki/LearnErrorHandling)

- [x] 11 [文件读取](https://studygolang.com/articles/14669) 读取整个文件ioutil 分块读取bus read 逐行读取  buf scanner [文件写入](https://studygolang.com/articles/19475) 并发写

- [x] 12  [Go FAQ](https://golang.org/doc/faq) 读后感: new make区别 一个allocate mem，一个还要继续 initialize ,struct type 不能和nil 做compare 操作，array是值类型。但是slice map channel都是引用类型，不过要强调一点是go都是按值传递 you see is what you want, 另外就是method receiver, where value reciever function, the same is pointer receiver，opposite not work! 需要说明的，如果value receiver 用在pointer receiver上，Go 会自动给她加上&。其他就是Go 语言设计思想了 goroutine 取代thread csp 模型, mutex sync包.取消泛型,同样需要GC。
Go memory model 另外叙述。

- [x] 13 [Err are values]( https://blog.golang.org/errors-are-values) Rob Pike 首先提出观点err are values,并展示一般的处理方法,区别于一般的try-catch-finally,再通过一个在tokyo会上的故事讲述如何优雅处理err的一个实践例子--->errWriter。这个最佳实践也被用在sdk bufio writer archive等包中。

- [x] 14 [Go Slices: usage and internals](https://blog.golang.org/go-slices-usage-and-internals) slice内部meta属性维护指向array关系，slice按需增长,不用担心index out of range。
另外，有个possible gotcha 就是 因为slice 指向array. 如果array 过大(通常发生在读取大文件时候)， 
这时候可以使用append函数。主要是丢弃array,或者说array-->slice ,slice之间用append合并。

- [x] 15 [JSON and Go](https://blog.golang.org/json-and-go) 讲述 json规范 marshal和unmarshal 以及go如何将复杂的json denote struct type. 
encode原则
>Only data structures that can be represented as valid JSON will be encoded:
>JSON objects only support strings as keys; to encode a Go map type it must be of the form map[string]T (where T is any Go type supported by the json package).
>Channel, complex, and function types cannot be encoded.
>Cyclic data structures are not supported; they will cause Marshal to go into an infinite loop.
>Pointers will be encoded as the values they point to (or 'null' if the pointer is nil).
decode原则
> -->tagName--->fieldName--->case insentive filedName

- [x] 16 [constant](https://blog.golang.org/constants) constant is untyped value with default type which is referred by syntax. constant give freedom ,对于Numeric 可以不用损失精度 比如PI. 另外叙述了如何表示最大值问题.取反操作(而非类型转换)

- [x] 17 [Gob Go自有序列化协议](https://blog.golang.org/gobs-of-data) 自描述  不需要额外维护字段信息 节省传输空间 ,proto buffer3个misfeature 1 只支持struct类型 不支持primitive type arrary 2 额外维护字段默认值 3字段分optional 和must不够灵活
- [x] 18 [Iota玩法](https://blog.learngoprogramming.com/golang-const-type-enums-iota-bc4befd096d3) [iota spec](https://golang.org/ref/spec#Iota)
iota的基本套路 1 create new Type ---> 2 Declare related constants for new Type  玩法 1 递增  2 递减 3 幂级扩大 4 重置 , iota+1 initialize

- [x] 19 [Go Map玩法](https://blog.golang.org/go-maps-in-action),map常用玩法 , key type 必须可以compare , 然后, 读写锁实现并发安全,不保证遍历顺序.

- [x] 20 [Go memory model](https://golang.org/ref/mem) 
1 happen before 规则讲述reorder 满足 某种partial order,那么 can read the desired value when write happening
2 synchronization 规则 
2.1 init()函数的顺序 
2.2  goroutine create happen before created goroutine begin execute
2.3  goroutine exit not guaranteed to happen before any event.if necessary,use lock or channel communication 
3  channel communication 
 Each send on a particular channel is matched to a corresponding receive from that channel, usually in a **different** goroutine. 
 tips: 注意 receive 和 send 是different goroutine

```go
var c = make(chan int, 10)
var a string

func f() {
	a = "hello, world"
	c <- 0
}

func main() {
	go f()
	<-c
	print(a)
}
```


> The closing of a channel happens before a receive that returns a zero value because the channel is closed.

> In the previous example, replacing c <- 0 with close(c) yields a program with the same guaranteed behavior.

高能预警.....

A receive from an unbuffered channel happens before the send on that channel completes.

注意是unbuffered channel哦 

demo如下

```go
var c = make(chan int)
var a string

func f() {
	a = "hello, world"
	<-c
}
func main() {
	go f()
	c <- 0
	print(a)
}
```


那么接下来,对于buffered channel呢,规则如下
The kth receive on a channel with capacity C happens before the k+Cth send from that channel completes


>This rule generalizes the previous rule to buffered channels. It allows a counting semaphore to be modeled by a buffered channel: the number of items in the channel corresponds to the number of active uses, the capacity of the channel corresponds to the maximum number of simultaneous uses, sending an item acquires the semaphore, and receiving an item releases the semaphore. This is a common idiom for limiting concurrency.
 
这段废话讲述的是 可以用buffered channel 构建semaphore 同步工具类.

锁(mutex)和读写锁(mutex)
对于锁,系铃还需解铃人(反过来了),那么这段话可以很好理解了.
> For any sync.Mutex or sync.RWMutex variable l and n < m, call n of l.Unlock() happens before call m of l.Lock() returns.

if m > n, n个unlock happen before 于 lock()

对于读写锁呢, 除了解锁Happen before加锁之外，还要做到读写互斥，写写互斥和写读互斥。



最后是once 起始就是用mutex 和done标记位加上double check实现只做一次function


- [x] [defer 配合匿名函数的闭包的几个坑 注意理解defer的3个规则](https://studygolang.com/articles/12061,https://studygolang.com/articles/12136,https://studygolang.com/articles/12319)
[Defer spec](https://golang.org/ref/spec#Defer_statements)
defer function  are executed after any result parameters are set by that return statement but before the function returns to its caller
还有这句话 
Each time a "defer" statement executes, the function value and parameters to the call are evaluated as usual and saved anew but the actual function is not invoked.
[关于defer执行时机的提问](http://bit.ly/2XAXhzM)
名词解释:Function literals is A function literal represents an anonymous function.
- [x] [SignalHandling](https://github.com/golang/go/wiki/SignalHandling)
主要利用chan 阻塞,直到signal 入列.最后调用defer函数,完成资源清理工作.

- [x] [Once实现单例模式](http://shanks.leanote.com/post/Untitled-55ca439338f41148cd000759-12)
- [x] [Effective_go cocurrency chapter](https://golang.org/doc/effective_go.html#concurrency)
原则 share memory by communicate , not communicate by shared memory 
- [x] [Go statements, Channel types, Send statements](https://golang.org/ref/spec#Channel_types)
1. channel的 len cap 和关闭 以及 channel 在何时blocked
2. nil channel , full channel, unbuffered channel will be blocked.
3. Channel type,len dynamic change,cap,close,

- [x] [????? Go单例模式](https://www.ardanlabs.com/blog/2013/07/singleton-design-pattern-in-go.html)
这篇主要讲如何实现goroutine safe的资源load,维护一个私有的变量map,读取只会一次,之后不会改变。没看懂，单例模式和java的单例模式完全概念不同。这篇文章讲的主要是单列模式管理资源如何设计。
- [x] [right handle nil ](https://golang.org/doc/faq#nil_error)

>nil 必须explicit return, 因为nil的定义是 type和value必须都是nil, 不能type是个类型,而value is not set 
```go
// 错误例子
func returnsError() error {
	var p *MyError = nil
	if bad() {
		p = ErrBad
	}
	return p // Will always return a non-nil error.
}
//正确例子
func returnsError() error {
	if bad() {
		return ErrBad
	}
	return nil // explicit return
}
```

channel 有无缓冲. channels of channel ,将每一次请求的数据,数据处理方法和一个存放数据的channel 抽象成一个请求 入列即可.

利用goroutine也可以写出parallelize 程序.

Miscellaneous
- [x] 1 [GO标准库中文文档](https://studygolang.com/static/pkgdoc/,https://golang.org/pkg/)，这个只能看看库的目录和大致功能，其实没啥用处，直接看英文稳定即可。
- [x] 2 [Go swagger 教程](https://gocn.vip/article/610) Go的swagger 其实代码耦合的挺严重的不好看。
- [x] 3 [interfaceSlice need explicit type conversion](https://github.com/golang/go/wiki/InterfaceSlice，https://golang.org/doc/faq#convert_slice_of_interface) 
- [x] 4 [Go data race检测](https://golang.org/doc/articles/race_detector.html,https://brantou.github.io/2017/05/23/go-race-detector/) 用build tag 排除某些包的检测 参考这个[链接戳我](https://golang.org/pkg/go/build/#hdr-Build_Constraints)
    go build run test -race 用build constraint 去除某些包的data race检测 .rac 会成倍性能损耗的.
- [x] 5 [从其他语言迁移到Go的Gopher需要维持的信仰Or原则](https://blog.rubylearning.com/best-practices-for-a-new-go-developer-8660384302fc)
 > Understand the power of interfaces, they are one of Go’s great gifts, potentially more important than channels or goroutines.
 > If you are coming from another language, be it a dynamic language like Python or Ruby, or a compiled language like Java or C#, leave your OO baggage at the door. Go is an object-oriented language, but it is not a class-based language and does not support inheritance.
 
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
- [x] [error-handling-and-go](https://blog.golang.org/error-handling-and-go) 这篇文章讲的一般,就是error封装 ,simplize and reduce error handle
- [x] [What happens with closures running as goroutines?](https://golang.org/doc/faq#closures_and_goroutines) 闭包的传参和值引用问题,这个问题好多次了,不想讲了
- [x] [Why doesn't my program run faster with more CPUs?](https://golang.org/doc/faq#parallel_slow) switch context at the cost of more cpu
--->[concurrent is not parallelism](https://blog.golang.org/concurrency-is-not-parallelism)
- [x] [How can I control the number of CPUs? ](https://golang.org/doc/faq#number_cpus) use func GOMAXPROCS(n int) int

[Rethinking Classical Concurrency Patterns](http://bit.ly/2XDt3vT)


- [X] [go interview](https://github.com/goquiz/goquiz.github.io) 不错哦 还是有点不错的比如Go启动执行main前干了什么
- [x] [新手中级常犯的错误](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/) 
- [x] [mutex,rwmutex的错误和正确示范](https://hackernoon.com/dancing-with-go-s-mutexes-92407ae927bf) 读写锁,写锁锁住一个变量V的修改,读锁可以对这个变量V加Rlock,Runlock,增加并发度,mutex不像Java一样不是可重入锁,意味着调用内层方法的时候还是需要先解锁的!!

PS: 有些需要编码加深记忆， 有些看看总结。

- [] [go database sql](http://go-database-sql.org/)非常值得一看
=========================================我是长长的分割线=====================================<br>
:whale2: :whale2: :whale2: 

- [x] [Go Map设计细节] [①](https://www.jianshu.com/p/aa0d4808cbb8) [②](https://www.cnblogs.com/qcrao-2018/p/10903807.html) 
> 先说一下设计map的几个key指标 loadFactor=6.5 ,noverflow 判断是否过多的hash collision ,maxKeySize=128 用来编译优化
整体结构是hmap表示map，bmap表示 新桶buckets数组的一个默认8key/value， oldBuckets旧桶，loadFactor*2^B表示可以最多能容纳的元素个数
扩容是在map 涉及put和delete操作,整体扩容可以认为是渐进式的过程，跟redis的map设计的类似
处理冲突是链表法，增加一个bucket节点，每一次collision，都会让noverflow++，用来判断扩容。
扩容的条件是 oldbucket为nil && （ > loadFactor(6.5) || overflow 过多） ，>6.5说明空间利用率过载了，overflow说明空间利用率太低，hash 冲突严重

数组和切片 [①](https://draveness.me/golang-array-and-slice), [②](https://halfrost.com/go_slice/)
parseArrayOrSliceType 解析数组或者slcie
typecheckcomplit 编译器完成静态检查
1. slice的copy函数, fm.len > to.len 要截断
2. slice内部表示
```go
type SliceHeader struct {
Data uintptr
Len int
Cap int
}
```

3. append，slice扩增策略 [growslice()](http://bit.ly/2CwCRPj)
只要扩容就是memmove（copy），生成一个新的数组。
>如果期望容量大于当前容量的两倍就会使用期望容量；
如果当前切片容量小于 1024 就会将容量翻倍；
如果当前切片容量大于 1024 就会每次增加 25% 的容量，直到新容量大于期望容量；

4. slice是延长还是扩容(当从array字面量构建的时候) [点我查看tricky code，辨析slice的扩容本质](https://play.golang.org/p/R1MtaHTxvHA)

- [x] [time递增单调性设计手稿](https://go.googlesource.com/proposal/+/master/design/12914-monotonic.md)
 [davecheney实现的单调时钟](http://bit.ly/32rGiBu)
 
> CLOCK_REALTIME CLOCK_MONOTONIC CLOCK_MONOTONIC_RAW CLOCK_BOOTTIME
CLOCK_REALTIME，就是 「 wall time 」，即是实际的时间。
CLOCK_MONTONIC，即单调时间，即从某个时间点开始到现在过去的时间。
用户不能修改这个时间，但是当系统进入休眠 ( suspend ) 时，CLOCK_MONOTONIC 是不会增加的
CLOCK_MONTONIC 其实就是一个计时器，当电脑开机的时候重新开始计时，当电脑睡眠的时候暂停计时，当电脑关机的时候停止计时
只不过这个计时器采用的不是 秒为单位，或者通常计时器上的 xx.xx ，而是采用的 CPU 时钟

timer计时器 需要 及时stop,停止过期 调用after function,[stoptimer源码](http://bit.ly/32yvd1o) 


#### go调度

- [x] [为什么要采用M:N模型,而不是1:1模型(利用多核,但是上下文切换很慢),N:1(多个user thread可以在1个os thread上快速切换,但是没有利用多核)的原因](https://morsmachine.dk/go-scheduler),就是加了引入一个Process概念 M:P:G = 1:1:N,那么对于M:G来说是1:N模型,对于P:M = 1:1, 不过就是设计的复杂啦

前面解释了原因,下面讲下如何设计这个模型([设计手稿在此](https://morsmachine.dk/go-scheduler)),达到
1去除signal global mutex
2 de-centralized 
3 M 附属的cache 能够 避免1:M模式下,excessive resource consumption. 
4 M之间hand off 传递G,会导致cache失效啊,最好绑定嘛
5 避免worker thread 因 频繁 syscall 造成的block/unblock 

未来可能要设计的: 
1 LIFO 来增强locality,在兼顾fairness情况下 
2 如果G创建了,不给他分配stack,这样创建成本就比较小了 
3 增强P对G和M对P的亲缘性, 
4 引入timer,调节M创建的数量


> [Go的 Work Steal模式](https://rakyll.org/scheduler/)
周期性的从调度计数是61的倍数的话就Global取,否则就无锁的(for cas)从local取,没有的话先从其他P 偷一半g过来,偷不到的话又从Global偷, 再偷不到的话从netPoll偷
最后从other P偷

[源码链接戳我](http://bit.ly/33uD3KW)
```go
// 保留核心的调度逻辑,略过无关的GC和trace代码
func schedule() {
	_g_ := getg() // 获取当前线程

	if _g_.m.locks != 0 {
		throw("schedule: holding locks")
	}

	if _g_.m.lockedg != 0 {
		stoplockedm()
		execute(_g_.m.lockedg.ptr(), false) // Never returns.
	}

	// 略 cgo 相关

top:
	if sched.gcwaiting != 0 {
		gcstopm()
		goto top
	}
	if _g_.m.p.ptr().runSafePointFn != 0 {
		runSafePointFn() // 安全点,关于安全点,你可以理解为 Safe-point means it is a safe suspension point for root set enumeration. 摘录自[李晓峰博客 GC safe point and region](http://bit.ly/34L3qN4)
	}

	var gp *g
	var inheritTime bool
	
    // 略 GC worker相关

	if gp == nil {
		// Check the global runnable queue once in a while to ensure fairness.
		// Otherwise two goroutines can completely occupy the local runqueue
		// by constantly respawning each other.
		if _g_.m.p.ptr().schedtick%61 == 0 && sched.runqsize > 0 {
			lock(&sched.lock)
			gp = globrunqget(_g_.m.p.ptr(), 1) //从 global queue拿
			unlock(&sched.lock)
		}
	}
	if gp == nil {
		gp, inheritTime = runqget(_g_.m.p.ptr()) //从 local queue拿
		if gp != nil && _g_.m.spinning {
			throw("schedule: spinning with local work")
		}
	}
	if gp == nil {
        // 这段也是work steal的核心思想!!,先无锁(for cas)
        // 从 其他P的local queue 偷一半(数目暂时存疑,设计手稿说是一半),
        // 偷不到,加锁,从global queue偷
        // 偷不到就从剥离了P,associated M 上偷
		gp, inheritTime = findrunnable() // blocks until work is available
	}

	// This thread is going to run a goroutine and is not spinning anymore,
	// so if it was marked as spinning we need to reset it now and potentially
	// start a new spinning M.
	if _g_.m.spinning {
		resetspinning() // spin状态重置false
	}
	// 略 移除了GC worker和 Tracer 的调度
	
    // 最终执行
	execute(gp, inheritTime)
}
```
go 关键字(本质上是调用newproc):
还有一个问题要回答?新建的G分配到上面去,这就需要回答go关键字
调用go关键字的G和P有自己的Local Queue,满了就放到Global 等到别人调用
[newproc源码戳](http://bit.ly/34IA3L8)
```go
func newproc(siz int32, fn *funcval) {
	argp := add(unsafe.Pointer(&fn), sys.PtrSize)
	gp := getg()
	pc := getcallerpc()
	systemstack(func() {
		newproc1(fn, (*uint8)(argp), siz, gp, pc)
	})
}

// Create a new g running fn with narg bytes of arguments starting
// at argp. callerpc is the address of the go statement that created
// this. The new g is put on the queue of g's waiting to run.
func newproc1(fn *funcval, argp *uint8, narg int32, callergp *g, callerpc uintptr) {
	_g_ := getg()

	if fn == nil {
		_g_.m.throwing = -1 // do not dump full stacks
		throw("go of nil func value")
	}
	// 
	acquirem() // disable preemption because it can be holding p in a local var
	
    //略 约束条件
	
	_p_ := _g_.m.p.ptr()
	// get gfreelist if not get from g global list
	newg := gfget(_p_)
	if newg == nil {
        //分配一个新G
		newg = malg(_StackMin)
		casgstatus(newg, _Gidle, _Gdead)
		allgadd(newg) // publishes with a g->status of Gdead so GC scanner doesn't look at uninitialized stack.
	}
	
	// 略 newG相关属性变量赋值
	//核心 直接摘录 ,放到 next的更新 
/**
if randomizeScheduler && next && fastrand()%2 == 0 {
		next = false
}
*/
    //// runqput tries to put g on the local runnable queue.
      // If next is false, runqput adds g to the tail of the runnable queue.
      // If next is true, runqput puts g in the _p_.runnext slot.
      // If the run queue is full, runnext puts g on the global queue.
      // Executed only by the owner P.
	runqput(_p_, newg, true)

	if atomic.Load(&sched.npidle) != 0 && atomic.Load(&sched.nmspinning) == 0 && mainStarted {
		wakep()
	}
    //enable preemption
	releasem(_g_.m)
}
```

- [x] [系统线程抢占式调度](https://studygolang.com/articles/14264) 要遇到的问题  cpu缓存  false share (
- [x] [go 在 1 create goroutine2  netpoller(异步调用) 3 GC ,4 sys call 需要做调度决策   5 atomic mutex同步调用](https://studygolang.com/articles/15316)
   
引入P M->P->GQueue 1:1:M, N:M模型. 利用多核以及避免上下文切换,  spin 的 M ,避免 unblock/block ,cpu intense.
先理解[Dmitry Vyukov 调度设计原稿](https://docs.google.com/document/d/1TTj4T2JO42uD5ID9e89oa0sLKhJYD0Y_kqxDv3I3XMw/edit)
>1 single global mutex to protect create complete and reschedule(中心化)

>2 goroutine 在 不同的m(worker thread ) 之间hand off goroutine 带来的系统开销

>3 M的mcache associate with all M, not just a specific M  running go code, 1:100 ,mcache up to 2m比较昂贵, 以及poor data locality

>>4 aggressive thread 频繁的系统调用 unblock/block

新版本 引入 P, 去中心化  以及  让每个m 保持负载, 而不至于IDLE , 通过工作窃取模式  lockOSThread  通过自旋spin G 而不是block/unblock,
让每一个P 和 每一个G  execute on last running P和M上面,那样可以可以cache line


下面2篇,常识快速阅读,没啥细节可以追究
- [x] [ardan lab:go schedule 3篇](https://studygolang.com/articles/14264),主要从模型语义上来阐述
- [x] [Go scheduler:Implementing language with lightweight concurrency](http://bit.ly/2Nur7SC)
<br>============================我是分割线====================

[go逃逸分析,去除编译inline优化,使用原生的,return指针类型,大对象分配的时候会从stack逃逸到heap](https://studygolang.com/articles/12443)
<br>=============================我是分割线===================
####go垃圾回收
1 对象如何分配 stack还是heap上,如果heap上,对象如何分配,tmalloc方法

内存分配规格
----16b---------------------32kb---
---mcache---mcentral------mheap----
stack原生变量,heap 调用返回指针对象,大对象

stack or heap? FAQ
>From a correctness standpoint, you don’t need to know. Each variable in Go exists as long as there are references to it. The storage location chosen by the implementation is irrelevant to the semantics of the language.
 
> The storage location does have an effect on writing efficient programs. When possible, the Go compilers will allocate variables that are local to a function in that function’s stack frame. However, if the compiler cannot prove that the variable is not referenced after the function returns, then the compiler must allocate the variable on the garbage-collected heap to avoid dangling pointer errors. Also, if a local variable is very large, it might make more sense to store it on the heap rather than the stack.
 
> In the current compilers, if a variable has its address taken, that variable is a candidate for allocation on the heap. However, a basic escape analysis recognizes some cases when such variables will not live past the return from the function and can reside on the stack.
 
 mmap申请虚拟地址,分成三级 mcache(唯一不加锁) <16b, mcentral , mheap(>32kb)
 
 回收闲置内存的  mcache 给 central 再给heap ,然后free syscall

2 对象如何进行三色标记  

三色标记 加 [write barrier](https://golang.org/src/runtime/mbarrier.go)
[三色标记过程动态阐释](https://making.pusher.com/golangs-real-time-gc-in-theory-and-practice/)

正如ardan说的一样,GC目前一直在改进,目前是标记清除的 Dijkstra的三色标记,
加上hybird write barrier(处理 goroutine stack上 black object pointer to white object)
这篇总结的比较全面,而且给出了很多有用的超链接,正是我想总结的,[戳我](http://legendtkl.com/2017/04/28/golang-gc/)

4 GC 性能分析工具

- [X] channel底层实现 [①](https://i6448038.github.io/2019/04/11/go-channel/) [②](https://draveness.me/golang/concurrency/golang-channel.html)

#### Go并发
1 支持并发的工具 select channel 
select和channel的实现原理

2 同步原语 mutex, once, RWMutex, WaitGroup, Cond, semaphore, singleFlight
mutex  CAS获取锁 spin (正常锁状态和非饥饿状态下)  ,
  如果饥饿状态下(等待时间超过1ms)或者不能spin,
 通过标记位运算判断去判断,以及是否要清楚饥饿状态
 [RWMutex](http://legendtkl.com/2016/10/23/golang-mutex/)
> 总结的不错:  Mutex 两种工作模式，
normal 正常模式，starvation 饥饿模式。
normal 情况下锁的逻辑与老版相似，休眠的 goroutine 以 FIFO 链表形式保存在 sudog 中，
被唤醒的 goroutine 与新到来活跃的 goroutine 竞解，
但是很可能会失败。如果一个 goroutine 等待超过 1ms，那么 Mutex 进入饥饿模式
饥饿模式下，解锁后，锁直接交给 waiter FIFO 链表的第一个，新来的活跃 goroutine 不参与竞争，并放到 FIFO 队尾
如果当前获得锁的 goroutine 是 FIFO 队尾，或是等待时长小于 1ms，那么退出饥饿模式
normal 模式下性能是比较好的，但是 starvation 模式能减小长尾 latency

[自选条件代码戳我](http://bit.ly/2qEWKRu) ----> 自选次数<4 && p>1 && Local Queue 为空
- [X] [mutex源码解释的也还不错,可以参考](https://colobu.com/2018/12/18/dive-into-sync-mutex/)

[mutex使用场景](https://studygolang.com/articles/12598) 只允许一个G访问临界区比较合适。

once: gmm定义的 Multiple threads can execute once.Do(f) for a particular f, 
but only one will run f(), and the other calls block until f() has returned.

>  // Note: Here is an incorrect implementation of Do:
>  // Note: Here i
```go
if atomic.CompareAndSwapUint32(&o.done, 0, 1) {
 			f()
}
```
>  //  Do guarantees that when it returns, f has finished.
  // This implementation would not implement that guarantee:
  // given two simultaneous calls, the winner of the cas would
  // call f, and the second would return immediately, without
  // waiting for the first's call to f to complete.
  // This is why the slow path falls back to a mutex, and why
  // the atomic.StoreUint32 must be delayed until after f returns


- [X] WaitGroup  [①](https://www.cnblogs.com/jiangz222/p/10348763.html) [②](http://bit.ly/33sNTku)
 

 为什么CSP 1 解耦啊 2 顺序一致性 3 消费者生产者语义明确 4 不需要加锁
 
 
- [x] [cpu 亲缘性](https://www.cnblogs.com/lubinlew/p/cpu_affinity.html),物理CPU,逻辑CPU,进程 bounded to last running Cpu
 
 > 理解顺序: [channel](https://draveness.me/golang-channel)(hcan entity)
// hchan struct

```go
type hchan struct {
	qcount   uint           // total data in the queue
	dataqsiz uint           // size of the circular queue
	buf      unsafe.Pointer // points to an array of dataqsiz elements
	elemsize uint16
	closed   uint32
	elemtype *_type // element type
	sendx    uint   // send index
	recvx    uint   // receive index
	recvq    waitq  // list of recv waiters
	sendq    waitq  // list of send waiters

	// lock protects all fields in hchan, as well as several
	// fields in sudogs blocked on this channel.
	//
	// Do not change another G's status while holding this lock
	// (in particular, do not ready a G), as this can deadlock
	// with stack shrinking.
	lock mutex
}
```
函数调用链: make-->walkexpr(OMAKECHAN)-->makechan-->chansend
makechan 根据size=0? 只分配hcan,不分配缓冲区-->不是指针类型,一起分配,-->hcan和缓冲区分开分配.
chansend (阻塞还是非阻塞很关键,也是实现channel 语义的全部逻辑实现), 
很简单. 对于nil channel, channel buffer和 unbuffer 以及 队列未满 队列满了 
send chan带和不带select都  还是表现出不同的行为的
分开讨论,

1   | 2 阻塞的send chan |  3 非阻塞的 send select |  
:-----------: | :-----------: | :-----------: 
| nil channel | gopark阻塞 | 返回false |
| 是不是阻塞|None| 非阻塞 && chan未close &&  (队列满 OR (队列为空&&(没有接受者)) 直接返回false|
|**加锁阶段**|||
| closed channel  | panic |  panic |
| 如果存在recv G | 直接发给recv,返回true|直接发给recv,返回true|
| 队列未满|enqueue 缓存区,返回true|enqueue 缓存区,返回true|
|队列满了,是否阻塞|G入队,阻塞,KeepAlive(ep)|返回true| 

总结:其实很简单,blocked send channel ,阻塞的情况下,在select下直接返回false即可,没有额外的多余动作

[chansend 完整源码戳我](http://bit.ly/34JiTwP)
```go
/*
 * generic single channel send/recv
 * If block is not nil,
 * then the protocol will not
 * sleep but return if it could
 * not complete.
 *
 * sleep can wake up with g.param == nil
 * when a channel involved in the sleep has
 * been closed.  it is easiest to loop and re-run
 * the operation; we'll see that it's now closed.
 */
func chansend(c *hchan, ep unsafe.Pointer, block bool, callerpc uintptr) bool {
	if c == nil { // nil channel的话,非阻塞 直接返回false,否则阻塞
		if !block {
			return false
		}
		gopark(nil, nil, waitReasonChanSendNilChan, traceEvGoStop, 2)
		throw("unreachable")
	}
    // 略
    
	// Fast path: check for failed non-blocking operation without acquiring the lock.
	//
	// After observing that the channel is not closed, we observe that the channel is
	// not ready for sending. Each of these observations is a single word-sized read
	// (first c.closed and second c.recvq.first or c.qcount depending on kind of channel).
	// Because a closed channel cannot transition from 'ready for sending' to
	// 'not ready for sending', even if the channel is closed between the two observations,
	// they imply a moment between the two when the channel was both not yet closed
	// and not ready for sending. We behave as if we observed the channel at that moment,
	// and report that the send cannot proceed.
	//
	// It is okay if the reads are reordered here: if we observe that the channel is not
	// ready for sending and then observe that it is not closed, that implies that the
	// channel wasn't closed during the first observation.
	if !block && c.closed == 0 && ((c.dataqsiz == 0 && c.recvq.first == nil) ||
		(c.dataqsiz > 0 && c.qcount == c.dataqsiz)) {
		return false
	}
    // 略
    // 加锁
	lock(&c.lock)
	// 调用close channel关闭之后,发送Queue就panic
	if c.closed != 0 {
		unlock(&c.lock)
		panic(plainError("send on closed channel"))
	}
    // 发现有个存在等待的G,发送给他并解锁
	if sg := c.recvq.dequeue(); sg != nil {
		// Found a waiting receiver. We pass the value we want to send
		// directly to the receiver, bypassing the channel buffer (if any).
		send(c, sg, ep, func() { unlock(&c.lock) }, 3)
		return true
	}
    // 队列未满,enqueue elem to send
	if c.qcount < c.dataqsiz {
		// Space is available in the channel buffer. Enqueue the element to send.
		qp := chanbuf(c, c.sendx)
        // copy to dst from src
		typedmemmove(c.elemtype, qp, ep)
		c.sendx++
		if c.sendx == c.dataqsiz {
			c.sendx = 0
		}
		c.qcount++
		unlock(&c.lock)
		return true
	}
    //以上条件都未满足, 如果非阻塞的话,就解锁了
	if !block {
		unlock(&c.lock)
		return false
	}
    // 以上条件都为满足,阻塞的话,就需要将 send G 入列了,并给他赋值相关信息,方便GC和trace
    // 并keep alive 这个 send G发送的消息Ep,等待 recv G

	// Block on the channel. Some receiver will complete our operation for us.
	gp := getg()
	mysg := acquireSudog()
	// No stack splits between assigning elem and enqueuing mysg
	// on gp.waiting where copystack can find it.
	mysg.elem = ep
	mysg.waitlink = nil
	mysg.g = gp
	mysg.isSelect = false
	mysg.c = c
	gp.waiting = mysg
	gp.param = nil
	c.sendq.enqueue(mysg)
	goparkunlock(&c.lock, waitReasonChanSend, traceEvGoBlockSend, 3)
	// Ensure the value being sent is kept alive until the
	// receiver copies it out. The sudog has a pointer to the
	// stack object, but sudogs aren't considered as roots of the
	// stack tracer.
	KeepAlive(ep)

	// someone woke us up.
	// 约束检查了
	return true
}

```

接收端的处理:
> 不管是chanrecv1还是chanrecv2,差了bool表达式,最终还是调用chanrecv
[chanrecv完整源码戳我](http://bit.ly/2pJDgvf)

表格对比暂时不做,就是阻塞的时候他直接返回就是了

```go
// 看这段注释就知道怎么回事了
// chanrecv receives on channel c and writes the received data to ep.
// ep may be nil, in which case received data is ignored.
// If block == false and no elements are available, returns (false, false).
// Otherwise, if c is closed, zeros *ep and returns (true, false).
// Otherwise, fills in *ep with an element and returns (true, true).
// A non-nil ep must point to the heap or the caller's stack.
func chanrecv(c *hchan, ep unsafe.Pointer, block bool) (selected, received bool) {
	if c == nil {
		if !block {
			return
		}
		gopark(nil, nil, waitReasonChanReceiveNilChan, traceEvGoStop, 2)
		throw("unreachable")
	}

	// Fast path: check for failed non-blocking operation without acquiring the lock.
	//
	// After observing that the channel is not ready for receiving, we observe that the
	// channel is not closed. Each of these observations is a single word-sized read
	// (first c.sendq.first or c.qcount, and second c.closed).
	// Because a channel cannot be reopened, the later observation of the channel
	// being not closed implies that it was also not closed at the moment of the
	// first observation. We behave as if we observed the channel at that moment
	// and report that the receive cannot proceed.
	//
	// The order of operations is important here: reversing the operations can lead to
	// incorrect behavior when racing with a close.
	if !block && (c.dataqsiz == 0 && c.sendq.first == nil ||
		c.dataqsiz > 0 && atomic.Loaduint(&c.qcount) == 0) &&
		atomic.Load(&c.closed) == 0 {
		return
	}
	

	lock(&c.lock)
	if c.closed != 0 && c.qcount == 0 {
		unlock(&c.lock)
		if ep != nil {
			typedmemclr(c.elemtype, ep)
		}
		return true, false
	}

	if sg := c.sendq.dequeue(); sg != nil {
		// Found a waiting sender. If buffer is size 0, receive value
		// directly from sender. Otherwise, receive from head of queue
		// and add sender's value to the tail of the queue (both map to
		// the same buffer slot because the queue is full).
		recv(c, sg, ep, func() { unlock(&c.lock) }, 3)
		return true, true
	}

	if c.qcount > 0 {
		// Receive directly from queue
		qp := chanbuf(c, c.recvx)
		if ep != nil {
			typedmemmove(c.elemtype, ep, qp)
		}
		typedmemclr(c.elemtype, qp)
		c.recvx++
		if c.recvx == c.dataqsiz {
			c.recvx = 0
		}
		c.qcount--
		unlock(&c.lock)
		return true, true
	}

	if !block {
		unlock(&c.lock)
		return false, false
	}

	// no sender available: block on this channel.
	gp := getg()
	mysg := acquireSudog()
	mysg.releasetime = 0
	if t0 != 0 {
		mysg.releasetime = -1
	}
	// No stack splits between assigning elem and enqueuing mysg
	// on gp.waiting where copystack can find it.
	mysg.elem = ep
	mysg.waitlink = nil
	gp.waiting = mysg
	mysg.g = gp
	mysg.isSelect = false
	mysg.c = c
	gp.param = nil
	// 入列并阻塞诶
	c.recvq.enqueue(mysg)
	goparkunlock(&c.lock, waitReasonChanReceive, traceEvGoBlockRecv, 3)

	// someone woke us up
	if mysg != gp.waiting {
		throw("G waiting list is corrupted")
	}
	gp.waiting = nil
	if mysg.releasetime > 0 {
		blockevent(mysg.releasetime-t0, 2)
	}
	closed := gp.param == nil
	gp.param = nil
	mysg.c = nil
	releaseSudog(mysg)
	return true, !closed
}
```

 附:ep 是传送的消息
 ##### Select
  1. [select的语义](https://segmentfault.com/a/1190000006815341) 
  2. [select实现原理](https://draveness.me/golang-select.html)

协程如何退出: for range , for select + 退出channel return/ ,ok:=chan判断chan是否关闭

Select底层调用的函数 case描述 scase
编译期间优化: [walkselectcases](http://bit.ly/2NWYY6S)
>我们在这里会分四种情况分别介绍优化的过程和结果
1 select 中不存在任何的 case；
2 select 中只存在一个 case；
3 select 中存在两个 case，其中一个 case 是 default 语句；
4 通用的 select 条件;
运行期调用顺序
> func Select(cases []SelectCase) (chosen int, recv Value, recvOK bool) // value.go 文件
  func rselect([]runtimeSelect) (chosen int, recvOK bool) //  value.go
  func reflect_rselect(cases []runtimeSelect) (int, bool) // select.go
  func selectgo(cas0 *scase, order0 *uint16, ncases int) (int, bool) // select.go文件
goroutine在select循环中的的唤醒和阻塞,以及case 编译器的优化,优化成if语句.然后运行期对send/receive/nil/default chan的实现

 ![select基本逻辑图](http://bit.ly/33wkBSj)
 
 scase结构体
 
 ```go
 type scase struct {
	c           *hchan
	elem        unsafe.Pointer
	kind        uint16
	pc          uintptr
	releasetime int64
}

const (
	caseNil = iota
	caseRecv
	caseSend
	caseDefault
)
// http://bit.ly/2NWYY6S

```
 
  mutex--->once---> WaitGroup(v/w,阻塞和v=0的时候逐个唤醒)

- [x] [Mutex、RWMutex、WaitGroup、Once 和 Cond  ErrGroup、Semaphore和 SingleFlight](https://draveness.me/golang-sync-primitives)

- [x] [Go定时器 Timer](https://draveness.me/golang-timer)
  > timer对象 根据pid,在64个分桶上找到自己的位置,然后,根据pid 定位timersBuckets(是个四叉堆),然后append) timer和ticker区别就是多了个period 以及在函数[timerproc](http://bit.ly/36Pib2Y)调用(2层for循环) 多了一层对period的处理,将他计算when,然后将heap index = 0 ,remove掉. timerbucket是一个四叉堆, 逻辑还是很简单的,到期的从堆顶(index=0)移除即可,最后出发调用sendtime或者 f function.
                               
- [x] Context包 [①](https://draveness.me/golang-context ) [②](https://www.cnblogs.com/qcrao-2018/p/11007503.html) [③官方blog context](https://blog.golang.org/context)
  总结:
  propagateCancel  1 done 是否Nil 2 err是否 nil 3 新建一个G 监听paren done channel, 并调用函数cancelCtx取消 close(done) 4 如果是定时的或者周期性的,cancelCtx赋值给timer f,有她负责取消 
  

- [x] [sync.pool 临时对象复用池](https://medium.com/a-journey-with-go/go-understand-the-design-of-sync-pool-2dde3024e277) 
private-->shared(g-p) -->other shared(for+cas)--->victim cache

new 

#### MISC
[GMP,讲的还可以,稍微了解下G和P的状态](https://draveness.me/golang-goroutine)

###以后再回来看看