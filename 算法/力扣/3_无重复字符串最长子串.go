package main

import "fmt"

/**
* @Author : awen
* @Date : 2020/2/17 11:52 上午
 */

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcab"))
}

func lengthOfLongestSubstring(s string) int {
	i := 0
	max := 0
	a := []rune(s)
	for m,c := range a {
		for n := i; n < m ; n++  {
			if a[n] == c {
				i = n + 1
			}
		}
		if m-i+1 > max {
			max = m - i + 1
		}
	}
	return max
}
