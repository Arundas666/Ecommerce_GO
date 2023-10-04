package helper

import (
	"errors"
	"firstpro/utils/models"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type AuthCustomClaims struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

func PasswordHashing(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", errors.New("internal server error")
	}
	hash := string(hashedPassword)
	return hash, nil
}
func GenerateTokenClients(user models.SignupDetailResponse) (string, error) {
	claims := &AuthCustomClaims{
		Id:    user.Id,
		Email: user.Email,
		Role:  "client",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("cosmetics"))
	if err != nil {
		fmt.Println("Token error is ", err)
		return "", err
	}
	return tokenString, nil

}
