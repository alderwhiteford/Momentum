package main

import (
	"fmt"
	"momentum/utilities"
	"os"
	"os/signal"
	"syscall"

	go_json "github.com/goccy/go-json"

	"github.com/gofiber/fiber/v2"
)

func createNewFiberApp() *fiber.App {
	app := fiber.New(fiber.Config{
		AppName: "Momentum",
		CaseSensitive: true,
		ErrorHandler: utilities.ErrorHandler,
		JSONEncoder: go_json.Marshal,
		JSONDecoder: go_json.Unmarshal,
	})

	return app
}

func setupHealthCheck(app *fiber.App) {
	app.Get("/health", func(ctx *fiber.Ctx) error {
		return ctx.SendString("OK")
	});
}

func gracefulShutdown(app *fiber.App) {
	quit := make(chan os.Signal, 1)
	// Listen for SIGINT and SIGTERM signals on the OS
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	if err := app.Shutdown(); err != nil {
		panic(fmt.Errorf("failed to shutdown the server: %s", err))
	}

	fmt.Println("Server is shutting down...")
}

func main() {
	app := createNewFiberApp()

	// Handler setup:
	setupHealthCheck(app)

	// TODO: Configure based on server configuration
	go func() {
		app.Listen(":8080")
	}()
	
	// Graceful shutdown
	gracefulShutdown(app)
}