package concurrent

import (
	"fmt"
	"sync"
)

var mu sync.Mutex

// 错误
func DoSth01() {

	for {
		mu.Lock()
		defer mu.Unlock()
		fmt.Println("哈哈哈")
		// some interesting code
		// <-- the defer is not executed here as one *may* think
	}
	// <-- it is executed here when the function exits
	fmt.Println("退出")
}

// 错误
func DoSth02() {

	for i := 0; i < 10; i++ {
		mu.Lock()
		defer mu.Unlock()
		fmt.Println("哈哈哈")
		// some interesting code
		// <-- the defer is not executed here as one *may* think
	}
	// <-- it is executed here when the function exits
	fmt.Println("退出")
}

// 正确
func DoSth03() {
	for {
		go func() {
			mu.Lock()
			defer mu.Unlock()
			fmt.Println("哈哈哈")
		}()
		// some interesting code
		// <-- the defer is not executed here as one *may* think
	}
	// <-- it is executed here when the function exits
	fmt.Println("退出")
}
