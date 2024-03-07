package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
	"github.com/kmdavidds/mager-spot-api/entity"
	"github.com/kmdavidds/mager-spot-api/model"
)

func (r *Rest) CreateKos(ctx *gin.Context) {
	param := model.KosCreate{}

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

	err = r.usecase.KosUsecase.CreateKos(param)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to create post",
			"error":   err,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{})
}

func (r *Rest) FetchAllKos(ctx *gin.Context) {
	koss, err := r.usecase.KosUsecase.GetAllKos()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed get koss",
			"error":   err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"koss": koss,
	})
}

func (r *Rest) FetchKos(ctx *gin.Context) {
	id := ctx.Param("id")
	parsedId, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to parse kos id",
			"error":   err,
		})
		return
	}

	kosWithAuthor, comments, err := r.usecase.KosUsecase.GetKos(model.KosParam{ID: parsedId})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to get kos",
			"error":   err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"kosWithAuthor": kosWithAuthor,
		"comments":         comments,
	})
}

func (r *Rest) ContactKos(ctx *gin.Context) {
	id := ctx.Param("id")
	parsedId, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to parse kos id",
			"error":   err,
		})
		return
	}

	asker, ok := ctx.Get("user")
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed get login user",
		})
		return
	}

	param := model.KosContact{
		ID: parsedId,
		AskerID: asker.(entity.User).ID,
	}

	contactLink, err := r.usecase.KosUsecase.ContactKos(param)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed get contact link",
			"error":   err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"contactLink": contactLink,
	})
}
