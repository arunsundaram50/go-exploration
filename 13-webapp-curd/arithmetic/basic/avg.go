package basic

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func AvgFunction(ctx *fiber.Ctx) error {
	var numbers []float64
	ctx.QueryParser(&numbers)

	sum := 0.0
	for _, number := range numbers {
		sum += number
	}

	avg := sum / float64(len(numbers))
	avgStr := fmt.Sprintf("%f", avg)

	return ctx.Status(200).SendString(avgStr)
}
