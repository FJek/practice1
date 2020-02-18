package main

import (
	"fmt"
	"io"
	"os"
)

/**
* 文件拷贝
* @Author : awen
* @Date : 2020/2/15 8:20 下午
 */
func main() {
	CopyFile("target.txt", "source.txt")
	fmt.Println("Copy done!")
}

/**
 * @Description: 拷贝文件
 * @Param:
 * @Return:
 * @Date: 2020/2/15
 * @Time: 10:23 下午
 */
func CopyFile(dst string, src string) (written int64, e error) {
	srcFile, e := os.Open(src)
	if e != nil {
		return
	}
	defer srcFile.Close()

	dstFile,e := os.Open(dst)
	if e != nil {
		return
	}
	return io.Copy(dstFile, srcFile)
}
