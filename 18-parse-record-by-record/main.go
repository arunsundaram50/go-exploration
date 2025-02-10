package main

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/jszwec/csvutil"
)

type Sep struct {
	Ticker string  `csv:"ticker"`
	Date   string  `csv:"date"`
	Open   string  `csv:"open"`
	Close  string  `csv:"close"`
	Volume float64 `csv:"volume"`
}

/*
os.ReadFile (reads entire file)
json.Unmarshal or csvutil.Unmarshal (reads entire data from the file)
json.Decode or csvutil.Decode (one record at a time)
*/
func main() {
	filename := "/Users/arunsundaram/data/investments/SHARADAR/SEP.csv"
	fileData, err := os.Open(filename) // name => fp (file details)
	if err != nil {
		panic(err)
	}
	characterDecoder := csv.NewReader(fileData)           // fp => file character decoder
	csvDecoder, _ := csvutil.NewDecoder(characterDecoder) // characters => csv record decoder

	i := 0
	for {
		var sep Sep
		err := csvDecoder.Decode(&sep) // ask csv decoder to decode the record
		if err != nil {
			fileData.Close()
			panic(err)
		}
		fmt.Println(i, sep)
		i++
	}
}
