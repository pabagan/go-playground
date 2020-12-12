package recipes

import (
	"fmt"
	"sync"
	"time"
)

type TokenCache struct {
	token      string
	expiration time.Time
	lock       *sync.Mutex
}

func (t *TokenCache) Start(done chan bool) {
	for {
		select {
		case <-done:
			fmt.Println("closing goroutine")
			return
		default:
			if time.Now().After(t.expiration) {
				t.lock.Lock()
				fmt.Println("getting new token")

				// do something
				t.token = "randomGeneratedToken"

				t.expiration = time.Now().Add(2 * time.Second)
				t.lock.Unlock()
			}
		}
	}
}

func (t *TokenCache) GetToken() string {
	t.lock.Lock()
	defer t.lock.Unlock()
	return t.token
}

func TestChannelServerLevelTask() {
	done := make(chan bool)
	t := TokenCache{
		token:      "randomGeneratedToken",
		expiration: time.Now(),
		lock:       &sync.Mutex{},
	}

	// spawn our process
	go t.Start(done)

	for i := 0; i < 3; i++ {
		fmt.Println(t.GetToken())
		time.Sleep(1 * time.Second)
	}

	close(done)
	fmt.Println("shutdown")
}
