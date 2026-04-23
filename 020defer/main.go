package main

import "fmt"

func main() {
	defer fmt.Println("world") // 延迟执行
	fmt.Println("hello")
}

// 输出：
// hello
// world
//多个 defer按 **LIFO（后进先出）顺序执行，就像栈一样。
//defer的参数在声明时就求值（确定值），而不是在调用时。
//即使函数 panic，defer 也会执行（可以用来捕获 panic 或清理资源）。
