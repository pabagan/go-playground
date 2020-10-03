package main

import (
	"log"

	"github.com/pabagan/go-playground/recipes"
)

func main() {
	now := recipes.GetTime()
	add1day := recipes.AddDays(now, 1)
	add1month := recipes.AddMonths(now, 1)
	add1monthAnd2Days := recipes.AddDaysAndMonths(now, 1, 2)

	recipes.LogTimes()

	log.Println("now:", now)
	log.Println("add 1 day:", add1day)
	log.Println("add 1 month:", add1month)
	log.Println("add 1 month and 2 days:", add1monthAnd2Days)

	recipes.StopExecutionSeconds(2)
}
