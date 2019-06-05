package semantics

import (
	"fmt"
	"sync"
)

// https://studygolang.com/articles/12719 实参取值 压入栈

func PrintBar(title string) {
	fmt.Println(title)
}

func Defer() {
	fmt.Println("当一个函数内多次调用 defer 时，Go 会把 defer 调用放入到一个栈中，随后按照后进先出（Last In First Out, LIFO）的顺序执行。")
	t1 := "Google"
	defer PrintBar(t1)
	// defer在执行时候传入实参值
	t1 = "Yahoo"
	t2 := "Facebook"
	defer PrintBar(t2)
	t3 := "Amazon"
	defer PrintBar(t3)

	fmt.Println("defer实际用处")

	var wg sync.WaitGroup
	r1 := rect{-67, 89}
	r2 := rect{5, -67}
	r3 := rect{8, 9}
	rects := []rect{r1, r2, r3}
	fmt.Println("如果你仔细观察，会发现 wg.Done() 只在 area 函数返回的时候才会调用。wg.Done() 应该在 area 将要返回之前调用，并且与代码流的路径（Path）无关，因此我们可以只调用一次 defer，来有效地替换掉 wg.Done() 的多次调用")
	for _, v := range rects {
		wg.Add(1)
		go v.area(&wg)
	}
	wg.Wait()
	fmt.Println("\nAll goroutine finish execute.....")

}

type rect struct {
	length int
	width  int
}

func (r rect) area(wg *sync.WaitGroup) {
	defer wg.Done()
	if r.length < 0 || r.width < 0 {
		fmt.Println("rect length or width should be greater than zero")
		return
	}
	area := r.length * r.width
	fmt.Printf("\nrect %v's area is %v\n", r, area)
}
