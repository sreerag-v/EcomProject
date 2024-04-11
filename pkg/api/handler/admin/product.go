package adminHandler

import (
	"Ecom/pkg/usecase/admin/interfaces"
	"Ecom/pkg/utils/models"
	response "Ecom/pkg/utils/res"
	"errors"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ProductHandler struct {
	usecase interfaces.ProductUsecase
}

func NewProductHandler(usecase interfaces.ProductUsecase) *ProductHandler {
	return &ProductHandler{
		usecase: usecase,
	}
}

func (pro *ProductHandler) AddProduct(ctx *gin.Context) {
	prodname := ctx.Request.FormValue("product_name")
	price := ctx.Request.FormValue("price")
	Price, _ := strconv.Atoi(price)
	color := ctx.Request.FormValue("color")
	stock := ctx.Request.FormValue("stock")
	Stock, _ := strconv.Atoi(stock)
	imagepath, _ := ctx.FormFile("image")
	extension := filepath.Ext(imagepath.Filename)
	image := uuid.New().String() + extension
	ctx.SaveUploadedFile(imagepath, "./public/images"+image)

	body := models.ProductAdd{
		ProductName: prodname,
		Price:       uint(Price),
		Color:       color,
		Stock:       uint(Stock),
		Image:       image,
	}

	err := pro.usecase.AddProduct(body)
	if err != nil {
		res := response.ErrResponse{Response: "Error While Creaing Product", Error: err.Error(), StatusCode: 500}
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	successRes := response.SuccResponse{Response: "Successfully Created New Product", StatusCode: 201}
	ctx.JSON(http.StatusCreated, successRes)
}

func (pro *ProductHandler) ViewProduct(ctx *gin.Context) {
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
	product, err := pro.usecase.ViewProduct(count, PageN, Sort)
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

func (pro *ProductHandler) ProductByName(ctx *gin.Context) {
	Proname := ctx.Query("Product_Name")
	Product, err := pro.usecase.ProductByName(Proname)
	if err != nil {
		res := response.ErrResponse{Response: "Error While Fetching Product", Error: err.Error(), StatusCode: 500}
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}
	successRes := response.SuccResponse{Response: Product, StatusCode: 200}
	ctx.JSON(http.StatusCreated, successRes)
}

func (pro *ProductHandler) UpdateProduct(ctx *gin.Context) {
	proId := ctx.Param("id")

	var EditProduct models.UpdateProduct
	if err := ctx.Bind(&EditProduct); err != nil {
		res := response.ErrResponse{Response: "Binding Error", Error: err.Error(), StatusCode: 400}
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	Intid, err := strconv.Atoi(proId)
	if err != nil {
		res := response.ErrResponse{Response: "String Converstion Error", Error: err.Error(), StatusCode: 400}
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	err = pro.usecase.EditProduct(Intid, EditProduct)
	if err != nil {
		res := response.ErrResponse{Response: "Error While Updating Product", Error: err.Error(), StatusCode: 500}
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	successRes := response.SuccResponse{Response: "Product Successfully Updated", StatusCode: 200}
	ctx.JSON(http.StatusCreated, successRes)

}

func (pro *ProductHandler) DeleteProduct(ctx *gin.Context) {
	proId := ctx.Param("id")

	Intid, err := strconv.Atoi(proId)
	if err != nil {
		res := response.ErrResponse{Response: "String Converstion Error", Error: err.Error(), StatusCode: 400}
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	err = pro.usecase.DeleteProduct(Intid)
	if err != nil {
		res := response.ErrResponse{Response: "Error While Deleting Product", Error: err.Error(), StatusCode: 500}
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	successRes := response.SuccResponse{Response: "Product Successfully Deleted", StatusCode: 200}
	ctx.JSON(http.StatusCreated, successRes)
}
