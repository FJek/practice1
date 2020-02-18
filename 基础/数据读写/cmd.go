package main

import (
	"fmt"
	"os"
	"strings"
)

/**
* 从命令行读取参数
* @Author : awen
* @Date : 2020/2/15 8:04 下午
 */

func main() {
	name := "fzw"
	if len(os.Args) > 1 {
		name += strings.Join(os.Args[1:]," ")
	}
	fmt.Println("Hi",name)
}
