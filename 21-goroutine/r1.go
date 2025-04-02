package goroutine

import (
	"fmt"
	"sync"
	"time"
)

var box sync.WaitGroup

func processFile(filename string) {
	fmt.Println(filename)
	time.Sleep(5 * time.Second)
	box.Done() // take token
}

func f() {
	for i := range 10 {
		box.Add(1) // put token
		go processFile(fmt.Sprintf("file_%d.json", i))
	}

	box.Wait() // check token: if empty quit, if not wait
}
