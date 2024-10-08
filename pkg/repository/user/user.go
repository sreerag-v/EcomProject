package userRepo

import (
	"Ecom/pkg/domain"
	"Ecom/pkg/repository/user/interfaces"

	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) interfaces.UserRepo {
	return &UserRepo{
		DB: db,
	}
}

func (usr *UserRepo) CheckPhoneNumberExist(num string) (bool, error) {
	var count int64
	if err := usr.DB.Table("users").Where("phone = ?", num).Count(&count).Error; err != nil {
		return true, err
	}

	// If count is greater than 0, it means a record with the given name exists
	return count > 0, nil
}

func (usr *UserRepo) SignUp(Body domain.User) error {
	err := usr.DB.Exec("INSERT INTO users(name,email,password,phone)VALUES($1,$2,$3,$4)", Body.Name, Body.Email, Body.Password, Body.Phone).Error
	if err != nil {
		return err
	}

	return nil
}

func (usr *UserRepo) CheckUsername(name string) (bool, error) {
	var count int64
	if err := usr.DB.Table("users").Where("name = ?", name).Count(&count).Error; err != nil {
		return true, err
	}

	// If count is greater than 0, it means a record with the given name exists
	return count > 0, nil
}

func (usr *UserRepo) GetUserDetails(name string) (domain.User, error) {
	var model domain.User
	if err := usr.DB.Table("users").Where("name = ?", name).Scan(&model).Error; err != nil {
		return domain.User{}, err
	}

	return model, nil
}
