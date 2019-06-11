package concurrent

import (
	"fmt"
	"math/rand"
	"time"
)

// [Google I/O 2012 - Go Concurrency Patterns](https://youtu.be/f6kdp27TYZs?list=PLdbCfH7zyIrRI2gAQ7KqTkF1jHYyNRr_O)

func Boring(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, "\t", i)
		// time.Sleep(time.Second)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

// channel passed
func boring(msg string, c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("%s %d", msg, i) // Expression to be sent can be any suitable value.
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

// function return a channel
func BoringChanReturn(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			sleepTime := time.Duration(rand.Intn(3e3)) * time.Millisecond
			fmt.Println("BoringChanReturn method, sleep time:\t", sleepTime.String())
			time.Sleep(sleepTime)
		}
	}()
	return c
}

// Multiplexing 多路复用
// 扇入扇出 fan-in and fan-out
func FanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}

type Message struct {
	str  string
	wait chan bool
}

/*
func a(c chan Message) {

	waitForIt := make(chan bool) // Shared between all messages.
	c <- Message{fmt.Sprintf("%s: %d", msg, i), waitForIt}
	time.Sleep(time.Duration(rand.Intn(2e3)) * time.Millisecond)
	<-waitForIt
}

func b(c chan Message) {
	for i := 0; i < 5; i++ {
		msg1 := <-c;
		fmt.Println(msg1.str)
		msg2 := <-c;
		fmt.Println(msg2.str)
		msg1.wait <- true
		msg2.wait <- true
	}
}*/

// case 7 Select
func faninWithSelect(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
			}
		}
	}()
	return c
}

/*func QuitChannel() {
	quit := make(chan bool)
	c := boring("Joe", quit)
	for i := rand.Intn(10); i >= 0; i-- { fmt.Println(<-c) }
	quit <- true
	select {
	case c <- fmt.Sprintf("%s: %d", msg, i):
		// do nothing
	case <-quit:
		return
	}


}*/
type Result string

var (
	Web    = fakeSearch("web")
	Image  = fakeSearch("image")
	Video  = fakeSearch("video")
	Web1   = fakeSearch("web")
	Image1 = fakeSearch("image")
	Video1 = fakeSearch("video")
	Web2   = fakeSearch("web")
	Image2 = fakeSearch("image")
	Video2 = fakeSearch("video")
)

type Search func(query string) Result

func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result(fmt.Sprintf("%s result for %q\n", kind, query))
	}
}

func Google(query string) (results []Result) {
	results = append(results, Web(query))
	results = append(results, Image(query))
	results = append(results, Video(query))
	return
}

//  Google 2
func Google2(query string) (results []Result) {
	c := make(chan Result)
	go func() { c <- Web(query) }()
	go func() { c <- Image(query) }()
	go func() { c <- Video(query) }()

	for i := 0; i < 3; i++ {
		result := <-c
		results = append(results, result)
	}
	return
}

func Google21(query string) (results []Result) {
	c := make(chan Result)
	go func() { c <- Web(query) }()
	go func() { c <- Image(query) }()
	go func() { c <- Video(query) }()

	timeout := time.After(80 * time.Millisecond)
	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timeout:
			fmt.Println("timed out")
			return
		}
	}
	return
}

func First(query string, replicas ...Search) Result {
	c := make(chan Result)
	searchReplica := func(i int) { c <- replicas[i](query) }
	for i := range replicas {
		go searchReplica(i)
	}
	return <-c
}

func Google3(query string) (results []Result) {
	c := make(chan Result)
	go func() { c <- First(query, Web1, Web2) }()
	go func() { c <- First(query, Image1, Image2) }()
	go func() { c <- First(query, Video1, Video2) }()
	timeout := time.After(80 * time.Millisecond)
	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timeout:
			fmt.Println("timed out")
			return
		}
	}
	return
}
