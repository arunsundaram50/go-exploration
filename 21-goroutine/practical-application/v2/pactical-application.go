package v2

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/jszwec/csvutil"
)

type Person struct {
	Name string `csv:"Name"`
	Age  int    `csv:"Age"`
}

func ProcessFile(filename string) float64 {
	fmt.Println("Scheduled file:", filename)
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	var people []Person
	err = csvutil.Unmarshal(data, &people)
	if err != nil {
		panic(err)
	}

	totalAge := 0.0
	for _, person := range people {
		totalAge += float64(person.Age)
		time.Sleep(time.Second)
	}

	averageAge := totalAge / float64(len(people))
	fmt.Println("                           Average age in file", filename, "is", averageAge)
	return averageAge
}

func ProcessFile3(filename string) (string, int) {
	fmt.Println("Scheduled file:", filename)
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	var people []Person
	err = csvutil.Unmarshal(data, &people)
	if err != nil {
		panic(err)
	}

	totalAge := 0.0
	for _, person := range people {
		totalAge += float64(person.Age)
		time.Sleep(time.Second)
	}

	averageAge := totalAge / float64(len(people))
	return filename, int(averageAge)
}

func ProcessFile2(filename string) int {
	fmt.Println("Scheduled file:", filename)
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	var people []Person
	err = csvutil.Unmarshal(data, &people)
	if err != nil {
		panic(err)
	}
	totalAge := 0.0
	for _, person := range people {
		totalAge += float64(person.Age)
		time.Sleep(time.Second)
	}
	averageAge := totalAge / float64(len(people))
	return int(averageAge)
}

func GetFilenames(exts ...string) []string {
	entries, err := os.ReadDir(".")
	if err != nil {
		panic(err)
	}

	var filenames []string
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		for _, ext := range exts {
			if strings.HasSuffix(entry.Name(), ext) {
				filenames = append(filenames, entry.Name())
			}
		}
	}

	return filenames
}
