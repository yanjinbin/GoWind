package semantics

import (
	"fmt"
)

// Go指针用法 https://studygolang.com/articles/12262

func PointerType() {
	// 指针的零值是Nil
	a := 25
	var b *int

	if b == nil {
		fmt.Println("b is", b)
		b = &a
		fmt.Println("b after initialization is", b)
	}
	fmt.Println(&Student{} == nil)
	// Todo 这个可以做篇文章
	// https://golang.org/ref/spec#Comparison_operators  In any comparison, the first operand must be assignable to the type of the second operand, or vice versa.
	// 参考这篇stackoverflow 了解到 nil 不是 struct type的 zero value http://bit.ly/2ETtECj
	// https://golang.org/ref/spec#The_zero_value
	// var s Student
	// fmt.Println(s == nil)

	// 指针的解引用
	c := 255
	d := &c
	fmt.Println("address of c is", c)
	fmt.Println("value of c is", *d)
	*d++
	fmt.Println("value of c is", c)

	// 不要向函数传递数组的指针，而应该使用切片
	e := [3]int{89, 90, 91}
	// tips 小技巧 数组--->指针 slice := arr[:]
	modify(e[:])
	fmt.Println(e)

	// Go不支持指针运算
	// d++

}

func modify(sls []int) {
	sls[0] = 90
}
