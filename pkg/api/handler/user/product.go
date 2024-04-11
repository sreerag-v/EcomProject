package userHandler

import (
	"Ecom/pkg/usecase/user/interfaces"
	response "Ecom/pkg/utils/res"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UProductHandler struct {
	usecase interfaces.UProductUsecase
}

func NewUProductHandler(usecase interfaces.UProductUsecase) *UProductHandler {
	return &UProductHandler{
		usecase: usecase,
	}
}

func (Upro *UProductHandler) ViewProduct(ctx *gin.Context) {
	count, err1 := strconv.Atoi(ctx.Query("count"))
	PageN, err2 := strconv.Atoi(ctx.Query("page"))
	Sort := ctx.Query("sort")
	err3 := errors.Join(err1, err2)
	if err3 != nil {
		ctx.JSON(400, gin.H{
			"Error": "Error in string conversion",
		})
		return
	}
	product, err := Upro.usecase.ViewProduct(count, PageN, Sort)
	if err != nil {
		res := response.ErrResponse{Response: "Error While Fetching Product", Error: err.Error(), StatusCode: 500}
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	if product == nil {
		res := response.ErrResponse{Response: "!!!Page Not Found!!!", Error: "Products Not found ", StatusCode: 200}
		ctx.JSON(http.StatusOK, res)
		return
	}

	successRes := response.SuccResponse{Response: product, StatusCode: 200}
	ctx.JSON(http.StatusCreated, successRes)
}

func (Upro *UProductHandler) ProductByName(ctx *gin.Context) {
	Proname := ctx.Query("Product_Name")
	Product, err := Upro.usecase.ProductByName(Proname)
	if err != nil {
		res := response.ErrResponse{Response: "Error While Fetching Product", Error: err.Error(), StatusCode: 500}
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}
	successRes := response.SuccResponse{Response: Product, StatusCode: 200}
	ctx.JSON(http.StatusCreated, successRes)
}
