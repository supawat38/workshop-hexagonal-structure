package handlers

import (
	"microserviceMOCK/pkg/validator"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type HTTPHandler struct {
	db        *gorm.DB
	validator validator.Validator
}

func New(db *gorm.DB) *HTTPHandler {
	return &HTTPHandler{
		db:        db,
		validator: validator.New(),
	}
}

//Check API
func (hdl *HTTPHandler) HealthCheck(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "fiber - ready to use",
	})
}
