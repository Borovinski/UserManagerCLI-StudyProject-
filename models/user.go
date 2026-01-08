package models

import "fmt"

//Инкапсуляция переменных структуры через camelcase
type User struct {
	id     int
	name   string
	age    int
	rating float64
}

func NewUser(idUser int, nameUser string, ageUser int, ratingUser float64) (*User, error) {
	validID := idUser > 0 && idUser <= 100
	validname := nameUser == ""
	validAge := ageUser > 0 && ageUser <= 150
	validRate := ratingUser >= 0.0 && ratingUser <= 10.0
	if !validID {
		return nil, fmt.Errorf("ID должен быть > 1 и <101!")
	}

	if validname {
		return nil, fmt.Errorf("Пустое имя!")
	}

	if !validAge {
		return nil, fmt.Errorf("Такого возраста нет!")
	}

	if !validRate {
		return nil, fmt.Errorf("Недопустиный рейтинг!")
	}

	return &User{
		id:     idUser,
		name:   nameUser,
		age:    ageUser,
		rating: ratingUser,
	}, nil
}

//Валидация при изменении имени
func (user *User) ChangeName(newName string) error {

	if user == nil {
		return fmt.Errorf("Пользователь пустой при изменении имени!")
	}
	if newName == "" {
		return fmt.Errorf("Пустое имя!")
	}
	user.name = newName
	return nil
}

//Валидация при изменении возраста
func (user *User) ChangeAge(newAge int) error {
	if user == nil {
		return fmt.Errorf("Пользователь пустой при изменении возраста!")
	}
	validAge := newAge > 0 && newAge <= 150
	if !validAge {
		return fmt.Errorf("Нельзя изменить возраст, возраст должен быть от 0 до 150 включительно")
	}
	user.age = newAge
	return nil
}

//Валидация при изменении рейтинга
func (user *User) ChangeRating(newRating float64) error {
	if user == nil {
		return fmt.Errorf("Пользователь пустой при изменении рейтинга!")
	}
	validRate := newRating >= 0.0 && newRating <= 10.0
	if !validRate {
		return fmt.Errorf("Нельзя изменить рейтинг, рейтинг должен быть от 0.0 до 10.0!")
	}

	user.rating = newRating
	return nil
}

//Геттер получения имени пользователя
func (user *User) GetName() string {
	if user == nil {
		return "Пользователь пустой при получении имени!"
	}
	return user.name
}

//Геттер получения ID пользователя
func (user *User) GetId() int {
	if user == nil {
		return 0
	}
	return user.id
}

//Геттер получения рейтинга пользователя
func (user *User) GetRating() float64 {
	if user == nil {
		return 0.0
	}
	return user.rating
}
