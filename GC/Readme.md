### Go 垃圾回收 garbage collect 专题

GC垃圾回收
1 https://blog.golang.org/ismmkeynote
2 https://blog.golang.org/go15gc
https://segmentfault.com/a/1190000010753702

Getting to Go: The Journey of Go's Garbage Collector
https://blog.golang.org/ismmkeynote

[**Hybird Write Barrier**](https://github.com/golang/proposal/blob/master/design/17503-eliminate-rescan.md)

>**Reasoning**<br>
 A full proof of the hybrid write barrier is  given at the end of this proposal. Here we give the high-level intuition behind the barrier.
 Unlike the Dijkstra write barrier, the hybrid barrier does not satisfy the strong tricolor invariant: for example, a black goroutine (a goroutine whose stack has been scanned) can write a pointer to a white object into a black object without shading the white object. However, it does satisfy the weak tricolor invariant [Pirinen '98]:
 Any white object pointed to by a black object is reachable from a grey object via a chain of white pointers (it is grey-protected).
 The weak tricolor invariant observes that it's okay for a black object to point to a white object, as long as some path ensures the garbage collector will get around to marking that white object.
>
>背景: dijkstra barrier 的mutator隐藏object行为,既 通过移动 white object 从heap 到stack上让他再次扫描,hybird则
> 如果stack一旦是黑的或者scan过了,那么就不会再次rescan.为了实现这个目标,并不严格遵守tricolor variant条件.
> 满足如果一个从white object通过grey chain链指向black object,只要确保GC标记了这些reachable object