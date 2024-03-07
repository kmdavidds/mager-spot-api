package usecase

import (
	"strings"

	"github.com/google/uuid"
	"github.com/kmdavidds/mager-spot-api/entity"
	"github.com/kmdavidds/mager-spot-api/internal/repository"
	"github.com/kmdavidds/mager-spot-api/model"
	"github.com/kmdavidds/mager-spot-api/pkg/bcrypt"
	"github.com/kmdavidds/mager-spot-api/pkg/jwt_auth"
	"github.com/kmdavidds/mager-spot-api/pkg/supabase"
)

type IUserUsecase interface {
	Register(param model.UserRegister) error
	Login(param model.UserLogin) (model.UserLoginResponse, error)
	GetUser(param model.UserParam) (entity.User, error)
	UpdateUser(param model.UserUpdates, user entity.User) error
}

type UserUsecase struct {
	ur       repository.IUserRepository
	bcrypt   bcrypt.Interface
	jwtAuth  jwt_auth.Interface
	supabase supabase.Interface
}

func NewUserUsecase(userRepository repository.IUserRepository, bcrypt bcrypt.Interface, jwtAuth jwt_auth.Interface, supabase supabase.Interface) IUserUsecase {
	return &UserUsecase{
		ur:      userRepository,
		bcrypt:  bcrypt,
		jwtAuth: jwtAuth,
		supabase: supabase,
	}
}

func (uu *UserUsecase) Register(param model.UserRegister) error {
	hashedPassword, err := uu.bcrypt.GenerateFromPassword(param.Password)
	if err != nil {
		return err
	}

	param.ID = uuid.New()
	param.Password = hashedPassword

	user := entity.User{
		ID:       param.ID,
		Username: param.Username,
		Email:    param.Email,
		Password: param.Password,
	}
	
	if (strings.Split(param.Email, "@")[1] == "student.ub.ac.id") {
		user.IsSeller = true
	}
	
	_, err = uu.ur.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (uu *UserUsecase) Login(param model.UserLogin) (model.UserLoginResponse, error) {
	result := model.UserLoginResponse{}

	user, err := uu.ur.GetUser(model.UserParam{
		Username: param.Username,
	})
	if err != nil {
		return result, err
	}

	err = uu.bcrypt.CompareHashAndPassword(user.Password, param.Password)
	if err != nil {
		return result, err
	}

	token, err := uu.jwtAuth.CreateJWTToken(user.ID)
	if err != nil {
		return result, err
	}

	result.Token = token

	return result, nil
}

func (uu *UserUsecase) GetUser(param model.UserParam) (entity.User, error) {
	return uu.ur.GetUser(param)
}

func (uu *UserUsecase) UpdateUser(param model.UserUpdates, user entity.User) error {
	err := uu.ur.UpdateUser(param, user)
	if err != nil {
		return err
	}

	return nil
}