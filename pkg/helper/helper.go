package helper

import (
	cfg "Ecom/pkg/config"
	"Ecom/pkg/domain"
	"Ecom/pkg/helper/interfaces"
	"Ecom/pkg/utils/models"
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type helper struct {
	config cfg.Config
}

func NewHelper(cfg cfg.Config) interfaces.Helper {
	return &helper{
		config: cfg,
	}
}

func (h *helper) CreateHashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", errors.New("internal server error")
	}

	hash := string(hashedPassword)
	return hash, nil
}

func (h *helper) CompareHashAndPassword(a string, b string) error {
	err := bcrypt.CompareHashAndPassword([]byte(a), []byte(b))
	if err != nil {
		return err
	}
	return nil
}

func (helper *helper) GenerateTokenAdmin(admin models.AdminDetailsResponse) (string, error) {
	tokenClaims := &models.AuthCustomClaims{
		Id:    admin.ID,
		Email: admin.Email,
		Role:  "admin",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 50).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)
	tokenString, err := token.SignedString([]byte("adminsecret")) //take this from runtime in future avoid hardcoding
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (helper *helper) GenerateTokenUser(details domain.User) (string, error) {
	accessTokenClaims := &models.AuthCustomClaims{
		Id:    details.ID,
		Email: details.Email,
		Role:  "user",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 90).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	accessTokenString, err := accessToken.SignedString([]byte("usersecret")) //take this from runtime in future avoid hardcoding
	if err != nil {
		return "", err
	}

	return accessTokenString, nil
}
