//go:build wireinject
// +build wireinject

package di

import (
	httpserver "Ecom/pkg/api"
	adminHandler "Ecom/pkg/api/handler/admin"
	userHandler "Ecom/pkg/api/handler/user"

	"Ecom/pkg/config"
	"Ecom/pkg/db"
	"Ecom/pkg/helper"
	adminRepo "Ecom/pkg/repository/admin"
	userRepo "Ecom/pkg/repository/user"
	adminUsecase "Ecom/pkg/usecase/admin"
	userUsecase "Ecom/pkg/usecase/user"

	"github.com/google/wire"
)

func InitializeAPI(cfg config.Config) (*httpserver.ServerHttp, error) {
	wire.Build(
		db.ConnectDatabase,
		adminRepo.NewAdminRepo,
		adminUsecase.NewAdminUsecase,
		adminHandler.NewAdminHandler,

		adminRepo.NewProductRepo,
		adminUsecase.NewProductUsecase,
		adminHandler.NewProductHandler,

		userRepo.NewUserRepo,
		userUsecase.NewUserUsecase,
		userHandler.NewUserHandler,

		userRepo.NewUProductRepo,
		userUsecase.NewUProductUsecase,
		userHandler.NewUProductHandler,

		userRepo.NewOrderRepo,
		userUsecase.NewOrderUsecase,
		userHandler.NewOrderHandler,

		httpserver.NewServerHttp,
		helper.NewHelper,
	)

	return &httpserver.ServerHttp{}, nil
}
