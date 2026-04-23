package main

import "fmt"

// 方法签名的集合叫做：接口(Interfaces)
type Animal interface { //定义接口
	Speak()
	//Eat()
}

type Cat struct { //定义结构体
	Name string
}

func (C Cat) Speak() { //Cat实现了Speak方法
	fmt.Printf("%s is speaking!!!\n", C.Name)
}

// 假设还有一个方法未实现
//func (c Cat) Eat() string {
//	return "嗷嗷嗷嗷嗷"
//}

func main() {
	cat := Cat{
		Name: "C1",
	}
	cat.Speak()

	var k Animal
	k = Cat{Name: "C2"} //可以直接把 Cat 赋值给 Animal类型的变量：
	k.Speak()
	//Go 编译器在编译时会检查赋值或传参的类型是否实现了接口。如果没有实现全部方法，会报错。

}
