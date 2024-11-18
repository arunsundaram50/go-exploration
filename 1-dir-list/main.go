package main

import (
	"os"
)

func main() {
	for i, dirName := range os.Args[1:] {
		println(i, dirName)
		println("============")
		entries, err := os.ReadDir(dirName)
		if err == nil {
			for _, e := range entries {
				println(e.Name())
			}
		} else {
			println("error:")
			println(err.Error())
		}
	}
}
