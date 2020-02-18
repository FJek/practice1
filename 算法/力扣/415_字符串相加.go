package main

import (
	"fmt"
	"strconv"
)

/**
* @Author : awen
* @Date : 2020/2/17 11:00 上午
 */

func main() {
	fmt.Println(len("123"))
	addStrings("123","789")
}
// 思路类似两数相加
func addStrings(num1 string, num2 string) string {
	var result string
	i,j,carry := len(num1)-1 , len(num2)-1,0
	sum := 0
	for i > 0 || j > 0 {
		s:= num1[i]
		x := int(s)
		y := int(num2[j])
		sum = x + y + carry
		carry = sum / 10
		result += strconv.Itoa(sum%10)+""
		i,j = i-1,j-1
	}
	fmt.Println(result)
	return result
}
