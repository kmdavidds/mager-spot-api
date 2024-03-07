package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
	"github.com/kmdavidds/mager-spot-api/entity"
	"github.com/kmdavidds/mager-spot-api/model"
)

func (r *Rest) CreateMakanan(ctx *gin.Context) {
	param := model.MakananCreate{}

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

	err = r.usecase.MakananUsecase.CreateMakanan(param)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to create post",
			"error":   err,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{})
}

func (r *Rest) FetchAllMakanan(ctx *gin.Context) {
	makanans, err := r.usecase.MakananUsecase.GetAllMakanan()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed get makanans",
			"error":   err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"makanans": makanans,
	})
}

func (r *Rest) FetchMakanan(ctx *gin.Context) {
	id := ctx.Param("id")
	parsedId, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to parse makanan id",
			"error":   err,
		})
		return
	}

	makananWithAuthor, comments, err := r.usecase.MakananUsecase.GetMakanan(model.MakananParam{ID: parsedId})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to get makanan",
			"error":   err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"makananWithAuthor": makananWithAuthor,
		"comments":         comments,
	})
}

func (r *Rest) ContactMakanan(ctx *gin.Context) {
	id := ctx.Param("id")
	parsedId, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to parse makanan id",
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

	param := model.MakananContact{
		ID: parsedId,
		AskerID: asker.(entity.User).ID,
	}

	contactLink, err := r.usecase.MakananUsecase.ContactMakanan(param)
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
