package practicalapplication

import (
	"os"
	"strings"
)

func GetFilenames() []string {
	entries, err := os.ReadDir(".")
	if err != nil {
		panic(err)
	}

	var filenames []string
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		if strings.HasSuffix(entry.Name(), ".csv") {
			filenames = append(filenames, entry.Name())
		}
	}

	return filenames
}
