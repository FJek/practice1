package product

/**
* @Author : awen
* @Date : 2020/2/18 10:48 上午
 */


// 华为产品
type HuaWeiMobile struct {
	Mobile
}
func (huawei *HuaWeiMobile) GetName() string{
	return huawei.name
}
func (huawei *HuaWeiMobile) GetPrice() int{
	return huawei.price
}
