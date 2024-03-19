package main

import (
	"github.com/kmdavidds/mager-spot-api/internal/delivery/rest"
	"github.com/kmdavidds/mager-spot-api/internal/repository"
	"github.com/kmdavidds/mager-spot-api/internal/usecase"
	"github.com/kmdavidds/mager-spot-api/pkg/bcrypt"
	"github.com/kmdavidds/mager-spot-api/pkg/config"
	"github.com/kmdavidds/mager-spot-api/pkg/database/postgresql"
	"github.com/kmdavidds/mager-spot-api/pkg/jwt_auth"
	"github.com/kmdavidds/mager-spot-api/pkg/middleware"
	"github.com/kmdavidds/mager-spot-api/pkg/supabase"
)

func init() {
	config.LoadEnvVariables()
	config.LoadMidtransConfig()
}

func main() {
	bcrypt := bcrypt.Init()

	jwtAuth := jwt_auth.Init()

	supabase := supabase.Init()

	db := postgresql.ConnectDatabase()

	postgresql.Migrate(db)

	repository := repository.NewRepository(db)

	usecase := usecase.NewUsecase(usecase.InitParam{
		Repository: repository,
		Bcrypt:     bcrypt,
		JWTAuth:    jwtAuth,
		Supabase:   supabase,
	})

	middleware := middleware.Init(usecase)

	rest := rest.NewRest(usecase, middleware)

	rest.MountEndpoint()

	rest.Serve()
}
