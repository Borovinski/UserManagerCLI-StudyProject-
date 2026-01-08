package menu

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"usermanagercli/storage"
)

var manager = storage.NewManager()

func Menu() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("------------------------")
		fmt.Println("----------Меню----------")
		fmt.Println("------------------------")
		fmt.Println("1. Добавить пользователя")
		fmt.Println("2. Просмотреть всех пользователей")
		fmt.Println("3. Найти пользователя по ID")
		fmt.Println("4. Найти пользователя по имени")
		fmt.Println("5. Найти пользователя по возрасту")
		fmt.Println("6. Найти пользователя по рейтингу")
		fmt.Println("7. Изменить ID пользователя")
		fmt.Println("8. Изменить имя пользователя")
		fmt.Println("9. Изменить возраст пользователя")
		fmt.Println("10. Изменить рейтинг пользователя")
		fmt.Println("11. Удалить пользователя")
		fmt.Println("0. Выйти")
		fmt.Print("Выберите пункт меню: ")

		choiceStr, _ := reader.ReadString('\n')
		choiceStr = strings.TrimSpace(choiceStr)
		choice, err := strconv.Atoi(choiceStr)
		if err != nil {
			fmt.Println("Ошибка: введите число")
			continue
		}

		switch choice {
		case 1:
			name, _ := ReadName()
			age, _ := ReadAge()
			rating, _ := ReadRating()
			if err := manager.AddUser(name, age, rating); err != nil {
				fmt.Println("Ошибка:", err)
			} else {
				fmt.Println("✅ Пользователь добавлен")
			}

		case 2:
			users := manager.ListUsers()
			for _, u := range users {
				fmt.Println(u)
			}

		case 3:
			id, _ := ReadID()
			user, err := manager.FindUserID(id)
			if err != nil {
				fmt.Println("Ошибка:", err)
				break
			}
			fmt.Println(user.InfoString())

		case 4:
			name, _ := ReadName()
			users, err := manager.FindByName(name)
			if err != nil {
				fmt.Println("Ошибка:", err)
				break
			}
			for _, u := range users {
				fmt.Println(u.InfoString())
			}

		case 5:
			age, _ := ReadAge()
			users, err := manager.FindByAge(age)
			if err != nil {
				fmt.Println("Ошибка:", err)
				break
			}
			for _, u := range users {
				fmt.Println(u.InfoString())
			}

		case 6:
			rate, _ := ReadRating()
			users, err := manager.FindByRate(rate)
			if err != nil {
				fmt.Println("Ошибка:", err)
				break
			}
			for _, u := range users {
				fmt.Println(u.InfoString())
			}

		case 7:
			oldID, _ := ReadID()
			newID, _ := ReadID()
			if err := manager.ChangeID(oldID, newID); err != nil {
				fmt.Println("Ошибка:", err)
			} else {
				fmt.Println("✅ ID пользователя изменён")
			}

		case 8:
			id, _ := ReadID()
			name, _ := ReadName()
			if err := manager.ChangeName(id, name); err != nil {
				fmt.Println("Ошибка:", err)
			} else {
				fmt.Println("✅ Имя пользователя изменено")
			}

		case 9:
			id, _ := ReadID()
			age, _ := ReadAge()
			if err := manager.ChangeAge(id, age); err != nil {
				fmt.Println("Ошибка:", err)
			} else {
				fmt.Println("✅ Возраст пользователя изменён")
			}

		case 10:
			id, _ := ReadID()
			rate, _ := ReadRating()
			if err := manager.ChangeRate(id, rate); err != nil {
				fmt.Println("Ошибка:", err)
			} else {
				fmt.Println("✅ Рейтинг пользователя изменён")
			}

		case 11:
			id, _ := ReadID()
			if err := manager.DeleteUser(id); err != nil {
				fmt.Println("Ошибка:", err)
			} else {
				fmt.Println("✅ Пользователь удалён")
			}

		case 0:
			fmt.Println("👋 Выход из программы")
			return

		default:
			fmt.Println("❌ Неверный пункт меню")
		}

		fmt.Println()
	}
}
