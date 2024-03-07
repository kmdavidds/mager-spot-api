package main

import (
	"github.com/kmdavidds/mager-spot-api/internal/delivery/rest"
	"github.com/kmdavidds/mager-spot-api/internal/repository"
	"github.com/kmdavidds/mager-spot-api/internal/usecase"
	"github.com/kmdavidds/mager-spot-api/pkg/bcrypt"
	"github.com/kmdavidds/mager-spot-api/pkg/config"
	"github.com/kmdavidds/mager-spot-api/pkg/database/postgresql"
	"github.com/kmdavidds/mager-spot-api/pkg/jwt"
	"github.com/kmdavidds/mager-spot-api/pkg/middleware"
)

func init() {
	config.LoadEnvVariables()
}

func main() {
	bcrypt := bcrypt.Init()

	jwt := jwt.Init()

	db := postgresql.ConnectDatabase()

	postgresql.Migrate(db)

	repository := repository.NewRepository(db)

	usecase := usecase.NewUsecase(usecase.InitParam{
		Repository: repository,
		Bcrypt:     bcrypt,
		JWT:        jwt,
	})

	middleware := middleware.Init(usecase)

	rest := rest.NewRest(usecase, middleware)

	rest.MountEndpoint()

	rest.Serve()
}
