package main

import "fmt"

/**
* @Author : awen
* @Date : 2020/2/15 4:35 下午
 */

func main() {
	primes := sieve2()
	for {
		fmt.Println(<- primes)
	}
}
// Send the sequence 2, 3_工厂方法, 4, ... to returned channel
func generatePrime2() chan int{
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()
	return ch
}

func filter2(in chan int,prime int) chan int {
	out := make(chan int)
	go func() {
		for  {
			if i := <-in;i%prime != 0 {
				out <- i
			}
		}
	}()
	return out
}

func sieve2() chan int{
	out := make(chan int)
	go func() {
		in := generatePrime2()
		for {
			prime := <- in
			in = filter2(in,prime)
			out <- prime
		}
	}()
	return out
}
