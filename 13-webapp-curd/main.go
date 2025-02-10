package main

import (
	b "webapp-curd/arithmetic/basic"

	"github.com/gofiber/fiber/v2"
)

func indexFunction(ctx *fiber.Ctx) error {
	return ctx.Status(200).SendString("hello")
}

func main() {
	app := fiber.New()
	app.Get("/", indexFunction)
	app.Get("/add", b.AddFunction)
	app.Get("/avg", b.AvgFunction)
	app.Listen(":8080")
}
