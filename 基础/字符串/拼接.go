package main

import (
	"bytes"
	"fmt"
	"strings"
)

/**
* 字符串拼接
* @Author : awen
* @Date : 2020/2/17 11:34 上午
 */

func main() {
	// +=
	str := "123-"
	str += "123123-"
	str += "qwerqwe"
	fmt.Println(str)

	// fmt.Sprintf
	str2:="hello"
	str3 := fmt.Sprintf("%s--%s", str2, str)
	fmt.Println(str3)

	// strings.Join() 字符串数组转字符串
	str4 := "str4"
	str4 = strings.Join([]string{str4, "hello"}, "---") // sep 是连接符号
	fmt.Println(str4) //str4---hello

	// []string和slice的append
	str5 := []string{"222"}
	str5 = append(str5,"123123")
	fmt.Println(strings.Join(str5, "-")) //222-123123

	// bytes.Buffer
	var buf bytes.Buffer
	buf.WriteString("12312")
	buf.WriteString("-")
	buf.WriteString("werwer")
	fmt.Println(buf.String())
}
