package main

import "fmt"

/**
* @Author : awen
* @Date : 2020/2/23 5:27 下午
 */

//现在要求输入一个整数n，请你输出斐波那契数列的第n项（从0开始，第0项为0）
//n<=39

// 递归
func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

// 迭代
func fibIteration(n int) int  {
	if n < 2 {
		return n
	}
	//f(0) = f(1), f(1) = f(2)
	a,b := 0,1
	for i := 2; i <= n ; i++ {
		a,b  = b ,a + b
	}
	return b
}


func main()  {
	fmt.Println(fibIteration(7))
	a,b := 2,1
	fmt.Println(a)
	fmt.Println(b)
}
