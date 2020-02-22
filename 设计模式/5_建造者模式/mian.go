package main

import (
	"fzw/practice1/设计模式/5_建造者模式/builder"
	"fzw/practice1/设计模式/5_建造者模式/director"
)

/**
* @Author : awen
* @Date : 2020/2/22 1:14 上午
 */

//比如说我是个老司机，但是除了开车还想造车，但是车的构造实在是太复杂了，那么我们就可以将车拆分...
//4个轮子、1个底盘、1个驾驶位...
//好了，为了简便，就造这三个吧，先造个爬犁出来...







func main() {
	// 创建一个指挥者
	director := new(director.Director)
	// 创建建造者
	builder := new(builder.CarBuilder)
	director.SetBuilder(builder)
	car := director.Generate()
	car.Show()
}

