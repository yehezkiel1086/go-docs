package main

import (
	"fmt"
	"time"
)

func createDate(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}

func parseDate(dateStr string) (time.Time, error) {
	var layout string = "2006-01-02 15:04:05"
	return time.Parse(layout, dateStr)
}

func sleepNSeconds(seconds int) {
	begin := time.Now()
	fmt.Println("Waiting start...")
	time.Sleep(time.Second * time.Duration(seconds))
	fmt.Println("Waiting done")

	end := time.Since(begin)
	fmt.Println("Sleep ends after", end)
}

func simpleScheduler() {
	fmt.Println("Scheduler start")

	for i := 1; i <= 3; i += 1 {
		fmt.Println("Doing something...")
		time.Sleep(time.Second * 2);
	}

	fmt.Println("Scheduler end")
}

func main() {
	var now time.Time = time.Now()
	var birthday time.Time = createDate(2002, 06, 17)
	var dateStr string = "2022-06-01 15:04:05"
	date, err := parseDate(dateStr)
	if err != nil {
		panic(err)
	}

	fmt.Println(now, birthday, date)

	sleepNSeconds(3)
	simpleScheduler()
}
