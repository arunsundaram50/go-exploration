package main

import (
	"fmt"
	"os"

	"github.com/jszwec/csvutil"
)

type Investment struct {
	Ticker       string  `csv:"ticker"`
	InvestorName string  `csv:"investorname"`
	SecurityType string  `csv:"securitytype"`
	CalendarDate string  `csv:"calendardate"`
	Value        float32 `csv:"value"`
	Units        float32 `csv:"units"`
	Price        float32 `csv:"price"`
}

func (i Investment) String() string {
	return fmt.Sprintf("%s, %s, %.2f", i.Ticker, i.InvestorName, i.Price)
}

func main() {
	var investments []Investment

	filename := "SHARADAR/SF3.csv"
	// filename := "x.csv"
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := csvutil.Unmarshal(data, &investments); err != nil {
		fmt.Println(err)
		return
	}

	for _, investment := range investments {
		if investment.Price > 100 {
			fmt.Println(investment)
		}
	}
}
