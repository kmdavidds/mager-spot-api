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

	routerGroup.GET("/login-user", r.middleware.AuthenticateUser, getLoginUser)

	routerGroup.POST("/register", r.Register)
	routerGroup.POST("/login", r.Login)

	routerGroup.PATCH("/update-user", r.middleware.AuthenticateUser, r.UpdateUser)
	routerGroup.PATCH("/update-photo", r.middleware.AuthenticateUser, r.UpdatePhoto)
	routerGroup.GET("/history", r.middleware.AuthenticateUser, r.ShowHistory)

	barang := routerGroup.Group("/barangs")
	barang.POST("", r.middleware.AuthenticateUser, r.middleware.OnlySeller, r.CreateBarang)
	barang.GET("", r.middleware.AuthenticateUser, r.FetchAllBarang)
	barang.GET("/:id", r.middleware.AuthenticateUser, r.FetchBarang)
	barang.POST("/:id/comment", r.middleware.AuthenticateUser, r.CreateComment)
	barang.GET("/:id/contact", r.middleware.AuthenticateUser, r.ContactBarang)

	kos := routerGroup.Group("/koss")
	kos.POST("", r.middleware.AuthenticateUser, r.middleware.OnlySeller, r.CreateKos)
	kos.GET("", r.middleware.AuthenticateUser, r.FetchAllKos)
	kos.GET("/:id", r.middleware.AuthenticateUser, r.FetchKos)
	kos.POST("/:id/comment", r.middleware.AuthenticateUser, r.CreateComment)
	kos.GET("/:id/contact", r.middleware.AuthenticateUser, r.ContactKos)
	
	makanan := routerGroup.Group("/makanans")
	makanan.POST("", r.middleware.AuthenticateUser, r.middleware.OnlySeller, r.CreateMakanan)
	makanan.GET("", r.middleware.AuthenticateUser, r.FetchAllMakanan)
	makanan.GET("/:id", r.middleware.AuthenticateUser, r.FetchMakanan)
	makanan.POST("/:id/comment", r.middleware.AuthenticateUser, r.CreateComment)
	makanan.GET("/:id/contact", r.middleware.AuthenticateUser, r.ContactMakanan)

	ojek := routerGroup.Group("/ojeks")
	ojek.POST("", r.middleware.AuthenticateUser, r.middleware.OnlySeller, r.CreateOjek)
	ojek.GET("", r.middleware.AuthenticateUser, r.FetchAllOjek)
	ojek.GET("/:id", r.middleware.AuthenticateUser, r.FetchOjek)
	ojek.POST("/:id/comment", r.middleware.AuthenticateUser, r.CreateComment)
	ojek.GET("/:id/contact", r.middleware.AuthenticateUser, r.ContactOjek)
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
