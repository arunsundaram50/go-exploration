package goroutine

import (
	"fmt"
	"sync"
	"time"
)

var box sync.WaitGroup

func processFile(num int, filename string) error {
	defer box.Done()
	time.Sleep(1 * time.Second)
	if num == 5 {
		return fmt.Errorf("5 is not accepted")
	}
	fmt.Println(filename)
	return nil
}

func f() {
	for i := range 10 {
		box.Add(1) // put token
		go func(i int) {
			err := processFile(i, fmt.Sprintf("file_%d.json", i))
			if err != nil {
				fmt.Println(err)
			}
		}(i)
	}

	box.Wait() // check token: if empty quit, if not wait
}
