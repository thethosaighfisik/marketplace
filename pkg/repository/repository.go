package database

import (
	"database/sql"
	"fmt"
	"log"
  _ "github.com/lib/pq"
  "marketplace/cmd/app/config"

)

type Database struct {
	DB *sql.DB
}

func NewDatabase(r config.Repository) *Database{
	sqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", r.Host, r.Port, r.User, r.Password, r.Name)

	db, err := sql.Open("postgres", sqlInfo)
	if err != nil{
		log.Fatal("database connect failed: %s", err)
	}
	return &Database{DB : db}
}