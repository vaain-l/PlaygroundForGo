package main

import "fmt"

func test1() {
	defer fmt.Println("defer") //可以打印
	fmt.Println("before panic")
	//arr := []int{1, 2, 3}
	//fmt.Println(arr[3]) //panic: runtime error: index out of range [3] with length 3

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("发生了panic 并且recover")
		}
	}()
	panic("panic")             //自定义panic
	fmt.Println("after panic") //在此函数中无法执行
}

// 子函数捕获panic
func child() {
	fmt.Println("child before panic")
	panic("child panic")
	fmt.Println("child after panic")
}
func parent() {
	fmt.Println("parent before panic")
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("child panic recover")
		}
	}()
	child()
	fmt.Println("parent after panic") //在此函数中无法执行
}

func main() {
	test1()
	fmt.Println("test main after panic") //可以打印
	parent()
	fmt.Println("parent main after panic") //可以打印
	fmt.Println("main before panic")

}
