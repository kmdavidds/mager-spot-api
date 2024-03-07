package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (m *middleware) OnlySeller(ctx *gin.Context) {
	user, err := m.jwtAuth.GetLoginUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "failed to get login user",
		})
		ctx.Abort()
		return
	}

	if !user.IsSeller {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "user is not a seller",
		})
		ctx.Abort()
		return
	}

	ctx.Next()
}
