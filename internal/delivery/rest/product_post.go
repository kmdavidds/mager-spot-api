package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
	"github.com/kmdavidds/mager-spot-api/entity"
	"github.com/kmdavidds/mager-spot-api/model"
)

func (r *Rest) CreateProductPost(ctx *gin.Context) {
	param := model.ProductPostCreate{}

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
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "failed get login user",
		})
		return
	}

	param.UserID = user.(entity.User).ID

	err = r.usecase.ProductPostUsecase.CreateProductPost(param)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to create post",
			"error":   err,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{})
}

func (r *Rest) GetProductPosts(ctx *gin.Context) {
	productPosts, err := r.usecase.ProductPostUsecase.GetProductPosts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to get product posts",
			"error":   err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"productPosts": productPosts,
	})
}

func (r *Rest) GetProductPost(ctx *gin.Context) {
	id := ctx.Param("id")
	parsedId, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to parse product post id",
			"error":   err,
		})
		return
	}

	productPost, err := r.usecase.ProductPostUsecase.GetProductPost(model.ProductPostKey{ID: parsedId})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to get product post",
			"error":   err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"productPost": productPost,
	})
}
