package semantics

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
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
	fmt.Println("a return:", a()) // 打印结果为 a return: 0
	fmt.Println("b return:", b()) // 打印结果为 b return: 2
	fmt.Println("c() return", c())
	fmt.Println("d() return", d())
	ErrDefer()
	DeferValueOrPointer()
	DeferClosure()
	fmt.Println("defer 返回值")
	release()
	releaseNew()

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

// 参见官方 连接 https://blog.golang.org/defer-panic-and-recover 解释的比较清楚
// 1. A deferred function's arguments are evaluated when the defer statement is evaluated.
// 2. Deferred function calls are executed in Last In First Out order after the surrounding function returns.
// 3. Deferred functions may read and assign to the returning function's named return values.
//
// return 2 not 1
func c() (i int) {
	defer func() { i++ }()
	return 1
}

func d() int {
	var i int = 4
	defer func() {
		i++
		fmt.Println("d  defer i值", i)
	}()
	return i
}

// https://studygolang.com/articles/14831
// defer的变量修改 返回值 辨析,本质是返回值的类型和声明 https://my.oschina.net/henrylee2cn/blog/505535
func a() int {
	var i int
	defer func() {
		i++
		fmt.Println("a defer2:", i) // 打印结果为 a defer2: 2
	}()
	defer func() {
		i++
		fmt.Println("a defer1:", i) // 打印结果为 a defer1: 1
	}()
	return i
}

func b() (i int) {
	defer func() {
		i++
		fmt.Println("b defer2:", i) // 打印结果为 b defer2: 2
	}()
	defer func() {
		i++
		fmt.Println("b defer1:", i) // 打印结果为 b defer1: 1
	}()
	return i // 或者直接 return 效果相同
}

// https://studygolang.com/articles/14831
// 先判断资源如IO 连接池 对象池 有没有err,有err 说明没有占用资源 ,不需要defer执行释放资源任务

func ErrDefer() error {
	resp, err := http.Get("www.google.com")
	// 先判断操作是否成功
	if err != nil {
		return err
	}
	// 如果操作成功，再进行Close操作
	defer resp.Body.Close()
	return nil
}

type Car struct {
	model string
}

func (c Car) PrintModel() {
	fmt.Println(c.model)
}

func DeferValueOrPointer() {
	c := Car{model: "DeLorean DMC-12"}
	c1 := &c
	defer c.PrintModel()
	defer c1.PrintModel()
	c.model = "Chevrolet Impala"
}

func DeferClosure() {
	// case 1
	for i := 0; i < 3; i++ {
		defer func() {
			i++
			fmt.Println("Case 1,i值", i) // i值是4 开始了 特别注意下哦!
		}()
	}
	defer fmt.Println("i值是4 开始了 特别注意下哦!")
	// case 2
	for i := 0; i < 3; i++ {
		defer func(i int) {
			i++
			fmt.Println("Case 2,i值", i)
		}(i) // 保存局部变量,闭包实参传入
	}
	// case 3
	for i := 0; i < 3; i++ {
		defer fmt.Println("Case 3, i值", i)
	}
}

// defer返回值改变调用者返回值
func release() error {
	defer func() error {
		return errors.New("error")
	}()
	return nil
}

func releaseNew() (err error) {
	defer func() error {
		err = errors.New("error")
		return err
	}()

	return nil
}

type myerror struct{}

func (myerror) String() string {
	return "myerror there!"
}

func errorly() {
	defer func() {
		fmt.Println(recover())
	}()

	if false {
		panic(myerror{})
	}
}

// 释放相同的资源,改写就是匿名函数传入实参
func do() error {
	f, err := os.Open("book.txt")
	if err != nil {
		return err
	}
	/*
		defer func() {
			if err := f.Close(); err != nil {
				// log etc
			}
		}()*/

	defer func(file io.Closer) {
		if err := f.Close(); err != nil {
			// log etc
		}
	}(f)

	// ..code...

	f, err = os.Open("another-book.txt")
	if err != nil {
		return err
	}

	/*defer func() {
		if err := f.Close(); err != nil {
			// log etc
		}
	}()*/

	defer func(file io.Closer) {
		if err := f.Close(); err != nil {
			// log etc
		}
	}(f)

	return nil
}

// 传的是实参还是形参有很大区别哦
func WhenParamCal01() {
	num := 42

	defer func() {
		// num: 13 because it sees the latest state of the variable
		// through its surrounding context before the func returns
		fmt.Println("defer func: num:", num)
	}()

	// overwrite num
	num = 13

	fmt.Println("main: num:", num)

	// deferred func is executed here
	// the recent value of num is 13
	// so, the defer will see it as 13, not 42
}

func printA(a int) {
	fmt.Println("value of a in deferred function", a)
}

func WhenParamCal02() {
	a := 5
	defer printA(a)
	a = 10
	fmt.Println("value of a before deferred function call", a)

}
