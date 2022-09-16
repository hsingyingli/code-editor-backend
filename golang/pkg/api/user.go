package api

import (
  "net/http"

  "github.com/lib/pq"
  "github.com/gin-gonic/gin"
  "github.com/hsingyingli/code-editor/pkg/db/sqlc"
)


type createUserRequest struct {
  Username string `json:"username" binding:"required"`
  Email    string `json:"email" binding:"required"`
  Password string `json:"password" binding:"required"`
}

func (server *Server) createUser(ctx *gin.Context) {
  var req createUserRequest 
  if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
  arg := db.CreateUserParams{
    Username: req.Username,
    Email: req.Email, 
    Password: req.Password,
	}

  user, err := server.store.CreateUser(ctx, arg)

  if err!=nil {
    if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique_violation":
				ctx.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()} )
		return
  }

  ctx.JSON(http.StatusOK, user)
}

func (server *Server) loginUser(ctx *gin.Context) {
  
}

func (server *Server) deleteUser(ctx *gin.Context) {
  
}

func (server *Server) updateUser(ctx *gin.Context) {
  
}
