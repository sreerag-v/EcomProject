package interfaces

import (
	"Ecom/pkg/domain"
	"Ecom/pkg/utils/models"
)

type UserUsecase interface {
	SignUp(Body domain.User) error
	Login(Body models.UserLogin) (string, error)
}
