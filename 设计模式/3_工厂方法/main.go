package main

import (
	"fmt"
	"fzw/practice1/设计模式/3_工厂方法/factory"
)

/**
* @Author : awen
* @Date : 2020/2/18 9:50 上午
 */
// 定义：工厂父类负责创建产品对象的公共接口，而工厂子类负责创建具体的产品对象
// 目的：将产品的实例化操作延迟到子类工厂中完成，一个子类工厂只创建一种产品，具有多态性

func main() {
	xiaoMifactory := &factory.XiaoMiFactory{}
	mobile := xiaoMifactory.CreateMobile()
	mobile.SetName("小米")
	mobile.SetPrice(3999)
	fmt.Printf("%s手机，价格:%d\n",mobile.GetName(),mobile.GetPrice())
	//华为
	huaweiFac := &factory.HuaWeiFactory{}
	hwm := huaweiFac.CreateMobile()
	hwm.SetPrice(4999)
	hwm.SetName("华为")
	fmt.Printf("%s手机，价格:%d\n",hwm.GetName(),hwm.GetPrice())
}
