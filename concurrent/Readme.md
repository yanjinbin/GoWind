### Go语言并发模型专题

https://github.com/golang/go/wiki/LearnConcurrency

 Go的内存模型 
 
 Context使用小鸡
 https://www.flysnow.org/2017/05/12/go-in-action-go-context.html
 
 go blog 中文翻译写得还不错的文章
 https://segmentfault.com/a/1190000006815341
 https://segmentfault.com/a/1190000006670880
 https://segmentfault.com/a/1190000006261218

 https://github.com/golang/go/wiki/Articles#concurrency--channels
 
 Concurrent
https://blog.golang.org/concurrency-is-not-parallelism
https://blog.golang.org/advanced-go-concurrency-patterns
https://blog.golang.org/context
https://talks.golang.org/2012/concurrency.slide
https://github.com/golang/go/wiki/MutexOrChannel


https://github.com/golang/go/wiki/LearnConcurrency



https://blog.golang.org/go-concurrency-patterns-timing-out-and

The Go Memory Model
https://golang.org/ref/mem

https://github.com/golang/go/wiki/Articles#concurrency--channels  concurrent 和 channel处理



https://github.com/golang/go/wiki/Timeouts

https://github.com/golang/go/wiki/LockOSThread

https://github.com/golang/go/wiki/MutexOrChannel


advance topic 高级主题

https://github.com/golang/go/wiki/LearnConcurrency

Share Memory By Communicating
https://blog.golang.org/share-memory-by-communicating
https://golang.org/doc/codewalk/sharemem/

Go Concurrency Patterns: Pipelines and cancellation
https://blog.golang.org/pipelines

Advanced Go Concurrency Patterns
https://blog.golang.org/advanced-go-concurrency-patterns

Go Concurrency Patterns: Timing out, moving on
https://blog.golang.org/go-concurrency-patterns-timing-out-and
context https://studygolang.com/articles/13866

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



https://talks.golang.org/2012/waza.slide#30

计算机世界里的定律
https://github.com/nusr/hacker-laws-zh

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
// go 语言设计思想
https://talks.golang.org/2012/splash.article
https://github.com/golang/go/wiki/LearnServerProgramming
https://github.com/golang/go/wiki/GoTalks
https://www.ardanlabs.com/blog/2014/01/concurrency-goroutines-and-gomaxprocs.html

https://www.ardanlabs.com/blog/2013/10/my-channel-select-bug.html

https://www.ardanlabs.com/blog/2013/09/pool-go-routines-to-process-task.html

https://www.ardanlabs.com/blog/2014/02/the-nature-of-channels-in-go.html

https://blog.golang.org/share-memory-by-communicating


- [x] Go性能分析工具玩具(https://blog.wolfogre.com/posts/go-ppof-practice/#%E4%BD%BF%E7%94%A8-pprof)

/debug下面 alloc mutex heap profile(cpu) 查看 cpu 内存 协程创建 锁争用  阻塞  
学习汇编 , 快速过一遍 计算机组成和结构原理(当做字典用),重点反复阅读 深入理解计算机系统()

火焰图生成步骤:

1 设置数据采样频率  
```go
 package main
 
 import (
 	"log"
 	"net/http"
 	"os"
 
 	// 略
 	_ "net/http/pprof" // 会自动注册 handler 到 http server，方便通过 http 接口获取程序运行采样报告
 	"runtime"
 
 	// 略
 )
 
 func main() {
 	// 略
 	runtime.GOMAXPROCS(1)              // 限制 CPU 使用数，避免过载
 	runtime.SetMutexProfileFraction(1) // 开启对锁调用的跟踪
 	runtime.SetBlockProfileRate(1)     // 开启对阻塞操作的跟踪
 
 	go func() {
 		// 启动一个 http server，注意 pprof 相关的 handler 已经自动注册过了
 		if err := http.ListenAndServe(":6060", nil); err != nil {
 			log.Fatal(err)
 		}
 		os.Exit(0)
 	}()
 
 	// 略
 }
```


```shell script
// 6060是你设置的端口号 
go tool pprof pprof http://127.0.0.1:6060/debug/pprof/profile -seconds 10
```

2 开启web, 命令:  go tool pprof -http=:[端口号,自己随意设置]   [路径名来自于第一步].pb.gz

go tool pprof -http=:8081 ~/pprof/pprof.samples.cpu.001.pb.gz

3 火焰图和ggprof差不多,火焰图web 点点就行 ggprof 命令行操作 


Go内存分配
https://zhuanlan.zhihu.com/p/29216091
http://goog-perftools.sourceforge.net/doc/tcmalloc.html
https://www.infoq.cn/article/IEhRLwmmIM7-11RYaLHR
大于 32K 的大对象直接从 mheap 分配。
小于 16B 的使用 mcache 的微型分配器分配
对象大小在 16B ~ 32K 之间的的，首先通过计算使用的大小规格，然后使用 mcache 中对应大小规格的块分配
如果对应的大小规格在 mcache 中没有可用的块，则向 mcentral 申请
如果 mcentral 中没有可用的块，则向 mheap 申请，并根据 BestFit 算法找到最合适的 mspan。如果申请到的 mspan 超出申请大小，将会根据需求进行切分，以返回用户所需的页数。剩余的页构成一个新的 mspan 放回 mheap 的空闲列表。
如果 mheap 中没有可用 span，则向操作系统申请一系列新的页（最小 1MB）。

[线程池](https://studygolang.com/articles/12512)

[内存泄漏](https://segmentfault.com/a/1190000019222661)