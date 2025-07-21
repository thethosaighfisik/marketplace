package usecases

import (
	"marketplace/internal/repository"
	"marketplace/pkg/logger"
	"marketplace/internal/domain"
	"marketplace/internal/auth"
	"errors"

	"github.com/AfterShip/email-verifier"
	"path/filepath"
)

type Usecases struct {
	repository *repository.RepositoryDB
	log *logger.SlogLogger
}

var Extensions = map[string]int{".png":0, ".pdf":0, ".jpg":0, ".raw":0,}

func NewUsecases(repository *repository.RepositoryDB, log *logger.SlogLogger) *Usecases{
	return &Usecases{repository:repository, log:log}
}

func (u *Usecases) Register(Email, Password string) error{
	if !emailverifier.IsAddressValid(Email){
		u.log.Error("email is not valid")
		return errors.New("email is not valid")
	}
	HashOfPassword, err := auth.GetHashOf(Password)
	if err != nil{
		u.log.Error("getting hash of password failed", err)
		return err
	}
	user := domain.CreateUser(Email, HashOfPassword)
	if err != nil {
		u.log.Error("create user failed")
		return err 
	}
	err = u.repository.Register(user)
	if err != nil{
		u.log.Error("registration failed")
		return err
	}
	return nil
}

func (u *Usecases) Login(Email, Password string) string{
	user, err := u.repository.IsUserExist(Email)
	if err != nil{
		u.log.Error("user does not exist, registration required")
		return ""
	}
	err = auth.ArePasswordsEqual(Password, user.Password)
	if err != nil{
		u.log.Error("password is not correct, try again")
		return ""
	}
	accessToken, err := auth.GetAccessToken(user)
	if err != nil{
		u.log.Error("creating access token failed", err)
		return ""
	}
	return accessToken
}

func (u *Usecases) PostCard(Header, Data, ImageAddress, accessToken string, Price int) error{
	user, err := auth.Parse(accessToken)
	if err != nil{
		u.log.Error("token parsing failed", err)
		return err 
	}
	card := domain.CreateCard(user.ID, user.Email, Header, Data, ImageAddress, Price)
	if err != nil{
		u.log.Error("create card failed")
		return err
	}
	err = u.isCardValid(card)
	if err != nil{
		u.log.Error("card is not valid", err)
		return err
	}
	err = u.repository.PostCard(card, user.ID, user.Email)
	if err != nil{
		return err 
	}
	return nil
}

func (u *Usecases) GetCards(accessToken, field, direction string, limit, offset, max_price, min_price int) []domain.Card{
	if accessToken == ""{
		return u.repository.GetCards("", field, direction, limit, offset, max_price, min_price)
	}
	user, err := auth.Parse(accessToken)
	if err != nil{
		u.log.Error("token parsing failed", err)
		return []domain.Card{}
	}
	return u.repository.GetCards(user.ID, field, direction, limit, offset, max_price, min_price)
}

func (u *Usecases) isCardValid(card domain.Card) error{
	if len(card.Header) < 1 || len(card.Header) > 50{
		return errors.New("Header of card is not valid")
	}
	if len(card.Data) < 1 || len(card.Data) > 200 {
		return errors.New("Data of card is not valid")
	}
	if card.Price < 1 || card.Price > 1000 {
		return errors.New("Price of card is not valid")
	}
	if len(card.ImageAddress) < 1 || len(card.ImageAddress) > 200 {
		return errors.New("Image Address of card is not valid")
	}
	_, ok := Extensions[filepath.Ext(card.ImageAddress)]
	if !ok {
		return errors.New("file extension is not valid")
	} 
	return nil
}