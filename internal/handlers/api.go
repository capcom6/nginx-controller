package handlers

import "github.com/gofiber/fiber/v2"

type Handler struct {
}

func (h *Handler) Register(app fiber.Router) {
	app.Get("/", h.Get)
}

func (h *Handler) Get(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func New() *Handler {
	return &Handler{}
}
