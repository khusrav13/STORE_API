package app

import (
	"HumoStore/api/services"
	"HumoStore/db"
	"github.com/jackc/pgx/pgxpool"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type MainServer struct {
	pool       *pgxpool.Pool
	router     *httprouter.Router
	initialize *db.DbPostgres
	UserSvc    *services.UserSvc
}

func NewMainServer(pool *pgxpool.Pool, router *httprouter.Router, initialize *db.DbPostgres, userService *services.UserSvc) *MainServer {
	return &MainServer{pool: pool, router: router, initialize: initialize, UserSvc: userService}
}

func (server *MainServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	writer.Header().Set("Access-Control-Allow-Origin", "GET, POST, PUT, OTPIONS")

	server.router.ServeHTTP(writer, request)
}

func (server *MainServer) Start() {
	err := server.initialize.DbInit()
	if err != nil {
		panic("server don't created")
	}

	server.InitRoutes()
}
