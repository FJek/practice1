package builder

import "fzw/practice1/设计模式/5_建造者模式/product"

/**
* @Author : awen
* @Date : 2020/2/22 1:30 上午
 */


//我需要一个大型项目构造者CarBuilder:
type CarBuilder struct {
	Car *product.Car
}
func (c *CarBuilder) NewProduct() {
	c.Car = new(product.Car)
}

func (c *CarBuilder) BuildWheels() {
	c.Car.Wheels = "build wheels"
}

func (c *CarBuilder) BuildChassis() {
	c.Car.Chassis = "build chassis"
}

func (c CarBuilder) BuildSeat() {
	c.Car.Seats = "build seats"
}

func (c CarBuilder) GetResult() interface{} {
	return c.Car
}
