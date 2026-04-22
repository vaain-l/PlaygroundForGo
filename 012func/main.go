package main

import (
	"errors"
	"fmt"
)

//	func 函数名(参数列表) (返回值列表) {
//		函数体
//	}
func add(x int, y int) int {
	return x + y //Go 需要明确的 return，也就是说，它不会自动 return 最后一个表达式的值

}

// 1.多返回值
func div(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// 这里通过多赋值操作来使用这两个不同的返回值
// 如果仅需要返回值的一部分的话,可以用空白标识符_

// 2.变参函数
func sum(numbers ...int) int {
	fmt.Print(numbers, " ")
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

//在调用时可以传递任意数量的参数

// 3.最简单的闭包
func outer() func() {
	message := "Hello" // 这是 outer 函数里的局部变量
	// 这个匿名函数用到了 message
	inner := func() {
		fmt.Println(message)
	}
	return inner
}

// outer函数执行完后，按理说它的局部变量 message应该被销毁。
// 但是因为inner函数还在使用它，Go会把message保留下来。
// inner函数和它保留的message一起，就是闭包。
// 包不仅能“记住”外部变量的值，还能修改它，而且每次调用闭包，都会基于上次修改的结果
func createCounter() func() int {
	count := 0 // 这个变量会被闭包记住
	return func() int {
		count++ // 修改记住的变量
		return count
	}
}

// 递归
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	res1 := add(1, 2)
	fmt.Println("1+2 =", res1)

	res2, _ := div(40, 2)
	fmt.Println("40/2 =", res2)

	sum([]int{4, 5, 6, 20, 30, 40, 60}...)
	fmt.Println("\n")

	//3.匿名函数和闭包
	//这个 func... 没有名字，叫匿名函数
	//把它赋值给变量 sum
	sum1 := func(a, b int) int {
		return a + b
	}
	// 现在可以用 sum1 来调用它
	fmt.Println(sum1(3, 5)) // 输出 8

	//有时像下面连变量都不想用，定义完就用。这种写法叫 IIFE（立即调用的函数表达式）
	//在 Go 里很常见，比如在 defer 或 `go` 后面。
	sum2 := func(a, b int) int {
		return a + b
	}(3, 9)
	fmt.Println(sum2)

	f := outer() // f 就是一个闭包
	f()          // 输出 "Hello"
	c1 := createCounter()
	fmt.Println(c1()) // 1
	fmt.Println(c1()) // 2
	fmt.Println(c1()) // 3

	c2 := createCounter()
	fmt.Println(c2()) // 1 （c2 有自己的独立 count）

	fmt.Println(fact(4))
	//闭包也可以是递归的，但这要求在定义闭包之前用类型化的var显式声明闭包。
	var fib func(n int) int //var显式声明闭包
	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}
	fmt.Println(fib(9))
}
