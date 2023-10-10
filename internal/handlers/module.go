package handlers

import (
	"github.com/capcom6/nginx-controller/internal/infra/http"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"handlers",
	fx.Provide(
		http.AsApiHandler(New),
	),
)
