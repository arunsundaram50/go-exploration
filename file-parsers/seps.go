package fileparsers

import (
	"fmt"
	records "goquandl/records/v2"

	"github.com/dustin/go-humanize"
)

type Sep struct {
	Ticker      string  `csv:"ticker"`
	Date        string  `csv:"date"`
	Open        float32 `csv:"open"`
	High        float32 `csv:"high"`
	Low         float32 `csv:"low"`
	Close       float32 `csv:"close"`
	Volume      float32 `csv:"volume"`
	LastUpdated string  `csv:"lastupdated"`
}

func (sep Sep) String() string {
	return fmt.Sprintf("%s, %s, %.2f, %.2f, %.0f", sep.Ticker, sep.Date, sep.Open, sep.Close, sep.Volume)
}

func init() {
	count := 0
	sepProcessor := func(pos int, sep *Sep, err error) bool {
		if sep.Ticker == "AAPL" {
			fmt.Printf("%s, %s, %+v\n", humanize.Comma(int64(pos)), humanize.Comma(int64(count)), *sep)
			count++
		}
		return false
	}
	records.Register("sep", sepProcessor)
}
