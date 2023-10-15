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

	v1.Put("/hosts/:hostname", h.Put)
	v1.Delete("/hosts/:hostname", h.Delete)
}

//	@Summary		Replace host upstreams
//	@Description	Replaces current configuration on hostname's upstreams
//	@Tags			Proxy
//	@Accept			json
//	@Produce		json
//	@Param			hostname	path		string			true	"Hostname"
//	@Param			request		body		PutHostname		true	"Upstreams configuration"
//	@Success		204			{object}	nil				"Success"
//	@Failure		400			{object}	ErrorResponse	"Validation error"
//	@Failure		500			{object}	ErrorResponse	"Internal server error"
//	@Router			/v1/hosts/:hostname [put]
//
// Replace host upstreams
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

//	@Summary		Delete host
//	@Description	Removes host from configuration
//	@Tags			Proxy
//	@Produce		json
//	@Param			hostname	path		string			true	"Hostname"
//	@Success		204			{object}	nil				"Success"
//	@Failure		400			{object}	ErrorResponse	"Validation error"
//	@Failure		404			{object}	ErrorResponse	"Host not found"
//	@Failure		500			{object}	ErrorResponse	"Internal server error"
//	@Router			/v1/hosts/:hostname [delete]
//
// Delete host
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
