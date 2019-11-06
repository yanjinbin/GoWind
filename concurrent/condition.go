package concurrent

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

func Condition() {
	c := sync.NewCond(&sync.Mutex{})

	for i := 0; i < 10; i++ {
		go listen(c, i)
	}

	time.Sleep(5 * time.Second)

	go broadcast(c)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
}

func broadcast(c *sync.Cond) {
	c.L.Lock()
	c.Broadcast()
	c.L.Unlock()
}

func listen(c *sync.Cond, i int) {
	c.L.Lock()
	// 不能假设condition条件为真啊
	// 所以需要lock,wait内部第一次先调用unlock 然后又再次lock
	c.Wait()
	fmt.Println("listen\t", i)
	c.L.Unlock()
}
