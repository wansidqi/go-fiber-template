package user

import (
	"github.com/gofiber/fiber/v2"
)

// SetupUserRoutes registers user feature routes under the given router group
func SetupUserRoutes(router fiber.Router) {
	userController := UserController(UserService())
	testUserController := UserController(TestUserService())

	userRouter := router.Group("/users")
	userRouter.Post("/", userController.create)
	userRouter.Get("/", userController.getAll)

	userRouterTest := router.Group("/users/test")
	userRouterTest.Get("/", testUserController.getAll)
}
