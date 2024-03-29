package helper

import (
	"firstpro/pkg/config"
	"fmt"

	"github.com/golang-jwt/jwt"
)

func GetTokenFromHeader(header string) string {
	// Example header format: "Bearer <token>"

	if len(header) > 7 && header[:7] == "Bearer " {
		return header[7:]
	}

	return header
	// return ""
}

func ExtractUserIDFromToken(tokenString string) (int, string, error) {
	cfg, _ := config.LoadConfig()

	token, err := jwt.ParseWithClaims(tokenString, &AuthCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		}

		return []byte(cfg.KEY), nil // Replace with your actual secret key
	})

	if err != nil {
		return 0, "", err
	}

	claims, ok := token.Claims.(*AuthCustomClaims)
	if !ok {
		return 0, "", fmt.Errorf("invalid token claims")
	}

	return claims.Id, claims.Email, nil
}
