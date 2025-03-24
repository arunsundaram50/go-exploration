package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
)

func unsafeListDir(relPath string) {
	entries, err := os.ReadDir("arun/" + relPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		fmt.Println(e.Name())
	}
}

func safeListDir(relPath string) {
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
}

func main() {
	dirToList := os.Args[1]
	safeListDir(dirToList)
}
