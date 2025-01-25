package api

import (
	"go.uber.org/fx"
)

var Module = fx.Module("api", ResolversModule, HandlersModule, ServerModule, RouterModule)
