package director

import (
	"fzw/practice1/设计模式/5_建造者模式/builder"
	"fzw/practice1/设计模式/5_建造者模式/product"
)

/**
* @Author : awen
* @Date : 2020/2/22 1:27 上午
 */

// 具体建造者传入指挥者
type Director struct {
	builder builder.Builder
}
func (d *Director) SetBuilder(builder builder.Builder) {
	d.builder = builder
}

func (d *Director) Generate() *product.Car {
	d.builder.NewProduct()
	d.builder.BuildChassis()
	d.builder.BuildSeat()
	d.builder.BuildWheels()
	return d.builder.GetResult().(*product.Car)
}
