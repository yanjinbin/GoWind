package concurrent

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

type Signal int

const (
	Exit Signal = iota
	Crash
	Done
	Abnormal
	Interrupt
	Doing
	Start
)

var sem = make(chan int, 10)

func Synchronous() {
	c := make(chan Signal)
	fmt.Println("①", time.Now())
	go func() {
		time.Sleep(5 * time.Second)
		fmt.Println("③sleep 5 second", time.Now())
		c <- Done
	}()
	fmt.Println("②", time.Now())
	// block until c has receive value
	<-c
}

/*func handle02(r *http.Request) {
	sem <- 1   // Wait for active queue to drain.
	process(r) // May take a long time.
	<-sem      // Done; enable next request to run.
}*/

func process(r *http.Request) {
	fmt.Println(r.Method)
	time.Sleep(5 * time.Second)
}

/*
func ServeBug(queue chan *http.Request) {
	for {
		req := <-queue
		go handle02(req) // Don't wait for handle02 to finish.
	}
}*/

func Serve(queue chan *http.Request) {
	for req := range queue {
		sem <- 1
		go func(req *http.Request) {
			process(req)
			<-sem
		}(req)
	}
}

func ServeQuit(clientRequests chan *http.Request, MaxOutstanding int, quit chan bool) {
	// Start handlers
	for i := 0; i < MaxOutstanding; i++ {
		// go handle02(clientRequests)
	}
	<-quit // Wait to be told to exit.
}

func Http() {
	queue := make(chan *http.Request, 100)
	Serve(queue)
}

type Request struct {
	args       []int
	f          func([]int) int
	resultChan chan int
}

func sum(a []int) (s int) {
	for _, v := range a {
		s += v
	}
	return
}

func ChannelsOfChannels() {
	clientRequests := make(chan *Request, 10)
	request := &Request{[]int{3, 4, 5}, sum, make(chan int)}
	// Send request
	clientRequests <- request
	handle02(clientRequests)
	// Wait for response.
	fmt.Printf("answer: %d\n", <-request.resultChan)
}

func handle02(queue chan *Request) {
	for req := range queue {
		req.resultChan <- req.f(req.args)
	}
}

type Buffer int

var freeList = make(chan *Buffer, 100)
var serverChan = make(chan *Buffer)

func client() {
	for {
		var b *Buffer
		// Grab a buffer if available; allocate if not.
		select {
		case b = <-freeList:
			// Got one; nothing more to do.
		default:
			// None free, so allocate a new one.
			b = new(Buffer)
		}
		//	load(b)              // Read next message from the net.
		serverChan <- b // Send to server.
	}
}

func server() {
	for {
		b := <-serverChan // Wait for work.
		// process(b)
		// Reuse buffer if there's room.
		select {
		case freeList <- b:
			// Buffer on free list; nothing more to do.
		default:
			// Free list full, just carry on.
		}
	}
	os.Getpid()
}
