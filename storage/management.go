package storage

import (
	"fmt"
	"usermanagercli/models"
)

func (m *Manager) AddUser(name string, age int, rating float64) error {
	id := m.nextID
	user, err := models.NewUser(id, name, age, rating)
	if err != nil {
		return fmt.Errorf("Ошибка создания пользователя менеджером!")
	}
	m.users[id] = user

	m.nextID++
	fmt.Printf("Log: Пользователь %s добавлен с ID %d\n", name, id)
	return nil
}

func (m *Manager) UpdateRating(id int, newRate float64) error {
	user, ok := m.users[id]
	if !ok {
		return fmt.Errorf("Пользователь с ID %d не найден", id)
	}

	err := user.ChangeRating(newRate)
	if err != nil {
		return err
	}
	fmt.Printf("Log: Рейтинг пользователя %s (ID %d) обновлён на %.2f\n",
		user.GetName(), user.GetId(), user.GetRating())
	return nil
}

func (m *Manager) ListUsers() []string {
	var result []string
	for _, user := range m.users {
		id := user.GetId()
		name := user.GetName()
		rating := user.GetRating()
		str := fmt.Sprintf("ID пользователя: %d\nИмя пользователя: %s\nРейтинг пользователя: %.2f", id, name, rating)
		result = append(result, str)
	}
	return result
}

// Здесь реализовавл с копированием, так как мне нужно просто поиск по имени
func (m *Manager) FindByName(name string) []models.User {
	var result []models.User
	for _, user := range m.users {
		if name == user.GetName() {
			result = append(result, *user)
		}
	}
	return result
}

func (m *Manager) DeleteUser(id int) error {
	if _, ok := m.users[id]; !ok {
		return fmt.Errorf("Пользователь с ID %d не найден", id)
	}
	delete(m.users, id)
	fmt.Printf("Log: Пользователь с ID %d удалён\n", id)
	return nil
}
