package factory

import "fzw/practice1/设计模式/3_工厂方法/product"

/**
* @Author : awen
* @Date : 2020/2/18 10:53 上午
 */

// 具体类
// 小米工厂
type XiaoMiFactory struct {}

func (mi *XiaoMiFactory) CreateMobile() product.IMobile{
	return &product.XiaoMiMobile{}
}
