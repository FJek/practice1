package factory

import "fzw/practice1/设计模式/工厂方法/product"

/**
* @Author : awen
* @Date : 2020/2/18 10:57 上午
 */

// VIVO工厂
type VivoFactory struct {}
func (vivo *VivoFactory) CreateMobile() product.IMobile{
	return &product.VivoMobile{}
}
