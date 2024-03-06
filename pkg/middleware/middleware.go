package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/kmdavidds/mager-spot-api/internal/usecase"
	"github.com/kmdavidds/mager-spot-api/pkg/jwt"
)

type Interface interface {
	AuthenticateUser(ctx *gin.Context)
}

type middleware struct {
	usecase *usecase.Usecase
	jwt     jwt.Interface
}

func Init(usecase *usecase.Usecase) Interface {
	return &middleware{
		usecase: usecase,
		jwt: jwt.Init(),
	}
}
