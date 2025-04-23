package main

import (
	"fmt"
	"time"
)

/*
  sq <- ========= <- res
*/

func f(n int) int {
	time.Sleep(time.Second)
	res := n * n
	return res
}

func fIntChan(n int, ch chan int) {
	res := f(n)
	ch <- res
}

type Square struct {
	Input  int
	Output int
}

func fStructChan(n int, ch chan Square) {
	res := f(n)
	square := Square{
		Input:  n,
		Output: res,
	}
	ch <- square
}

func mainSync() {
	ch := make(chan int)
	for i := range 10 {
		go fIntChan(i, ch)
		sq := <-ch
		fmt.Println("Square of", i, "is", sq)
	}
}

func mainBuggy() {
	ch := make(chan int)
	for i := range 10 {
		go fIntChan(i, ch)
	}

	for i := range 10 {
		sq := <-ch
		fmt.Println("Square of", i, "is", sq)
	}
}

func mainUsingStruct() {
	ch := make(chan Square)

	for i := range 10 {
		go fStructChan(i, ch)
	}

	for range 10 {
		val := <-ch
		fmt.Println("square of", val.Input, "is", val.Output)
	}
}
