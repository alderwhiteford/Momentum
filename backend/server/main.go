package main

import (
	"fmt"
	"momentum/server/services/user"
	"momentum/server/storage"
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

func initializeHealthCheck(app *fiber.App) {
	app.Get("/health", func(ctx *fiber.Ctx) error {
		return ctx.SendString("OK")
	});
}

func initializeServices(app *fiber.App, db *storage.PostgresDB) {
	// User service:
	user.NewUserService(db).InitializeRoutes(app);
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
	// Initialize the app and settings:
	app := createNewFiberApp();
	settings := utilities.Settings("config");
	
	// Initialize the database
	db := storage.NewPostgresDB(settings.Database);
	defer db.Close()

	// Initialize the handlers:
	initializeHealthCheck(app);
	initializeServices(app, db);

	// Start the server:
	go func() {
		app.Listen(":8080")
	}()
	
	// Graceful shutdown
	gracefulShutdown(app)
}