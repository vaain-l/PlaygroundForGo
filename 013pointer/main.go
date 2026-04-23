package main

import "fmt"

// 通过两个函数：zeroval和 zeroptr 来比较指针和值
func zeroval(ival int) {
	ival = 0
}
func zeroptr(iptr *int) {
	*iptr = 0
}
func main() {
	var x int = 42
	var p *int = &x // p 是指向 x 的指针

	fmt.Println(*p) // 42，解引用得到原值
	*p = 100        // 修改 x 的值
	fmt.Println(x)  // 100
	//通过指针可以间接访问或修改它指向的值

	//zeroval在main函数中不能改变i的值 但是zeroptr可以，因为它有这个变量的内存地址的引用。
	i := 1
	fmt.Println("initial:", i)

	zeroval(i) // 修改的是副本，原始变量 i 不受影响
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)
}
