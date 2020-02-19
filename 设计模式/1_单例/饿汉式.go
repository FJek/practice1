package main

import "fmt"

/**
* @Author : awen
* @Date : 2020/2/17 3_工厂方法:13 下午
 */

//饿汉模式将在包加载的时候就创建单例对象，当程序中用不到该对象时，浪费了一部分空间
//和懒汉模式相比，更安全，但是会减慢程序启动速度

// 构建一个结构体，用来实例化单例
type example2 struct {
	name string
}

// 声明一个私有变量，作为单例
var instance2 *example2

// init函数将在包初始化时执行，实例化单例
func init() {
	instance2 = new(example2)
	instance2.name = "juefulin"
}

func GetInstance2() *example2 {
	return instance2
}

func main() {
	instance2 := GetInstance2()
	fmt.Println(instance2.name)
	// 二次测试
	instance3 := GetInstance2()
	fmt.Println(instance3.name)

}
