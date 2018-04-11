package main

import (
	"fmt"
	"time"
)


func main() {
	currTime := time.Now()
	fmt.Println("current time is ", currTime)

	weekday := currTime.Weekday()
	fmt.Println("Current week day is ", weekday.String())

	if weekday.String() == "Wednesday" {
		fmt.Println("today is ", weekday.String())
	}
}
