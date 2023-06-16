package pkg

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
)

type JWTService interface {
	GenerateToken(userID string) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtCustomClaim struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

func NewJWTService() JWTService {
	return &jwtService{
		issuer:    getIssuer(),
		secretKey: getSecretKey(),
	}
}

func getIssuer() string {
	issuer := os.Getenv("ISSUER")
	return issuer
}

func getSecretKey() string {
	secretKey := os.Getenv("JWT_SECRET")
	return secretKey
}

func (j *jwtService) GenerateToken(UserID string) (string, error) {

	expire := jwt.Time{
		Time: time.Now().AddDate(0, 0, 1),
	}

	claims := &jwtCustomClaim{
		UserID,
		jwt.StandardClaims{
			ExpiresAt: &expire,
			Issuer:    j.issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		return "", err
	}
	return t, nil
}

func (j *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("metode masuk yang tidak terduga %v", t_.Header["alg"])
		}
		return []byte(j.secretKey), nil
	})
}
