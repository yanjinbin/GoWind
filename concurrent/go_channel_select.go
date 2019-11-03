package concurrent

import (
	"fmt"
)

func ChanInChan() {

}

func Chan1() {
	chanInt := make(chan int, 1)
	chanInt <- 11
	fmt.Println(<-chanInt)
}

func Chan3() {
	<-c
	c <- 0
	fmt.Println("done")
}
func Chan4() {
	c <- 0
	<-c
	fmt.Println("done")
}

var c = make(chan int)
var helloStr string

func f() {
	//time.After();
	helloStr = "hello, world" //1
	<-c                       //2
}

func Chan2() {
	go f()              //3
	c <- 0              //4
	fmt.Print("打印:", a) //5
}
