package main

import (
	"fmt"
	"time"
)

/**
* @Author : awen
* @Date : 2020/2/15 2:17 下午
 */

func main() {
	ch := make(chan string) // 无缓冲通道
	go sendDataToChan(ch)
	go getDataFromChan(ch)
	time.Sleep(2*1e9) // 休息两秒
}

func sendDataToChan(ch chan string) {
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokyo"
}

func getDataFromChan(ch chan string) {
	var input string
	for  {
		input = <-ch
		fmt.Println(input)
	}
}
