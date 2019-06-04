package semantics

import "fmt"

func FirstClassFunction() {
	// 第一种玩法
	a := func() {
		fmt.Println("函数是第一公民")
	}
	a()
	// 第二种写法
	func() {
		fmt.Println("不用赋值,和a比多了一个括号")
	}()
	// 第三种写法 和第二种没啥本质区别
	func(name string) {
		fmt.Printf("名字是%s\n", name)
	}("小斌斌")

	// 进阶玩法
	type add func(a, b int) int

	var sum add = func(a, b int) int {
		return a + b
	}
	s := sum(5, 6)

	fmt.Printf("Sum:%d\n", s)

	// 高阶玩法
	// 高阶函数

	// wiki(https://en.wikipedia.org/wiki/Higher-order_function)
	// 把高阶函数（Hiher-order Function）定义为：满足下列条件之一的函数：
	//
	// 接收一个或多个函数作为参数
	// 返回值是一个函数

	c := func(a, b int) int {
		return a * b
	}
	simple(c)

	// closure 闭包.  这里有一个trick 就是闭包在匿名函数中传递的时候 到底是传的是引用(相当于指针)还是值(对象副本)
	// 可以看看这篇文章 Go语言参数传递是传值还是传引用(https://mp.weixin.qq.com/s/iKhhrIQ7zEMhJiVuegzPYg)
	// 当一个匿名函数所访问的变量定义在函数体的外部时，就称这样的匿名函数为闭包。
	// 每一个闭包都会绑定自己的外围变量surrounding variable
	// Go的函数作为第一公民 有什么作用呢

	d := hello()
	e := hello()
	fmt.Println(e("golanger"))

	fmt.Println(d("gopher"))
	fmt.Println(e("!"))

	// 实际用途

	s1 := Student{
		"A",
		12,
		11.2,
	}
	s2 := Student{
		"B", 13, 9.1,
	}
	data := []Student{s1, s2}
	rs := filterStudent(data, func(student Student) bool {
		if student.Age > 12 && student.Height < 10 {
			return true
		}
		return false
	})
	fmt.Println(rs)

	// 函数第一公民在Gin中间件中的使用案列
	// gin web 上应用  [Gin中间件使用](https://gin-gonic.com/zh-cn/docs/examples/using-middleware/)
	//`
	//	// 你可以为每个路由添加任意数量的中间件。
	//	r.GET("/benchmark", MyBenchLogger(), benchEndpoint)`

	//  Go里面传递的是值拷贝呢还是引用呢 就是值  官方The Go language specification 解释-->https://golang.org/ref/spec#Calls

	// Why are maps, slices, and channels references while arrays are values?
	// 数组和slice 区别 https://www.cnblogs.com/howDo/archive/2013/04/25/GoLang-Array-Slice.html
	ints_slice := []int{1, 2, 3, 4}
	ints_slice_bak := ints_slice
	ints_slice[1] = 10
	fmt.Println(ints_slice, "\n", ints_slice_bak)
	ints_array := [3]int{10, 9, 12}
	ints_array_bak := ints_array
	ints_array[1] = 100
	fmt.Println(ints_array, "\n", ints_array_bak)

	// range trick
	// https://xiaozhou.net/something-about-range-of-go-2016-04-10.html#more
	list := []Foo{
		{"A"},
		{"B"},
		{"C"},
	}

	list2 := make([]*Foo, len(list))

	for i, value := range list {
		list2[i] = &value
	}

	fmt.Println(list[0], list[1], list[2])
	fmt.Println(list2[0], list2[1], list2[2])

	var value Foo
	for i := 0; i < len(list); i++ {
		value = list[i]
		list2[i] = &value
		// 观察i和value的区别
		fmt.Printf("i:%d \t,value:%p\n", i, &value)
	}
	fmt.Println(list2[0], list2[1], list2[2])

	for i := 0; i < len(list); i++ {
		list2[i] = &list[i]
	}
	fmt.Println(list2[0], list2[1], list2[2])

	// http://zhige.me/2016/11/29/2016/201611/goForRange-%E9%97%AD%E5%8C%85/
	// 闭包的坑(trick)
	fmt.Println("闭包坑")
	ss := []int{987, 654, 321}

	// 第一种情况，很正常，输出0，1，2，3，4
	for i, v := range ss {
		fmt.Println(i, "①", v)
	}

	//第二种情况，典型的Go语言闭包，它捕获了变量i，但是要注意的是它持有的是引用不是拷贝，当for循环结束时，i=4
	//所以闭包输出的结果都是4
	for i, v := range ss {
		defer func() {
			fmt.Println(i, "②", v)
		}()
	}

	//第三种情况，这种情况下的闭包，它并没有捕获变量i，而是通过传参的方式，这种情况下它得到的是拷贝
	//所以其结果是4，3，2，1，0
	for i, v := range ss {
		defer func(i int, v int) {
			fmt.Println(i, "③", v)
		}(i, v)
	}

}

func simple(a func(a, b int) int) {
	fmt.Println(a(2, 60))
}

func hello() func(name string) string {
	t := "hello"
	// t 在匿名函数中更改了自己的变量,但是换成 A处 代码 就变了
	c := func(name string) string {
		t = t + "\t" + name + "\n"
		// A处
		// s := t + "\t" + name + "\n"
		return t
	}
	return c
}

type Student struct {
	Name   string
	Age    int
	Height float64
}

func filterStudent(students []Student, condition func(Student) bool) []Student {
	var r []Student
	for _, v := range students {
		if condition(v) {
			r = append(r, v)
		}
	}
	return r
}

type Foo struct {
	bar string
}
