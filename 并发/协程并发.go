package main

import (
	"fmt"
	"time"
)

/**
* @Author : awen
* @Date : 2020/2/15 1:35 下午
 */

// 主协程
func main() {
	for i:=0;i<10 ;i++  {
		go create(i)
	}
	time.Sleep(2*1e9)
}

func create(i int)  {
	fmt.Println(i)
}
