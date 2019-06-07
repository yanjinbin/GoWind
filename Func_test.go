package main

import (
	"GoWind/sdk"
	"GoWind/semantics"
	"fmt"
	"testing"
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
