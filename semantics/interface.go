package semantics

import "fmt"

type SalaryCalculator interface {
	CalculateSalary() int
}

type Permanent struct {
	empId    int
	basicpay int
	pf       int
}

type Contract struct {
	empId    int
	basicpay int
}

//salary of permanent employee is sum of basic pay and pf
func (p Permanent) CalculateSalary() int {
	return p.basicpay + p.pf
}

//salary of contract employee is the basic pay alone
func (c Contract) CalculateSalary() int {
	return c.basicpay
}

/*
total expense is calculated by iterating though the SalaryCalculator slice and summing
the salaries of the individual employees
*/
func totalExpense(s []SalaryCalculator) {
	expense := 0
	for _, v := range s {
		expense = expense + v.CalculateSalary()
	}
	fmt.Printf("Total Expense Per Month $%d", expense)
}

func SumExpense() {
	fmt.Println("和其他语言（如 Java）很不同，其他一些语言要求一个类使用 implement 关键字，来显式地声明该类实现了接口。而在 Go 中，并不需要这样。如果一个类型包含了接口中声明的所有方法，那么它就隐式地实现了 Go 接口")
	pemp1 := Permanent{1, 5000, 20}
	pemp2 := Permanent{2, 6000, 30}
	cemp1 := Contract{3, 3000}
	employees := []SalaryCalculator{pemp1, pemp2, cemp1}
	totalExpense(employees)
	fmt.Println("接口的内部表示:我们可以把接口看作内部的一个元组 (type, value)。 type 是接口底层的具体类型（Concrete Type），而 value 是具体类型的值。")
	describe(pemp1)
	describe(cemp1)
	fmt.Println("空接口:没有包含方法的接口称为空接口。空接口表示为 interface{}。由于空接口没有方法，因此所有类型都实现了空接口。")
	describeEmptyInterface(1)
	describeEmptyInterface(3.2)
	describeEmptyInterface(pemp1)
	describeEmptyInterface(cemp1)

	fmt.Println("类型断言")
	assertInt(12)
	assertInt(12.3)
	fmt.Println("类型选择")

	switchType(123.3)
	switchType(pemp1)

}

func describe(t SalaryCalculator) {
	fmt.Printf("Interface type %T value %v\n", t, t)
}

func describeEmptyInterface(i interface{}) {
	fmt.Printf("Type = %T, value = %v\n", i, i)
}

func assertInt(t interface{}) {
	if v, ok := t.(int); ok {
		fmt.Printf("int值:%v\n", v)
	}
	fmt.Printf("type 是 %T ，value 是 %v\n", t, t)
}

func switchType(t interface{}) {
	switch v := t.(type) {
	case string:
		fmt.Printf("string值:%v\n", t)
	case int:
		fmt.Printf("int值:%v\n", t)
	case SalaryCalculator:
		fmt.Println("SalaryCalculator类型", v.CalculateSalary())
	default:
		fmt.Printf("实际类型是%T,值:%v\n", t, t)

	}
}
