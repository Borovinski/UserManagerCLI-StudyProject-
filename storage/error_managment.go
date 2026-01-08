package storage

import (
	"fmt"
	"usermanagercli/models"
)

// Возвращает ошибку если уже есть существующий ID
func (m *Manager) IsIDTaken(id int) bool {
	_, exists := m.users[id]
	return exists
}

// Возвращает пользователя по ID или ошибку, если не найден
func (m *Manager) getUserByID(id int) (*models.User, error) {
	user, ok := m.users[id]
	if !ok {
		return nil, fmt.Errorf("Пользователь с ID %d не найден", id)
	}
	return user, nil
}

func wrapError(msg string, err error) error {
	if err != nil {
		return fmt.Errorf("%s: %w", msg, err)
	}
	return nil
}
