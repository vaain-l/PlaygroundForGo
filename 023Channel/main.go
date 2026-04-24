package main

import (
	"fmt"
	"sync"
)

// 不通过共享内存来通信，而是通过通信来共享内存
func add(a, b int, ch *chan int) {
	*ch <- a + b
}

var wg sync.WaitGroup

func work(num *int, ch *chan struct{}) {
	*ch <- struct{}{}
	*num++ //非原子操作
	<-*ch
	wg.Done()
}

func main() {
	//无缓存channel
	ch1 := make(chan int)
	go func() {
		data1 := <-ch1
		fmt.Println("data1:", data1)
	}()
	ch1 <- 1023

	//有缓存channel
	ch2 := make(chan int, 3)
	ch2 <- 1
	ch2 <- 2
	ch2 <- 3
	go func() {
		data1 := <-ch2
		fmt.Println("data1:", data1)
	}()
	ch2 <- 4

	ch3 := make(chan int)
	close(ch3)
	v1, ok1 := <-ch3
	fmt.Println(v1, ok1)

	ch4 := make(chan int, 3)
	ch4 <- 4
	close(ch4)
	v2, ok2 := <-ch4
	fmt.Println(v2, ok2)
	v2, ok2 = <-ch4
	fmt.Println(v2, ok2)

	ch5 := make(chan int)
	go add(5, 9, &ch5)
	fmt.Println(<-ch5)

	ch6 := make(chan struct{}, 1)
	wg.Add(1000)
	Num := 0
	for i := 0; i < 1000; i++ {
		go work(&Num, &ch6)
	}
	wg.Wait()
	fmt.Println(Num) //不是1000

}
