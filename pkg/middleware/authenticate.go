package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kmdavidds/mager-spot-api/model"
)

func (m *middleware) AuthenticateUser(ctx *gin.Context) {
	bearer := ctx.GetHeader("Authorization")
	if bearer == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "empty token",
		})
		ctx.Abort()
		return
	}

	token := strings.Split(bearer, " ")[1]
	userId, err := m.jwtAuth.ValidateToken(token)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "failed to validate token",
		})
		ctx.Abort()
		return
	}

	user, err := m.usecase.UserUsecase.GetUser(model.UserParam{
		ID: userId,
	})
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "failed to get user",
		})
		ctx.Abort()
		return
	}

	ctx.Set("user", user)

	ctx.Next()
}
