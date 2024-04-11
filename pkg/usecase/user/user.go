package userUsecase

import (
	"Ecom/pkg/domain"
	helper "Ecom/pkg/helper/interfaces"
	"Ecom/pkg/repository/user/interfaces"
	services "Ecom/pkg/usecase/user/interfaces"
	"Ecom/pkg/utils/models"
	"errors"
)

type UserUsecase struct {
	repo   interfaces.UserRepo
	helper helper.Helper
}

func NewUserUsecase(repo interfaces.UserRepo, helper helper.Helper) services.UserUsecase {
	return &UserUsecase{
		repo:   repo,
		helper: helper,
	}
}

func (usr *UserUsecase) SignUp(Body domain.User) error {
	// check the number exist or not
	exists, err := usr.repo.CheckPhoneNumberExist(Body.Phone)
	if err != nil {
		return err
	}

	if exists {
		return errors.New("phone number already exists")
	}
	hashed, err := usr.helper.CreateHashPassword(Body.Password)
	if err != nil {
		return err
	}

	Body.Password = hashed

	if err := usr.repo.SignUp(Body); err != nil {
		return err
	}

	return nil
}

func (usr *UserUsecase) Login(Body models.UserLogin) (string, error) {
	exists, err := usr.repo.CheckUsername(Body.Username)
	if err != nil {
		return "", err
	}

	if !exists {
		return "", errors.New("check username again")
	}

	details, err := usr.repo.GetUserDetails(Body.Username)
	if err != nil {
		return "", err
	}

	err = usr.helper.CompareHashAndPassword(details.Password, Body.Password)
	if err != nil {
		return "", errors.New("password mismatch")
	}

	token, err := usr.helper.GenerateTokenUser(details)
	if err != nil {
		return token, err
	}

	return token, nil
}
