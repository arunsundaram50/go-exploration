package main

import (
	"fmt"

	"github.com/jszwec/csvutil"
)

type UserPii struct {
	Year   int    `csv:"year"`
	Salary int    `csv:"salary"`
	Hobby  string `csv:"hobby"`
}

var xyz = []UserPii{
	{2021, 1000, "Basket ball"},
	{2022, 2000, "Cricket"},
	{2023, 3000, "Cricket"},
	{2024, 4000, "Soccer"},
	{2025, 5000, "Volleyball"},
}

func main() {
	data, err := csvutil.Marshal(xyz)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("CSV = %s\n", string(data))
}
