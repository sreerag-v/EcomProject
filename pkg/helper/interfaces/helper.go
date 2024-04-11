package interfaces

import (
	"Ecom/pkg/domain"
	"Ecom/pkg/utils/models"
)

type Helper interface {
	CreateHashPassword(string) (string, error)
	CompareHashAndPassword(a string, b string) error
	GenerateTokenAdmin(admin models.AdminDetailsResponse) (string, error)
	GenerateTokenUser(details domain.User) (string, error)
}
