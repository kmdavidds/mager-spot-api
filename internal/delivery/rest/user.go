package rest

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/kmdavidds/mager-spot-api/entity"
	"github.com/kmdavidds/mager-spot-api/model"
	"gorm.io/gorm"
)

func (r *Rest) Register(ctx *gin.Context) {
	param := model.UserRegister{}

	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to bind request body",
			"error":   err,
		})
		return
	}

	err = r.usecase.UserUsecase.Register(param)
	if err != nil {
		if err.(*pgconn.PgError).Code == "23505" {
			ctx.JSON(http.StatusConflict, gin.H{
				"message": "user already exists",
				"error":   err,
			})
		} else if errors.Is(err, gorm.ErrInvalidData) {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "invalid input data",
				"error":   err,
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "failed to register user",
				"error":   err,
			})
		}
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
			"error":   err,
		})
		return
	}

	token, err := r.usecase.UserUsecase.Login(param)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "user not found or invalid credentials",
				"error":   err,
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "failed to log in user",
				"error":   err,
			})
		}
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
			"error":   err,
		})
		return
	}

	user, ok := ctx.Get("user")
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "failed to get login user",
		})
		return
	}

	err = r.usecase.UserUsecase.UpdateUser(param, user.(entity.User))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": "user not found",
				"error":   err,
			})
		} else if errors.Is(err, gorm.ErrInvalidData) {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "invalid input data",
				"error":   err,
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "failed to update user",
				"error":   err,
			})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}

func (r *Rest) UpdatePhoto(ctx *gin.Context) {
	param := model.PhotoUpdate{}

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
	param.PhotoLink = user.(entity.User).ProfilePhotoLink

	err = r.usecase.UserUsecase.UpdatePhoto(param)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": "user not found",
				"error":   err,
			})
		} else if errors.Is(err, gorm.ErrInvalidData) {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "invalid input data",
				"error":   err,
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "failed to update user photo",
				"error":   err,
			})
		}
	}

	ctx.JSON(http.StatusOK, gin.H{})
}

func (r *Rest) ShowHistory(ctx *gin.Context) {
	user, ok := ctx.Get("user")
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "failed to get login user",
		})
		return
	}

	historyAll, err := r.usecase.UserUsecase.ShowHistory(user.(entity.User))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to fetch histories",
			"error":   err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"histories": historyAll,
	})
}

func (r *Rest) GetContactLink(ctx *gin.Context) {
	user, ok := ctx.Get("user")
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "failed get login user",
		})
		return
	}

	id := ctx.Param("id")
	parsedId, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to parse post id",
			"error":   err,
		})
		return
	}

	param := model.SellerContact{
		User: user.(entity.User),
	}

	category := ctx.Param("category")
	switch category {
	case "apartment-post":
		apartmentPost, err := r.usecase.ApartmentPostUsecase.GetApartmentPost(model.ApartmentPostKey{ID: parsedId})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "failed to get apartment post",
				"error":   err,
			})
			return
		}
		param.ApartmentPost = apartmentPost
		param.Seller = apartmentPost.User
		param.Category = "apartment-post"
	case "food-post":
		foodPost, err := r.usecase.FoodPostUsecase.GetFoodPost(model.FoodPostKey{ID: parsedId})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "failed to get food post",
				"error":   err,
			})
			return
		}
		param.FoodPost = foodPost
		param.Seller = foodPost.User
		param.Category = "food-post"
	case "product-post":
		productPost, err := r.usecase.ProductPostUsecase.GetProductPost(model.ProductPostKey{ID: parsedId})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "failed to get product post",
				"error":   err,
			})
			return
		}
		param.ProductPost = productPost
		param.Seller = productPost.User
		param.Category = "product-post"
	case "shuttle-post":
		shuttlePost, err := r.usecase.ShuttlePostUsecase.GetShuttlePost(model.ShuttlePostKey{ID: parsedId})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "failed to get shuttle post",
				"error":   err,
			})
			return
		}
		param.ShuttlePost = shuttlePost
		param.Seller = shuttlePost.User
		param.Category = "shuttle-post"
	}

	contactLink, err := r.usecase.UserUsecase.GetContactLink(param)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": "seller not found",
				"error":   err,
			})
		} else if errors.Is(err, gorm.ErrInvalidData) {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "invalid input data",
				"error":   err,
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "failed to get contact link",
				"error":   err,
			})
		}
		return
	}

	err = r.usecase.UserUsecase.CreateHistoryRecord(param)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to create a history record",
			"error":   err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"contactLink": contactLink,
	})
}

func (r *Rest) AuthenticateEmail(ctx *gin.Context) {
	param := model.EmailAuth{}

	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to bind request body",
			"error":   err,
		})
		return
	}

	_, err = r.usecase.UserUsecase.GetUser(model.UserParam{Email: param.Email})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": "user not found",
				"error":   err,
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "failed to get user",
				"error":   err,
			})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}

func (r *Rest) SearchAllPosts(ctx *gin.Context) {
	query := ctx.Query("query")
	if query == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "bad query",
		})
		return
	}

	apartmentPosts, err := r.usecase.ApartmentPostUsecase.SearchApartmentPosts(model.ApartmentPostKey{Title: query})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to get apartment posts",
			"error":   err,
		})
		return
	}
	foodPosts, err := r.usecase.FoodPostUsecase.SearchFoodPosts(model.FoodPostKey{Title: query})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to get food posts",
			"error":   err,
		})
		return
	}
	productPosts, err := r.usecase.ProductPostUsecase.SearchProductPosts(model.ProductPostKey{Title: query})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to get product posts",
			"error":   err,
		})
		return
	}
	shuttlePosts, err := r.usecase.ShuttlePostUsecase.SearchShuttlePosts(model.ShuttlePostKey{Title: query})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to get shuttle posts",
			"error":   err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"apatmentPosts": apartmentPosts,
		"foodPosts":     foodPosts,
		"productPosts":  productPosts,
		"shuttlePosts":  shuttlePosts,
	})
}