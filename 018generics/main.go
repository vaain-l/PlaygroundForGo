package main

import "fmt"

// 泛型的核心是类型参数（type parameter）把类型也当成一个“变量”来使用。
// T 是一个类型参数，any 是约束（constraint），表示 T 可以是任何类型
func PrintSlice[T any](s []T) {
	for _, v := range s {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

func main() {
	PrintSlice([]int{1, 2, 3})     // T 被推断为 int
	PrintSlice([]string{"a", "b"}) // T 被推断为 string
	//[T any]：方括号内声明类型参数T，any是约束（Go1.18+中any等价于interface{}）
	//调用时可以不显式指定类型参数，编译器会自动推断。
}
