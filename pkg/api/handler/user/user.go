package userHandler

import (
	"Ecom/pkg/domain"
	"Ecom/pkg/usecase/user/interfaces"
	"Ecom/pkg/utils/models"
	response "Ecom/pkg/utils/res"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type UserHandler struct {
	usecase interfaces.UserUsecase
}

func NewUserHandler(usecase interfaces.UserUsecase) *UserHandler {
	return &UserHandler{
		usecase: usecase,
	}
}

func (usr *UserHandler) UserSignup(c *gin.Context) {
	var Body domain.User

	if err := c.Bind(&Body); err != nil {
		res := response.ErrResponse{Response: "Binding Error ", Error: err.Error(), StatusCode: 400}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	validate := validator.New()
	if err := validate.Struct(Body); err != nil {
		res := response.ErrResponse{Response: "Struct Validation Error", Error: err.Error(), StatusCode: 400}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if err := usr.usecase.SignUp(Body); err != nil {
		res := response.ErrResponse{Response: "Error From Creating User", Error: err.Error(), StatusCode: 400}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := response.SuccResponse{Response: "Signup completed successfully", StatusCode: 200}
	c.JSON(http.StatusCreated, res)
}

func (usr *UserHandler) UserLogin(c *gin.Context) {
	var Body models.UserLogin

	if err := c.Bind(&Body); err != nil {
		res := response.ErrResponse{Response: "Binding Error", Error: err.Error(), StatusCode: 400}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	token, err := usr.usecase.Login(Body)
	if err != nil {
		res := response.ErrResponse{Response: "Error From UserLogin", Error: err.Error(), StatusCode: 500}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	//return result
	res := response.LoginRes{TokenString: token, StatusCode: 201}
	c.JSON(http.StatusCreated, res)
}
