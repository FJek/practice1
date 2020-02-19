package main

import "fmt"

/**
* @Author : awen
* @Date : 2020/2/19 9:50 上午
 */

// 抽象工厂模式
// 定义：
//

// 使用公司招coder

func main() {
	var f IFactory
	f = new(TencentFactory)
	goCoder := f.CreateGoCoder()
	javaCoder := f.CreateJavaCoder()

	goCoder.Code()
	javaCoder.Code()
}
