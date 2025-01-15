package main

import (
	"webapp-curd/arithmetic"

	"github.com/gofiber/fiber/v2"
)

func indexFunction(ctx *fiber.Ctx) error {
	return ctx.Status(200).SendString("hello")
}

func main() {
	app := fiber.New()
	app.Get("/", indexFunction)
	app.Get("/add", arithmetic.AddFunction)
	app.Get("/avg", arithmetic.AvgFunction)
	app.Listen(":8080")
}
