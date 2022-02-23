package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	"github.com/vapor05/financeview/pkg/store"
)

type Resolver struct {
	Db *store.Database
}
