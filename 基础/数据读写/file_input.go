package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/**
* 读文件
* @Author : awen
* @Date : 2020/2/15 10:29 下午
 */

func main() {
	inputFile, err := os.Open("数据读写/input.dat")
	if err != nil {
		fmt.Println(err.Error())
		fmt.Printf("An error occurred on opening the inputfile\n" +
			"Does the file exist?\n" +
			"Have you got acces to it?\n")
		return // exit code
	}
	defer inputFile.Close() // release resource before exit

	inputReader := bufio.NewReader(inputFile)

	for  {
		inputString, err := inputReader.ReadString('\n')
		fmt.Printf("The input was: %s", inputString)
		if err == io.EOF { // when read end of file
			return
		}
	}
}


