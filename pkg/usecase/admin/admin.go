package adminUsecase

import (
	"Ecom/pkg/domain"
	helper "Ecom/pkg/helper/interfaces"
	"Ecom/pkg/repository/admin/interfaces"
	services "Ecom/pkg/usecase/admin/interfaces"
	"Ecom/pkg/utils/models"
	"context"
	"errors"
	"time"
)

type AdminUsecase struct {
	repo   interfaces.AdminRepo
	helper helper.Helper
}

func NewAdminUsecase(repo interfaces.AdminRepo, helper helper.Helper) services.AdminUsecase {
	return &AdminUsecase{
		repo:   repo,
		helper: helper,
	}
}

func (adm *AdminUsecase) AdminSignup(body domain.Admin) error {
	exist, err := adm.repo.FindAdminByEmail(body.Email)
	if err != nil {
		return err
	}

	if exist != 0 {
		return errors.New("admin email already exists")
	}

	hash, err := adm.helper.CreateHashPassword(body.Password)

	if err != nil {
		return err
	}
	// assign hashed password
	body.Password = hash
	// give the previlage of admin
	body.Previlege = "admin"

	// create new admin
	if err := adm.repo.AdminSignup(body); err != nil {
		return err
	}

	return nil
}

func (adm *AdminUsecase) AdminLogin(ctx context.Context, Body models.AdminLogin) (string, error) {
	if ctx.Err() != nil {
		return "", errors.New("request time out")
	}

	//check the email already exist or not

	admDetails, err := adm.repo.GetAdminDetailsByEmail(ctx, Body.Email)
	if err != nil {
		return "", err
	}

	//  compare the password
	err = adm.helper.CompareHashAndPassword(admDetails.Password, Body.Password)
	if err != nil {
		return "", err
	}

	var AdminResponse models.AdminDetailsResponse

	AdminResponse.ID = int(admDetails.ID)
	AdminResponse.Email = admDetails.Email
	AdminResponse.Name = admDetails.Name

	token, err := adm.helper.GenerateTokenAdmin(AdminResponse)

	if err != nil {
		return "", err
	}

	return token, nil
}

func (adm *AdminUsecase) ViewAllOrders(count, page int) ([]models.AdminViewOrder, error) {
	AllOrders, err := adm.repo.ViewAllOrders(count, page)
	if err != nil {
		return []models.AdminViewOrder{}, err
	}

	return AllOrders, nil
}

func (adm *AdminUsecase) FetchOrderDates(fromtime, totime time.Time) ([]models.SalesReport, error) {
	SalesReport, err := adm.repo.FetchOrderDates(fromtime, totime)
	if err != nil {
		return []models.SalesReport{}, err
	}
	return SalesReport, nil
}
