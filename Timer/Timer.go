package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	log("<---- See how the current time is logged? Pretty cool, right? \n\nThis was my first project in Go!\n\nPress ENTER to exit the console.")
	fmt.Scanln()

}

func log(value string) {
	var currentTime time.Time = time.Now()

	var hour int = currentTime.Hour()
	var minute int = currentTime.Minute()
	var second int = currentTime.Second()

	var hourstr string = addZeroAtBeginning(hour)
	var minutestr string = addZeroAtBeginning(minute)
	var secondstr string = addZeroAtBeginning(second)

	fmt.Println("[" + hourstr + ":" + minutestr + ":" + secondstr + "]" + " " + value)

}

func addZeroAtBeginning(value int) string {
	var valuetoreturn string

	if value <= 9 {
		valuetoreturn = "0" + strconv.Itoa(value)
	} else {
		valuetoreturn = strconv.Itoa(value)
	}

	return valuetoreturn
}
