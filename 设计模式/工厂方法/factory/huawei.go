package factory

import "fzw/practice1/设计模式/工厂方法/product"

/**
* @Author : awen
* @Date : 2020/2/18 10:56 上午
 */

// 华为工厂
type HuaWeiFactory struct {}

func (wei *HuaWeiFactory) CreateMobile() product.IMobile{
	return &product.HuaWeiMobile{}
}
