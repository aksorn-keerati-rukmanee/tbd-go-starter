package http

import (
	"github.com/aksorn-keerati-rukmanee/tbd-go-starter/internal/core/entities"
	"github.com/aksorn-keerati-rukmanee/tbd-go-starter/internal/core/usecases"
	"github.com/gofiber/fiber/v2"
)

type HttpLanguageHandler struct {
	languageUseCase usecases.LanguageUseCase
}

func NewHttpLanguageHandler(useCase usecases.LanguageUseCase) *HttpLanguageHandler {
	return &HttpLanguageHandler{languageUseCase: useCase}
}

func (h *HttpLanguageHandler) CreateLanguage(c *fiber.Ctx) error {
	var language entities.Language
	if err := c.BodyParser(&language); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	if err := h.languageUseCase.CreateLanguage(language); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(language)
}

func (h *HttpLanguageHandler) GetAllLanguage(c *fiber.Ctx) error {

	res, err := h.languageUseCase.GetAllLanguage()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
