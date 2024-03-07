package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kmdavidds/mager-spot-api/entity"
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

func (r *Rest) UpdateUser(ctx *gin.Context) {
	param := model.UserUpdates{}

	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to bind request body",
			"error": err,
		})
		return
	}

	user, ok := ctx.Get("user")
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to get login user",
		})
		return
	}

	err = r.usecase.UserUsecase.UpdateUser(param, user.(entity.User))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to update user",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}
