package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/pgxpool"
)

type DbPostgres struct {
	pool *pgxpool.Pool
}

func NewDbInit(pool *pgxpool.Pool) *DbPostgres {
	return &DbPostgres{pool: pool}
}

func (receiver *DbPostgres) DbInit() (err error) {
	DDls := []string{
		UserTable,
		CardTable,
		CategoryTable,
		ProductTable,
		OrderTable,
		CreateRolesTable,
		CreateUserRoleTable,
	}
	for _, ddl := range DDls {
		_, err := receiver.pool.Exec(context.Background(), ddl)
		if err != nil {
			fmt.Println("error cannot init DB:", err)
			return err
		}
	}
	return nil
}
