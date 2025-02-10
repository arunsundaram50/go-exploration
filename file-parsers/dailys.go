package fileparsers

import (
	"fmt"
	records "goquandl/records/v2"

	"github.com/dustin/go-humanize"
)

type Daily struct {
	Ticker          string `csv:"ticker"`
	Date            string `csv:"date"`
	LastUpdated     string `csv:"lastupdated"`
	EnterpriseValue string `csv:"ev"`
	EveBit          string `csv:"evebit"`
	EveBitda        string `csv:"evebitda"`
	MarketCap       string `csv:"marketcap"`
	PriceBook       string `csv:"pb"`
	PriceEarnings   string `csv:"pe"`
	PriceSales      string `csv:"ps"`
}

func (d Daily) String() string {
	return fmt.Sprintf("%s, %s, %s", d.Ticker, d.Date, d.EnterpriseValue)
}

func init() {
	count := 0
	dailyProcessor := func(pos int, daily *Daily, err error) bool {
		if daily.Ticker == "AAPL" {
			fmt.Printf("%s, %s, %+v\n", humanize.Comma(int64(pos)), humanize.Comma(int64(count)), *daily)
			count++
		}
		return false
	}
	records.Register("daily", dailyProcessor)
}
