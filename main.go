package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func main() {
	stringTime := "2017-08-30 16:40:41"
	loc, _ := time.LoadLocation("Local")
	the_time, err := time.ParseInLocation("2006-01-02 15:04:05", stringTime, loc)
	if err == nil {
		//unix_time := the_time.Unix() //1504082441
		fmt.Println(the_time)
	}

}
