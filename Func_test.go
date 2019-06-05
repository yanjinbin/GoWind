package GoWind

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
