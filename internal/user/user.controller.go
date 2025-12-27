package user

import (
	"fmt"
	"go-fiber-app/pkg/response"

	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	service UserServiceInterface
}

func UserController(service UserServiceInterface) *Controller {
	return &Controller{service: service}
}

func (c *Controller) create(ctx *fiber.Ctx) error {
	var user User
	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse("Invalid request body", err.Error()))
	}

	if err := c.service.Create(&user); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse("Failed to create user", err.Error()))
	}

	return ctx.Status(fiber.StatusCreated).JSON(response.SuccessResponse("User created successfully", user))
}

func (c *Controller) getAll(ctx *fiber.Ctx) error {
	users, err := c.service.GetAllUsers()
	// fmt.Printf("users: %+v\n", users)
	fmt.Printf("err: %v\n", err)
	
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse("Failed to fetch users", err.Error()))
	}
	return ctx.Status(fiber.StatusOK).JSON(response.SuccessResponse("Users fetched", users))
}
