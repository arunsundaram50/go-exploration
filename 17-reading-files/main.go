package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/jszwec/csvutil"
)

type Person struct {
	Name string
	Age  int
}

func readCsv() {
	var persons []Person
	data, _ := os.ReadFile("my.csv")
	csvutil.Unmarshal(data, &persons)
	fmt.Println(persons)
}

func readJson() {
	var persons []Person
	data, _ := os.ReadFile("my.json")
	json.Unmarshal(data, &persons)
	fmt.Println(persons)
}

func readText() {
	var lines []string
	data, _ := os.ReadFile("my.txt")
	s := string(data)
	lines = strings.Split(s, "\n")
	fmt.Println(lines[0])
	fmt.Println(lines[1])
}

func main() {
	readCsv()
	readJson()
	readText()
}
