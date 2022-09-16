package api

func (server *Server) setupRouter() {

  v1 := server.router.Group("/v1")
  v1.Use(GinMiddleware("http://localhost:3000"))
  {
    user := v1.Group("/user")
    {
      user.POST("/", server.createUser)
      user.DELETE("/me", server.deleteUser)
      user.PATCH("/me", server.updateUser)
    }
    

    ws := v1.Group("ws")
    {
      ws.GET("", server.wss.ServeWss)
      ws.POST("", server.wss.ServeWss)
    }
  }
}
