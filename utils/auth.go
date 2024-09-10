package utils

import (
	"errors"
	"fmt"
	"strings"
	"time"

	bcrypt "golang.org/x/crypto/bcrypt"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.RegisteredClaims
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func CheckPassword(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func GenerateToken(email, role string) (string, error) {
	key := "secret"
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"role":  role,
		"exp":   time.Hour * 24,
	})

	return token.SignedString([]byte(key))
}

func VerifyToken(tokenString string) (*Claims, error) {

	splitToken := strings.Split(tokenString, "Bearer ")
	key := "secret"

	if len(splitToken) != 2 {
		return nil, errors.New("Invalid token")
	}

	token, err := jwt.Parse(splitToken[1], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(key), nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("Invalid token")
	}

	var claim Claims

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		claim = Claims{
			Email: claims["email"].(string),
			Role:  claims["role"].(string),
		}
	} else {
		fmt.Println(err)
	}

	return &claim, nil
}
