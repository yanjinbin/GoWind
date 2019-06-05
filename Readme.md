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
- [x] 3 [Go 指针](https://studygolang.com/articles/12262) 1 指针的表示和解引用(取指针指向的对象数据) 2 指针类型的zero value是零值 ，但是唯独 struct type 没有零值 所以 struct type  不能判断是否nil 参考这2篇文章 Ref: A:https://golang.org/ref/spec#Comparison_operators B: Go各种数据类型的zero value 默认值 https://golang.org/ref/spec#The_zero_value

- [x] 4 [Go方法](https://studygolang.com/articles/12264)  1 为什么要有方法，因为Go不是纯粹的面向对象系统 2 基于类型的方法是一个实现面向OOP的好方法 2 值接收器 ∈ 指针接收器，也就是值接收器的方法，可以用指针接收器。如果是值接收器，那么不能代替，而是Go自动加一个&指针运算符 3 什么时候用值接收器OR指针接收器 参考Go FAQ 章节:Should I define methods on values or pointers?  一般实践原则:用值接收器，除非你想改遍指针的数据  4  结构体成员-->匿名字段 的方法可以直接调用 5 非结构体的方法定义需要和非结构体在同一个包中.

- [x] 5 [Go 接口①](https://studygolang.com/articles/12266) [Go接口②](https://studygolang.com/articles/12325)
1.接口内部表示(type,value) 2. 接口类型断言和switch选择 3. 实现多个接口 4 接口组合嵌套 reader writer

- [x] 6  [组合](https://studygolang.com/articles/12680) [多态](https://studygolang.com/articles/12681)
Golang的组合和多态玩法,没有继承特性  
 
- [x] 7 [defer](https://studygolang.com/articles/12719) 1 defer用法 入栈, defer 实参的值在函数执行入栈时传入,而非方法正常体结束玩时候在调用(也就是说不在出栈时候调用) 2 入栈遵循Last In First Out 3 defer在sync包 waitgroup的使用

- [x] 8 [panic recover 实践](https://studygolang.com/articles/12785)  总结:函数/方法 发生panic后,不会再执行正常流程,执行完毕所有defer出栈之后,程序控制会转移给调用方,直到当前协程的所有函数退出,程序打印处panic 堆栈信息.可以用recover恢复同一个协程的panic,但是要注意此时的panic信息,要用debug.PrintStack方法打印.

#### Doing
- [ ] 7 go 错误处理
- [ ] 8 Go反射 https://studygolang.com/articles/13178
- [ ] 9 Go读写文件 https://studygolang.com/articles/14669 https://studygolang.com/articles/19475
- [ ] 10 Go select 关键字 https://studygolang.com/articles/12522
- [ ] 11 Go channel https://studygolang.com/articles/12402
- [ ] 12 Go mutex 和 WaitGroup 用法
- [ ] 13 Go mutex 和 WaitGroup 源码分析
- [ ] 14 GO Ratelimiter https://github.com/golang/go/wiki/RateLimiting

ToDo

PS: 有些需要编码加深记忆， 有些看看总结。
:whale2: :whale2: :whale2: 

