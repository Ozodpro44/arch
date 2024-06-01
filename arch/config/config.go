package config

import "fmt"

type Config struct {
	DbUser		string
	DbPassword	string
	DbHost 		string
	DbPort		int
	DbName		string
}

func NewConfig()Config {
	
	cfg:=Config {}

	fmt.Print("Enter DBUser:")
	fmt.Scanln(&cfg.DbUser)
	fmt.Print("Enter DBPassword:")
	fmt.Scanln(&cfg.DbPassword)
	fmt.Print("Enter Host:")
	fmt.Scanln(&cfg.DbHost)
	fmt.Print("Enter Port:")
	fmt.Scanln(&cfg.DbPort)
	fmt.Print("Enter DBName:")
	fmt.Scanln(&cfg.DbName)
	return cfg
}