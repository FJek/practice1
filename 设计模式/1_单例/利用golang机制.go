package main

import (
	"fmt"
	"sync"
)

/**
* @Author : awen
* @Date : 2020/2/17 3_工厂方法:23 下午
 */

type example4 struct {
	name string
}

var instance4 *example4
var once sync.Once

func GetInstance4() *example4 {

	once.Do(func() {
		instance4 = new(example4)
		instance4.name = "第一次赋值单例"
	})
	return instance4
}

func main() {
	e1 := GetInstance4()
	fmt.Println(e1.name)

	e2 := GetInstance4()
	fmt.Println(e2.name)
}

