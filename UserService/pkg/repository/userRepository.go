package repository

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"gitlab.com/AlimKenzza/authorization/interfaces"
	"gitlab.com/AlimKenzza/authorization/pkg/dataUser"
	"log"
)

type UserRepository struct {
	pool pgxpool.Pool
}

func NewUserRepository(conn *pgxpool.Pool) interfaces.IUsersRepository {
	return &UserRepository{*conn}
}
func (r *UserRepository) GetUserById(id int) *dataUser.User {
	stmt := "SELECT * FROM users WHERE user_id = $1"
	o := &dataUser.User{}
	err := r.pool.QueryRow(context.Background(), stmt, id).Scan(&o.Id, &o.Username, &o.Password, &o.Email)
	if err != nil {
		log.Println("Didn't find user with id ", id)
		return nil
	}
	return o
}

func (r UserRepository) GetAllUsers() []*dataUser.User {
	stmt := "SELECT * FROM users"
	rows, err := r.pool.Query(context.Background(), stmt)
	if err != nil {
		log.Fatal("Failed to SELECT: %v", err)
		return nil
	}
	defer rows.Close()
	users := []*dataUser.User{}
	for rows.Next() {
		o := &dataUser.User{}
		err = rows.Scan(&o.Id, &o.Username, &o.Password, &o.Email)
		if err != nil {
			log.Fatalf("Failed to scan: %v", err)
			return nil
		}
		users = append(users, o)
	}
	if err = rows.Err(); err != nil {
		return nil
	}
	return users
}

func (r *UserRepository) DeleteUser(user dataUser.User) bool {
	_, err := r.pool.Exec(context.Background(),
		"DELETE FROM users WHERE user_id = $1", user.Id)
	if err != nil {
		return false
	}
	return true
}
func (r UserRepository) UpdateUser(user dataUser.User) bool {
	_, err := r.pool.Exec(context.Background(),
		"UPDATE users SET username = $1, user_password = $2, email = $3 WHERE user_id = $4",
		user.Username, user.Password, user.Email, user.Id)
	if err != nil {
		return false
	}
	return true
}
