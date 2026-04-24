package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup //

func test1() {
	go fmt.Println("gorutine2")
}

func work(id int) {
	fmt.Printf("第%d个协程正在工作\n", id)
	wg.Done() //结束一个并发协程
}
func main() {
	//同步调用,并发执行
	go fmt.Println("gorutine1")
	go test1()
	go func() { fmt.Println("gorutine3") }()
	fmt.Println("hello")
	time.Sleep(1 * time.Second)

	wg.Add(5) //并发任务数量
	for i := 1; i <= 5; i++ {
		go work(i)

	}
	wg.Wait()

}
