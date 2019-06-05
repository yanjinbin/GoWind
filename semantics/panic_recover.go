package semantics

import (
	"fmt"
	"runtime/debug"
)

// panic 退出逻辑
// 但在有些情况，当程序发生异常时，无法继续运行。在这种情况下，我们会使用 panic 来终止程序。当函数发生 panic 时，它会终止运行，在执行完所有的延迟函数后，程序控制返回到该函数的调用方。这样的过程会一直持续下去，直到当前协程的所有函数都返回退出，然后程序会打印出 panic 信息，接着打印出堆栈跟踪（Stack Trace），

func fullName(firstName *string, lastName *string) {
	defer recoverName()
	defer fmt.Println("deferred call in fullName")
	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		go panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
}

func PanicRecover() {
	fmt.Println("panic 在 程序无法继续运行了,比如端口占用了这种系统错误 或者 编程错误 除数是0 ")
	defer fmt.Println("deferred call in main")
	firstName := "Yan"
	lastName := "JinBin"

	fullName(&firstName, nil)
	fullName(nil, &lastName)
	fmt.Println("returned normally from main")
}

func recoverName() {
	fmt.Println("recover 是一个内建函数，用于重新获得 panic 协程的控制,注意是同一个协程,recvoer不了不是同一个协程的!")
	if r := recover(); r != nil {
		fmt.Println("recovered from ", r)
		fmt.Println("打印同一个协程 panic 后,被recover恢复之后的panic堆栈消息")
		debug.PrintStack()
	}
}
