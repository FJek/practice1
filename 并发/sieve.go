package main

import "fmt"

/**
* @Author : awen
* @Date : 2020/2/15 4:29 下午
 */

// 选择器算法
func main() {
	ch := make(chan int)
	go generatePrime(ch)
	for  {
		prime := <- ch
		fmt.Print(prime," ")
		ch1 := make(chan int)
		go filter(ch ,ch1,prime)
		ch = ch1
	}
}
// Send the sequence 2, 3, 4, ... to returned channel
func generatePrime( ch chan int) {
	for i := 2; ; i++ {
		ch <- i
	}
}

func filter(in,out chan int, prime int) {
	for  {
		i := <-in
		if i%prime != 0 {
			out <- i
		}
	}
}

