package resolvers

import (
	"github.com/emoss08/trenova/internal/core/ports/repositories"
	"github.com/emoss08/trenova/internal/pkg/logger"
	"go.uber.org/fx"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	orgRepo repositories.OrganizationRepository
	l       *logger.Logger
}

type ResolverParams struct {
	fx.In

	// Repositories
	OrgRepo repositories.OrganizationRepository

	// Logger
	Logger *logger.Logger
}

func NewResolver(p ResolverParams) *Resolver {
	return &Resolver{
		orgRepo: p.OrgRepo,
		l:       p.Logger,
	}
}
