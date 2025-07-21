package repository

import (
	"marketplace/internal/domain"
	"marketplace/pkg/logger"
	"marketplace/pkg/repository"
	"fmt"
)

type Repository interface{
	Register(user domain.User)
	IsUserExist(Email string) (domain.User, error)
	PostCard(Card domain.Card, user_id string)
	GetCards(user_id string) []domain.Card
}

type RepositoryDB struct {
	db *database.Database
	log *logger.SlogLogger
}

func NewRepositoryDB(db *database.Database, log *logger.SlogLogger) *RepositoryDB{
	return &RepositoryDB{db:db, log:log}
}

func (r *RepositoryDB) Register(user domain.User) error{
	_, err := r.db.DB.Exec("insert into users (Email, Password) values ($1, $2)", user.Email, user.Password)
	if err != nil {
		r.log.Error("insert into table failed", err)
		return err
	}
	// r.log.Info("add user is successful")
	return nil
}

func (r *RepositoryDB) IsUserExist(Email string) (domain.User, error){
	var user domain.User
	err := r.db.DB.QueryRow("select id, email, password from users where Email = $1", Email).Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		r.log.Error("", err)
		return domain.User{}, err
	}
	return user, nil
}

func (r *RepositoryDB) PostCard(Card domain.Card, user_id, email string) error{
	_, err := r.db.DB.Exec("insert into cards (User_id, Email, Header, Data, ImageAddress, Price) values ($1, $2, $3, $4, $5, $6)", user_id, email, Card.Header, Card.Data, Card.ImageAddress, Card.Price)
	if err != nil {
		r.log.Error("", err)
		return err
	}
	return nil
	// r.log.Info("add card is successful")
}

var safeFields = map[string]string{
	"price":"price",
	"created_at":"created_at",
}
var safeDirections = map[string]string{
	"asc":"asc",
	"desc":"desc",
}

func (r *RepositoryDB) GetCards(user_id, field, direction string, limit, offset, max_price, min_price int) []domain.Card{
	if user_id == ""{
		query := fmt.Sprintf("select Header, Data, ImageAddress, Price, Created_at from cards WHERE (Price > %d AND Price < %d) order by %s %s LIMIT $1 OFFSET $2", min_price, max_price, safeFields[field], safeDirections[direction])
		rows, err := r.db.DB.Query(query, limit, offset)
		if err != nil {
			r.log.Error("getting cards failed", err)
		}
		defer rows.Close()
		var cards []domain.Card
		for rows.Next(){
			card := domain.Card{}
			err = rows.Scan(&card.Header, &card.Data, &card.ImageAddress, &card.Price, &card.Created_at)
			if err != nil {
				r.log.Error("%s", err)
			}
			cards = append(cards, card)
		}
		return cards
	} else {
		query := fmt.Sprintf("select Header, Data, ImageAddress, Price, Email, Created_at from cards WHERE (Price > %d AND Price < %d) order by %s %s LIMIT $1 OFFSET $2", min_price, max_price, safeFields[field], safeDirections[direction])
		rows, err := r.db.DB.Query(query, limit, offset)
		if err != nil {
			r.log.Error("getting cards failed", err)
		}
		defer rows.Close() 
		var cards []domain.Card
		for rows.Next(){
			card := domain.Card{}
			err = rows.Scan(&card.Header, &card.Data, &card.ImageAddress, &card.Price, &card.Email, &card.Created_at)
			if err != nil {
				r.log.Error("%s", err)
			}
			cards = append(cards, card)
		}
		return cards
	}
	return []domain.Card{}
	
}