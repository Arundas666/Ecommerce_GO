package helper

import (
	"firstpro/config"
	"firstpro/utils/models"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type authCustomClaimsShippingCoordinator struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	jwt.StandardClaims
}

func GenerateTokenShippingCoordinator(shippingCoordinator models.ShippingCoordinatorDetailsResponse) (string, error) {
	cfg, _ := config.LoadConfig()

	claims := &authCustomClaimsShippingCoordinator{
		Firstname: shippingCoordinator.Firstname,
		Lastname:  shippingCoordinator.Lastname,
		Email:     shippingCoordinator.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(cfg.KEY_FOR_SHIPPING_COORDINATOR))

	if err != nil {
		fmt.Println("Error is ", err)
		return "", err
	}

	return tokenString, nil
}

func ValidateTokenShippingCoordinator(tokenString string) (*authCustomClaimsShippingCoordinator, error) {
	cfg, _ := config.LoadConfig()

	token, err := jwt.ParseWithClaims(tokenString, &authCustomClaimsShippingCoordinator{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(cfg.KEY_FOR_SHIPPING_COORDINATOR), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*authCustomClaimsShippingCoordinator); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}
