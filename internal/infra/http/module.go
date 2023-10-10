package http

import "go.uber.org/fx"

var Module = fx.Module(
	"http",
	fx.Provide(
		fx.Annotate(
			New,
			fx.ParamTags(`group:"api-routes"`),
		),
	),
)
