package nginx

import "go.uber.org/fx"

var Module = fx.Module(
	"nginx",
	fx.Provide(New),
)
