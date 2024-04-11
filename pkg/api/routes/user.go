package routes

import (
	userHandler "Ecom/pkg/api/handler/user"
	"Ecom/pkg/api/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(engine *gin.RouterGroup,
	userHandler *userHandler.UserHandler,
	orderHandler *userHandler.OrderHandler, uproductHandler *userHandler.UProductHandler) {

	engine.POST("/signup", userHandler.UserSignup)
	engine.POST("/login", userHandler.UserLogin)

	engine.Use(middleware.UserAuthMiddleware)
	{
		product := engine.Group("/product")
		{
			product.GET("/view", uproductHandler.ViewProduct)
			product.GET("/name", uproductHandler.ProductByName)

			order := product.Group("/order")
			{
				order.POST("/place/:id", orderHandler.PlaceOrder)
				order.GET("/view", orderHandler.ViewOrders)
			}
		}
	}
}
