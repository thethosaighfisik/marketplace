package service

import (
	"marketplace/internal/usecases"
	"marketplace/pkg/logger"
	"marketplace/pkg/app"
)

type Service struct{
	usecases *usecases.Usecases
	log *logger.SlogLogger

}

func NewService(usecases *usecases.Usecases, log *logger.SlogLogger) *Service{
	return &Service{usecases:usecases, log:log}
}

func (s *Service) Register(user dto.User) error{
	return s.usecases.Register(user.Email, user.Password)
}

func (s *Service) Login(user dto.User) string{
	return s.usecases.Login(user.Email, user.Password)
}

func (s *Service) PostCard(card dto.Card, accessToken string) error{
	return s.usecases.PostCard(card.Header, card.Data, card.ImageAddress, accessToken, card.Price,)
}

func (s *Service) GetCards(accessToken, field, direction string, limit, offset, max_price, min_price int) []dto.Card{
	cards := s.usecases.GetCards(accessToken, field, direction, limit, offset, max_price, min_price)
	dtoCards := make([]dto.Card, 0)
	for _, card := range cards{
		var dtoCard dto.Card
		dtoCard.Header = card.Header
		dtoCard.Data = card.Data
		dtoCard.ImageAddress = card.ImageAddress
		dtoCard.Price = card.Price
		dtoCard.Email = card.Email
		dtoCard.Created_at = card.Created_at
		dtoCards = append(dtoCards, dtoCard)
	}
	return dtoCards
}