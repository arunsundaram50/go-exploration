package main

import (
	"fmt"
	"os"
	"time"
)

func f1(j int) { // func(int)
	fmt.Printf("Hello, f1 world\n")
}

func f2(i int) { // func(int)
	fmt.Printf("Hello, f2 world\n")
}

func f3(i int) { // func(int)
	fmt.Printf("Hello, %d world\n", i)
}

func main() {

	var fx func(int)

	if os.Args[1] == "f1" {
		fx = f1
	} else if os.Args[1] == "f2" {
		fx = f2
	} else if os.Args[1] == "f3" {
		fx = f3
	}

	go fx(1)
	go fx(2)
	go fx(3)

	time.Sleep(1 * time.Second)
}
