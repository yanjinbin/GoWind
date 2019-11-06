package concurrent

import (
	"fmt"
	"sync"
)

// https://segmentfault.com/a/1190000006261218
// todo 需要修改变量命名 以及 为什么dead lock的原因

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func main() {
	/*
		in := gen(2, 3)
		// 启动两个运行 sq 的goroutine
		// 两个goroutine的数据均来自于 in
		c1 := sq(in)
		c2 := sq(in)

		// 消耗 output 生产的第一个值
		done := make(chan struct{}, 2)
		out := merge(done, c1, c2)
		fmt.Println(<-out) // 4 or 9

		// 告诉其他发送者，我们将要离开
		// 不再接收它们的数据
		done <- struct{}{}
		done <- struct{}{}*/

	// 设置一个 全局共享的 done channel，
	// 当流水线退出时，关闭 done channel
	// 所有 goroutine接收到 done 的信号后，
	// 都会正常退出。
	done := make(chan struct{})
	defer close(done)

	in := gen02(done, 2, 3)

	// 将 sq 的工作分发给两个goroutine
	// 这两个 goroutine 均从 in 读取数据
	c1 := sq01(done, in)
	c2 := sq01(done, in)

	// 消费 output 生产的第一个值
	out := merge(done, c1, c2)
	fmt.Println(<-out) // 4 or 9

	// defer 调用时，done channel 会被关闭。
}

func gen02(done <-chan struct{}, nums ...int) <-chan int {
	out := make(chan int, len(nums))
	go func() {
		defer close(out)
		for n := range nums {
			select {
			case out <- n:
			case <-done:
				return
			}
		}
	}()
	return out
}

func sq01(done <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			select {
			case out <- n * n:
			case <-done:
				return
			}
		}
	}()
	return out
}

// fan-in fan-out
func merge(done <-chan struct{}, cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	// 为 cs 的的每一个 输入channel
	// 创建一个goroutine。output函数将
	// 数据从 c 拷贝到 out，直到c关闭，
	// 或者接收到 done 信号；
	// 然后调用 wg.Done()
	output := func(c <-chan int) {
		for n := range c {
			select {
			case out <- n:
			case <-done:
			}
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// 启动一个 goroutine，用于所有 output goroutine结束时，关闭 out
	// 该goroutine 必须在 wg.Add 之后启动
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func GoodMerge02(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)
	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func gen01(nums ...int) <-chan int {
	out := make(chan int, len(nums))
	for _, v := range nums {
		out <- v
	}
	close(out)
	return out
}

func GoodMerge01(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int, 1)
	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

// bad practice
func BadMerge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)
	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		defer wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// deadlock
	wg.Wait()
	close(out)

	return out
}
