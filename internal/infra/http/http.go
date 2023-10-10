package http

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"go.uber.org/fx"
)

func New(apiHandlers []ApiHanlder, lc fx.Lifecycle) (*fiber.App, error) {
	app := fiber.New()

	app.Use(logger.New())
	app.Use(recover.New())

	api := app.Group("/api")
	for _, handler := range apiHandlers {
		handler.Register(api)
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go app.Listen(":4000")

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return app.ShutdownWithContext(ctx)
		},
	})

	return app, nil
}
