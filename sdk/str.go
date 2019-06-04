package sdk

import (
	"fmt"
	"unicode/utf8"
)

// https://studygolang.com/articles/12261
func PrintBytes(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x\t", s[i])
	}
}

func PrintChars(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c\t", s[i])
	}
}

// rune可以用来处理utf8占据2个字节的单字符
func PrintRune(s string) {
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c\t", runes[i])
	}
}

func Convert() {
	// 字符串 byte rune 互转

	// utf8编码之后的16进制
	byteSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}
	str := string(byteSlice)
	fmt.Println(str)
	// unicode代码点
	runeSlice := []rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072}
	str = string(runeSlice)
	fmt.Println(str)

}

func CalLen(s string) {
	fmt.Printf("length of %s is %d , len:%d", s, utf8.RuneCountInString(s), len(s))
}
