package semantics

import "fmt"

type Employee struct {
	name string
	age  int
}

/*
使用值接收器的方法。
*/
func (e Employee) changeName(newName string) {
	e.name = newName
}

/*
使用指针接收器的方法。
*/
func (e *Employee) changeAge(newAge int) {
	e.age = newAge
}

func Method() {
	// FAQ https://golang.org/doc/faq  Why do T and *T have different method sets? As the Go specification says, the method set of a type T consists of all methods with receiver type T, while that of the corresponding pointer type *T consists of all methods with receiver *T or T. That means the method set of *T includes that of T, but not the reverse.

	// 	Should I define methods on values or pointers? ¶

	e := &Employee{
		name: "Mark Andrew",
		age:  50,
	}
	fmt.Printf("Employee name before change: %s", e.name)
	e.changeName("Michael Andrew")
	fmt.Printf("\nEmployee name after change: %s", e.name)

	fmt.Printf("\n\nEmployee age before change: %d", e.age)
	(e).changeAge(51)
	fmt.Printf("\nEmployee age after change: %d\n", e.age)

	fmt.Println("匿名字段的方法:属于结构体的匿名字段的方法可以被直接调用，就好像这些方法是属于定义了匿名字段的结构体一样。")
	p := person{
		firstName: "Elon",
		lastName:  "Musk",
		address: address{
			city:  "Los Angeles",
			state: "California",
		},
	}
	// 相当于访问 p.address.fullAddress()
	p.fullAddress() //访问 address 结构体的 fullAddress 方法

	fmt.Println("定义在非结构体接收器的方法原则:为了在一个类型上定义一个方法，方法的接收器类型定义和方法的定义应该在同一个包中。")
	num1 := myInt(5)
	num2 := myInt(10)
	sum := num1.add(num2)
	fmt.Println("Sum is", sum)

}

type address struct {
	city  string
	state string
}

func (a address) fullAddress() {
	fmt.Printf("Full address: %s, %s\n", a.city, a.state)
}

type person struct {
	firstName string
	lastName  string
	address
}

/*
func (a int) add(b int) {
}
*/

type myInt int

func (a myInt) add(b myInt) myInt {
	return a + b
}
