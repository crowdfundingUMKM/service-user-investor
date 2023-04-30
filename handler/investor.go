package handler

import (
	"io/ioutil"
	"net/http"
	"os"
	"service-user-investor/auth"
	"service-user-investor/helper"
	"service-user-investor/investor"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type userInvestorHandler struct {
	userService investor.Service
	authService auth.Service
}

func NewUserHandler(userService investor.Service, authService auth.Service) *userInvestorHandler {
	return &userInvestorHandler{userService, authService}
}

func (h *userInvestorHandler) GetLogtoAdmin(c *gin.Context) {
	// check id admin
	id := os.Getenv("ADMIN_ID")
	if c.Param("id") == id {
		content, err := ioutil.ReadFile("./log/gin.log")
		if err != nil {
			response := helper.APIResponse("Failed to get log", http.StatusBadRequest, "error", nil)
			c.JSON(http.StatusBadRequest, response)
			return
		}
		c.String(http.StatusOK, string(content))
	} else {
		response := helper.APIResponse("Your not Admin, cannot Access", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusNotFound, response)
		return
	}
}

// for admin get env
func (h *userInvestorHandler) ServiceHealth(c *gin.Context) {
	// check env open or not
	errEnv := godotenv.Load()
	if errEnv != nil {
		response := helper.APIResponse("Failed to get env for service investor", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	id := os.Getenv("ADMIN_ID")
	if c.Param("id") != id {
		response := helper.APIResponse("Your not Admin, cannot Access", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusNotFound, response)
		return
	}
	db_user := os.Getenv("DB_USER")
	db_pass := os.Getenv("DB_PASS")
	db_name := os.Getenv("DB_NAME")
	db_port := os.Getenv("DB_PORT")
	instance_host := os.Getenv("INSTANCE_HOST")
	service_port := os.Getenv("PORT")
	jwt_secret := os.Getenv("JWT_SECRET")
	status_account := os.Getenv("STATUS_ACCOUNT")
	admin_id := os.Getenv("ADMIN_ID")

	data := map[string]interface{}{
		"db_user":          db_user,
		"db_pass":          db_pass,
		"db_name":          db_name,
		"db_port":          db_port,
		"db_instance_host": instance_host,
		"service_port":     service_port,
		"jwt_secret":       jwt_secret,
		"status_account":   status_account,
		"admin_id":         admin_id,
	}
	errService := c.Errors
	if errService != nil {
		response := helper.APIResponse("Service investor is not running", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response := helper.APIResponse("Service investor is running", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}

// deactive account
func (h *userInvestorHandler) DeactiveUser(c *gin.Context) {
	var input investor.DeactiveUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("User Not Found", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	// check id admin
	id := os.Getenv("ADMIN_ID")
	if c.Param("id") == id {
		// get id user

		// deactive user
		deactive, err := h.userService.DeactivateAccountUser(input)

		data := gin.H{
			"success_deactive": deactive,
		}

		if err != nil {
			response := helper.APIResponse("Failed to deactive user", http.StatusBadRequest, "error", data)
			c.JSON(http.StatusBadRequest, response)
			return
		}
		response := helper.APIResponse("User has been deactive", http.StatusOK, "success", data)
		c.JSON(http.StatusOK, response)
	} else {
		response := helper.APIResponse("Your not Admin, cannot Access", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusNotFound, response)
		return
	}
}

func (h *userInvestorHandler) ActiveUser(c *gin.Context) {
	var input investor.DeactiveUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("User Not Found", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	// check id admin
	id := os.Getenv("ADMIN_ID")
	if c.Param("id") == id {
		// get id user

		// deactive user
		active, err := h.userService.ActivateAccountUser(input)

		data := gin.H{
			"success_deactive": active,
		}

		if err != nil {
			response := helper.APIResponse("Failed to active user", http.StatusBadRequest, "error", data)
			c.JSON(http.StatusBadRequest, response)
			return
		}
		response := helper.APIResponse("User has been active", http.StatusOK, "success", data)
		c.JSON(http.StatusOK, response)
	} else {
		response := helper.APIResponse("Your not Admin, cannot Access", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusNotFound, response)
		return
	}
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
	token, err := h.authService.GenerateToken(newUser.UnixID)
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
	token, err := h.authService.GenerateToken(loggedinUser.UnixID)

	if err != nil {
		if err != nil {
			response := helper.APIResponse("Login failed", http.StatusBadRequest, "error", nil)
			c.JSON(http.StatusBadRequest, response)
			return
		}
	}

	if loggedinUser.StatusAccount == "Deactive" {
		errorMessage := gin.H{"errors": "Your account is deactive by admin"}
		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	formatter := investor.FormatterUser(loggedinUser, token)

	response := helper.APIResponse("Succesfuly loggedin", http.StatusOK, "success", formatter)

	// check role acvtive and not send massage your account deactive

	c.JSON(http.StatusOK, response)
}

func (h *userInvestorHandler) CheckEmailAvailability(c *gin.Context) {
	var input investor.CheckEmailInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Email checking failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	isEmailAvailable, err := h.userService.IsEmailAvailable(input)
	if err != nil {
		errorMessage := gin.H{"errors": "Server error"}
		response := helper.APIResponse("Email checking failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{
		"is_available": isEmailAvailable,
	}

	metaMessage := "Email has been registered"

	if isEmailAvailable {
		metaMessage = "Email is available"
	}

	response := helper.APIResponse(metaMessage, http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}

func (h *userInvestorHandler) CheckPhoneAvailability(c *gin.Context) {
	var input investor.CheckPhoneInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Phone checking failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	isPhoneAvailable, err := h.userService.IsPhoneAvailable(input)

	if err != nil {
		errorMessage := gin.H{"errors": "Server error"}
		response := helper.APIResponse("Phone checking failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{
		"is_available": isPhoneAvailable,
	}

	metaMessage := "Phone has been registered"

	if isPhoneAvailable {
		metaMessage = "Phone is available"
	}

	response := helper.APIResponse(metaMessage, http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}
