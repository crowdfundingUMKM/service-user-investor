package handler

import (
	"io/ioutil"
	"net/http"
	"os"
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

func (h *userInvestorHandler) GetLogtoAdmin(c *gin.Context) {
	// get id from params
	// get user from service
	// format user
	// response
	// id := c.Param("id")
	// user, err := h.userService.GetUserByID(id)
	// if err != nil {
	// 	response := helper.APIResponse("Failed to get user", http.StatusBadRequest, "error", nil)
	// 	c.JSON(http.StatusBadRequest, response)
	// 	return
	// }

	// formatter := investor.FormatterUser(user, "")
	// response := helper.APIResponse("Success to get user", http.StatusOK, "success", formatter)
	// c.JSON(http.StatusOK, response)
	id := os.Getenv("ADMIN_ID")
	if c.Param("id") == id {
		content, err := ioutil.ReadFile("./log/gin.log")
		if err != nil {
			c.String(500, "Failed to read log file admin: %v", err)
			return
		}
		c.String(http.StatusOK, string(content))
	} else {
		c.String(http.StatusNotFound, "Not found")
	}
}

// for admin get env
func (h *userInvestorHandler) ServiceHealth(c *gin.Context) {
	db_user := os.Getenv("DB_USER")
	db_pass := os.Getenv("DB_PASS")
	db_name := os.Getenv("DB_NAME")
	db_port := os.Getenv("DB_PORT")
	instance_host := os.Getenv("INSTANCE_HOST")

	data := map[string]interface{}{
		"db_user":          db_user,
		"db_pass":          db_pass,
		"db_name":          db_name,
		"db_port":          db_port,
		"db_instance_host": instance_host,
	}
	err := c.Errors
	if err != nil {
		response := helper.APIResponse("Service investor is not running", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response := helper.APIResponse("Service investor is running", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
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

func (h *userInvestorHandler) Login(c *gin.Context) {

	var input investor.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedinUser, err := h.userService.Login(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	token, err := h.authService.GenerateToken(loggedinUser.ID)

	if err != nil {
		if err != nil {
			response := helper.APIResponse("Login failed", http.StatusBadRequest, "error", nil)
			c.JSON(http.StatusBadRequest, response)
			return
		}
	}

	formatter := investor.FormatterUser(loggedinUser, token)

	response := helper.APIResponse("Succesfuly loggedin", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}
