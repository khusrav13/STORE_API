package services

import (
	"HumoStore/db"
	"HumoStore/hashes"
	"HumoStore/models"
	"context"
	"fmt"
	"github.com/jackc/pgx/pgxpool"
	"log"
)

type UserSvc struct {
	pool *pgxpool.Pool
}

func NewUserSvc(pool *pgxpool.Pool) *UserSvc {
	return &UserSvc{pool: pool}
}

func (receiver *UserSvc) CheckUser(Login string) (result bool, err error) {
	conn, err := receiver.pool.Acquire(context.Background())
	if err != nil {
		log.Printf("Cannot connect, please check it.", err)
		return
	}
	defer conn.Release()
	var exists string
	err = conn.QueryRow(context.Background(), db.GetUserByLogin, Login).Scan(&exists)
	if err != nil {
		return false, err
	}
	if exists != "" {
		return true, nil
	}
	return
}

func (receiver *UserSvc) SignUp(User models.User) (err error) {
	conn, err := receiver.pool.Acquire(context.Background())
	if err != nil {
		log.Printf("Cannot connect %e", err)
		return
	}

	defer conn.Release()
	password := hashes.HashIt(User.Password)

	_, err = conn.Exec(context.Background(), db.SignUp, User.Name, User.Surname, password, User.Login, User.Email, User.Phone, User.Status)
	if err != nil {
		fmt.Printf("Cannot Sign Up, err %e", err)
		return
	}
	return nil
}

func GetUser(User models.UserRole) (responseUSER models.ResponseUser) {

	responseUSER.Name = User.Name
	responseUSER.Surname = User.Surname
	responseUSER.Login = User.Login
	responseUSER.Email = User.Email
	responseUSER.Phone = User.Phone
	responseUSER.Role = User.Role
	responseUSER.Status = User.Status
	responseUSER.CreatedAt = User.CreatedAt

	return responseUSER

}

func (receiver *UserSvc) GetUserByLogin(userLogin string) (userId int, responseUser models.ResponseUser, err error) {
	conn, err := receiver.pool.Acquire(context.Background())

	if err != nil {
		log.Printf("Cannot get connection %e", err)
		return
	}
	defer conn.Release()
	var user models.UserRole
	err = conn.QueryRow(context.Background(), db.GetLogin, userLogin).Scan(
		&user.ID,
		&user.Name,
		&user.Surname,
		&user.Login,
		&user.Password,
		&user.Phone,
		&user.Role,
		&user.Status,
		&user.CreatedAt,
	)
	if err != nil {
		return
	}
	responseUser = GetUser(user)
	return user.ID, responseUser, nil
}

func (receiver *UserSvc) UserRole(roleID int, Login string) (err error) {
	conn, err := receiver.pool.Acquire(context.Background())
	if err != nil {
		log.Printf("Cannot get user %e", err)
		return
	}
	defer conn.Release()

	_,err = conn.Exec(context.Background(),db.AddUserRole, roleID, Login)
	return nil
}
