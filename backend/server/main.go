package main

import (
	"fmt"
	"momentum/middleware"
	"momentum/server/services/auth"
	"momentum/server/services/user"
	"momentum/server/storage"
	"momentum/utilities"
	"os"
	"os/signal"
	"syscall"

	go_json "github.com/goccy/go-json"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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

func initializeBaseMiddleware(app *fiber.App, settings utilities.ApplicationSettings) {
	// Setup CORS middleware:
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	// Middleware for validating tokens:
	app.Use(middleware.NewToken(settings.AuthSettings))
}

func initializeHealthCheck(app *fiber.App) {
	app.Get("/health", func(ctx *fiber.Ctx) error {
		return ctx.SendString("OK")
	});
}

func initializeServices(app *fiber.App, db *storage.PostgresDB, settings utilities.ApplicationSettings) {
	// User service:
	userService := userService.NewUserService(db);
	userMiddleware := middleware.NewUser(settings.AuthSettings);
	userService.InitializeRoutes(app, userMiddleware);
	
	// Auth service:
	authService := authService.NewAuthService(db, settings.AuthSettings);
	authService.InitializeRoutes(app);
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

	// Initialize the middleware:
	initializeBaseMiddleware(app, settings);
	
	// Initialize the database
	db := storage.NewPostgresDB(settings.Database);
	defer db.Close()

	// Initialize the handlers:
	initializeHealthCheck(app);
	initializeServices(app, db, settings);

	// Start the server:
	go func() {
		app.Listen(":8080")
	}()
	
	// Graceful shutdown
	gracefulShutdown(app)
}