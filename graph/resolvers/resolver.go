//go:generate go run github.com/99designs/gqlgen generate

package resolvers

import (
	"github.com/go-pg/pg"
	"trackingApp/database"
	"trackingApp/services"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Services  *services.Services
	DB        *pg.DB
	UsersRepo *database.UsersRepo
}
