package gorouties

import (
	"errors"
	"fmt"
	"runtime"
	"time"
)

func test6() {
	// 在主线程创建了一个chan
	errCh := make(chan error)
	fmt.Println("make Err")
	go func() {
		errCh <- errors.New("err")
	}()
	fmt.Println("Finish", <-errCh)
}

func test4() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	c := make(chan bool, 10)
	for i := 0; i < 10; i++ {
		go Go4(c, i)
	}

	for i := 0; i < 9; i++ {
		<-c
	}

}
func Go4(c chan bool, index int) {
	a := 1
	for i := 0; i < 10000; i++ {
		a += i
	}
	fmt.Println(index, a)
	c <- true
}

func test3() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	c := make(chan bool, 10)

	for i := 0; i < 10; i++ {
		go func(c chan bool, index int) {
			a := 1
			for i := 0; i < 10000; i++ {
				a += i
			}
			fmt.Println(index, a)
			if index == 9 {
				c <- true
			}
		}(c, i)
	}
	<-c

}

func test2() {
	// range关键字
	c := make(chan bool)

	// go关键字,开启协程.
	go func() {
		fmt.Println("GO!...")
		c <- true
		time.Sleep(time.Second * 3)
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("Hey!....")

}
func test1() {
	// 使用channal实现通信.
	c := make(chan bool)

	go func() {
		fmt.Println("GO! ...!")

		// 向chan中写入, 在c不可写之前阻塞.
		c <- true
	}()

	// 从c中读取数据,在c不可读之前阻塞.
	<-c
	fmt.Println("Hey!....")
}
