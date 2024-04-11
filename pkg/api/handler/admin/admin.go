package adminHandler

import (
	"Ecom/pkg/domain"
	"Ecom/pkg/usecase/admin/interfaces"
	"Ecom/pkg/utils/models"
	response "Ecom/pkg/utils/res"
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type AdminHandler struct {
	usecase interfaces.AdminUsecase
}

func NewAdminHandler(usecase interfaces.AdminUsecase) *AdminHandler {
	return &AdminHandler{
		usecase: usecase,
	}
}

func (adm *AdminHandler) AdminSignUp(ctx *gin.Context) {
	var Body domain.Admin
	if err := ctx.Bind(&Body); err != nil {
		res := response.ErrResponse{Response: "Binding Error", Error: err.Error(), StatusCode: 400}

		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	err := adm.usecase.AdminSignup(Body)
	if err != nil {
		res := response.ErrResponse{Response: "Error While Creaing Admin", Error: err.Error(), StatusCode: 500}
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	successRes := response.SuccResponse{Response: "Successfully created new admin", StatusCode: 200}
	ctx.JSON(http.StatusCreated, successRes)

}

func (adm *AdminHandler) AdminLogin(ctx *gin.Context) {
	var Body models.AdminLogin

	if err := ctx.Bind(&Body); err != nil {
		res := response.ErrResponse{Response: "Binding Error", Error: err.Error(), StatusCode: 400}
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	validate := validator.New()
	if err := validate.Struct(Body); err != nil {
		res := response.ErrResponse{Response: "Struct Error", Error: err.Error(), StatusCode: 400}
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	ctxtime, cance := context.WithTimeout(ctx.Request.Context(), 30*time.Second)
	defer cance()

	Token, err := adm.usecase.AdminLogin(ctxtime, Body)
	if err != nil {
		errRes := response.ErrResponse{Response: "Error In AdminLogin", Error: err.Error(), StatusCode: 500}
		ctx.JSON(http.StatusInternalServerError, errRes)
		return
	}

	successRes := response.LoginRes{TokenString: Token, StatusCode: 200}
	ctx.JSON(http.StatusOK, successRes)
}

func (adm *AdminHandler) ViewAllOrders(ctx *gin.Context) {
	count, err1 := strconv.Atoi(ctx.Query("count"))
	PageN, err2 := strconv.Atoi(ctx.Query("page"))
	err3 := errors.Join(err1, err2)
	if err3 != nil {
		ctx.JSON(400, gin.H{
			"Error": "Error in string conversion",
		})
		return
	}
	AllOrders, err := adm.usecase.ViewAllOrders(count, PageN)
	if err != nil {
		res := response.ErrResponse{Response: "Error While Fetching Orders", Error: err.Error(), StatusCode: 500}
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}
	if AllOrders == nil {
		res := response.ErrResponse{Response: "!!!Page Not Found!!!", Error: "Orders Not found ", StatusCode: 200}
		ctx.JSON(http.StatusOK, res)
		return
	}

	successRes := response.SuccResponse{Response: AllOrders, StatusCode: 200}
	ctx.JSON(http.StatusCreated, successRes)
}

func (adm *AdminHandler) SalesReport(ctx *gin.Context) {
	// want to fetch the dates from the url
	startDate := ctx.Query("startDate")
	endDate := ctx.Query("endDate")

	// converting the date string to time.time
	fromtime, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		res := response.ErrResponse{Response: "Invalid Start Date", Error: err.Error(), StatusCode: 400}
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	totime, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		res := response.ErrResponse{Response: "Invalid End Date", Error: err.Error(), StatusCode: 400}
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	// Fetch the dates from Order Table
	SalesReport, err := adm.usecase.FetchOrderDates(fromtime, totime)
	if err != nil {
		res := response.ErrResponse{Response: "Error While Fetching Dates From Orders", Error: err.Error(), StatusCode: 500}
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	// Creating Excel file
	ex := excelize.NewFile()

	// Creating a new sheet
	SheetName := "SalesReport"
	index := ex.NewSheet(SheetName)

	// Setting values of headers
	ex.SetCellValue(SheetName, "A1", "Order Date")
	ex.SetCellValue(SheetName, "B1", "Order ID")
	ex.SetCellValue(SheetName, "C1", "Product name")
	ex.SetCellValue(SheetName, "D1", "Price")
	ex.SetCellValue(SheetName, "E1", "User Name")
	ex.SetCellValue(SheetName, "F1", "Order Status")
	ex.SetCellValue(SheetName, "G1", "Quantity")

	for i, order := range SalesReport {
		row := i + 2
		ex.SetCellValue(SheetName, fmt.Sprintf("A%d", row), order.CreatedAt.Format("01/02/2006"))
		ex.SetCellValue(SheetName, fmt.Sprintf("B%d", row), order.OrderID)
		ex.SetCellValue(SheetName, fmt.Sprintf("C%d", row), order.ProductName)
		ex.SetCellValue(SheetName, fmt.Sprintf("D%d", row), order.Price)
		ex.SetCellValue(SheetName, fmt.Sprintf("F%d", row), order.UserName)
		ex.SetCellValue(SheetName, fmt.Sprintf("G%d", row), order.OrderStatus)
		ex.SetCellValue(SheetName, fmt.Sprintf("H%d", row), order.Quantity)
	}

	// Setting active sheet of the workbook
	ex.SetActiveSheet(index)

	if err := ex.SaveAs("./public/SalesReport.xlsx"); err != nil {
		fmt.Println(err)
	}

	successRes := response.SuccResponse{Response: "Successfully created Sales Report", StatusCode: 200}
	ctx.JSON(http.StatusCreated, successRes)
}
