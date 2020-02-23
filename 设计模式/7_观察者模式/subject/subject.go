package subject

import "fzw/practice1/设计模式/7_观察者模式/observer"

/**
* @Author : awen
* @Date : 2020/2/23 11:14 上午
 */

// 被观察者--公众号
type IOfficeAccount interface {
	AddCustomer(customer observer.Customer)       // 新增用户
	Publish()           // 发布信息
	NotifyAllCustomer() // 通知观察者(订阅者)
}
// 具体公众号A
type OfficeAccount struct {
	Name string
	customers []observer.Customer
}
func (oa *OfficeAccount) AddCustomer(customer observer.Customer) {
	oa.customers = append(oa.customers,customer)
}
func (oa *OfficeAccount) Publish() {
	oa.NotifyAllCustomer()
}
func (oa *OfficeAccount) NotifyAllCustomer() {
	for _, cus := range oa.customers {
		cus.Update(oa.Name)
	}
}
