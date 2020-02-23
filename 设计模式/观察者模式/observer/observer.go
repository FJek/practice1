package observer

import "fmt"

/**
* @Author : awen
* @Date : 2020/2/23 11:14 上午
 */

// 观察者--用户
type Customer interface {
	Update(subject string)
}

// 具体观察者A
type CustomerA struct{}
func (a *CustomerA) Update(subject string) {
	fmt.Println("我是客户A,我接收到了"+subject+"的新推送")
}

// 具体观察者B
type CustomerB struct{}
func (b *CustomerB) Update(subject string) {
	fmt.Println("我是客户B,我接收到了"+subject+"的新推送")
}
