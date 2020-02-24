package main

import "fmt"

/**
* @Author : awen
* @Date : 2020/2/24 3:59 下午
 */



/*
	问题：一只青蛙一次可以跳上1级台阶，也可以跳上2级……它也可以跳上n级。求该青蛙跳上一个n级的台阶总共有多少种跳法。

	找规律：
	1阶：1种；
	2阶：2种；
	3阶：4种；
	4阶：8种；
	n阶：2f(n-1)种；

	或者：
	n-1阶：f(n-2)+f(n-3)+...f(1)+f(0)
	n阶：f(n-1)+f(n-2)+...f(1)+f(0) => 2f(n-1)
	得出一个斐波那契函数。
*/

// 迭代
func jumpFloorSuper(n int) int {
	if n <= 2{
		return n
	}
	return 2*jumpFloorSuper(n-1)
}
// 递归
func jumpFloorSuperIterator( n int ) int {
	if n <= 2{
		return n
	}
	result := 2
	for i := 3; i <= n; i++ {
		result = 2*result
	}

	return result
}
func main() {
	N := jumpFloorSuperIterator(6)
	fmt.Println(N)

}
