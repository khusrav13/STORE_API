package main

import (
	"HumoStore/api/services"
	"HumoStore/app"
	"HumoStore/db"
	"context"
	"fmt"
	"github.com/jackc/pgx/pgxpool"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	router := httprouter.New()
	pool, err := pgxpool.Connect(context.Background(), `postgres://postgres:3477@localhost:5432/Humo?sslmode=disable`)
	if err != nil {
		log.Printf("Error - %e", err)
		log.Fatal("Cannot connect to Database")
	} else {
		fmt.Println("Connection to DB is success")
	}

	UserService := services.NewUserSvc(pool)
	dbInit := db.NewDbInit(pool)

	server := app.NewMainServer(pool, router, dbInit, UserService)

	server.Start()
	panic(http.ListenAndServe(":3000", server))

}
