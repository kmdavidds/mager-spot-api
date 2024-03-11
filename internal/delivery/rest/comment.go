package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kmdavidds/mager-spot-api/entity"
	"github.com/kmdavidds/mager-spot-api/model"
)

func (r *Rest) CreateComment(ctx *gin.Context) {
	param := model.CommentCreate{}

	err := ctx.ShouldBindJSON(&param)
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

	id := ctx.Param("id")
	parsedId, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to parse post id",
			"error":   err,
		})
		return
	}

	category := ctx.Param("category")

	switch category {
	case "apartment-post":
		param.ApartmentPostID = parsedId
	case "food-post":
		param.FoodPostID = parsedId
	case "product-post":
		param.ProductPostID = parsedId
	case "shuttle-post":
		param.ShuttlePostID = parsedId
	}

	err = r.usecase.CommentUsecase.CreateComment(param)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to create comment",
			"error":   err,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{})
}
