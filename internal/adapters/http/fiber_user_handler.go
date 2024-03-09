package http

import (
	"github.com/aksorn-keerati-rukmanee/tbd-go-starter/internal/core/entities"
	"github.com/aksorn-keerati-rukmanee/tbd-go-starter/internal/core/usecases"
	"github.com/gofiber/fiber/v2"
)

type HttpUserHandler struct {
	userUseCase usecases.UserUseCase
}

func NewHttpUserHandler(useCase usecases.UserUseCase) *HttpUserHandler {
	return &HttpUserHandler{userUseCase: useCase}
}

func (h *HttpUserHandler) CreateUser(c *fiber.Ctx) error {
	var user entities.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	if err := h.userUseCase.CreateUser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

func (h *HttpUserHandler) GetAllUser(c *fiber.Ctx) error {

	res, err := h.userUseCase.GetAllUser()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
