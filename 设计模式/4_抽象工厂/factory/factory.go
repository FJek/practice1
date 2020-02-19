package factory

import "fzw/practice1/设计模式/4_抽象工厂/product"

/**
* @Author : awen
* @Date : 2020/2/19 10:26 上午
 */

// 抽象工厂 HR
type ICoderFactory interface {
	CreateJavaCoder() product.ICoder
	CreateGoCoder() product.ICoder
}
