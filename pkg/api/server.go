package httpserver

import (
	adminHandler "Ecom/pkg/api/handler/admin"
	userHandler "Ecom/pkg/api/handler/user"

	"Ecom/pkg/api/routes"
	"log"

	"github.com/gin-gonic/gin"
)

type ServerHttp struct {
	engine *gin.Engine
}

func NewServerHttp(adminHandler *adminHandler.AdminHandler, userHandler *userHandler.UserHandler,
	productHandler *adminHandler.ProductHandler, orderHandler *userHandler.OrderHandler,
	uproductHandler *userHandler.UProductHandler) *ServerHttp {
	engine := gin.New()
	engine.Use(gin.Logger())

	routes.AdminRoutes(engine.Group("/admin"), adminHandler, productHandler)
	routes.UserRoutes(engine.Group("/user"), userHandler, orderHandler, uproductHandler)
	return &ServerHttp{engine: engine}
}

func (server *ServerHttp) Start() {
	err := server.engine.Run(":8080")
	if err != nil {
		log.Fatal("unable to Start Server")
	}
}
