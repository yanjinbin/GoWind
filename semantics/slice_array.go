package semantics

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

// [Go Slices: usage and internals](https://blog.golang.org/go-slices-usage-and-internals)
// A slice is a descriptor of an array segment. It consists of a pointer to the array, the length of the segment, and its capacity (the maximum length of the segment).
// 第一个格子是ptr 指向array数据起始address ,紧随其后是连续的一块数据array

// [append make操作说明](https://golang.org/ref/spec#Making_slices_maps_and_channels)
// [The length of a nil slice, map or channel is 0. The capacity of a nil slice or channel is 0.](https://golang.org/ref/spec#Length_and_capacity)

// [Array](https://golang.org/doc/effective_go.html#arrays) 特被注意下面这句话:除了按值传递之外,也就是copy,长度也是array的meta属性
// The size of an array is part of its type. The types [10]int and [20]int are distinct.
// 参考链接: https://golang.org/doc/effective_go.html#slices

func Slice() {
	orders := [3]int{9, 10, 11}
	orders = [...]int{13, 53, 34}
	fmt.Println(orders)
	letters := []string{"a", "b", "c"}
	letters = make([]string, 3, 10)
	fmt.Println(letters, "len", len(letters), "cap", cap(letters))
	// 小技巧 array convert into slice
	sliceLetter := letters[:]
	fmt.Println(sliceLetter)
	var classes []Class
	fmt.Println(classes == nil)
	var seniorClass Class
	// struct type zero value is not nil 见Go FAQ
	// fmt.Println(seniorClass==nil)
	seniorClass = Class{
		"日语1", 40,
	}
	fmt.Println(seniorClass)
	var ptrClass *Class
	fmt.Println(ptrClass == nil)

}

type Class struct {
	Name     string
	Capacity int
}

// how built-in append func work
func AppendByte(slice []byte, data ...byte) []byte {
	m := len(slice)
	n := m + len(data)
	// resize for future growth
	if n > cap(slice) {
		newSlice := make([]byte, (n+1)*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	// copy dst
	slice = slice[0:n]
	// append data
	copy(slice[m:n], data)
	return slice
}

func Filter(s []int, condition func(int) bool) []int {
	var p []int
	for _, v := range s {
		if condition(v) {
			p = append(p, v)
		}

	}
	return p
}

// tips 如果一个slice point to a large array ,而slice本身很小的话,GC 不会释放这个array
var digitRegexp = regexp.MustCompile("[0-9]+")

func FindDigits(filename string) ([]byte, error) {
	b, err := ioutil.ReadFile(filename)
	return digitRegexp.Find(b), err
}

// 改写①
func CopyDigits(filename string) ([]byte, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("ReadFile err", err)
		return nil, err
	}
	b = digitRegexp.Find(b)
	c := make([]byte, len(b))
	copy(c, b)
	return c, nil
}

// Exercise 改写②
func AppendDigits(filename string) ([]byte, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("ReadFile err", err)
		return nil, err
	}
	b = digitRegexp.Find(b)
	var digits []byte
	digits = append(digits, b...)
	return digits, nil
}

func SliceGrowTrick01() {
	array := [4]int{10, 20, 30, 40}
	fmt.Println(array)
	slice := array[0:2]
	fmt.Println(slice)
	newSlice := append(slice, 50)
	fmt.Printf("before array=%v, pointer=%p, len=%d\n", array, &array, len(array))
	fmt.Printf("Before slice = %v, Pointer = %p, len = %d, cap = %d\n", slice, &slice, len(slice), cap(slice))
	fmt.Printf("Before newSlice = %v, Pointer = %p, len = %d, cap = %d\n", newSlice, &newSlice, len(newSlice), cap(newSlice))
	newSlice[1] += 10
	fmt.Printf("After slice = %v, Pointer = %p, len = %d, cap = %d\n", slice, &slice, len(slice), cap(slice))
	fmt.Printf("After newSlice = %v, Pointer = %p, len = %d, cap = %d\n", newSlice, &newSlice, len(newSlice), cap(newSlice))
	fmt.Printf("After array = %v\n", array)
}

func SliceGrowTrick02() {
	array := [4]int{10, 20, 30, 40}
	fmt.Println(array)
	slice := array[0:2]
	fmt.Println(slice)
	newSlice := append(slice, 50, 99, 22, 33, 44, 77, 88, 11)
	fmt.Printf("before array=%v, pointer=%p, len=%d\n", array, &array, len(array))
	fmt.Printf("Before slice = %v, Pointer = %p, len = %d, cap = %d\n", slice, &slice, len(slice), cap(slice))
	fmt.Printf("Before newSlice = %v, Pointer = %p, len = %d, cap = %d\n", newSlice, &newSlice, len(newSlice), cap(newSlice))
	newSlice[1] += 10
	fmt.Printf("After slice = %v, Pointer = %p, len = %d, cap = %d\n", slice, &slice, len(slice), cap(slice))
	fmt.Printf("After newSlice = %v, Pointer = %p, len = %d, cap = %d\n", newSlice, &newSlice, len(newSlice), cap(newSlice))
	fmt.Printf("After array = %v\n", array)
	fmt.Println("上面slice的修改并不会对array有影响")
}

func SliceGrowTrick03() {
	array := []int{10, 20, 30, 40}
	fmt.Println(array)
	slice := array[0:2]
	fmt.Println(slice)
	newSlice := append(slice, 50)
	fmt.Printf("before array=%v, pointer=%p, len=%d\n", array, &array, len(array))
	fmt.Printf("Before slice = %v, Pointer = %p, len = %d, cap = %d\n", slice, &slice, len(slice), cap(slice))
	fmt.Printf("Before newSlice = %v, Pointer = %p, len = %d, cap = %d\n", newSlice, &newSlice, len(newSlice), cap(newSlice))
	newSlice[1] += 10
	fmt.Printf("After slice = %v, Pointer = %p, len = %d, cap = %d\n", slice, &slice, len(slice), cap(slice))
	fmt.Printf("After newSlice = %v, Pointer = %p, len = %d, cap = %d\n", newSlice, &newSlice, len(newSlice), cap(newSlice))
	fmt.Printf("After array = %v\n", array)
	fmt.Println("更有意思了，修改了array变量")
}
