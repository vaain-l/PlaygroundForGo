package main

import "fmt"

func main() {
	m1 := make(map[string]string)
	m1["a"] = "apple"
	m1["b"] = "banana"
	for _, v := range m1 {
		fmt.Printf("%s\n", v)

	}

	arr := [3]int{10, 20, 30}
	for i, v := range arr {
		fmt.Printf("索引:%d, 值:%d\n", i, v)
	}
	// 输出：索引:0, 值:10 ...

	for i, c := range "go语言" {
		fmt.Println(i, c)
	}
	
}
