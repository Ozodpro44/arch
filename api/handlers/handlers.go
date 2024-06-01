package handlers

import (
	"app/models"
	"app/storage/postgres"
	"context"
	"fmt"
	"log"
)

type Handlers struct {
	UserRepo postgres.UserRepo
}

func NewHandlers(UserRepo postgres.UserRepo) Handlers {
	return Handlers{UserRepo: UserRepo}
}

func (h *Handlers) GetUserByUsername(UserRepo postgres.UserRepo) {
	var username string
	fmt.Print("Enter User name which you want to get:")
	fmt.Scanln(&username)

	user, err := UserRepo.GetUserByUsername(context.Background(), username)

	if err != nil {
		log.Println("Error on getting user:", err)
		return
	}

	fmt.Println(user.Create_at.Date())

}

func (h *Handlers) CreateUser(UserRepo postgres.UserRepo) {
	var user models.Users
	var year int
	var month int
	var day int
	fmt.Print("Enter user name:")
	fmt.Scanln(&user.User_name)
	fmt.Print("Enter Gmail:")
	fmt.Scanln(&user.Gmail)
	fmt.Println("Enter DATE OF BIRTH")
	fmt.Print("YYYY:")
	fmt.Scan(&year)
	fmt.Print("MM:")
	fmt.Scanln(&month)
	fmt.Print("DD:")
	fmt.Scanln(&day)

	user.Date = user.Date.AddDate(year-1, month-1, day-1)

	err := UserRepo.CreateUser(context.Background(), user)

	if err != nil {
		log.Println("Error on Create user ", err)
		return
	}

	fmt.Println("User created")

}

func (h *Handlers) UpdateUser(UserRepo postgres.UserRepo) {

	var username string
	var username2 string

	fmt.Print("enter Updating username:")
	fmt.Scanln(&username)
	fmt.Print("enter New username:")
	fmt.Scanln(&username2)

	err := UserRepo.UpdateUser(context.Background(), username, username2)
	if err != nil {
		log.Println("error on updating", err)
		return
	}

}

func (h *Handlers) DeleteUser(UserRepo postgres.UserRepo) {
	var username string
	fmt.Print("Enter Deleting username:")
	fmt.Scanln(&username)

	err := UserRepo.DeleteUser(context.Background(), username)

	if err != nil {
		log.Println("error on deleting", err)
		return
	}
}

func (h *Handlers) GetAllUSers(UserRepo postgres.UserRepo) {

	// var limit int
	// fmt.Print("Enter rows limit:")
	// fmt.Scanln(&limit)
	users, err := UserRepo.GetAllUSers(context.Background())

	if err != nil {
		log.Println("Error on get all users:", err)
		return
	}

	fmt.Println("rows | user_id | username | date_of_birth |            gmail               |         created_at          ")
	fmt.Println("-----+---------+----------+---------------+--------------------------------+----------------------------")
	for i, v := range users {
		fmt.Printf("%4d | %7d | %8s | %13s | %-30s | %26s\n", i+1, v.User_id, v.User_name, v.Date.Format("2006-02-01"), v.Gmail, v.Create_at)

	}
}
