package factory

import "fzw/practice1/设计模式/3_工厂方法/product"

/**
* @Author : awen
* @Date : 2020/2/18 10:54 上午
 */

// 父类工厂
type IMobileFactory interface {
	CreateMobile() product.Mobile
}
