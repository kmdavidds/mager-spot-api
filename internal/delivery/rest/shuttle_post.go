package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
	"github.com/kmdavidds/mager-spot-api/entity"
	"github.com/kmdavidds/mager-spot-api/model"
)

func (r *Rest) CreateShuttlePost(ctx *gin.Context) {
	param := model.ShuttlePostCreate{}

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

	err = r.usecase.ShuttlePostUsecase.CreateShuttlePost(param)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to create post",
			"error":   err,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{})
}

func (r *Rest) GetShuttlePosts(ctx *gin.Context) {
	shuttlePosts, err := r.usecase.ShuttlePostUsecase.GetShuttlePosts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to get shuttle posts",
			"error":   err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"shuttlePosts": shuttlePosts,
	})
}

func (r *Rest) GetShuttlePost(ctx *gin.Context) {
	id := ctx.Param("id")
	parsedId, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to parse shuttle post id",
			"error":   err,
		})
		return
	}

	shuttlePost, err := r.usecase.ShuttlePostUsecase.GetShuttlePost(model.ShuttlePostKey{ID: parsedId})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to get shuttle post",
			"error":   err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"shuttlePost": shuttlePost,
	})
}

func (r *Rest) SearchShuttlePosts(ctx *gin.Context) {
	query := ctx.Query("query")
	if query == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "bad query",
		})
		return
	}

	param := model.ShuttlePostKey{
		Title: query,
	}
	shuttlePosts, err := r.usecase.ShuttlePostUsecase.SearchShuttlePosts(param)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to get shuttle posts",
			"error":   err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"shuttlePosts": shuttlePosts,
	})
}