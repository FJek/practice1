package main

import (
	"fmt"
	"time"
)

/**
* @Author : awen
* @Date : 2020/2/15 4:19 下午
 */

func main() {
	stream := pump3()
	go suck3(stream)
	time.Sleep(1e9)
}

func pump3() chan int{
	ch := make(chan int)
	go func() {
		for i := 0; ; i++ {
			ch <-i
		}
	}()
	return ch
}

func suck3(ch chan int) {
	for  {
		fmt.Println(<-ch)
	}
}
