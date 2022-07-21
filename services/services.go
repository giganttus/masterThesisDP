package services

import (
	"errors"
	"github.com/go-pg/pg"
	"trackingApp/database"
)

type Services struct {
	ItemsRepo database.ItemsRepo
	UsersRepo database.UsersRepo
	DB        *pg.DB
}

func NewService(usersRepo database.UsersRepo, itemsRepo database.ItemsRepo) *Services {
	return &Services{UsersRepo: usersRepo, ItemsRepo: itemsRepo}
}

var Role string

var (
	ErrBadCredentials        = errors.New("email/password combination don't work")
	ErrDatabaseValueExists   = errors.New("value exists in database")
	ErrDatabaseValueNotExist = errors.New("value doesn't exist in database")
	ErrRelationExists        = errors.New("relation exists")
	ErrUnauthenticated       = errors.New("unauthenticated")
	ErrForbidden             = errors.New("unauthorized")
	ErrPasswordHash          = errors.New("hash failed")
	ErrLengthLimit           = errors.New("the allowable length is not met")
)
