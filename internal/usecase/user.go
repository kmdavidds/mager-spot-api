package usecase

import (
	"fmt"
	"net/smtp"
	"net/url"
	"os"
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
	UpdatePhoto(param model.PhotoUpdate) error
	ShowHistory(user entity.User) ([]entity.History, error)
	CreateHistoryRecord(param model.SellerContact) error
	GetContactLink(param model.SellerContact) (string, error)
	GetSellerInvoices(user entity.User) ([]entity.Invoice, error)
	GetBuyerInvoices(user entity.User) ([]entity.Invoice, error)
	SendPayoutsEmail(user entity.User) error
}

type UserUsecase struct {
	ur      repository.IUserRepository
	bcrypt  bcrypt.Interface
	jwtAuth jwt_auth.Interface
	sb      supabase.Interface
}

func NewUserUsecase(userRepository repository.IUserRepository, bcrypt bcrypt.Interface, jwtAuth jwt_auth.Interface, supabase supabase.Interface) IUserUsecase {
	return &UserUsecase{
		ur:      userRepository,
		bcrypt:  bcrypt,
		jwtAuth: jwtAuth,
		sb:      supabase,
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

	if strings.Split(param.Email, "@")[1] == "student.ub.ac.id" {
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

func (uu *UserUsecase) UpdatePhoto(param model.PhotoUpdate) error {
	param.Image.Filename = fmt.Sprintf("%s.%s", param.UserID.String(), strings.Split(param.Image.Filename, ".")[1])

	if param.PhotoLink != "" {
		err := uu.sb.Delete(param.PhotoLink)
		if err != nil {
			return err
		}
	}

	imageLink, err := uu.sb.Upload(param.Image)
	if err != nil {
		return err
	}

	param.PhotoLink = imageLink

	err = uu.ur.UpdatePhoto(param)
	if err != nil {
		return err
	}

	return nil
}

func (uu *UserUsecase) ShowHistory(user entity.User) ([]entity.History, error) {
	return uu.ur.ShowHistory(user)
}

func (uu *UserUsecase) CreateHistoryRecord(param model.SellerContact) error {
	history := entity.History{
		ID:             uuid.New(),
		Category:       param.Category,
		UserID:         param.User.ID,
		SellerUsername: param.Seller.Username,
	}

	switch param.Category {
	case "apartment-post":
		history.PostID = param.ApartmentPost.ID
		history.Title = param.ApartmentPost.Title
		history.Price = param.ApartmentPost.Price
		history.ImageLink = param.ApartmentPost.PictureLink
	case "food-post":
		history.PostID = param.FoodPost.ID
		history.Title = param.FoodPost.Title
		history.Price = param.FoodPost.Price
		history.ImageLink = param.FoodPost.PictureLink
	case "product-post":
		history.PostID = param.ProductPost.ID
		history.Title = param.ProductPost.Title
		history.Price = param.ProductPost.Price
		history.ImageLink = param.ProductPost.PictureLink
	case "shuttle-post":
		history.PostID = param.ShuttlePost.ID
		history.Title = param.ShuttlePost.Title
		history.Price = param.ShuttlePost.Price
		history.ImageLink = param.ShuttlePost.PictureLink
	}

	return uu.ur.CreateHistoryRecord(history)
}

func (uu *UserUsecase) GetContactLink(param model.SellerContact) (string, error) {
	var postTitle string
	message := ""

	switch param.Category {
	case "product-post":
		postTitle = param.ProductPost.Title
	case "food-post":
		postTitle = param.FoodPost.Title
	case "apartment-post":
		postTitle = param.ApartmentPost.Title
		message += fmt.Sprintf("Tanggal Kedatangan Kos: %s", param.Date)
	case "shuttle-post":
		postTitle = param.ShuttlePost.Title
	}

	param.User.DisplayName = strings.ReplaceAll(param.User.DisplayName, " ", "%20")
	postTitle = strings.ReplaceAll(postTitle, " ", "%20")

	message += fmt.Sprintf("Halo! nama saya %s.\nSaya tertarik dengan postingan anda yang berjudul %s.\nApakah masih tersedia?", param.User.DisplayName, postTitle)

	contactLink := fmt.Sprintf("https://wa.me/%s?text=%s", param.Seller.PhoneNumber, url.QueryEscape(message))

	return contactLink, nil
}

func (uu *UserUsecase) GetSellerInvoices(user entity.User) ([]entity.Invoice, error) {
	return uu.ur.GetSellerInvoices(user)
}

func (uu *UserUsecase) GetBuyerInvoices(user entity.User) ([]entity.Invoice, error) {
	return uu.ur.GetBuyerInvoices(user)
}

func (uu *UserUsecase) SendPayoutsEmail(user entity.User) error {
	auth := smtp.PlainAuth(
		"",
		os.Getenv("COMPANY_EMAIL"),
		os.Getenv("EMAILER_PASSWORD"),
		"smtp.gmail.com",
	)

	emailBody := fmt.Sprintf("ID: %s\n"+
		"Username: %s\n"+
		"Email: %s\n"+
		"Phone Number: %s\n"+
		"Balance: %d\n", user.ID, user.Username, user.Email, user.PhoneNumber, user.Balance)

	msg := fmt.Sprintf("To: magerspot@gmail.com\r\n"+
		"Subject: Request Payout\r\n"+
		"\r\n"+
		"%s\r\n", emailBody)

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		user.Email,
		[]string{os.Getenv("COMPANY_EMAIL")},
		[]byte(msg),
	)
	if err != nil {
		return err
	}

	return nil
}
