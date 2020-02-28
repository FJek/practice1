package main

import (
	"fmt"
	"strconv"
	"strings"
)

/**
* @Author : awen
* @Date : 2020/2/28 10:25 下午
 */

func main() {
	v1 := `1.10.2`
	v2 := `1.2.10`
	res := compareVersion(v1, v2)
	fmt.Println(res)
}

func compareVersion(v1,v2 string ) int {
	v1s := strings.Split(v1, ".")
	v2s := strings.Split(v2,".")
	for i := 0; i < len(v1s); i++ {
		fmt.Println(v1s[i],v2s[i])
		a, _ := strconv.Atoi(v1s[i])
		b, _ := strconv.Atoi(v2s[i])
		if a > b {
			return 1
		} else if a < b {
			return -1
		} else {
			continue
		}
	}
	fmt.Println(v1s,v2s)
	return 0
}
