package main

import (
	"github.com/capcom6/nginx-controller/internal/config"
	"github.com/capcom6/nginx-controller/internal/handlers"
	"github.com/capcom6/nginx-controller/internal/infra/http"
	"github.com/capcom6/nginx-controller/internal/infra/logger"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func main() {
	fx.New(
		config.Module,
		logger.Module,
		http.Module,
		handlers.Module,
		fx.Invoke(func(cfg config.Config, log *zap.Logger, h *fiber.App) {
			log.Info("Config", zap.Any("config", cfg))
		}),
	).Run()
}
