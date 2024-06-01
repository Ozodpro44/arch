package api

import (
	"app/api/handlers"
	"app/storage/postgres"
	"fmt"
)

func Api(UserRepo postgres.UserRepo){

	h:=handlers.NewHandlers(UserRepo)
	x:=-1 

	for x!=0{
		fmt.Println(`
		1.CreateUser
		2.UpdateUser
		3.Delete User
		4.Get All users
		5.Get user By Username
		0.Exit
		`)
		fmt.Scanln(&x)

		switch x{
		case 1:
			h.CreateUser(UserRepo)
		case 2:
			h.UpdateUser(UserRepo)
		case 3:
			h.DeleteUser(UserRepo)
		case 4:
			h.GetAllUSers(UserRepo)
		case 5:
			h.GetUserByUsername(UserRepo)
		case 0:
			return
		}
	}
}