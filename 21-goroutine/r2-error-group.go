package goroutine

import (
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

type X struct {
	Age int
}

type Pf2 func(int, string) error
type PfX func(int, int) error

func processFile2(num int, filename string) error {
	time.Sleep(2 * time.Second)
	if num%2 == 0 {
		return fmt.Errorf("even number %d is not accepted", num)
	}
	fmt.Println(filename)
	return nil
}

func g() {
	eg := errgroup.Group{}

	for i := range 10 {
		f := func() error {
			err := processFile2(i, fmt.Sprintf("file_%d.json", i))
			return err
		}

		eg.Go(f)
	}

	err := eg.Wait()
	if err != nil {
		fmt.Printf("program failed: %v\n", err)
	}
}
