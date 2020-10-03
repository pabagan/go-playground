package recipes

// https://golang.org/pkg/time/

import (
	"fmt"
	"time"
)

/*
GetTime get current Golang timestamp
*/
func GetTime() time.Time {
	return time.Now()
}

/*
AddDaysToDate add days to date
*/
func AddDaysToDate(
	time time.Time,
	days int,
) time.Time {
	return time.AddDate(0, 0, days)
}

/*
AddMonthsToDate add months to date
*/
func AddMonthsToDate(
	time time.Time,
	months int,
) time.Time {
	return time.AddDate(0, months, 0)
}

/*
AddDaysAndMonthsToDate add days to date
*/
func AddDaysAndMonthsToDate(
	time time.Time,
	months int,
	days int,
) time.Time {
	return time.AddDate(0, months, days)
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
