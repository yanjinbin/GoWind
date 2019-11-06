package main

import (
	. "GoWind/concurrent"
	"GoWind/sdk"
	"GoWind/semantics"
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestFirstClassFunction(t *testing.T) {
	semantics.FirstClassFunction()
}

func TestStr(t *testing.T) {
	name1 := "hello world"
	name2 := "Señor"
	sdk.PrintBytes(name1)
	fmt.Printf("\n")
	sdk.PrintBytes(name2)
	fmt.Printf("\n")
	sdk.PrintChars(name1)
	fmt.Printf("\n")
	sdk.PrintChars(name2)
	fmt.Printf("\nrune可以用来处理utf8占据2个字节的单字符\n")
	sdk.PrintRune(name1)
	fmt.Printf("\n")
	sdk.PrintRune(name2)
	fmt.Printf("\n字符串unicode utf8编码互转\n")

	sdk.Convert()
	// 出一篇小文 你真的用对了len内置函数嘛?
	sdk.CalLen("Señor")
}

func TestPointer(t *testing.T) {
	semantics.PointerType()
}

func TestMethod(t *testing.T) {
	semantics.Method()
}

func TestInterface(t *testing.T) {
	semantics.SumExpense()
}

func TestPolymorphismComposite(t *testing.T) {
	// 组合  多态
	fmt.Println("组合")
	semantics.Composite()
	fmt.Println("多态")
	semantics.Polymorphism()
}

func TestDefer(t *testing.T) {
	semantics.Defer()
}

func TestPanicRecover(t *testing.T) {
	semantics.PanicRecover()
}

func TestReflect(t *testing.T) {
	o := semantics.Order{
		OrdId:      456,
		CustomerId: 56,
	}
	semantics.CreateQuery(o)
	e := semantics.Employee{
		12,
		"Naveen",
		565,
		"Coimbatore",
		90000,
		"India",
	}
	semantics.CreateQuery(e)
	i := 90
	semantics.CreateQuery(i)
}

func TestError(t *testing.T) {
	semantics.Err()
}

func TestFileIO(t *testing.T) {
	sdk.FileIO()
}

func TestBufio(t *testing.T) {
	sdk.Bufio()
}

func TestScanner(t *testing.T) {
	sdk.BufScanner()
}

func TestConcurrentWrite(t *testing.T) {
	sdk.ConcurrentWrite()
}

func TestSlice(t *testing.T) {
	semantics.Slice()
	semantics.AppendByte([]byte{'a', 'm', 'e'}, 'r', 'i', 'c', 'a', 'n')
	filter := semantics.Filter([]int{1, 6, 7}, func(i int) bool {
		if i > 10 {
			return false
		}
		return true
	})
	fmt.Println("filters:", filter)
	// 这里err 处理的有问题 可以参考 Go Blog err are values --> https://blog.golang.org/errors-are-values
	path := "./semantics/sample.txt"
	digits, err := semantics.FindDigits(path)
	digits, err = semantics.CopyDigits(path)
	digits, err = semantics.AppendDigits(path)
	if err != nil {
		return
	}
	fmt.Println("digits:", digits)
}

func TestConst(t *testing.T) {
	semantics.Const()
}

func TestGob(t *testing.T) {
	sdk.Gob()
}

func TestIoat(t *testing.T) {
	semantics.Iota()
}

func TestMap(t *testing.T) {
	sdk.Map()
}

func TestConcurrent(t *testing.T) {
	// case 1
	/*
		go concurrent.Boring("easy boy")
		fmt.Println("调用main 结束")*/

	// case 2

	/*go concurrent.Boring("easy boy")
	fmt.Println("I'm listening")
	time.Sleep(2 * time.Second)
	fmt.Println("You are boring, I'm leaving !")
	*/
	// case 3

	/*c := make(chan string)
	go concurrent.BoringChan("boring!", c)
	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-c) // Receive expression is just a value.
	}
	fmt.Println("You're boring; I'm leaving.")*/

	// case 4

	d := BoringChanReturn("boring!") // Function returning a channel.
	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-d)
	}
	fmt.Println("You're boring; I'm leaving.")

	// case 5 more instances
	// lockstep解释:a way of marching with each person as close as possible to the one in front.
	joe := BoringChanReturn("Joe")
	ann := BoringChanReturn("Ann")
	for i := 0; i < 5; i++ {
		fmt.Println(<-joe)
		fmt.Println(<-ann)
	}
	fmt.Println("You're both boring; I'm leaving.")

	//case 6 fan-in and fan-out
	c := FanIn(BoringChanReturn("xiaoming"), BoringChanReturn("Lileir"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("You're both boring; I'm leaving.")

	// case 7 use select to fan-in fan-out
	fmt.Println("===case 7===")
	c7 := BoringChanReturn("joe")
Loop:
	for {
		select {
		case s := <-c7:
			fmt.Println("s:", s, time.Now())
		case <-time.After(1 * time.Second):
			fmt.Println("you are too slow")
			// 加不加return 需要重视
			fmt.Println("跳出循环")
			break Loop

		}
	}

	// case 8 tag:select关键字 whole conversation time out
	fmt.Println("===case 8====", time.Now())
	c8 := BoringChanReturn("Paul")
	timeout := time.After(5 * time.Second)
	for {
		select {
		case s := <-c8:
			fmt.Println("s:", s, time.Now())
		case <-timeout:
			fmt.Println("you talk too much")
			//加不加return 可是不一样的哦
			// 不加 执行一次之后,就失效了 不起作用了
			return
		}
	}
}

func TestSynchronous(t *testing.T) {
	Synchronous()
}

func TestEffectiveGo(t *testing.T) {
	// ChannelsOfChannels()
	ChanInChan()
}

func TestChan(t *testing.T) {
	// StackOverFlow上面的解答 http://bit.ly/2XH3RVy
	// chan用法辨析
	fmt.Println("比较chan1 chan2 chan3 chan4 ,这是什么原因呢?")
	Chan1()
	Chan2()
	Chan3()
	Chan4()
}

func TestChannelReceiveSend(t *testing.T) {
	Channel()
	// ints := make(chan int, 1)
	// SelectSample(nil, ints, nil)
}

func TestAppend(t *testing.T) {
	x := []string{"start"}

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		y := append(x, "hello", "world")
		t.Log(cap(y), len(y))
	}()
	go func() {
		defer wg.Done()
		z := append(x, "goodbye", "bob")
		t.Log(cap(z), len(z))
	}()
	wg.Wait()
}

func TestAppend1(t *testing.T) {
	x := make([]string, 0, 6)

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		y := append(x, "hello", "world")
		t.Log(len(y))
	}()
	go func() {
		defer wg.Done()
		z := append(x, "goodbye", "bob")
		t.Log(len(z))
	}()
	wg.Wait()
}

func TestMutex(t *testing.T) {
	// Mutex()
	Mutex01()
}

func TestCond(t *testing.T) {
	Cond()
}

func TestOnceDo(t *testing.T) {
	OnceDo()
}

func TestFastRandom(t *testing.T) {
	var N uint64 = 100
	var U uint64 = rand.Uint64()
	ret := (U * N) >> 32

	fmt.Println(ret)
	fmt.Println(4 >> 1)
}

func TestCtx(t *testing.T) {
	Ctx()
}

func TestSelect(t *testing.T) {
	SelectTry()
}

func TestCondition(t *testing.T) {
	Condition()
}
