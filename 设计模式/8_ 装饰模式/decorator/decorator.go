package decorator

import (
	"fmt"
	"fzw/practice1/设计模式/8_ 装饰模式/component"
)

/**
* @Author : awen
* @Date : 2020/2/23 5:15 下午
 */

// 具体装饰类
type ShoesDecorator struct {
	component.IAppearance
}
type GlassesDecorator struct {
	component.IAppearance
}
func (sd *ShoesDecorator) Display() {
	fmt.Println("穿上了鞋子")
	sd.IAppearance.Display()
}
func (gd *GlassesDecorator) Display() {
	fmt.Println("戴上了眼镜")
	gd.IAppearance.Display()
}

// 定义工厂函数产出具体装饰类
func NewDecorator(s string, decorator component.IAppearance) component.IAppearance {
	switch s {
	case "shoes" :
		return &ShoesDecorator{
			IAppearance:decorator,
		}
	case "glasses":
		return &GlassesDecorator{
			IAppearance:decorator,
		}
	default:
		return nil
	}

}
