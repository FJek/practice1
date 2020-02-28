package main

import "fmt"

/**
* @Author : awen
* @Date : 2020/2/28 11:27 下午
 */

const (
	WIDTH  = 1920
	HEIGHT = 1080
)

type pixel int
var screen [WIDTH][HEIGHT]pixel

func main() {
	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			screen[x][y] = 0
		}
	}
	fmt.Println(screen)
}
