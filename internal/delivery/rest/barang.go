package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
	"github.com/kmdavidds/mager-spot-api/entity"
	"github.com/kmdavidds/mager-spot-api/model"
)

func (r *Rest) CreateBarang(ctx *gin.Context) {
	param := model.BarangCreate{}

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

	err = r.usecase.BarangUsecase.CreateBarang(param)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to create post",
			"error":   err,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{})
}

func (r *Rest) FetchAllBarang(ctx *gin.Context) {
	barangs, err := r.usecase.BarangUsecase.GetAllBarang()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed get barangs",
			"error":   err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"barangs": barangs,
	})
}

func (r *Rest) FetchBarang(ctx *gin.Context) {
	id := ctx.Param("id")
	parsedId, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to parse barang id",
			"error":   err,
		})
		return
	}

	barangWithAuthor, comments, err := r.usecase.BarangUsecase.GetBarang(model.BarangParam{ID: parsedId})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to get barang",
			"error":   err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"barangWithAuthor": barangWithAuthor,
		"comments":         comments,
	})
}

func (r *Rest) ContactBarang(ctx *gin.Context) {
	id := ctx.Param("id")
	parsedId, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to parse barang id",
			"error":   err,
		})
		return
	}

	param := model.BarangContact{
		ID: parsedId,
	}

	contactLink, err := r.usecase.BarangUsecase.ContactBarang(param)
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
