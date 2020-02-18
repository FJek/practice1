package main

import "fmt"

// 累加器
func AddUpper() func(int) (int,string) {
	var n  = 10
	var str  =  "h"
	return func(x int) (int,string) {
		n = n+x
		str = str+string(36)
		return  n,str
	}
}
func main() {
	f := AddUpper()
	fmt.Println(f(1))
	fmt.Println(f(2))
	fmt.Println(f(3))
}
