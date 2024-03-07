package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
	"github.com/kmdavidds/mager-spot-api/entity"
	"github.com/kmdavidds/mager-spot-api/model"
)

func (r *Rest) CreateOjek(ctx *gin.Context) {
	param := model.OjekCreate{}

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

	err = r.usecase.OjekUsecase.CreateOjek(param)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to create post",
			"error":   err,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{})
}

func (r *Rest) FetchAllOjek(ctx *gin.Context) {
	ojeks, err := r.usecase.OjekUsecase.GetAllOjek()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed get ojeks",
			"error":   err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"ojeks": ojeks,
	})
}

func (r *Rest) FetchOjek(ctx *gin.Context) {
	id := ctx.Param("id")
	parsedId, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to parse ojek id",
			"error":   err,
		})
		return
	}

	ojekWithAuthor, comments, err := r.usecase.OjekUsecase.GetOjek(model.OjekParam{ID: parsedId})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to get ojek",
			"error":   err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"ojekWithAuthor": ojekWithAuthor,
		"comments":         comments,
	})
}

func (r *Rest) ContactOjek(ctx *gin.Context) {
	id := ctx.Param("id")
	parsedId, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to parse ojek id",
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

	param := model.OjekContact{
		ID: parsedId,
		AskerID: asker.(entity.User).ID,
	}

	contactLink, err := r.usecase.OjekUsecase.ContactOjek(param)
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
