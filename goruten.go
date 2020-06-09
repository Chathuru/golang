package main

import (
	"fmt"
	"time"
	"sync"
)

var wg sync.WaitGroup

func say(s string) {
	for i := 0; i < 4; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond*100)
	}
	wg.Done()
}

func main() {
	wg.Add(1)
	go say("blt")
	wg.Add(1)
	go say("cms")
	wg.Wait()
}