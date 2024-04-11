package interfaces

import "Ecom/pkg/domain"

type UserRepo interface {
	CheckPhoneNumberExist(string) (bool, error)
	SignUp(domain.User) error
	CheckUsername(string) (bool, error)
	GetUserDetails(string) (domain.User, error)
}
