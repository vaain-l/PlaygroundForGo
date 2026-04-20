package main

import "fmt"

func main() {
	//if 分支条件1{
	//	分支体1
	//}else{
	//	分支体2，不满足分支条件1时就会执行
	//}
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

}
