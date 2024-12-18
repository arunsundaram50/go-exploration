package fibercloud

import (
	"net/http"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
)

// Initialize the Fiber app and define routes
var app *fiber.App

func init() {
	app = fiber.New()

	// Define a route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World from Fiber!")
	})
}

// FiberHandler is the entry point for the Cloud Function
func FiberHandler(w http.ResponseWriter, r *http.Request) {
	// Use the adaptor to convert Fiber's handler to an HTTP-compatible handler
	adaptor.FiberApp(app)(w, r)
}
