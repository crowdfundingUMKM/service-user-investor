package auth

import (
	"errors"
	"os"

	"github.com/dgrijalva/jwt-go"
)

type Service interface {
	GenerateToken(userUnixID string) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct {
}

var secret = os.Getenv("JWT_SECRET")
var SECRET_KEY = []byte(secret)

func NewService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(userUnixID string) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userUnixID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signedToken, err
	}
	return signedToken, nil

}

// ValidateToken is used to validate token from user input and return the token if it is valid or return error if it is invalid
func (s *jwtService) ValidateToken(endcodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(endcodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("Invalid token")
		}

		return []byte(SECRET_KEY), nil
	})
	if err != nil {
		return token, err
	}
	return token, nil
}
