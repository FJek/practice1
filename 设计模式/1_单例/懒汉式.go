package main

import "fmt"

/**
* @Author : awen
* @Date : 2020/2/17 3:04 下午
 */


var instance  *Instance
type Instance struct {
	name string
}
func getInstance() *Instance{

	if instance == nil {  // 懒汉模式单例，当并发时，都会检测到 instance == nil
		 instance =  new(Instance)
	}
	return instance
}

func main() {
	ins := getInstance()
	ins.name = "juefulin"
	fmt.Println(ins.name)
	// 第二次初始化 还是第一次的内容
	s2 := getInstance()
	fmt.Println(s2.name)
}
