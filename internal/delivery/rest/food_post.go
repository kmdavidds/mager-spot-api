package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
	"github.com/kmdavidds/mager-spot-api/entity"
	"github.com/kmdavidds/mager-spot-api/model"
)

func (r *Rest) CreateFoodPost(ctx *gin.Context) {
	param := model.FoodPostCreate{}

	err := ctx.ShouldBindWith(&param, binding.FormMultipart)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to bind request body",
			"error":   err,
		})
		return
	}

	user, ok := ctx.Get("user")
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed get login user",
		})
		return
	}

	param.UserID = user.(entity.User).ID

	err = r.usecase.FoodPostUsecase.CreateFoodPost(param)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to create post",
			"error":   err,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{})
}

func (r *Rest) GetFoodPosts(ctx *gin.Context) {
	foodPosts, err := r.usecase.FoodPostUsecase.GetFoodPosts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to get food posts",
			"error":   err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"foodPosts": foodPosts,
	})
}

func (r *Rest) GetFoodPost(ctx *gin.Context) {
	id := ctx.Param("id")
	parsedId, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to parse food post id",
			"error":   err,
		})
		return
	}

	foodPost, err := r.usecase.FoodPostUsecase.GetFoodPost(model.FoodPostKey{ID: parsedId})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to get food post",
			"error":   err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"foodPost": foodPost,
	})
}
