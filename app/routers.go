package app

func (server *MainServer) InitRoutes() {
	server.router.POST("/api/checkUser", server.CheckUserHandler)
	server.router.POST("/api/signUp", server.SignUp)
}