package data

import (
	"context"
)

type Store interface {
	SaveUser(ctx context.Context, u *User) error
	GetUser(ctx context.Context, id string) (*User, error)
}

func NewStore() Store {
	return &memStore{map[string]*User{}}
}
