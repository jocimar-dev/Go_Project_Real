package main

import (
	"awesomeProject/application/repositories"
	"awesomeProject/domain"
	"awesomeProject/framework/utils"
	"fmt"
	"log"
)

func main() {

	db := utils.ConnectDB()

	user := domain.User{
		Name:     "Jocimar",
		Email:    "jocimarneres@gmail.com",
		Password: "321",
	}

	userRepo := repositories.UserRepositoryDb{Db: db}
	result, err := userRepo.Insert(&user)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}
