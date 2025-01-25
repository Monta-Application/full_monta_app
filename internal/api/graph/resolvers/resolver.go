package resolvers

import "github.com/emoss08/trenova/internal/core/ports/db"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DB db.Connection
}
