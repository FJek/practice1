package factory

import (
	"fzw/practice1/设计模式/4_抽象工厂/product"
	"fzw/practice1/设计模式/4_抽象工厂/product/coder"
)

/**
* @Author : awen
* @Date : 2020/2/19 10:27 上午
 */

// 具体工厂 不同的公司
type TencentFactory struct {}
func (tf *TencentFactory) CreateJavaCoder() product.ICoder {
	return &coder.JavaCoder{}
}
func (tf *TencentFactory) CreateGoCoder() product.ICoder {
	return &coder.GoCoder{}
}
