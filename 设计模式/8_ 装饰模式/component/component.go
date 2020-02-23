package component

import "fmt"

/**
* @Author : awen
* @Date : 2020/2/23 5:14 下午
 */

// 抽象组件
type IAppearance interface {
	Display()
}

// 具体组件
type Hat struct{}
type Coat struct{}
type Pants struct{}

func (hat *Hat) Display() {
	fmt.Println("显示帽子")
}
func (coat *Coat) Display() {
	fmt.Println("显示衣服")
}
func (pants *Pants) Display() {
	fmt.Println("显示裤子")
}
