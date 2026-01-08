package models

import "fmt"

//Вся логика валидации данных находится на managment
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

//Валидация при изменении ID
func (user *User) ChangeID(newId int) error {
	user.id = newId
	return nil
}

//Валидация при изменении имени
func (user *User) ChangeName(newName string) error {

	if newName == "" {
		return fmt.Errorf("Пустое имя!")
	}
	user.name = newName
	return nil
}

//Валидация при изменении возраста
func (user *User) ChangeAge(newAge int) error {
	validAge := newAge > 0 && newAge <= 150
	if !validAge {
		return fmt.Errorf("Нельзя изменить возраст, возраст должен быть от 1 до 150 включительно")
	}
	user.age = newAge
	return nil
}

//Валидация при изменении рейтинга
func (user *User) ChangeRating(newRating float64) error {
	validRate := newRating >= 0.0 && newRating <= 10.0
	if !validRate {
		return fmt.Errorf("Нельзя изменить рейтинг, рейтинг должен быть от 0.0 до 10.0!")
	}

	user.rating = newRating
	return nil
}

//Геттер получения имени пользователя
func (user *User) GetName() string {
	return user.name
}

//Геттер получения ID пользователя
func (user *User) GetId() int {
	return user.id
}

//геттер получения возраста пользователя
func (user *User) GetAge() int {
	return user.age
}

//Геттер получения рейтинга пользователя
func (user *User) GetRating() float64 {
	return user.rating
}

//Вывод всех пользователей
func (user *User) InfoString() string {
	return fmt.Sprintf(
		"ID пользователя: %d\nИмя пользователя: %s\nВозраст пользователя: %d\nРейтинг пользователя: %.2f",
		user.GetId(), user.GetName(), user.GetAge(), user.GetRating(),
	)
}
