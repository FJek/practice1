package main

import (
	//"fzw/practice1/设计模式/8_ 装饰模式/component"
	//"fzw/practice1/设计模式/8_ 装饰模式/decorator"
	"fzw/practice1/设计模式/8_ 装饰模式/component"
	"fzw/practice1/设计模式/8_ 装饰模式/decorator"
)

/**
* @Author : awen
* @Date : 2020/2/23 11:32 上午
 */

// 装饰模式  https://blog.csdn.net/cloudUncle/article/details/83515130
// 我们想拓展一个类或者对象的功能无外乎两种方式：
//   1 继承父类，在子类拓展功能
//   2 关联机制，将一个类的对象嵌入另一个对象，我们称这个嵌入的对象称为装饰器

// 定义：
//  动态地给一个对象增加一些额外的职责(Responsibility)，就增加对象功能来说，装饰模式比生成子类实现更为灵活2

// 模式结构：
//  Component: 抽象构件
//  ConcreteComponent: 具体构件
//  Decorator: 抽象装饰类
//  ConcreteDecorator: 具体装饰类

// 以装饰人为例子
func main() {
	hat := &component.Hat{}
	decoratorA := decorator.NewDecorator("shoes", hat)
	decoratorA.Display()

	newDecorator := decorator.NewDecorator("glasses", hat)
	newDecorator.Display()

}
