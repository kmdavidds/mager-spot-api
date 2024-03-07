package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/kmdavidds/mager-spot-api/internal/usecase"
	"github.com/kmdavidds/mager-spot-api/pkg/jwt_auth"
)

type Interface interface {
	AuthenticateUser(ctx *gin.Context)
	OnlySeller(ctx *gin.Context)
}

type middleware struct {
	usecase *usecase.Usecase
	jwtAuth jwt_auth.Interface
}

func Init(usecase *usecase.Usecase) Interface {
	return &middleware{
		usecase: usecase,
		jwtAuth: jwt_auth.Init(),
	}
}
