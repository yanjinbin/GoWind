package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	fmt.Println("执行main函数")
	runtime.GOMAXPROCS(runtime.NumCPU())
	start := time.Now()
	var t *time.Timer
	t = time.AfterFunc(randomDuration(), func() {
		fmt.Println(time.Now().Sub(start))
		t.Reset(randomDuration())
	})
	time.Sleep(5 * time.Second)
}
func randomDuration() time.Duration {
	return time.Duration(rand.Int63n(1e9))
}
