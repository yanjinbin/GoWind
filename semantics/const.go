package semantics

import (
	"fmt"
	"reflect"
)

// https://blog.golang.org/constants
// 不支持转义字符 这跟shell语法差不多
// raw string literals
const nameBackquote = `hello\tworld`

// string constan
const nameDoubleQuote = "hello\tworld"

// 单引号，双引号和反引号区别 ---> https://segmentfault.com/a/1190000004850183
func Const() {
	fmt.Println("nameDoubleQuote", nameDoubleQuote)
	fmt.Println("nameBackquote", nameBackquote)
	fmt.Println(nameBackquote == nameDoubleQuote)
	fmt.Println(reflect.TypeOf(nameBackquote))

	// What type does this string constant have? The obvious answer is string, but that is wrong.
	// This is an untyped string constant, which is to say it is a constant textual value that does not yet have a fixed type. Yes, it's a string, but it's not a Go value of type string. It remains an untyped string constant even when given a name:
	const hello = "Hello, 世界"

	var u uint
	const v = -1
	// u = uint(v) // Error: negative value
	fmt.Println("u\t", u)
	// const MaxUint uint = ^0 // Error: overflow

	const MaxUint = ^uint(0)
	fmt.Printf("%x\n", MaxUint)
	// The concept of untyped constants in Go means that all the numeric constants, whether integer, floating-point, complex, or even character values, live in a kind of unified space. It's when we bring them to the computational world of variables, assignments, and operations that the actual types matter. But as long as we stay in the world of numeric constants, we can mix and match values as we like. All these constants have numeric value

	// It is this notion of an untyped constant that makes it possible for us to use constants in Go with great freedom.
	//
	//So what, then, is a typed string constant? It's one that's been given a type, like this:
	//
	const typedHello string = "Hello, 世界"

	fmt.Println(hello, typedHello)

	//Default type determined by syntax

	//The default type of an untyped constant is determined by its syntax. For string constants, the only possible implicit type is string. For numeric constants, the implicit type has more variety. Integer constants default to int, floating-point constants float64, rune constants to rune (an alias for int32), and imaginary constants to complex128. Here's our canonical print statement used repeatedly to show the default types in action:
	fmt.Printf("%T %v\n", 0, 0)
	fmt.Printf("%T %v\n", 0.0, 0.0)
	fmt.Printf("%T %v\n", 'x', 'x')
	fmt.Printf("%T %v\n", 0i, 0i)

	// once given a type ,variable type can not be changed
	type MyBool bool
	const True = true
	const TypedTrue bool = true
	var mb MyBool
	mb = true // OK
	mb = True // OK
	// mb = TypedTrue // Bad
	fmt.Println(mb)

	var f = 'a' * 1.5
	fmt.Println(f)

	//Numeric constants live in an arbitrary-precision numeric space; they are just regular numbers. But when they are assigned to a variable the value must be able to fit in the destination. We can declare a constant with a very large value:

	const Huge = 1e1000
	// —that's just a number, after all—but we can't assign it or even print it. This statement won't even compile:

	// fmt.Println(Huge)

}
