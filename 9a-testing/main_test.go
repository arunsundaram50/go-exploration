package atesting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	sum := add(1, 2)
	assert.Equalf(t, 3, sum, "unexpected value")
}

func TestGetGreeting(t *testing.T) {
	greet := getGreeting()
	assert.Equalf(t, greet, "Hello, world", "unexpected value")
}

func TestMain(t *testing.T) {
	main()
}
