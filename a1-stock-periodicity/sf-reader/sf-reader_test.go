package sfreader

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestForEachDaily(t *testing.T) {
	count := 0
	dailyObserver := func(i int, d *Daily) bool {
		count++
		fmt.Printf("Daily %d: %+v\n", i+1, d)
		return i < 99 // stop after 100 records
	}
	err := ForEachDaily(dailyObserver)
	require.NoError(t, err)
	require.Equal(t, 100, count, "ForEachDaily should stop after 100 records")
}
