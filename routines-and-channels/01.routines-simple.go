package routinesAndChannels

import (
	"fmt"
	"time"
)

func TestRoutinesSimple() {
	fmt.Println("TestRoutinesSimple")
	go countSimple("Paella")
	countSimple("Cocido")
}

func countSimple(thing string) {
	for i := 1; true; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}
