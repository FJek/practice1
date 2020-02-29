package main

import "fmt"

/**
* @Author : awen
* @Date : 2020/2/29 10:59 上午
 */

func main() {
	// 输入
	str := ""
	fmt.Scan(&str)
	for i := 0; i <= len(str)/2; i++ {
		if str[i] != str[len(str)-(i+1)] {
			fmt.Println("False")
			return
		}
	}
	fmt.Println("True")
}

func isCircle(str string) bool {

	for i := 0; i <= len(str)/2; i++ {
		if str[i] != str[len(str)-(i+1)] {
			return false
		}
	}
	return true
}
