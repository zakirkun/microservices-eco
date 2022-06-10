package jwt

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JWTFields struct {
	UserID int
	Name   string
	Email  string
}

type JWTServices interface {
	Generate(data JWTFields) (map[string]string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct {
}

var (
	SECRET_KEY = []byte(os.Getenv("JWT_KEY"))
	EXP_HOUR   = 24
	EXP_MINUTE = 15
)

func New() *jwtService {
	return &jwtService{}
}

func (s *jwtService) Generate(data JWTFields) (map[string]string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claim := jwt.MapClaims{}
	claim["userID"] = data.UserID
	claim["name"] = data.Name
	claim["email"] = data.Email
	claim["exp"] = time.Now().Add(time.Minute * time.Duration(EXP_MINUTE)).Unix()

	t, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return nil, err
	}

	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["userID"] = data.UserID
	rtClaims["exp"] = time.Now().Add(time.Hour * time.Duration(EXP_HOUR)).Unix()

	rt, err := refreshToken.SignedString([]byte("secret"))
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"access_token":  t,
		"refresh_token": rt,
	}, nil
}

func (s *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	valid, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("Invalid Token")
		}

		return []byte(SECRET_KEY), nil
	})

	return valid, err
}
