package main

import (
	"fmt"
	"time"
)

/**
* @Author : awen
* @Date : 2020/2/15 4:16 下午
 */


func sum(x, y int, c chan int) {
	time.Sleep(2*1e9)
	c <- x + y

}

func main() {
	c := make(chan int)
	go sum(12, 13, c)
	fmt.Println(<-c) // 25
}
