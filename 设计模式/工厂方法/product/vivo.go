package product

/**
* @Author : awen
* @Date : 2020/2/18 10:48 上午
 */

type VivoMobile struct {
	Mobile
}
func (mobile *VivoMobile) GetName() string{
	return mobile.name
}
func (mobile *VivoMobile) GetPrice() int {
	return mobile.price
}

