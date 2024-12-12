package main

import "github.com/gofiber/fiber/v2"

func index(ctx *fiber.Ctx) error {
	return ctx.Status(200).SendString("hello, world")
}

func greetMe(ctx *fiber.Ctx) error {
	s := ctx.Query("name")
	if s == "John" {
		return ctx.Status(400).SendString("I do not like this name")
	} else {
		return ctx.Status(200).SendString("hello, " + s)
	}
}

func main() {
	webApp := fiber.New()
	webApp.Get("/", index)
	webApp.Get("/greet-me", greetMe)
	webApp.Listen(":8080")
}
