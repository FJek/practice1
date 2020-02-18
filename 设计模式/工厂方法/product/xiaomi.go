package product

/**
* @Author : awen
* @Date : 2020/2/18 10:45 上午
 */

// 小米产品
type XiaoMiMobile struct {
	Mobile
}
func (xiaomi *XiaoMiMobile) GetName() string{
	return xiaomi.name
}
func (xiaomi *XiaoMiMobile) GetPrice() int{
	return xiaomi.price
}
