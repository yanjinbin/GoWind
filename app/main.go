package main

import (
	"fmt"
	"sync"
	"time"
)

/*func main() {
	fmt.Println("Hello Goer")
}*/

/*func main() {
	stop := make(chan bool)

	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("监控退出，停止了...")
				return
			default:
				fmt.Println("goroutine监控中...")
				time.Sleep(2 * time.Second)
			}
		}
	}()

	time.Sleep(10 * time.Second)
	fmt.Println("可以了，通知监控停止")
	stop<- true
	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)
}*/

/*func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("监控退出，停止了...")
				return
			default:
				fmt.Println("goroutine监控中...")
				time.Sleep(2 * time.Second)
			}
		}
	}(ctx)

	time.Sleep(10 * time.Second)
	fmt.Println("可以了，通知监控停止")
	cancel()
	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)
}*/

/*func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go watch(ctx, "【监控1】")
	go watch(ctx, "【监控2】")
	go watch(ctx, "【监控3】")

	time.Sleep(10 * time.Second)
	fmt.Println("可以了，通知监控停止")
	cancel()
	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)
}

func watch(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "监控退出，停止了...")
			return
		default:
			fmt.Println(name, "goroutine监控中...")
			time.Sleep(2 * time.Second)
		}
	}
}*/

var rw sync.RWMutex
var wg sync.WaitGroup

func main() {
	N := 15
	wg.Add(N)
	for i := 0; i < N; i++ {
		if i&1 == 1 {
			go EXRead(i)
		} else {
			go SRead(i)
		}
	}
	wg.Wait()
	time.Sleep(20*time.Second)

}

func EXRead(i int) {
	defer wg.Done()
	rw.Lock()
	fmt.Printf("EX打印%d\n", i)
	rw.Unlock()
}

func SRead(i int) {
	defer wg.Done()
	rw.RLock()
	fmt.Printf("SR打印%d\n", i)
	rw.RUnlock()
}
