package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

/**
* 一次性读取文件里的所有字符串
* @Author : awen
* @Date : 2020/2/15 10:51 下午
 */

func main() {
	//inputFile := "数据读写/input.dat"
	inputFile, e := os.Open("数据读写/input.dat")
	if e != nil {
		fmt.Println(e.Error())
		return
	}
	outputFile := "数据读写/products_copy.txt"
	buf := make([]byte ,1024)
	inputReader := bufio.NewReader(inputFile)
	n, e := inputReader.Read(buf)
	if n == 0 {
		fmt.Fprintf(os.Stderr, "File Error: %s\n", e)
		return
	}
	fmt.Printf("%s\n", string(buf)) // print string of file
	// write data in data to outputFile
	e = ioutil.WriteFile(outputFile, buf, 0644) // oct, not hex
	if e != nil {
		panic(e.Error())
	}
}
