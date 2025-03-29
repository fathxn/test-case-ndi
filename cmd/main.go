package main

import (
	"fmt"
	"log"
	"test-case-ndi/internal/config"
	"test-case-ndi/internal/delivery/http"
	"test-case-ndi/internal/middleware"
	"test-case-ndi/internal/repository"
	"test-case-ndi/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// load config
	cfg := config.LoadConfig()

	// dependency injection
	userRepo := repository.NewUserRepository()
	userUseCase := usecase.NewUserUsecase(userRepo, cfg.JWTSecret)
	userHandler := http.NewUserHandler(userUseCase, cfg.JWTSecret)

	// setup middleware
	authMiddleware := middleware.NewAuthMiddleware(middleware.AuthConfig{
		JWTSecret: cfg.JWTSecret,
	})

	app := fiber.New(fiber.Config{
		AppName: "Bank App",
	})

	app.Use(logger.New())

	userHandler.SetupRoutes(app, authMiddleware)

	serverAddr := fmt.Sprintf(":%s", cfg.ServerPort)
	log.Printf("Server starting on %s", serverAddr)
	if err := app.Listen(serverAddr); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
