package main

import (
	"fmt"
	"math"
	"sort"
)

func NbYear(p0 int, percent float64, aug int, p int) int {
	// your code
	var year int = 0
	var population float64 = float64(p0)
	var targetPopulation float64 = float64(p)
	var peopleLeavingOrStaying float64 = float64(aug)
	var percentage float64 = float64(percent / 100)

	for population < targetPopulation {
		year++
		population = population + (population*percentage + peopleLeavingOrStaying)

		// println(year, int(population))
	}

	return year
}

func Gps(s int, x []float64) int {
	var averageSpeed []int
	if len(x) < 2 {
		return 0
	}

	for index := range x {
		if index == 0 {
			continue
		}

		deltaDistance := int(math.Floor(3600*(x[index]-x[index-1]))) / s
		averageSpeed = append(averageSpeed, int(deltaDistance))
	}

	sort.Ints(averageSpeed)

	return averageSpeed[len(averageSpeed)-1]
}

// func init() {
// 	fmt.Println("Init has run")
// }
func NbDig(n int, d int) int {
	var count int
	var listSquared []int
	for i := 1; i <= n; i++ {
		if i*i > n {
			break
		}
		// if i < strings.Contains(
		// 	fmt.Sprintf("%d", i*i),
		// 	fmt.Sprintf("%d", d),
		// ) {
		// 	count++
		// }

		listSquared = append(listSquared, i*i)
	}

	fmt.Println(listSquared)
	return count
}
func FindOutlier(integers []int) int {
	var odd, even []int
	for _, v := range integers {
		if v%2 == 0 {
			odd = append(odd, v)
		} else {
			even = append(even, v)
		}
	}

	if len(odd) == 1 {
		return odd[0]
	}

	return even[0]
}

func Factorial(n int) int {
	// put your code here

	var total int = n

	for i := 1; i < n; i++ {
		total = total * i
	}

	return total
}
func main() {
	// recipes.TestLogs()
	// recipes.TestTimes()
	// recipes.TestGenerics()
	// recipes.TestSeparateHouseNumber()
	// recipes.TestTypeAssertions()
	// recipes.TestChannelRequestLevelTask()
	// recipes.TestChannelServerLevelTask()
	// randomString := recipes.RandomString(10)
	// fmt.Println("Random string", randomString)

	// recipes.RequestBackoff("https://dev.klikkie.com")
	// recipes.GetAccount("https://dev.klikkie.com")
	// recipes.TestPassByValue()
	// recipes.TestTMapCheckProperty()
	// recipes.TestTags()
	// recipes.TestEqualCheck()
	// recipes.
	// recipes.TestCompareInterfaces()
	// recipes.TestAppCLIFlags()
	// routinesAndChannels.TestRoutinesAndChannels()
	// fmt.Println(NbYear(1500000, 2.5, 10000, 2000000))

	// x := []float64{0.0, 0.19, 0.5, 0.75, 1.0, 1.25, 1.5, 1.75, 2.0, 2.25}
	// const s int = 15

	// fmt.Println(Gps(s, x))
	// fmt.Println(NbDig(25, 1))
	// fmt.Println(strings.Count("aa", "a"))
	// fmt.Println(strings.End)

	// listA, endList := strings.Split("abc"), strings.Split("bd")

	// lastChars := strings.Split
	// for i := len(endList); i <= 0; i-- {

	// }
	// fmt.Println(isDigitPresent(101010100, 1))
	// word := "Hola"
	// end := "la"

	// fmt.Println(word[len(word)-len(end) : len(word)])
	// fmt.Println(strings.HasSuffix(word, end))

	// integers := []int{1, 3, 5, 2}

	// fmt.Println(FindOutlier(integers))
	fmt.Println(Factorial(3))
}
