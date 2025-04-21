package practicalapplication

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetFilenamesLength(t *testing.T) {
	filenames := GetFilenames()
	require.Equal(t, 3, len(filenames))
}

func TestGetFilenamesNotEmpty(t *testing.T) {
	filenames := GetFilenames()
	require.Equal(t, 3, len(filenames))
	require.NotEqual(t, "", filenames[0])
	require.NotEqual(t, "", filenames[1])
	require.NotEqual(t, "", filenames[2])
}

func TestGetFilenamesAreValid(t *testing.T) {
	filenames := GetFilenames()
	require.Equal(t, 3, len(filenames))

	for _, filename := range filenames {
		_, err := os.Stat(filename)
		if os.IsNotExist(err) {
			require.NoError(t, err, "file "+filename+" does not exist")
		}
	}
}

func TestGetFilenamesAreUnique(t *testing.T) {
	filenames := GetFilenames()
	require.Equal(t, 3, len(filenames))

	uniqueFilenames := make(map[string]struct{})
	for _, filename := range filenames {
		uniqueFilenames[filename] = struct{}{}
	}

	require.Equal(t, len(uniqueFilenames), len(filenames), "filenames are not unique")
}

func TestGetFilenamesAreCSV(t *testing.T) {
	filenames := GetFilenames()
	require.Equal(t, 3, len(filenames))

	for _, filename := range filenames {
		require.True(t, isCSV(filename), filename+" is not a CSV file")
	}
}

func isCSV(filename string) bool {
	// Check if the file has a .csv extension
	if len(filename) < 4 {
		return false
	}
	return filename[len(filename)-4:] == ".csv"
}
