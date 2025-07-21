package handlers

import (
	"encoding/json"
	"net/http"
	"marketplace/internal/service"
	"marketplace/pkg/app"
	"strconv"
)

type Handlers struct {
	service *service.Service
}

func NewHandlers(service *service.Service) *Handlers{
	return &Handlers{service:service}
}

func (h *Handlers) Register(w http.ResponseWriter, r *http.Request){
	var user dto.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil{
		http.Error(w, "invalid requets", http.StatusBadRequest)
		return
	}
	err = h.service.Register(user)
	if err != nil{
		http.Error(w, "registration failed", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user.Email)
}

func (h *Handlers) Login(w http.ResponseWriter, r *http.Request){
	var user dto.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil{
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	accessToken := h.service.Login(user)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(accessToken)
}

func (h *Handlers) PostCard(w http.ResponseWriter, r *http.Request){
	accessToken := r.Header.Get("Authorization")
	if accessToken == ""{
		http.Error(w, "missing authorization token", http.StatusUnauthorized)
		return
	}
	accessToken = accessToken[7:]
	var card dto.Card
	err := json.NewDecoder(r.Body).Decode(&card)
	if err != nil{
		http.Error(w, "invalid data for request", http.StatusBadRequest)
		return
	}
	err = h.service.PostCard(card, accessToken)
	if err != nil{
		http.Error(w, "create post failed", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(card)
}

func (h *Handlers) GetCards(w http.ResponseWriter, r *http.Request){
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
    if err != nil || page < 1 {
        page = 1 
    }
    limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
    if err != nil || limit < 1 {
        limit = 10
    }
	field := r.URL.Query().Get("field")
	if len(field) == 0{
		field = "created_at"
	}
	direction := r.URL.Query().Get("direction")
	if len(direction) == 0{
		direction = "desc"
	}
	max_price, err := strconv.Atoi(r.URL.Query().Get("max_price"))
    if err != nil || max_price < 1 {
        max_price = 1000 
    }
	min_price, err := strconv.Atoi(r.URL.Query().Get("min_price"))
    if err != nil || min_price < 1 {
        min_price = 1 
    }

    offset := (page - 1) * limit

	accessToken := r.Header.Get("Authorization")
	var cards []dto.Card
	if accessToken == ""{
		cards = h.service.GetCards(accessToken, field, direction, limit, offset, max_price, min_price)
	} else {
		accessToken = accessToken[7:]
		cards = h.service.GetCards(accessToken, field, direction, limit, offset, max_price, min_price)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cards)
}