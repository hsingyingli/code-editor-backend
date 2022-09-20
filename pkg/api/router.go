package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) setupRouter() {

	server.router.Use(GinMiddleware("http://localhost:3000"))

	v1 := server.router.Group("/v1")
	{
		v1.POST("/user", server.createUser)
		v1.POST("/user/login", server.loginUser)
		v1.POST("/tokens/renew_access", server.renewAccessToken)
		v1.GET("/ws", server.wss.ServeWss)
		v1.POST("/ws", server.wss.ServeWss)

		v1.POST("/test", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, "hello world")
		})

		auth := v1.Group("/")
		auth.Use(authMiddleware(server.tokenMaker))
		{
			//auth.DELETE("/user/me", server.deleteUser)
			//auth.PATCH("/user/me", server.updateUser)
		}
	}
}
