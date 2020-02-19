package main

import (
	"fmt"
	"sync"
)

/**
* @Author : awen
* @Date : 2020/2/17 3_工厂方法:17 下午
 */


var instance3  *example3
var mux sync.Mutex

type example3 struct {
	name string
}
func getInstance3() *example3{
	if instance3 == nil {  // 懒汉式+锁机制
		mux.Lock()
		defer mux.Unlock()
		if instance3 == nil{
			instance3 = &example3{}
		}
	}
	return instance3
}

func main() {
	ins := getInstance3()
	ins.name = "juefulin"
	fmt.Println(ins.name)
	// 第二次初始化 还是第一次的内容
	s2 := getInstance3()
	fmt.Println(s2.name)
}
