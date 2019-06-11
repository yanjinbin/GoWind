package concurrent

import "fmt"

// https://golang.org/ref/spec#Channel_types
// channel的 len cap 和关闭 以及 channel 在何时blocked
// nil channel , full channel, unbuffered channel will be blocked.
var channels chan int

func CloseChannel(channels chan int) chan int {
	for i := 0; i < 10; i++ {
		fmt.Println("i:", i, "len(channels)", len(channels), "cap(channels)", cap(channels))
		if i == 5 {
			close(channels)
			return channels
		}
		channels <- i
	}
	return channels
}

func Channel() {
	channels = make(chan int, 10)
	fmt.Println("initialize channel length:\t", len(channels), "cap(channel)", cap(channels))
	CloseChannel(channels)
	fmt.Println("关闭channel之后,步骤3")
	fmt.Println("len(channels)", len(channels), "cap(channel)", cap(channels))
	if rs, ok := <-channels; ok {
		// 输出chan type的zero value了
		fmt.Println("输出chan type的zero value了--->\t", rs, "len:", len(channels), "cap:", cap(channels))
	}
	fmt.Println("可以看到channel len长度变小了0-1-2-3-4-5-4,cap 依旧是10")
}
