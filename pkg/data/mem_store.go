package data

import (
	"context"
	"fmt"
)

type memStore struct {
	users map[string]*User
}

func (m *memStore) SaveUser(ctx context.Context, u *User) error {
	m.users[u.ID] = u
	return nil
}

func (m *memStore) GetUser(ctx context.Context, id string) (*User, error) {
	u, ok := m.users[id]
	if ok {
		return u, nil
	}

	return nil, fmt.Errorf("no user with ID: %s", id)
}
