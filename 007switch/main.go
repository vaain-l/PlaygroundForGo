package main

import (
	"fmt"
	"time"
)

func main() {
	//switch 判断对象{
	//case 分支条件1:
	//	分支体1//判断对象满足条件1时执行
	//case 分支条件2:
	//	分支体2//判断对象满足条件2时执行
	//}

	i := 2
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	//switch 判断对象 {
	//case 分支条件1, 分支条件2:
	//	分支体1//判断对象满足条件1或者条件2时执行
	//default:
	//	分支体3//判断对象不满足条件1或者条件2时执行，类似else
	//}
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	//switch {
	//case 分支条件1:
	//	分支体1
	//default:
	//	分支体2
	//}
	////上下代码块效果一样
	//if 分支条件1{
	//	分支体1
	//}else{
	//	分支体2
	//}
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	if t.Hour() < 12 {
		fmt.Println("It's before noon")
	} else {
		fmt.Println("It's after noon")
	}

}
