package concurrent

import (
	"fmt"
	"sync"
)

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

func gen_1(nums ...int) <-chan int {
	out := make(chan int, len(nums))
	for _, n := range nums {
		out <- n
	}
	close(out)
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
	done <- struct{}{}
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

func good_merge(cs ...<-chan int) <-chan int {
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

func good_merge_1(cs ...<-chan int) <-chan int {
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
func bad_merge(cs ...<-chan int) <-chan int {
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
