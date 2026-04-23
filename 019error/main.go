package main

import (
	"errors"
	"fmt"
)

// Go 没有异常（try-catch-finally）而是将错误作为普通的返回值
// 在 Go 中，错误是一个内置接口：
//
//	type error interface {
//		Error() string
//	}
func divide1(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	} else {
		return a / b, nil
	}
}
func divide2(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("%d / %d cannot divide by zero", a, b)
	} else {
		return a / b, nil
	}
}
func main() {
	//err1 := errors.New("something went wrong")
	//// 带格式化的错误
	//err2 := fmt.Errorf("user %s not found", username)
	res1, err1 := divide1(10, 0)
	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println(res1)
	}

	res2, err2 := divide2(10, 0)
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println(res2)
	}
}
