package storage

import "usermanagercli/models"

type Manager struct {
	users  map[int]*models.User
	nextID int
}

func NewManager() *Manager {
	return &Manager{
		users:  make(map[int]*models.User),
		nextID: 1,
	}
}
