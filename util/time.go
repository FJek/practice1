package util

import "time"

// 计算时间差 day d2-d1
func SubDay (d2 time.Time,d1 time.Time) int{
	time := d2.Sub(d1)
	return int(time)/68400000
}