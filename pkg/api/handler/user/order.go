package userHandler

import (
	"Ecom/pkg/usecase/user/interfaces"
	"Ecom/pkg/utils/models"
	response "Ecom/pkg/utils/res"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	usecase interfaces.OrderUsecase
}

func NewOrderHandler(usecase interfaces.OrderUsecase) *OrderHandler {
	return &OrderHandler{
		usecase: usecase,
	}
}

func (ord *OrderHandler) PlaceOrder(ctx *gin.Context) {
	PId := ctx.Param("id")
	ProId, err := strconv.Atoi(PId)
	if err != nil {
		ctx.JSON(400, gin.H{
			"Error": "Error in string conversion",
		})
		return
	}
	Qu := ctx.Query("quantity")
	Quantity, err := strconv.Atoi(Qu)
	if err != nil {
		ctx.JSON(400, gin.H{
			"Error": "Error in string conversion",
		})
		return
	}
	Body := models.PlaceOrderData{
		ProductId: uint(ProId),
		Quantity:  Quantity,
	}

	if err := ctx.Bind(&Body); err != nil {
		res := response.ErrResponse{Response: "Binding Error ", Error: err.Error(), StatusCode: 400}
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	Uid := ctx.GetInt("id")

	err = ord.usecase.PlaceOrder(Body, Uid)
	if err != nil {
		res := response.ErrResponse{Response: "Error While Placing Order", Error: err.Error(), StatusCode: 500}
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	successRes := response.SuccResponse{Response: "Order Placed Sucessfully", StatusCode: 201}
	ctx.JSON(http.StatusCreated, successRes)
}

func (ord *OrderHandler) ViewOrders(ctx *gin.Context) {
	count, err1 := strconv.Atoi(ctx.Query("count"))
	PageN, err2 := strconv.Atoi(ctx.Query("page"))
	err3 := errors.Join(err1, err2)
	if err3 != nil {
		ctx.JSON(400, gin.H{
			"Error": "Error in string conversion",
		})
		return
	}
	Uid := ctx.GetInt("id")
	Order, err := ord.usecase.ViewOrders(Uid, count, PageN)
	if err != nil {
		res := response.ErrResponse{Response: "Error While Viewing Order", Error: err.Error(), StatusCode: 500}
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}
	if Order == nil {
		res := response.ErrResponse{Response: "!!!Page Not Found!!!", Error: "Orders Not found ", StatusCode: 200}
		ctx.JSON(http.StatusOK, res)
		return
	}

	successRes := response.SuccResponse{Response: Order, StatusCode: 200}
	ctx.JSON(http.StatusCreated, successRes)
}
