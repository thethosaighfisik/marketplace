package domain

import "time"

type Card struct{
	ID string
	Header string 
	Data string 
	ImageAddress string 
	Price int
	User_id string
	Email string
	Created_at time.Time
}

func CreateCard(user_id, email, Header, Data, ImageAddress string, Price int) Card{
	return Card{User_id:user_id, Email:email, Header:Header, Data:Data, ImageAddress:ImageAddress, Price:Price}
}