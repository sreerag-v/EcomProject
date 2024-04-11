package routes

import (
	adminHandler "Ecom/pkg/api/handler/admin"
	"Ecom/pkg/api/middleware"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(engine *gin.RouterGroup,
	adminHanlder *adminHandler.AdminHandler, productHandler *adminHandler.ProductHandler) {

	engine.POST("/signup", adminHanlder.AdminSignUp)
	engine.POST("/login", adminHanlder.AdminLogin)

	engine.Use(middleware.AdminAuthMiddleware)
	{
		{
			product := engine.Group("/product")

			product.POST("/create", productHandler.AddProduct)
			product.GET("/view", productHandler.ViewProduct)
			product.GET("/name", productHandler.ProductByName)
			product.PATCH("/update/:id", productHandler.UpdateProduct)
			product.DELETE("/delete/:id", productHandler.DeleteProduct)

			order := product.Group("/order")
			{
				order.GET("/view_all", adminHanlder.ViewAllOrders)
			}
			sales := order.Group("/sales")
			{
				sales.GET("/report", adminHanlder.SalesReport)
			}
		}

	}

}
