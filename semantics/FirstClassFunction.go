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

	// https://segmentfault.com/a/1190000015135105
	// https://segmentfault.com/a/1190000015246182
	// https://www.jianshu.com/p/8210dbdc1ff1




	// Go的函数作为第一公民 有什么作用呢


}

func simple(a func(a, b int) int) {
	fmt.Println(a(2, 60))
}
