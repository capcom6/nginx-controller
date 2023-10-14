package main

import (
	"github.com/capcom6/nginx-controller/internal/config"
	"github.com/capcom6/nginx-controller/internal/handlers"
	"github.com/capcom6/nginx-controller/internal/infra/http"
	"github.com/capcom6/nginx-controller/internal/infra/logger"
	"github.com/capcom6/nginx-controller/internal/infra/validator"
	"github.com/capcom6/nginx-controller/internal/services/nginx"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		config.Module,
		logger.Module,
		http.Module,
		validator.Module,
		handlers.Module,
		nginx.Module,
		fx.Invoke(func(h *fiber.App) {

		}),
	).Run()
}
