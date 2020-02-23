package main

/**
* @Author : awen
* @Date : 2020/2/23 5:57 下午
 */

// 一只青蛙一次可以跳上1级台阶，也可以跳上2级。求该青蛙跳上一个n级的台阶总共有多少种跳法（先后次序不同算不同的结果）。

//1阶：共1种跳法；
//2阶：共2种跳法；
//3阶：共3种跳法；
//n阶：先跳1级，还剩n-1级，跳法是f(n-1)；先跳2级，还剩n-2级，跳法是f(n-2)，共f(n-1)+f(n-2)种跳法；
//得出一个斐波那契函数。

func jumpFloor(N int) int {
	if N <= 0 {
		return 0
	}

	if N == 1 || N == 2 {
		return N
	}

	return jumpFloor(N-1) + jumpFloor(N-2)
}

// 迭代
func jumpFloorIterator(N int) int {
	if N <= 0 {
		return 0
	}

	if N == 1 || N == 2 {
		return N
	}

	a, b := 1, 2
	for i := 3; i <= N; i++ {
		a, b = b, a+b
	}
	return b
}
func main() {

}
