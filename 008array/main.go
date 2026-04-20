package main

import "fmt"

func main() {
	//var 数组名 [数组长度]数组元素类型
	var a [5]int
	fmt.Println(a)

	a[4] = 100 //指定了数组a的4号元素为100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])
	fmt.Println("len:", len(a)) //返回数组的长度

	b := [5]int{1, 2, 3, 4, 5} //一行内声明并初始化一个数组
	fmt.Println("dcl:", b)

	//数组类型是一维的，但是你可以组合构造多维的数据结构
	var twoD [4][3]int
	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)
	
}
