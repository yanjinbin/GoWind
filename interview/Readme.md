
来自于[Go夜读面试专题](https://reading.developerlearning.cn/interview/)
```go
/*func main() {
	var peo People = Stduent{}
	think := "bitch"
	fmt.Println(peo.Speak(think))
}
*/
/*type People interface {
	Show()
}

type Student struct{}

func (stu *Student) Show() {

}

func live() People {
	var stu *Student
	fmt.Println(stu==nil)
	return stu
}

func main() {
	if live() == nil {
		fmt.Println("AAAAAAA")
	} else {
		i := live()
		fmt.Printf("Type = %T, value = %v\n", i, i)
		fmt.Println("BBBBBBB")
	}
}*/

/*
func main() {

	/*wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("A: ", i)
			wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("B: ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()*/

// concurrent.Chan3()
/*fmt.Println("执行main函数")
runtime.GOMAXPROCS(runtime.NumCPU())
start := time.Now()
var t *time.Timer
t = time.AfterFunc(randomDuration(), func() {
	fmt.Println(time.Now().Sub(start))
	t.Reset(randomDuration())
})
time.Sleep(5 * time.Second)

// https://github.com/golang/go/wiki/SignalHandling
// 这个场景常见于http server的启动
f, err := ioutil.TempFile("", "test")
if err != nil {
	panic(err)
}
defer os.Remove(f.Name())
sig := make(chan os.Signal, 1)
signal.Notify(sig, os.Interrupt)
// 一直阻塞,直到出列之后,调用defer
<-sig*/
/*
	c := make(chan int, 3)
	go printNumber(1, 3, c)
	go printNumber(4, 6, c)
	_ = <-c
	_ = <-c
	// fatal error: all goroutines are asleep - deadlock!
	// _ = <-c
*/

func printNumber(from, to int, c chan int) {
	for x := from; x <= to; x++ {
		fmt.Printf("%d\n", x)
		time.Sleep(1 * time.Millisecond)
	}
	c <- 0
}

func randomDuration() time.Duration {
	return time.Duration(rand.Int63n(1e9))
}

/*type People interface {
	Speak(string) string
}

type Stduent struct{}

func (stu Stduent) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}

*/

/*func main() {
	i := GetValue()

	switch i.(type) {
	case int:
		println("int")
	case string:
		println("string")
	case interface{}:
		println("interface")
	default:
		println("unknown")
	}

}

func GetValue() int {
	return 1
}*/

/*func main() {

	println(DeferFunc1(1))
	println(DeferFunc2(1))
	println(DeferFunc3(1))
}

func DeferFunc1(i int) (t int) {
	t = i
	defer func() {
		t += 3
	}()
	return t
}

func DeferFunc2(i int) int {
	t := i
	defer func() {
		t += 3
	}()
	return t
}

func DeferFunc3(i int) (t int) {
	defer func() {
		t += i
	}()
	return 2
}*/

/*func Foo(x interface{}) {
	if x == nil {
		fmt.Println("empty interface")
		return
	}
	fmt.Println("non-empty interface")
}
func main() {
	var x *int = nil
	Foo(x)
}*/

//func main() {
//	list := make([]int, 0)
//
//	//
//	// list := new([]int)
//	list = append(list, 1)
//	fmt.Println(list)
//	list1 := new([]int)
//	fmt.Printf("%T ,%v", list1, list1)
//
//	s1 := []int{1, 2, 3}
//	s2 := []int{4, 5}
//	s1 = append(s1, s2...)
//	fmt.Println(s1)
//
//	sn1 := struct {
//		age  int
//		name string
//	}{age: 11, name: "qq"}
//	sn2 := struct {
//		age  int
//		name string
//	}{age: 11, name: "qq"}
//	/*
//		sn3 := struct {
//			name string
//			age  int
//		}{age: 11, name: "qq"}*/
//
//	if sn1 == sn2 {
//		fmt.Println("sn1 == sn2")
//	}
//	/*
//
//		if sn1 == sn3 {
//			fmt.Println("sn1 == sn2")
//		}
//	*/
//	/*
//		sm1 := struct {
//			age int
//			m   map[string]string
//		}{age: 11, m: map[string]string{"a": "1"}}
//		sm2 := struct {
//			age int
//			m   map[string]string
//		}{age: 11, m: map[string]string{"a": "1"}}
//
//		if sm1 == sm2 {
//			fmt.Println("sm1 == sm2")
//		}*/
//}

/*func GetValue(m map[int]string, id int) (string, bool) {
	if _, exist := m[id]; exist {
		return "存在数据", true
	}
	return "", false
}
func main() {
	intmap := map[int]string{
		1: "a",
		2: "bb",
		3: "ccc",
	}

	v, err := GetValue(intmap, 3)
	fmt.Println(v, err)
}*/

/*
var(
	size :=1024
	max_size = size*2
)
func main()  {
	println(size,max_size)
}*/

const cl = 100

var bl = 123

//
//func main()  {
//
//	/*var x *int = nil
//	//*x   // causes a run-time panic
//	// &*x  // causes a run-time panic
//	fmt.Println(*x,&*x)
//
//	println(&bl,bl)
//	println(&cl,cl)*/
//
///*
//	type MyInt1 int
//	// type alias
//	type MyInt2 = int
//	var i int =9
//	var z *int=new(int)
//	 var i1 MyInt1 = MyInt1(i) // ok
//	// var i1 MyInt1 =i // not ok
//	var i2 MyInt2 = i
//	fmt.Println(i1,i2,z)*/
//}

/*
type User struct {
}
type MyUser1 User
type MyUser2 = User
func (i MyUser1) m1(){
	fmt.Println("MyUser1.m1")
}
func (i User) m2(){
	fmt.Println("User.m2")
}

func main() {
	var i1 MyUser1
	var i2 MyUser2
	i1.m1()
	// i1.m2()
	i2.m2()
}*/

/*var ErrDidNotWork = errors.New("did not work")

func DoTheThing(reallyDoIt bool) (err error) {
	// var result string
	if reallyDoIt {
		result, err := tryTheThing()
		fmt.Println("1")
		if err != nil || result != "it worked" {
			fmt.Println("2",err==nil)
			err = ErrDidNotWork
		}
	}
	return
}

func tryTheThing() (string,error)  {
	return "",ErrDidNotWork
}

func main() {
	fmt.Println(DoTheThing(true))
	fmt.Println("====")
	fmt.Println(DoTheThing(false))
}*/

/*func test(x int) (func(),func())  {
	return func() {
		println(x)
		x+=10
	}, func() {
		println(x)
	}
}

func main()  {
	a,b:=test(100)
	a()
	b()
}*/

/*
func main()  {
	defer func() {
		if err:=recover();err!=nil{
			fmt.Println(err)
		}else {
			fmt.Println("fatal")
		}
	}()

	defer func() {
		panic("defer panic")
	}()
	panic("panic")
}*/
/*
func main()  {
	defer func() {
		fmt.Println("1")
		if err:=recover();err!=nil{
			fmt.Println("++++")
			f:=err.(func()string)
			fmt.Println(err,f(),reflect.TypeOf(err).Kind().String())
		}else {
			fmt.Println("fatal")
		}
	}()

	defer func() {
		panic(func() string {
			return  "defer panic"
		})
	}()
	panic("panic")
}*/

/*
func main() {
	defer func() {
		if err:=recover();err!=nil{
			fmt.Println("执行恢复")
		}
	}()
	defer_call()
}

func defer_call() {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()
	panic("触发异常")
}*/

/*func main() {
	ch := make(chan int, 1000)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()
	go func() {
		for {
			a, ok := <-ch
			if !ok {
				fmt.Println("close")
				return
			}
			fmt.Println("a: ", a)
		}
	}()
	// close(ch)
	fmt.Println("ok")
	time.Sleep(time.Second * 5)
	close(ch)
}
*/

/*type ConfigOne struct {
	Daemon string
}

func (c *ConfigOne) String() string {
	return fmt.Sprintf("print: %v", p)
}

func main() {
	c := &ConfigOne{}
	c.String()
}*/
/*
func main(){
	fmt.Println(len("你好bj!"))
}*/

/*type Test struct {
	Name string
}

var list map[string]Test

func main() {
	// https://golang.org/ref/spec#Index_expressions
	list = make(map[string]Test)
	name := Test{"xiaoming"}
	list["name"] = name
	test := list["name"]
	test.Name="哈哈哈"
	fmt.Println(list["name"])
}*/

/*
type S struct {
}

func f(x interface{}) {
}

func g(x *interface{}) {
}

func main() {
	s := S{}
	p := &s
	f(s) //A
	//g(s) //B
	f(p) //C
	//g(p) //D

}*/


func main() {
	defer ( func() {
		if err := recover();err !=nil{
			fmt.Println("err",err)
		}
	}())
	defer recover()
	// defer  len("aa")
	fmt.Println("1")
	panic("panic herer")
}

func TestMisc(t *testing.T) {
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
}
func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

```