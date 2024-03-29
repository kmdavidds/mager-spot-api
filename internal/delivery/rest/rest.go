package rest

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/kmdavidds/mager-spot-api/internal/usecase"
	"github.com/kmdavidds/mager-spot-api/pkg/middleware"
)

type Rest struct {
	router     *gin.Engine
	usecase    *usecase.Usecase
	middleware middleware.Interface
}

func NewRest(usecase *usecase.Usecase, middleware middleware.Interface) *Rest {
	return &Rest{
		router:     gin.Default(),
		usecase:    usecase,
		middleware: middleware,
	}
}

func (r *Rest) MountEndpoint() {
	routerGroup := r.router.Group("/api/v1")

	routerGroup.POST("/register", r.Register)
	routerGroup.POST("/login", r.Login)
	routerGroup.GET("/login-user", r.middleware.AuthenticateUser, getLoginUser)
	routerGroup.POST("/auth-email", r.AuthenticateEmail)
	routerGroup.GET("/seller-invoice", r.middleware.AuthenticateUser, r.middleware.OnlySeller, r.GetSellerInvoices)
	routerGroup.GET("/buyer-invoice", r.middleware.AuthenticateUser, r.middleware.OnlySeller, r.GetBuyerInvoices)
	routerGroup.POST("/email-payouts", r.middleware.AuthenticateUser, r.middleware.OnlySeller, r.SendPayoutsEmail)

	routerGroup.PATCH("/update-user", r.middleware.AuthenticateUser, r.UpdateUser)
	routerGroup.PATCH("/update-photo", r.middleware.AuthenticateUser, r.UpdatePhoto)
	routerGroup.GET("/history", r.middleware.AuthenticateUser, r.ShowHistory)
	routerGroup.GET("/search", r.middleware.AuthenticateUser, r.SearchAllPosts)
	routerGroup.POST("/verify-payment", r.Verify)
	routerGroup.POST("/:category/:id/comment", r.middleware.AuthenticateUser, r.CreateComment)
	routerGroup.GET("/:category/:id/contact", r.middleware.AuthenticateUser, r.GetContactLink)
	routerGroup.GET("/:category/:id/purchase", r.middleware.AuthenticateUser, r.Purchase)

	apartmentPost := routerGroup.Group("/apartment-post")
	apartmentPost.POST("", r.middleware.AuthenticateUser, r.middleware.OnlySeller, r.CreateApartmentPost)
	apartmentPost.GET("", r.middleware.AuthenticateUser, r.GetApartmentPosts)
	apartmentPost.GET("/:id/", r.middleware.AuthenticateUser, r.GetApartmentPost)
	apartmentPost.GET("/search", r.middleware.AuthenticateUser, r.SearchApartmentPosts)

	foodPost := routerGroup.Group("/food-post")
	foodPost.POST("", r.middleware.AuthenticateUser, r.middleware.OnlySeller, r.CreateFoodPost)
	foodPost.GET("", r.middleware.AuthenticateUser, r.GetFoodPosts)
	foodPost.GET("/:id/", r.middleware.AuthenticateUser, r.GetFoodPost)
	foodPost.GET("/search", r.middleware.AuthenticateUser, r.SearchFoodPosts)

	productPost := routerGroup.Group("/product-post")
	productPost.POST("", r.middleware.AuthenticateUser, r.middleware.OnlySeller, r.CreateProductPost)
	productPost.GET("", r.middleware.AuthenticateUser, r.GetProductPosts)
	productPost.GET("/:id/", r.middleware.AuthenticateUser, r.GetProductPost)
	productPost.GET("/search", r.middleware.AuthenticateUser, r.SearchProductPosts)

	shuttlePost := routerGroup.Group("/shuttle-post")
	shuttlePost.POST("", r.middleware.AuthenticateUser, r.middleware.OnlySeller, r.CreateShuttlePost)
	shuttlePost.GET("", r.middleware.AuthenticateUser, r.GetShuttlePosts)
	shuttlePost.GET("/:id/", r.middleware.AuthenticateUser, r.GetShuttlePost)
	shuttlePost.GET("/search", r.middleware.AuthenticateUser, r.SearchShuttlePosts)
}

func (r *Rest) Serve() {
	addr := os.Getenv("APP_ADDRESS")
	port := os.Getenv("APP_PORT")

	err := r.router.Run(fmt.Sprintf("%s:%s", addr, port))
	if err != nil {
		log.Fatalf("Error while serving: %v", err)
	}
}

func getLoginUser(ctx *gin.Context) {
	user, ok := ctx.Get("user")
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed get login user",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
