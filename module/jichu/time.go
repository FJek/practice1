package main

import (
	"fmt"
	"time"
)

func main() {
	startTime := time.Now().Add(-time.Duration(5) * time.Hour * 24).Format("2006-01-02")
	fmt.Println(startTime)
}
