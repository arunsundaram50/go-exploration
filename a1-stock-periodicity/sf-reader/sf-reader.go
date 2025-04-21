package sfreader

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
)

/*
CSV format:
ticker,date,lastupdated,ev,evebit,evebitda,marketcap,pb,pe,ps
A,2021-09-01,2021-09-01,54542.7,42.8,34.4,53112.7,10.7,53.6,8.6
AA,2021-09-01,2021-09-01,9086.1,10.2,5.8,8521.1,2.3,19.8,0.8
*/
type Daily struct {
	Ticker      string  `csv:"ticker"`
	Date        string  `csv:"date"`
	LastUpdated string  `csv:"lastupdated"`
	EV          float64 `csv:"ev"`
	EVEbit      float64 `csv:"evebit"`
	EVEbitDA    float64 `csv:"evebitda"`
	MarketCap   float64 `csv:"marketcap"`
	PB          float64 `csv:"pb"`
	PE          float64 `csv:"pe"`
	PS          float64 `csv:"ps"`
}

// Create a decoder for the Daily struct since PS can be an empty string
func (d *Daily) UnmarshalCSV(record []string) error {
	if len(record) != 10 {
		return fmt.Errorf("expected 10 fields, got %d", len(record))
	}
	d.Ticker = record[0]
	d.Date = record[1]
	d.LastUpdated = record[2]
	d.EV = parseFloat(record[3])
	d.EVEbit = parseFloat(record[4])
	d.EVEbitDA = parseFloat(record[5])
	d.MarketCap = parseFloat(record[6])
	d.PB = parseFloat(record[7])
	d.PE = parseFloat(record[8])
	d.PS = parseFloat(record[9])
	return nil
}

func (d *Daily) String() string {
	return fmt.Sprintf("Ticker: %s, Date: %s, LastUpdated: %s, EV: %.2f, EVEbit: %.2f, EVEbitDA: %.2f, MarketCap: %.2f, PB: %.2f, PE: %.2f, PS: %.2f",
		d.Ticker, d.Date, d.LastUpdated, d.EV, d.EVEbit, d.EVEbitDA, d.MarketCap, d.PB, d.PE, d.PS)
}

type DailyObserver func(int, *Daily) bool

func ReadCSVRecords(filePath string, processRecord func([]string) bool) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	csvReader := csv.NewReader(bytes.NewReader(data))
	csvReader.FieldsPerRecord = -1 // Allow variable number of fields
	csvReader.ReuseRecord = true
	_, err = csvReader.Read() // Read the header
	if err != nil {
		return err
	}

	for {
		record, err := csvReader.Read()
		if err != nil {
			break
		}
		if !processRecord(record) {
			// Stop processing when processRecord returns false
			return nil
		}
	}

	return nil
}

func ForEachDaily(o DailyObserver) error {
	filePath := filepath.Join(PREFIX, "DAILY.csv")
	recordNo := -1
	var errCaptured error

	errCaptured = ReadCSVRecords(filePath, func(record []string) bool {
		daily := Daily{}
		errCaptured = daily.UnmarshalCSV(record)
		if errCaptured != nil {
			// Return the UnmarshalCSV error to the caller
			return false // Stop processing
		}
		recordNo++
		return o(recordNo, &daily)
	})

	return errCaptured
}
