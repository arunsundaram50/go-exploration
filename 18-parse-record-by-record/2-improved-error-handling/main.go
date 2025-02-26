package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"reflect"

	"github.com/jszwec/csvutil"
)

type Sep struct {
	Ticker string  `csv:"ticker"`
	Date   string  `csv:"date"`
	Open   string  `csv:"open"`
	Close  string  `csv:"close"`
	Volume float64 `csv:"volume"`
}

// error of a specific type
type SepError struct {
	errorText    string
	recordNumber int
}

func (e SepError) Error() string {
	return fmt.Sprintf("SEP error at record %d, %s", e.recordNumber, e.errorText)
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

	var ute *csvutil.UnmarshalTypeError
	var mce *csvutil.MissingColumnsError
	i := 0
	for {
		var sep Sep
		err := csvDecoder.Decode(&sep) // ask csv decoder to decode the record
		if err == nil {
			//fmt.Println(i, sep)
			i++
			continue
		}

		if err == io.EOF {
			break
		}

		if errors.As(err, &ute) {
			fmt.Printf("%d could not be unmarshalled %v: [%s] [%s]\n", i, err, ute.Type, ute.Value)
			continue
		}

		if errors.As(err, &mce) {
			fmt.Printf("%d was missing one or more columns\n", i)
			continue
		}

		// simple way of creating an error object
		err1 := errors.New("hello")
		fmt.Printf("err: %v\n", err1)

		// specialized error
		var err2 = SepError{"unknown error", i}
		fmt.Printf("err: %v\n", err2)

		/// wrapped error

		typeName := reflect.TypeOf(err).Name()
		fmt.Printf("Typename=%s\n", typeName)
		fmt.Printf("%d unexpected error for record %+v %+t\n", i, err, err)
		fileData.Close()
		panic(err)
	}
}
