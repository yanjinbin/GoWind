package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"os/signal"
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

	// https://github.com/golang/go/wiki/SignalHandling
	// 这个场景常见于http server的启动
	f, err := ioutil.TempFile("", "test")
	if err != nil {
		panic(err)
	}
	defer os.Remove(f.Name())
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	// 一直阻塞,直到出列之后,调用defer
	<-sig
}
func randomDuration() time.Duration {
	return time.Duration(rand.Int63n(1e9))
}
