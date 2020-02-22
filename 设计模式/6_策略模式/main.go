package main

import "fmt"

/**
* @Author : awen
* @Date : 2020/2/22 4:12 下午
 */
// 策略模式
// 生活中 根据环境或者条件的不同选择不同的策略来完成任务
// 开发中，实现某一个功能有多种途径，使用此模式可以方便的新增新的途径，比如实现排序可以有很多方法
// 我们可以把这些方法封装到一个类里，这就是一个策略

// 定义：
// 	我们定义一系列算法，将每一个算法封装起来，并让它们可以相互替换

// 包含角色：
// 	Context:环境
//	Strategy:策略
//	ConcreteStrategy:具体策略

// 以商品销售为例：满100-20、打八折、无优惠
// 1 所有策略接口
type CashSuper interface {
	AcceptMoney(money float64) float64  // 计算需要付的金额
}

// 2 定义三个具体策略
// 无优惠
type CashNormal struct{}

func NewCashNormal() *CashNormal {
	return &CashNormal{}
}
func (normal *CashNormal) AcceptMoney(money float64) float64 {
	return money
}

// 八折优惠
type CashRebate struct {
	Rebate float64 // 折扣
}
// 传入折扣
func NewCashRebate(rebate float64) *CashRebate {
	return &CashRebate{
		Rebate:rebate,
	}
}
// 折扣后价格
func (rebate *CashRebate) AcceptMoney(money float64) float64 {
	return money * rebate.Rebate // ✖️折扣
}

// 满100 - 20
type CashReturn struct {
	ReturnCondition float64 // 返利需要达到的条件
	ReturnMoney     float64 // 返利金额
}
func NewCashReturn(returnCondition float64, returnMoney float64) *CashReturn {
	return &CashReturn{
		ReturnCondition: returnCondition,
		ReturnMoney:     returnMoney,
	}
}
func (cr *CashReturn) AcceptMoney (money float64) float64 {
	if money > cr.ReturnCondition{ // 达到满减条件
		moneyMinus :=  int(money / cr.ReturnCondition)
		return money - (float64(moneyMinus) * cr.ReturnMoney)
	}
	return money
}


// 3 最重要的，定义CashContext结构（环境）,做满减筛选
type CashContext struct {
	Strategy CashSuper
}
func NewCashContext(cashType string) *CashContext {
	cashContext := new(CashContext)
	switch cashType {
	case "打八折":
		cashContext.Strategy = NewCashRebate(0.8)
	case "满100-20":
		cashContext.Strategy = NewCashReturn(100,20)
	default:
		cashContext.Strategy = NewCashNormal()
	}
	return cashContext
}
//在策略生产成功后，我们就可以直接调用策略的函数。
// money 购物总金额
func (cc *CashContext) GetMoney(money float64) float64 {
	return cc.Strategy.AcceptMoney(money)
}

// 4 测试
func main() {
	var (
		actualMoney float64
		money float64
	)
	// 没优惠
	money = 200.0
	cashNormal := NewCashContext("没优惠")
	actualMoney = cashNormal.GetMoney(money)
	fmt.Println("没优惠的价格：",actualMoney)

	// 打八折
	money = 200
	cashRebate := NewCashContext("打八折")
	actualMoney = cashRebate.GetMoney(money)
	fmt.Println("打八折后的价格：",actualMoney)

	// 满100-20
	money = 200
	cashReturn := NewCashContext("满100-20")
	actualMoney = cashReturn.GetMoney(money)
	fmt.Println("满100-20后的价格：",actualMoney)
}
