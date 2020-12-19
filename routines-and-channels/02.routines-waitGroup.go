package routinesAndChannels

import (
	"fmt"
	"sync"
	"time"
)

func TestRoutinesWaitGroup() {
	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		count("paella")
		wg.Done()
	}()

	go func() {
		count("cocido")
		wg.Done()
	}()

	wg.Wait()
}

func count(thing string) {
	for i := 1; true; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}
