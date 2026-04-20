package main

import "fmt"

func main() {
	//切片名 := make([]切片元素类型,切片长度)
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	s := make([]string, 3)
	fmt.Println("emp:", s)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))

	//向切片尾部追加同类型元素或者切片
	s = append(s, "d")
	s = append(s, "d")      //s = ....接收其返回值
	s = append(s, "e", "f") //s = ....接收其返回值
	s = append(s, []string{"g", "h"}...)

	fmt.Println("s:", s)

}
