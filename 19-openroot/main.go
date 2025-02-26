package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
)

func f(relPath string) {
	entries, err := os.ReadDir("arun/" + relPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		fmt.Println(e.Name())
	}
}

func g(relPath string) {
	arunsRoot, err := os.OpenRoot("arun")
	if err != nil {
		log.Fatalf("error opening root for arun: %v", err)
	}

	entries, err := fs.ReadDir(arunsRoot.FS(), relPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		fmt.Println(e.Name())
	}

	fp, err := arunsRoot.Create("globe.txt")
	if err != nil {
		log.Fatal(err)
	}

	fp.WriteString("Hello globe")
	fp.Close()
}

func main() {
	dirToList := os.Args[1]
	g(dirToList)
}
