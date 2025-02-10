package fileparsers

import (
	"fmt"
	records "goquandl/records/v2"
)

type Ticker struct {
	Ticker       string `csv:"ticker"`
	Name         string `csv:"name"`
	IsDelisted   string `csv:"isdelisted"`
	Category     string `csv:"category"`
	SicIndustry  string `csv:"sicindustry"`
	FamaSector   string `csv:"famasector"`
	FamaOndustry string `csv:"famaindustry"`
	Sector       string `csv:"sector"`
	Industry     string `csv:"industry"`
	MarketCap    string `csv:"scalemarketcap"`
	Revenue      string `csv:"scalerevenue"`
	CompanySite  string `csv:"companysite"`
	Location     string `csv:"location"`
	FirstAdded   string `csv:"firstadded"`
	LastUpdated  string `csv:"lastupdated"`
}

func (t Ticker) String() string {
	return fmt.Sprintf("%s, %s, %s, %s, %s, %s, %s, %s, %s, %s",
		t.Ticker, t.Name, t.IsDelisted, t.Category, t.MarketCap, t.Revenue, t.Category,
		t.Location, t.FirstAdded, t.LastUpdated)
}

func init() {
	count := 0
	tickerProcessor := func(pos int, t *Ticker, err error) bool {
		if err != nil {
			return false
		}
		if t.IsDelisted == "Y" {
			return false
		}
		fmt.Println(pos, count, *t)
		count++
		return false
	}

	records.Register("tickers", tickerProcessor)
}
