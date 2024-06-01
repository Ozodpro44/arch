package main

import (
	"app/api"
	"app/config"
	"app/storage/postgres"
	"context"
	"fmt"
	"log"
)

func main() {
	cfg := config.NewConfig()

	conn, err := postgres.Conn(cfg)

	if err != nil {
		log.Println("Error on Conn", conn)
		return
	}
	defer conn.Close(context.Background())

	fmt.Println(conn)
	UserRepo := postgres.NewUserRepo(conn)

	api.Api(UserRepo)
}
