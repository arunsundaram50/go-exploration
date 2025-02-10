package efficientbasic

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func AddFunction(ctx *fiber.Ctx) error {
	a := ctx.Query("a")
	b := ctx.Query("b")
	aa, _ := strconv.Atoi(a)
	bb, _ := strconv.Atoi(b)

	if aa == 100 || bb == 100 {
		return ctx.Status(500).SendString("100 is not allowed")
	}

	sum := aa + bb
	sumStr := strconv.Itoa(sum)
	return ctx.Status(200).SendString(sumStr)
}
