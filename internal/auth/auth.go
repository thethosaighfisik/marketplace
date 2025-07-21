package auth

import (
	"golang.org/x/crypto/bcrypt"
	"github.com/golang-jwt/jwt/v4"
	"marketplace/internal/domain"
	"errors"
	"time"
)

const (
	secretKey = "mysecreykey"
)


func GetHashOf(password string) (string, error){
	if len(password) < 8 || len(password) > 25{
		return "", errors.New("password is not valid")
	}
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func ArePasswordsEqual(password string, hashedPassword string) error{
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func GetAccessToken(user domain.User) (string, error){
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.ID,
		"email":    user.Email,
		"exp":      time.Now().Add(time.Minute * 1).Unix(),
	   })
	  
	accessTokenString, err := accessToken.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return accessTokenString, nil
}



func Parse(accessTokenString string) (domain.User, error) {
	accessToken, err := jwt.Parse(accessTokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		 return nil, errors.New("token parse failed")
		}
		return []byte(secretKey), nil
	})
	  
	var user domain.User
	if errors.Is(err, jwt.ErrTokenExpired) || !accessToken.Valid{
		claims, _ := accessToken.Claims.(jwt.MapClaims)
		user.ID = claims["id"].(string)
		user.Email = claims["email"].(string)
		return user, err//errors.New("Timing is everything")
	}
	if err != nil || !accessToken.Valid {
		return user, err
	}


	
	claims, ok := accessToken.Claims.(jwt.MapClaims)
	if ok && accessToken.Valid {
		user.ID = claims["id"].(string)
		user.Email = claims["email"].(string)
	}
	
	return user, nil
}