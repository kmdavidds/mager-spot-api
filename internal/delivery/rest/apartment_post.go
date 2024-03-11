package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
	"github.com/kmdavidds/mager-spot-api/entity"
	"github.com/kmdavidds/mager-spot-api/model"
)

func (r *Rest) CreateApartmentPost(ctx *gin.Context) {
	param := model.ApartmentPostCreate{}

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

	err = r.usecase.ApartmentPostUsecase.CreateApartmentPost(param)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to create post",
			"error":   err,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{})
}

func (r *Rest) GetApartmentPosts(ctx *gin.Context) {
	apartmentPosts, err := r.usecase.ApartmentPostUsecase.GetApartmentPosts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to get apartment posts",
			"error":   err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"apartmentPosts": apartmentPosts,
	})
}

func (r *Rest) GetApartmentPost(ctx *gin.Context) {
	id := ctx.Param("id")
	parsedId, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to parse apartment post id",
			"error":   err,
		})
		return
	}

	apartmentPost, err := r.usecase.ApartmentPostUsecase.GetApartmentPost(model.ApartmentPostKey{ID: parsedId})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to get apartment post",
			"error":   err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"apartmentPost": apartmentPost,
	})
}
