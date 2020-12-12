package recipes

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func Channel() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}

func ChannelSRequestLevelTaskimple() {
	resCh := make(chan bool)

	go func(ch chan<- bool) {
		time.Sleep(5 * time.Second)
		ch <- true
	}(resCh)

	res := <-resCh
	fmt.Println(res)
}

func ChannelsRequestLevelTaskWithWaitGroupStatic() {
	resCh := make(chan string, 2)

	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	go func(ch chan<- string, wg *sync.WaitGroup) {
		defer wg.Done()
		time.Sleep(1 * time.Second)
		ch <- "Hello"
	}(resCh, &waitGroup)

	go func(ch chan<- string, wg *sync.WaitGroup) {
		defer wg.Done()
		time.Sleep(1 * time.Second)
		ch <- "World"
	}(resCh, &waitGroup)

	waitGroup.Wait()

	for len(resCh) > 0 {
		res := <-resCh
		fmt.Println(res)
	}
}

func ChannelsRequestLevelTaskWaitGroupDynamic(bufferSize int) {
	wordCh := make(chan string, bufferSize)
	words := strings.Split("the quick brown fox jumped over the lazy dog", " ")

	var waitGroup sync.WaitGroup

	// what is wrong here?
	for _, word := range words {
		waitGroup.Add(1)
		go func(ch chan<- string, wg *sync.WaitGroup, currentWord string) {
			defer wg.Done()
			time.Sleep(1 * time.Second)
			ch <- currentWord
		}(wordCh, &waitGroup, word)
	}

	done := make(chan bool)
	go func(d chan bool) {
		defer close(d)
		for res := range wordCh {
			fmt.Printf("%s ", res)
		}
	}(done)

	waitGroup.Wait()
	close(wordCh)
	<-done
}

func ChannelRequestLevelTaskWithContext() {
	// we ignore the cancelFunc here but in a real scenario, we'd want to invoke the cancelFunc if we finish within the deadline
	newCtx, _ := context.WithDeadline(
		context.Background(),
		time.Now().Add(1*time.Second),
	)

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("close goroutine")
				return
			}
			// do something else
		}
	}(newCtx)

	<-newCtx.Done()
	fmt.Println("finished")
}

func TestChannelRequestLevelTask() {
	ChannelRequestLevelTaskWithContext()
}
