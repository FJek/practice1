package main

import (
	"fmt"
	"sync"
	"time"
)

/**
* @Author : awen
* @Date : 2020/2/16 9:29 下午
 */

var exitChan = make(chan bool,1)
var wg sync.WaitGroup
func main() {
	wg.Add(1)
	go f()
	time.Sleep(time.Second*5)
	exitChan<-true
	wg.Wait()
}

func f() {
	defer wg.Done()
	FORLOOP:
	for  {
		fmt.Println("hello")
		time.Sleep(time.Second*1)
		select {
		case <- exitChan:
			break FORLOOP
		default:
		}
	}
}
