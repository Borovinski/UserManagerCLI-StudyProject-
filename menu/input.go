package input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Ввод имени
func ReadName() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите имя: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	if name == "" {
		return "", fmt.Errorf("Имя не может быть пустым!")
	}
	return name, nil
}

// Ввод возраста
func ReadAge() (int, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите возраст: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	age, err := strconv.Atoi(input)
	if err != nil || age <= 0 {
		return 0, fmt.Errorf("Возраст должен быть положительным числом!")
	}
	return age, nil
}

// Ввод рейтинга
func ReadRating() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите рейтинг: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	rating, err := strconv.ParseFloat(input, 64)
	if err != nil || rating < 0 || rating > 10 {
		return 0, fmt.Errorf("Рейтинг должен быть числом от 0 до 10!")
	}
	return rating, nil
}

// Ввод ID пользователя
func ReadID() (int, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите ID пользователя: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	id, err := strconv.Atoi(input)
	if err != nil || id < 1 {
		return 0, fmt.Errorf("ID должен быть целым числом, начиная с 1!")
	}

	return id, nil
}
