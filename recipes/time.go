package recipes

// https://golang.org/pkg/time/

import (
	"fmt"
	"log"
	"time"
)

/*
  GetTime get current Golang timestamp
*/
func GetTime() time.Time {
	return time.Now()
}

/*
  AddDays add days to date
*/
func AddDays(from time.Time, days int) time.Time {
	return from.AddDate(0, 0, days)
}

/*
  AddMonths add months to date
*/
func AddMonths(from time.Time, months int) time.Time {
	return from.AddDate(0, months, 0)
}

/*
  AddDaysAndMonths add days to date
*/
func AddDaysAndMonths(from time.Time, months, days int) time.Time {
	return from.AddDate(0, months, days)
}

/*
  Timer demo function
*/
func Timer() {
	start := time.Now()
	t := time.Now()

	time.Sleep(3 * time.Second)

	since := time.Since(start)
	elapsed := t.Sub(start)
	// since := time.Since(start)

	fmt.Println("Elapsed: ", elapsed)
	fmt.Println("Since: ", since)
}

/*
  StopExecutionSeconds for a duration of seconds
*/
func StopExecutionSeconds(seconds time.Duration) {
	start := time.Now()
	t := time.Now()

	time.Sleep(seconds * time.Second)

	since := time.Since(start)
	elapsed := t.Sub(start)
	// since := time.Since(start)

	fmt.Println("Elapsed: ", elapsed)
	fmt.Println("Since: ", since)
	fmt.Println("Since seconds: ", since.Seconds())
	fmt.Println("Since minutes: ", since.Minutes())
}

/*
  LogTimes in diferent measures
*/
func LogTimes() {
	fmt.Printf("Minute: %s\n", time.Minute)
	fmt.Printf("Second: %s\n", time.Second)
	fmt.Printf("Hour: %s\n", time.Hour)
	fmt.Printf("Millisecond: %s\n", time.Millisecond)
	fmt.Printf("Microsecond: %s\n", time.Microsecond)
	fmt.Printf("Nanosecond: %s\n", time.Nanosecond)
}

func TestTimes() {
	now := GetTime()
	add1day := AddDays(now, 1)
	add1month := AddMonths(now, 1)
	add1monthAnd2Days := AddDaysAndMonths(now, 1, 2)

	LogTimes()

	log.Println("now:", now)
	log.Println("add 1 day:", add1day)
	log.Println("add 1 month:", add1month)
	log.Println("add 1 month and 2 days:", add1monthAnd2Days)

	StopExecutionSeconds(2)
}
