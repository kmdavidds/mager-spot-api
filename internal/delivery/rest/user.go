package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kmdavidds/mager-spot-api/model"
)

func (r *Rest) Register(ctx *gin.Context) {
	param := model.UserRegister{}

	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to bind request body",
			"error": err,
		})
		return
	}

	err = r.usecase.UserUsecase.Register(param)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to register user",
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{})
}

func (r *Rest) Login(ctx *gin.Context) {
	param := model.UserLogin{}

	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to bind request body",
			"error": err,
		})
		return
	}

	token, err := r.usecase.UserUsecase.Login(param)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to log in user",
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, token)
}
