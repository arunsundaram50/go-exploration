package main

import practicalapplication "goroutine/practical-application/v2"

func main() {
	filenames := practicalapplication.GetFilenames(".csv", ".tsv")
	for _, filename := range filenames {
		println(filename)
	}
}
