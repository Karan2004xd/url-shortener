package internal

import (
	"errors"
	"strings"
	"time"
	"url-shortner/config"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"user_id": userId,
		"exp": time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte (config.GetJwtSecretKey()))
}

func VerifyToken(token string) (int64, error) {
	index := strings.Index(token, "Bearer ")

	if index != -1 {
		token = strings.Replace(token, "Bearer ", "", 1)
	}

	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("Unexpected signing method.")
		}

		return []byte(config.GetJwtSecretKey()), nil
	})

	if err != nil {
		return 0, errors.New("Could not parse the jwt token")
	}

	if !parsedToken.Valid {
		return 0, errors.New("Invalid JWT token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return 0, errors.New("Invalid token claims.")
	}

	userId := int64(claims["user_id"].(float64))
	return userId, nil
}
