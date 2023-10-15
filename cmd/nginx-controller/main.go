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

//	@title			Nginx Controller
//	@version		1.0.0
//	@description	API for controlling nginx reverse proxy

//	@contact.name	Aleksandr Soloshenko
//	@contact.email	i@capcom.me

//	@host		localhost:3000
//	@schemes	http
//	@BasePath	/api

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
