package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time study of golang")
	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006"))
	fmt.Println(presentTime.Format("02-01-2006"))
	fmt.Println(presentTime.Format("02-01-2006 15:04:05 Monday"))

	createDate := time.Date(2020, time.August, 11, 9, 12, 34, 0, time.UTC)
	fmt.Println(createDate)
	fmt.Println(createDate.Format("02-01-2006 15:04:05 Monday"))
}
