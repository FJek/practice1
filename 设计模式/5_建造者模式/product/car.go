package product

import "fmt"

/**
* @Author : awen
* @Date : 2020/2/22 1:28 上午
 */

// 车子
type Car struct {
	Wheels  string // 轮子
	Chassis string // 地盘
	Seats   string //座位
}

func (c *Car) Show() {
	fmt.Println(c.Wheels,c.Seats,c.Chassis)
}
