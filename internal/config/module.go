package config

import (
	"github.com/capcom6/nginx-controller/internal/services/nginx"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"config",
	fx.Provide(GetConfig),
	fx.Provide(func(cfg Config) nginx.Config {
		return nginx.Config{
			ConfigPath:     cfg.Nginx.Location,
			ConfigTemplate: cfg.Nginx.Template,
		}
	}),
)
