package api

import (
	"github.com/emoss08/trenova/internal/api/graph/resolvers"
	"go.uber.org/fx"
)

var ResolversModule = fx.Module("api.graphql.resolvers", fx.Provide(resolvers.NewResolver))
