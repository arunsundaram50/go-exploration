package sfreader

import (
	"os"
	"path/filepath"
)

var PREFIX = ""

func init() {
	// Set the PREFIX variable to the directory where the CSV files are located
	PREFIX = os.Getenv("SF_READER_PREFIX")
	if PREFIX == "" {
		home, _ := os.UserHomeDir()
		PREFIX = filepath.Join(home, "data/investments/SHARADAR")
	}
}
