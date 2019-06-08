package sdk

import (
	"fmt"
	"sync"
)

// go map源码剖析 并发安全
//  https://blog.golang.org/go-maps-in-action

func Map() {

	likes := make(map[string][]*Person)
	for _, p := range people {
		for _, l := range p.Likes {
			likes[l] = append(likes[l], p)
		}
	}
	fmt.Println(likes)

	// map not goroutine safe 设计原因---> https://golang.org/doc/faq#atomic_maps
	var counter = struct {
		sync.RWMutex
		m map[string]int
	}{m: make(map[string]int)}

	counter.RLock()
	n := counter.m["some_key"]
	counter.RUnlock()
	fmt.Println("some_key:", n)

	counter.Lock()
	counter.m["some_key"]++
	counter.Unlock()

}

type Person struct {
	Name  string
	Likes []string
}

var people []*Person
