package main

import "fmt"

/**
* @Author : awen
* @Date : 2020/2/24 4:32 下午
 */

//我们可以用2*1的小矩形横着或者竖着去覆盖更大的矩形。请问用n个2*1的小矩形无重叠地覆盖一个2*n的大矩形，总共有多少种方法？
//先放2*1，则f(n-1)，先放1*2，则f(n-2)。

func rectCover(n int) int {
	if n <= 2 {
		return n
	}
	return rectCover(n-1) + rectCover(n-2)
}

func main() {
	fmt.Println(rectCover(5))
}
