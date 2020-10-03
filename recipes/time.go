package recipes

// https://golang.org/pkg/time/

import (
	"fmt"
	"time"
)

func getTime() time.Time {
	return time.Now()
}

func addDaysToDate(days int) time.Time {
	now := time.Now()

	return now.AddDate(0, 0, days)
}

func addMonthsToDate(months int) time.Time {
	now := time.Now()

	return now.AddDate(0, months, 0)
}

func addDaysAndMonthsToDate(months, days int) time.Time {
	now := time.Now()

	return now.AddDate(0, months, days)
}

func timer() {
	start := time.Now()
	t := time.Now()

	time.Sleep(3 * time.Second)

	since := time.Since(start)
	elapsed := t.Sub(start)
	// since := time.Since(start)

	println("Elapsed: ", elapsed)
	println("Since: ", since)
}

// func stopExecutionSeconds(seconds int) {
// 	time.Sleep(seconds * time.Second)
// }
func logTimes() {
	fmt.Printf("Minute: %s", time.Minute)
	fmt.Printf("Second: %s", time.Second)
	fmt.Printf("Hour: %s", time.Hour)
	fmt.Printf("Millisecond: %s", time.Millisecond)
	fmt.Printf("Microsecond: %s", time.Microsecond)
	fmt.Printf("Nanosecond: %s", time.Nanosecond)
}

func dateDemo() {
	// https://golang.org/pkg/time/
	now := time.Now()
	// Prints duration until t
	time.Sleep(3 * time.Second)

	fmt.Println("Duration until t:", time.Until(now))
	println(now.String())

	// time.Now().Before(deadline)

	start := time.Now()
	// ... operation that takes 20 milliseconds ...
	t := time.Now()
	since := time.Since(start)
	elapsed := t.Sub(start)
	// since := time.Since(start)

	println("Elapsed: ", elapsed)
	println("Since: ", since)

	fmt.Println("\nToday:", now)

	after := now.AddDate(1, 0, 0)
	fmt.Println("\nAdd 1 Year:", after)

	after = now.AddDate(0, 1, 0)
	fmt.Println("\nAdd 1 Month:", after)

	after = now.AddDate(0, 0, 1)
	fmt.Println("\nAdd 1 Day:", after)

	after = now.AddDate(2, 2, 5)
	fmt.Println("\nAdd multiple values:", after)

	after = now.Add(10 * time.Minute)
	fmt.Println("\nAdd 10 Minutes:", after)

	after = now.Add(10 * time.Second)
	fmt.Println("\nAdd 10 Second:", after)

	after = now.Add(10 * time.Hour)
	fmt.Println("\nAdd 10 Hour:", after)

	after = now.Add(10 * time.Millisecond)
	fmt.Println("\nAdd 10 Millisecond:", after)

	after = now.Add(10 * time.Microsecond)
	fmt.Println("\nAdd 10 Microsecond:", after)

	after = now.Add(10 * time.Nanosecond)
	fmt.Println("\nAdd 10 Nanosecond:", after)
}
