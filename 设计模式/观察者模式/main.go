package main

import (
	"fzw/practice1/设计模式/观察者模式/observer"
	"fzw/practice1/设计模式/观察者模式/subject"
)

/**
* @Author : awen
* @Date : 2020/2/23 12:30 上午
 */

// 观察者模式
// 模式动机：
// 	 建立一种对象与对象之间的依赖关系，一个对象发生改变时将自动通知其他对象，其他对象将相应做出反应。
// 	 在此，发生改变的对象称为观察目标，而被通知的对象称为观察者，一个观察目标可以对应多个观察者，
// 	 而且这些观察者之间没有相互联系，可以根据需要增加和删除观察者，使得系统更易于扩展，这就是观察者模式的模式动机。

// 定义：
//	 定义对象间的一种一对多依赖关系，使得每当一个对象状态发生改变时，其相关依赖对象皆得到通知并被自动更新。
//	 观察者模式又叫做发布-订阅（Publish/Subscribe）模式、模型-视图（Model/View）模式、源-监听器（Source/Listener）模式或从属者（Dependents）模式。

// 结构组成
//   Subject: 目标
//   ConcreteSubject: 具体目标
//   Observer: 观察者
//   ConcreteObserver: 具体观察者

// 以用户订阅公众号为例

func main() {
	customerA := &observer.CustomerA{}
	customerB := &observer.CustomerB{}

	// 模拟用户订阅公众号A
	officeA := &subject.OfficeAccount{
		Name:"冲出宇宙",
	}
	officeA.AddCustomer(customerA)
	officeA.AddCustomer(customerB)
	officeA.Publish()

	// 模拟用户订阅公众号B
	officeB := &subject.OfficeAccount{
		Name:"火山爆发",
	}
	officeB.AddCustomer(customerA)
	officeB.AddCustomer(customerB)
	officeB.Publish()


}
