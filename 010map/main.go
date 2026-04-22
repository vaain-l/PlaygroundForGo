package main

import "fmt"

func main() {
	//映射map是 Go 内建的关联数据类型
	m1 := make(map[string]int)
	//设置键值对
	m1["k1"] = 7
	m1["k2"] = 13
	//也可以通过下面边的语法在一行代码中声明并初始化一个新的 map。
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
	//获取一个键的值
	v1 := m1["k1"]
	fmt.Println(v1) //7
	//返回一个 map 的键值对数量
	fmt.Println("len:", len(m1))
	//移除键值对
	delete(m1, "k1")
	fmt.Println(m1)

	//当从一个 map 中取值时，还可以选择是否接收第二个返回值，
	//该值表明了 map 中是否存在这个键，用来消除键不存在和键的值为零值产生的歧义
	scores := make(map[string]int)
	scores["bob1"] = 0
	score1 := scores["bob1"]
	score2 := scores["bob2"]
	fmt.Println(score1)
	fmt.Println(score2)
	//score有两种可能：
	//1. 如果 bob 存在，且他的score是 0。
	//2. 如果 bob 不存在score类型为int，零值为0
	//无法区分“bob 不存在”和“bob 的score 为 0”。

	//解决方案：第二个返回值（comma-ok idiom）
	//value, ok := mapName[key]
	scores = map[string]int{
		"alice": 100,
		"bob":   0, // 真实积分为 0
	}
	score1, ok := scores["bob"]
	if ok {
		fmt.Println("bob 存在，积分是", score1) // 输出 0
	} else {
		fmt.Println("bob 不存在")
	}

	score2, ok2 := scores["charlie"]
	if ok2 {
		fmt.Println("charlie 存在，积分是", score2)
	} else {
		fmt.Println("charlie 不存在") // 输出这个
	}
}
