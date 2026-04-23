package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Score int
}

type Teacher struct {
	Name   string
	Age    int
	Suject string
}

// 在构造函数中封装创建新的结构实例是一种习惯用法
func newTeacher(name string, suject string) Teacher {
	s1 := Teacher{Name: name, Suject: suject}
	s1.Age = 42
	return s1
}
func main() {
	var stu1 Student
	fmt.Println(stu1)

	var stu2 Student
	stu2.Name = "Jack"
	stu2.Age = 21
	stu2.Score = 100
	fmt.Println(stu2)

	var stu3 Student = Student{
		Name:  "Reacher",
		Age:   21,
		Score: 99,
	}
	fmt.Println(stu3)

	stu4 := Student{
		Name:  "Leon",
		Age:   20,
		Score: 59,
	}
	fmt.Println(stu4)

	stu5 := &Student{
		Name:  "Ada",
		Age:   29,
		Score: 200,
	}
	stu5.Score = 101
	fmt.Println(stu5)

	t1 := newTeacher("Tom", "math")
	fmt.Println(t1)

}
