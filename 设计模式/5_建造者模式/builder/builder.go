package builder

/**
* @Author : awen
* @Date : 2020/2/22 1:25 上午
 */

type Builder interface {
	NewProduct()            // 创建一个空产品
	BuildWheels()           // 建造轮子
	BuildChassis()          // 建造底盘
	BuildSeat()             // 建造驾驶位

	GetResult() interface{} // 获取建造好的产品
}
