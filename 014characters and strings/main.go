package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	var a byte = 'a'
	var b rune = '牛'
	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("%c\n", a)
	fmt.Printf("%c\n", b)

	s1 := "Hello中国"
	for i := 0; i < len(s1); i++ {
		fmt.Printf("%c ", s1[i]) // 按字节打印,中文被拆成 UTF-8 的byte，出现乱码
	}

	//原理：for range遇到字符串时，Go会隐式按UTF-8解码，每次迭代得到一个rune，且索引会按字节偏移量跳跃。
	for _, r := range "Hello中国" { //自动解码为rune
		fmt.Printf("%c ", r)
	}
	fmt.Print("\n")

	s2 := "Go语言"
	for i := 0; i < len(s2); i++ {
		fmt.Printf("%x ", s2[i]) // 47 6f e8 af ad e8 a8 80 十六进制
		//8
	}
	fmt.Print("\n")
	fmt.Println(len(s2))

	for _, r := range s2 {
		fmt.Printf("%c ", r) // G o 语 言
	}
	fmt.Print("\n")
	fmt.Println(utf8.RuneCountInString(s2)) //4

	//字符串拼接性能陷阱
	var s3 string
	for i := 0; i < 10000; i++ {
		s3 += "a" // 每次迭代都会分配新内存并复制旧数据
	}
	//better way
	var builder strings.Builder
	builder.Grow(10000) // 预分配容量，减少扩容次数
	for i := 0; i < 10000; i++ {
		builder.WriteByte('a')
	}

}
