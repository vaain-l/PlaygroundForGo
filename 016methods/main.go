package main

import "fmt"

type Teacher struct {
	Name    string
	Age     int
	Subject string
}

func (t Teacher) SayHello() {
	fmt.Printf("Hello,my name is %s,my age is %d,i teach %s\n", t.Name, t.Age, t.Subject)
}

func newTeacher(name string, age int) Teacher {
	s1 := Teacher{Name: name, Age: age}
	return s1
}

// 语法糖：调用方法时Go会自动处理值和指针之间的转换。
// 想要避免在调用方法时产生一个拷贝，或者想让方法可以修改接受结构体的值,可以使用指针来调用方法。
func (t Teacher) SetAge1(age int) {
	t.Age = age
}
func (t *Teacher) SetAge2(age int) {
	t.Age = age
}

//Go支持为结构体类型定义方法(methods)但是内置类型不允许加方法
//func (i int) Double() int {
//	k := i * 2
//	return k
//}

// 继承
type Animal struct {
	Name string
}

func (a Animal) Speaking() {
	fmt.Printf("%s is speaking!!!", a.Name)
}

type Cat struct {
	Animal //相当于嵌套一个字段，只需要写类型，像Name string等，不需要名字-匿名字段
}

func main() {
	t1 := newTeacher("Jack", 20)
	t1.Subject = "math"
	t1.SayHello()

	t2 := newTeacher("Tom", 82)
	t2.SetAge1(100) //在调用方法时产生一个拷贝,对原结构体没有变化
	fmt.Println(t2)
	t2.SetAge2(200) //在调用方法时指向原结构体的内存地址并且修改
	//(&t1).SetAge2(200)//SetAge2()的接收器是指针，但已自动处理值和指针之间的转换
	fmt.Println(t2)

	cat := Cat{}
	cat.Name = "耄耋1"
	cat.Animal.Name = "耄耋2" //cat结构体中没有name字段，但是继承了animal中的name字段

	cat.Speaking()
	cat.Animal.Speaking() //显式调用，上下两种方式相同

}
