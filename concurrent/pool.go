package concurrent

import (
	"fmt"
	"sync"
)

func Pool() {
	p := &sync.Pool{
		New: func() interface{} {
			return 0
		},
	}

	a := p.Get().(int)
	p.Put(1)
	b := p.Get().(int)
	fmt.Print(a, b)
}
