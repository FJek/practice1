package main

import (
	"fmt"
	"time"
)

/**
* @Author : awen
* @Date : 2020/2/15 2:33 下午
 */

func main() {
	ch1 := make(chan int)
	go pump(ch1)       // pump hangs
	go suck(ch1) // prints only 0
	time.Sleep(2*1e9)
}
func suck(ch chan int) {
	for  {
		fmt.Println(<-ch)
	}
}
//func pump(ch chan int) {
//	for i := 0; ; i++ {
//		ch <- i
//	}
//}
