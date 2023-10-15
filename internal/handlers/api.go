package handlers

import (
	"errors"
	"os"

	"github.com/capcom6/nginx-controller/internal/services/nginx"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	validator *validator.Validate
	nginx     *nginx.Nginx
}

func (h *Handler) Register(app fiber.Router) {
	v1 := app.Group("/v1")

	v1.Get("/", h.Get)
	v1.Put("/hosts/:hostname", h.Put)
	v1.Delete("/hosts/:hostname", h.Delete)
}

func (h *Handler) Get(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func (h *Handler) Put(c *fiber.Ctx) error {
	hostname := c.Params("hostname")
	if err := h.validator.Var(hostname, "hostname"); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	req := PutHostname{}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := h.validator.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	context := putHostnameToContext(c.Params("hostname"), req)

	if err := h.nginx.Apply(context); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusOK)
}

func (h *Handler) Delete(c *fiber.Ctx) error {
	hostname := c.Params("hostname")
	if err := h.validator.Var(hostname, "hostname"); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := h.nginx.Remove(hostname); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return c.SendStatus(fiber.StatusNotFound)
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func New(v *validator.Validate, n *nginx.Nginx) *Handler {
	return &Handler{
		validator: v,
		nginx:     n,
	}
}
