package postgres

import (
	"app/models"
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

type UserRepo struct {
	db *pgx.Conn
}

func NewUserRepo(db *pgx.Conn) UserRepo {

	return UserRepo{db: db}
}
func (u *UserRepo) GetUserByUsername(ctx context.Context, username string) (models.Users, error) {
	var user models.Users
	query := "SELECT * FROM users WHERE username=$1"
	err := u.db.QueryRow(ctx, query, username).Scan(&user.User_id, &user.User_name, &user.Date, &user.Gmail, &user.Create_at)
	if err != nil {
		log.Println("err on GetUsersByID ", err)
		return user, err
	}

	return user, nil
}

func (u *UserRepo) CreateUser(ctx context.Context, user models.Users) error {

	query := "INSERT INTO users(username,gmail,date_of_birth)VALUES($1,$2,$3)"
	_, err := u.db.Exec(ctx, query, user.User_name, user.Gmail, user.Date)
	if err != nil {
		log.Println("error on CreateUser ", err)
		return err
	}
	return nil
}

func (u *UserRepo) UpdateUser(ctx context.Context, name1 string, name2 string) error {

	query := "UPDATE users SET username = $1 WHERE username = $2;"
	_, err := u.db.Exec(ctx, query, name2, name1)
	if err != nil {
		log.Println("error on Updating User ", err)
		return err
	}
	fmt.Println("Updated")
	return nil
}

func (u *UserRepo) DeleteUser(ctx context.Context, username string) error {

	query := "DELETE FROM users WHERE username=$1"

	_, err := u.db.Exec(ctx, query, username)

	if err != nil {
		log.Println("Error on Deleting user", err)
		return err
	}

	fmt.Println("Deleted")
	return nil
}

func (u *UserRepo) GetAllUSers(ctx context.Context) ([]models.Users, error) {
	var users []models.Users
	var user models.Users
	query := "SELECT * FROM users;"
	row, err := u.db.Query(ctx, query)
	if err != nil {
		log.Println("Error on get USers", err)
	}
	for row.Next() {

		err := row.Scan(&user.User_id, &user.User_name, &user.Date, &user.Gmail, &user.Create_at)
		if err != nil {
			log.Println("Error on Scan all users", err)
			return users, err
		}
		users = append(users, user)
	}

	return users, nil
}
