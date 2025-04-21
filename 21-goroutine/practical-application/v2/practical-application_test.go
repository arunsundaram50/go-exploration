package v2

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestProcessFileBasic(t *testing.T) {
	t1 := time.Now()
	limit := 8
	filenames := GetFilenames(".csv")

	wg := sync.WaitGroup{}
	ch := make(chan bool, limit)
	for _, filename := range filenames {
		wg.Add(1)
		ch <- true
		go func(fn string) {
			defer wg.Done()
			defer func() { <-ch }()
			_ = ProcessFile(fn)
		}(filename)
	}

	wg.Wait()
	t2 := time.Now()
	elapsed := t2.Sub(t1)
	fmt.Println("Total time taken:", elapsed)
}

func TestProcessFileBuggy(t *testing.T) {
	t1 := time.Now()
	filenames := GetFilenames(".csv")
	wg := sync.WaitGroup{}

	tunnelOfAvg := make(chan int, 9)
	for _, filename := range filenames {
		wg.Add(1)
		go func(fn string) {
			defer wg.Done()
			tunnelOfAvg <- ProcessFile2(fn)
		}(filename)
	}

	wg.Wait()
	for range tunnelOfAvg {
		avg := <-tunnelOfAvg
		fmt.Println("                           Average age in file is", avg)
	}

	t2 := time.Now()
	elapsed := t2.Sub(t1)
	fmt.Println("Total time taken:", elapsed)
}

func TestProcessFileClean(t *testing.T) {
	t1 := time.Now()
	filenames := GetFilenames(".csv")
	wg := sync.WaitGroup{}

	type info struct {
		filename string
		avg      int
	}
	tunnelOfAvg := make(chan info, len(filenames))
	for _, filename := range filenames {
		wg.Add(1)
		go func(fn string) {
			defer wg.Done()
			fn, avg := ProcessFile3(fn)
			tunnelOfAvg <- info{filename: fn, avg: avg}
		}(filename)
	}

	go func() {
		wg.Wait()
		close(tunnelOfAvg) // Close the channel after all goroutines are done
	}()

	for avg := range tunnelOfAvg {
		fmt.Println("                           Average in file", avg.filename, "is", avg.avg)
	}

	t2 := time.Now()
	elapsed := t2.Sub(t1)
	fmt.Println("Total time taken:", elapsed)
}
