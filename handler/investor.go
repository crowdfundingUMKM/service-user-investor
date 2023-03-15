package handler

import (
	"net/http"
	"service-user-investor/auth"
	"service-user-investor/helper"
	"service-user-investor/investor"

	"github.com/gin-gonic/gin"
)

type userInvestorHandler struct {
	userService investor.Service
	authService auth.Service
}

func NewUserHandler(userService investor.Service, authService auth.Service) *userInvestorHandler {
	return &userInvestorHandler{userService, authService}
}

func (h *userInvestorHandler) RegisterUser(c *gin.Context) {
	// tangkap input dari user
	// map input dari user ke struct RegisterUserInput
	// struct di atas kita passing sebagai parameter service

	var input investor.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	// generate token
	token, err := h.authService.GenerateToken(newUser.ID)
	if err != nil {
		if err != nil {
			response := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", nil)
			c.JSON(http.StatusBadRequest, response)
			return
		}
	}

	formatter := investor.FormatterUser(newUser, token)

	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}
