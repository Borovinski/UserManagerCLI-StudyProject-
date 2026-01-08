package main

import (
	"fmt"
	"usermanagercli/models"
)

func main() {
	user, err := models.NewUser(1, "", 18, 5.0)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	fmt.Println(*user)
}
