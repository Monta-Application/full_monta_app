package api

import (
	"github.com/emoss08/trenova/internal/api/graph/resolvers"
	"go.uber.org/fx"
)

var Module = fx.Module("api", resolvers.ResolversModule, HandlersModule, ServerModule, RouterModule)
