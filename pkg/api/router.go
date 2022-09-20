package api

func (server *Server) setupRouter() {

	v1 := server.router.Group("/v1")
	v1.Use(GinMiddleware("http://localhost:3000"))
	{
		user := v1.Group("/user")
		{
			user.POST("/", server.createUser)
			user.POST("login", server.loginUser)
		}

		auth := v1.Group("/")
		auth.Use(authMiddleware(server.tokenMaker))
		{
			auth.DELETE("/user/me", server.deleteUser)
			auth.PATCH("/user/me", server.updateUser)
			auth.GET("/ws", server.wss.ServeWss)
			auth.POST("/ws", server.wss.ServeWss)
		}
	}
}
