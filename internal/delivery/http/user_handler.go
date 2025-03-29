package http

import (
	"test-case-ndi/internal/domain"

	"github.com/gofiber/fiber/v2"
)

// struct untuk user handler
type UserHandler struct {
	userUsecase domain.UserUsecase
}

// constuctor untuk UserHandler
func NewUserHandler(userUsecase domain.UserUsecase, jwtSecret string) *UserHandler {
	return &UserHandler{
		userUsecase: userUsecase,
	}
}

func (h *UserHandler) GetBalance(c *fiber.Ctx) error {
	userID := c.Locals("userID").(int)

	user, err := h.userUsecase.GetUserBalance(userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "user not found",
		})
	}

	type BalanceResponse struct {
		Status  string  `json:"status"`
		User    string  `json:"user"`
		Balance float64 `json:"balance"`
	}

	response := BalanceResponse{
		Status:  "success",
		User:    user.Username,
		Balance: user.Balance,
	}

	return c.JSON(response)
}

func (h *UserHandler) Login(c *fiber.Ctx) error {
	var creds domain.LoginRequest

	if err := c.BodyParser(&creds); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	if creds.Username == "" || creds.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "username and password are required",
		})
	}

	authResponse, err := h.userUsecase.Login(creds)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"token":  authResponse.Token,
		"user":   authResponse.User,
	})
}

func (h *UserHandler) GetUserByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid user ID",
		})
	}

	userResponse, err := h.userUsecase.GetUserByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "user not found",
		})
	}

	return c.JSON(fiber.Map{
		"username": userResponse.Username,
	})
}

// setup routes
func (h *UserHandler) SetupRoutes(app *fiber.App, authMiddleware fiber.Handler) {
	// public endpoints
	app.Post("/login", h.Login)
	app.Get("/user/:id", h.GetUserByID)

	// protected endpoint
	app.Get("/balance", authMiddleware, h.GetBalance)
}
