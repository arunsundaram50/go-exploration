package atesting

import "fmt"

func getGreeting() string {
	return "Hello, world"
}

func add(a, b int) int {
	return a + b
}

func main() {
	greeting := getGreeting()
	fmt.Printf("Greeing is %s\n", greeting)

	x := 1
	y := 2

	sum := add(x, y)
	fmt.Printf("Sum is of %d + %d = %d\n", x, y, sum)
}
