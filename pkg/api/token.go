package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type renewAccessTokenResponse struct {
	AccessToken          string    `json:"access_token"`
	AccessTokenExpiresAt time.Time `json:"access_token_expires_at"`
}

func (server *Server) renewAccessToken(ctx *gin.Context) {
	refreshToken, err := ctx.Cookie("refresh_token")

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err})
		return
	}

	refreshPayload, err := server.tokenMaker.VerifyToken(refreshToken)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err})
		return
	}

	session, err := server.store.GetSession(ctx, refreshPayload.ID)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if session.RefreshToken != refreshToken {
		err := fmt.Errorf("mismatch refresh token")
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err})
	}

	accessToken, accessTokenPayload, err := server.tokenMaker.CreateToken(
		refreshPayload.Username,
		server.config.ACCESS_TOKEN_DURATION,
	)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rsp := renewAccessTokenResponse{
		AccessToken:          accessToken,
		AccessTokenExpiresAt: accessTokenPayload.ExpiredAt,
	}

	ctx.JSON(http.StatusOK, rsp)
}
