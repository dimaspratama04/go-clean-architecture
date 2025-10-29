package controller

import (
	"context"
	"time"

	"golang-redis/internal/delivery/http/request"
	"golang-redis/internal/delivery/http/response"
	"golang-redis/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

type UsersController struct {
	UsersUC *usecase.UsersUseCase
}

func NewUsersController(usersUC *usecase.UsersUseCase) *UsersController {
	return &UsersController{UsersUC: usersUC}
}

func (uc *UsersController) Register(c *fiber.Ctx) error {
	var payload request.RegisterRequest
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "invalid credentials.",
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	users, err := uc.UsersUC.Register(ctx, payload.Username, payload.Email)
	if err != nil {
		switch err.Error() {
		case "EMAIL_ALREADY_REGISTERED":
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{
				"status":  "error",
				"message": "email already registered",
			})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":  "error",
				"message": "ups, something went wrong",
			})
		}
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "ok",
		"message": "register success",
		"data": fiber.Map{
			"user": fiber.Map{
				"id":    users.ID,
				"name":  users.Username,
				"email": users.Email,
			},
		},
	})
}

func (uc *UsersController) Login(c *fiber.Ctx) error {
	var payload request.LoginRequest
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "invalid request payload.",
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	token, user, err := uc.UsersUC.Login(ctx, payload.Email)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.Response{
		Status:  "ok",
		Message: "succesfully logged in.",
		Data: fiber.Map{
			"access_token": token,
			"users": fiber.Map{
				"id":    user.ID,
				"name":  user.Username,
				"email": user.Email,
			},
		},
	})
}
