package main

import "fmt"

/**
* @Author : awen
* @Date : 2020/2/17 3_工厂方法:34 下午
 */

// 定义： 又叫静态工厂方法，根据参数的不同返回不同类的实例，专门定义一个类来创建其他类（共同父类）的实例
// 角色：
//	Factory: 工厂觉得负责创建所有实例的内部逻辑
//	Product: 抽象产品是所创建的所有产品的父类，负责描述实例产品的所有公共接口
//	ConcreteProduct: 具体产品角色是创建目标，所有创建的对象都充当这个角色的某个具体类的实例

// 适用环境：
//工厂类负责创建的对象比较少：由于创建的对象较少，不会造成工厂方法中的业务逻辑太过复杂。
//客户端只知道传入工厂类的参数，对于如何创建对象不关心：客户端既不需要关心创建细节，甚至连类名都不需要记住，只需要知道类型所对应的参数。

// 公共接口
type Product interface {
	create()
}

// A 实现Product接口
type ConcreteProductA struct {
}
func (pA ConcreteProductA)create()  {
	fmt.Println("create product A")
}

// 产品B
type ConcreteProductB struct {
}
func (pB ConcreteProductB)create()  {
	fmt.Println("create product B")
}

// 产品C
type ConcreteProductC struct {
}
func (pC ConcreteProductC)create()  {
	fmt.Println("create product C")
}

// 工厂结构体
type Factory struct {
}
// 3_工厂方法
// 返回具体产品，具体产品再调用实现的create方法
func (f Factory)Generate(name string) Product{
	switch name {
	case "A":
		return ConcreteProductA{}
	case "B":
		return ConcreteProductB{}
	default:
		return ConcreteProductC{}
	}
}

// 客户端调用工厂方法 来生产不同的产品
func main() {
	factory := new(Factory)
	pA := factory.Generate("A")
	pA.create()

	pB := factory.Generate("B")
	pB.create()

	pC := factory.Generate("C")
	pC.create()

}
