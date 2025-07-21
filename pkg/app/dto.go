package dto

import (
	"time"
)

type User struct{
	Password string `json: "password"`
	Email string `json: "email"`
}

type Card struct{
	Header string `json: "header"`
	Data string `json: "data"`
	ImageAddress string `json: "imageaddress"`
	Price int `json: "price"`
	Email string `json: "email"`
	Created_at time.Time `json: "created_at"`
}