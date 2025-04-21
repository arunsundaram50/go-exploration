package sqldbstorage

import (
	"fmt"
	"os"
	"testing"
)

func divide(num, den int) (int, error) {
	if den == 0 {
		return 0, fmt.Errorf("den can't be zero!")
	}

	answer := num / den
	return answer, nil
}

func TestDefer(t *testing.T) {

	type operands struct {
		num int
		den int
	}

	arOperands := []operands{
		{300, 10},
		{400, 10},
		{500, 50},
	}

	var err error
	fp, err := os.OpenFile("x.txt", os.O_CREATE|os.O_RDWR, 0o644)
	if err != nil {
		panic(err)
	}

	defer func() {
		fp.Close()
		if err != nil {
			os.Remove("x.txt")
		}
	}()

	var answer int
	for _, operand := range arOperands {
		answer, err = divide(operand.num, operand.den)
		if err != nil {
			return
		}

		_, err = fp.WriteString(fmt.Sprintf("answer = %d\n", answer))
		if err != nil {
			return
		}
	}

}
