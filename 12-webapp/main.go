package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func index(ctx *fiber.Ctx) error {
	return ctx.Status(201).SendString("hello, world")
}

func greetMe(ctx *fiber.Ctx) error { // func(*Ctx) error
	s := ctx.Query("name")
	if s == "John" {
		return ctx.Status(400).SendString("I do not like this name")
	} else {
		return ctx.Status(200).SendString("hello, " + s)
	}
}

func add(ctx *fiber.Ctx) error {
	aStr := ctx.Query("a")
	bStr := ctx.Query("b")

	aInt, err := strconv.Atoi(aStr)
	if err != nil {
		return err
	}
	bInt, err := strconv.Atoi(bStr)
	if err != nil {
		return err
	}

	totalInt := aInt + bInt
	totalStr := strconv.Itoa(totalInt)

	return ctx.Status(200).SendString(totalStr)
}

func main() {
	webApp := fiber.New()
	webApp.Get("/", index)
	webApp.Get("/greet-me", greetMe)
	webApp.Get("/add", add)
	webApp.Listen(":8080")
}
