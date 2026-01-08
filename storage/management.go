package storage

import (
	"fmt"
	"usermanagercli/models"
)

func (m *Manager) AddUser(name string, age int, rating float64) error {
	id := m.nextID
	//Валидация уникальности ID
	if m.IsIDTaken(id) {
		return fmt.Errorf("ID %d уже занят", id)
	}

	//Валидация создания пользователя
	user, err := models.NewUser(id, name, age, rating)
	if wrapErr := wrapError("Ошибка создания пользователя", err); wrapErr != nil {
		return wrapErr
	}
	//Добавления пользователя
	m.users[id] = user

	m.nextID++
	fmt.Printf("Log: Пользователь добавлен — ID: %d, Имя: %s, Возраст: %d, Рейтинг: %.2f\n", user.GetId(), user.GetName(), user.GetAge(), user.GetRating())
	return nil
}

func (m *Manager) UpdateRating(id int, newRate float64) error {
	user, err := m.getUserByID(id)
	if err != nil {
		return err
	}

	if err := wrapError("Ошибка обновления рейтинга", user.ChangeRating(newRate)); err != nil {
		return err
	}

	fmt.Printf(
		"Log: Пользователь обновлён (Рейтинг) — ID: %d, Имя: %s, Возраст: %d, Рейтинг: %.2f\n",
		user.GetId(), user.GetName(), user.GetAge(), user.GetRating(),
	)

	return nil
}

func (m *Manager) ListUsers() []string {
	var result []string
	for _, user := range m.users {
		result = append(result, user.InfoString())
	}
	return result
}

// ID уникален, значит у нас будет один пользователь, то есть мапу указателей (users) можно не перебирать
func (m *Manager) FindUserID(id int) (*models.User, error) {
	return m.getUserByID(id)
}

// Общая функция поиска
func (m *Manager) findUsersBy(filter func(*models.User) bool, errMsg string) ([]*models.User, error) {
	var result []*models.User
	for _, user := range m.users {
		if filter(user) {
			result = append(result, user)
		}
	}

	if len(result) == 0 {
		return nil, fmt.Errorf(errMsg)
	}

	return result, nil
}

// Поиск по имени
func (m *Manager) FindByName(name string) ([]*models.User, error) {
	return m.findUsersBy(
		func(u *models.User) bool { return u.GetName() == name },
		fmt.Sprintf("Пользователи с именем '%s' не найдены", name),
	)
}

// Поиск по возрасту
func (m *Manager) FindByAge(age int) ([]*models.User, error) {
	return m.findUsersBy(
		func(u *models.User) bool { return u.GetAge() == age },
		fmt.Sprintf("Пользователи с возрастом %d не найдены", age),
	)
}

// Поиск по рейтингу
func (m *Manager) FindByRate(rate float64) ([]*models.User, error) {
	return m.findUsersBy(
		func(u *models.User) bool { return u.GetRating() == rate },
		fmt.Sprintf("Пользователи с рейтингом %.2f не найдены", rate),
	)
}

// Удаление пользователя
func (m *Manager) DeleteUser(id int) error {
	// Проверяем, есть ли пользователь
	_, err := m.getUserByID(id)
	if err != nil {
		return err
	}

	// Удаляем пользователя
	delete(m.users, id)

	fmt.Printf("Log: Пользователь с ID %d удалён\n", id)
	return nil
}

// Изменение ID пользователя
func (m *Manager) ChangeID(oldID int, newID int) error {
	// Получаем пользователя по старому ID
	user, err := m.getUserByID(oldID)
	if err != nil {
		return err
	}

	// Проверяем, свободен ли новый ID
	if m.IsIDTaken(newID) {
		return fmt.Errorf("ID %d уже занят", newID)
	}

	// Удаляем старую запись и обновляем ID
	delete(m.users, oldID)
	if err := user.ChangeID(newID); err != nil {
		return fmt.Errorf("Не удалось изменить ID пользователя: %w", err)
	}
	m.users[newID] = user

	// Обновляем nextID
	if newID >= m.nextID {
		m.nextID = newID + 1
	}

	fmt.Printf("Log: Изменён ID пользователя: %s\n", user.InfoString())
	return nil
}

// Изменение возраста пользователя
func (m *Manager) ChangeAge(id int, age int) error {
	// Получаем пользователя по ID
	user, err := m.getUserByID(id)
	if err != nil {
		return err
	}

	// Меняем возраст и пробрасываем ошибку, если она есть
	if err := user.ChangeAge(age); err != nil {
		return fmt.Errorf("Не удалось изменить возраст пользователя: %w", err)
	}

	fmt.Printf("Log: Изменён возраст пользователя: %s\n", user.InfoString())
	return nil
}

// Изменение рейтинга пользователя
func (m *Manager) ChangeRate(id int, rate float64) error {
	// Получаем пользователя по ID
	user, err := m.getUserByID(id)
	if err != nil {
		return err
	}

	// Меняем рейтинг и пробрасываем ошибку, если она есть
	if err := user.ChangeRating(rate); err != nil {
		return fmt.Errorf("Не удалось изменить рейтинг пользователя: %w", err)
	}

	fmt.Printf("Log: Изменён рейтинг пользователя: %s\n", user.InfoString())
	return nil
}

// Изменение имени пользователя
func (m *Manager) ChangeName(id int, newName string) error {
	// Получаем пользователя по ID
	user, err := m.getUserByID(id)
	if err != nil {
		return err
	}

	// Меняем имя и пробрасываем ошибку из модели
	if err := user.ChangeName(newName); err != nil {
		return fmt.Errorf("Не удалось изменить имя пользователя: %w", err)
	}

	fmt.Printf("Log: Изменено имя пользователя: %s\n", user.InfoString())
	return nil
}
