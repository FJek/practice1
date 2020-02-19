package product

/**
* @Author : awen
* @Date : 2020/2/18 10:46 上午
 */

// 抽象产品
type Mobile struct {
	name string
	price int
}
type IMobile interface {
	SetName(string)
	SetPrice(int)

	GetName() string
	GetPrice() int
}

func (mobile *Mobile) SetName(name string) {
	mobile.name = name
}
func (mobile *Mobile) SetPrice(price int)  {
	mobile.price = price
}
