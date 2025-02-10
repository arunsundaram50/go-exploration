package records

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"sync"

	"github.com/jszwec/csvutil"
)

type RecordProcessor[T any] func(int, *T, error) bool
type DATASET_NAME string

func NewDatasetNameFromString(str string) DATASET_NAME {
	return DATASET_NAME(strings.ToLower(str))
}

func NewDatasetNameFromType[T any]() DATASET_NAME {
	return DATASET_NAME(strings.ToLower(reflect.TypeOf((*T)(nil)).Elem().Name()))
}

func (dsn DATASET_NAME) GetDataFilename() string {
	h, _ := os.UserHomeDir()
	dataFilepath := filepath.Join(h, fmt.Sprintf("data/investments/SHARADAR/%s.csv", strings.ToUpper(string(dsn))))
	if _, err := os.Stat(dataFilepath); !os.IsNotExist(err) {
		return dataFilepath
	}
	// try singular
	if strings.HasSuffix(string(dsn), "s") {
		singular := dsn[0 : len(dsn)-2]
		dataFilepath := filepath.Join(h, fmt.Sprintf("data/investments/SHARADAR/%s.csv", strings.ToUpper(string(singular))))
		if _, err := os.Stat(dataFilepath); !os.IsNotExist(err) {
			return dataFilepath
		}
	}
	// try plural
	dataFilepath = filepath.Join(h, fmt.Sprintf("data/investments/SHARADAR/%s.csv", strings.ToUpper(string(dsn)+"s")))
	if _, err := os.Stat(dataFilepath); !os.IsNotExist(err) {
		return dataFilepath
	}

	return ""
}

var (
	processors              = make(map[DATASET_NAME]any)
	typeRegistry            = make(map[DATASET_NAME]reflect.Type)
	genericWrapperFunctions = make(map[DATASET_NAME]func() error)
	mu                      sync.RWMutex
)

// Register function to add new processors dynamically and store type info
func Register[T any](datasetName DATASET_NAME, processor RecordProcessor[T]) {
	mu.Lock()
	defer mu.Unlock()
	lowerDatasetName := DATASET_NAME(strings.ToLower(string(datasetName)))
	processors[lowerDatasetName] = processor
	typeRegistry[lowerDatasetName] = reflect.TypeOf((*T)(nil)).Elem()

	// Store a callable wrapper function that calls ForEach2[T]
	genericWrapperFunc := func() error {
		return ForEach2[T]()
	}
	genericWrapperFunctions[lowerDatasetName] = genericWrapperFunc
}

func ForEach[T any](filename string, datasetName DATASET_NAME) error {
	mu.RLock()
	proc, exists := processors[datasetName]
	if !exists { // try plural
		proc, exists = processors[datasetName+"s"]
	}
	mu.RUnlock()

	if !exists {
		return fmt.Errorf("processor not found: %s", datasetName)
	}

	typedProc, ok := proc.(RecordProcessor[T])
	if !ok {
		return fmt.Errorf("processor type mismatch for: %s", datasetName)
	}

	file, errOpening := os.Open(filename)
	if errOpening != nil {
		return fmt.Errorf("failed to open file: %w", errOpening)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	decoder, errCreatingDecoder := csvutil.NewDecoder(reader)
	if errCreatingDecoder != nil {
		if typedProc(-1, nil, fmt.Errorf("failed to create decoder: %w", errCreatingDecoder)) {
			return fmt.Errorf("stopped by recordProcessor while initializing")
		}
	}

	recordNumber := -1
	var record T
	for {
		recordNumber++

		errRecordDecoding := decoder.Decode(&record)

		if errRecordDecoding == nil {
			if stop := typedProc(recordNumber, &record, nil); stop {
				break
			} else {
				continue
			}
		}

		if errRecordDecoding.Error() == "EOF" {
			break
		}

		if stop := typedProc(recordNumber, &record, errRecordDecoding); stop {
			break
		}
	}
	return nil
}

// ForEach2 - Automatically infers processor name from type T
func ForEach2[T any]() error {
	datasetName := NewDatasetNameFromType[T]()
	dataFilepath := datasetName.GetDataFilename()
	return ForEach[T](dataFilepath, datasetName)
}

func ForEach3(datasetNameAsStr string) error {
	datasetName := NewDatasetNameFromString(datasetNameAsStr)
	mu.RLock()
	wrapperFunc, exists := genericWrapperFunctions[datasetName]
	if !exists { // try plural
		wrapperFunc, exists = genericWrapperFunctions[datasetName+"s"]
	}
	mu.RUnlock()

	if !exists {
		return fmt.Errorf("unknown processor: %s", datasetName)
	}

	// Call the pre-registered ForEach2[T] function dynamically
	return wrapperFunc()
}
