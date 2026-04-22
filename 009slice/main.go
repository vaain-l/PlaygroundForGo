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

	//slice copy
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("c:", c)

	//“切片”操作
	l1 := s[2:5] //这个 slice 包含从 s[0] 到 s[5]（不包含 5）的元素,
	l2 := s[2:]  //这个 slice 包含从 s[2]（包含 2）之后的元素。
	fmt.Println("l1:", l1)
	fmt.Println("l2:", l2)

	//Slice 可以组成多维数据结构。
	//内部的 slice 长度可以不一致，这一点和多维数组不同
	var matrix [4][3]int // 4行，每行3列，所有行长度固定为4
	matrix[0][2] = 5
	// 初始化时，每行必须给足4个元素
	grid := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("matrix:", matrix)
	fmt.Println("grid:", grid)
}
