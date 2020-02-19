package factory

import (
	"fzw/practice1/设计模式/4_抽象工厂/product"
	"fzw/practice1/设计模式/4_抽象工厂/product/coder"
)

/**
* @Author : awen
* @Date : 2020/2/19 10:28 上午
 */

type ByteFactory struct {}
func (bf *ByteFactory) CreateJavaCoder() product.ICoder {
	return &coder.JavaCoder{}
}
func (bf *ByteFactory) CreateGoCoder() product.ICoder {
	return &coder.GoCoder{}
}
