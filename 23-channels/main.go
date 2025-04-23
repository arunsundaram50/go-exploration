package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Square struct {
	Input  int
	Output int
}

func fStructChan(n int, ch chan Square) {
	time.Sleep(time.Second)
	s := Square{
		Input:  n,
		Output: n * n,
	}
	ch <- s
}

func fStructChanWg(n int, ch chan Square, wg *sync.WaitGroup) {
	time.Sleep(time.Second)
	s := Square{
		Input:  n,
		Output: n * n,
	}
	ch <- s
	wg.Done()
}

func mainUsingStruct() {
	ch := make(chan Square)

	n := rand.Intn(5) + 5
	for i := range n {
		go fStructChan(i, ch)
	}

	for range n {
		val := <-ch
		fmt.Println("square of", val.Input, "is", val.Output)
	}
}

func mainUnknownNBuggy() {
	ch := make(chan Square)

	for i := range rand.Intn(5) + 5 {
		go fStructChan(i, ch)
	}

	for val := range ch {
		fmt.Println("square of", val.Input, "is", val.Output)
	}
}

func mainUnknownNUsingWg() {
	tunnel := make(chan Square)
	tokenBox := sync.WaitGroup{}

	for i := range rand.Intn(50) + 5 {
		tokenBox.Add(1)
		go fStructChanWg(i, tunnel, &tokenBox)
	}

	go func() {
		tokenBox.Wait()
		close(tunnel)
	}()

	for val := range tunnel {
		fmt.Println("square of", val.Input, "is", val.Output)
	}

	fmt.Println("complete")
}
