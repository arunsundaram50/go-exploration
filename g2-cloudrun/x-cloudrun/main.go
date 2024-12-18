package main

import (
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

var logger *zap.Logger

func main() {
	var err error
	logger, err = zap.NewProduction() // Use zap.NewDevelopment() for human-readable logs
	if err != nil {
		panic("Failed to initialize Zap logger: " + err.Error())
	}
	defer logger.Sync() // Flushes buffer, if any

	// Create a new Fiber app
	app := fiber.New()

	// Define the root endpoint
	app.Get("/", func(c *fiber.Ctx) error {
		logger.Info("Root endpoint accessed")
		return c.SendString("Hello, World!")
	})

	// Define the /add/:num1/:num2 endpoint
	app.Get("/add/:num1/:num2", func(c *fiber.Ctx) error {
		logger.Info("Add endpoint accessed")
		// Extract parameters
		num1, err1 := strconv.Atoi(c.Params("num1"))
		num2, err2 := strconv.Atoi(c.Params("num2"))

		// Handle invalid parameters
		if err1 != nil || err2 != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Parameters must be integers")
		}

		// Compute the sum and return the result
		result := num1 + num2
		return c.SendString(strconv.Itoa(result))
	})

	// Start the server on the specified port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default to 8080 if PORT is not set
	}
	logger.Info("Server starting", zap.String("port", port))
	if err := app.Listen(":" + port); err != nil {
		logger.Fatal("Server failed to start", zap.Error(err))
	}
}
