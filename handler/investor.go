package handler

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	api_admin "service-user-investor/api/admin"
	"service-user-investor/auth"
	"service-user-investor/core"
	"service-user-investor/database"
	"service-user-investor/helper"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type userInvestorHandler struct {
	userService core.Service
	authService auth.Service
}

func NewUserHandler(userService core.Service, authService auth.Service) *userInvestorHandler {
	return &userInvestorHandler{userService, authService}
}

func (h *userInvestorHandler) GetLogtoAdmin(c *gin.Context) {
	// check id admin
	adminID := c.Param("admin_id")
	adminInput := api_admin.AdminIdInput{UnixID: adminID}
	getAdminValueId, err := api_admin.GetAdminId(adminInput)

	if err != nil {
		response := helper.APIResponse(err.Error(), http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentAdmin := c.MustGet("currentUserAdmin").(api_admin.AdminId)

	if c.Param("admin_id") == getAdminValueId && currentAdmin.UnixAdmin == getAdminValueId {
		content, err := ioutil.ReadFile("./tmp/gin.log")
		if err != nil {
			response := helper.APIResponse("Failed to get log", http.StatusBadRequest, "error", nil)
			c.JSON(http.StatusBadRequest, response)
			return
		}
		// download with browser
		if c.Query("download") == "true" {
			c.Header("Content-Disposition", "attachment; filename=gin.log")
			c.Data(http.StatusOK, "application/octet-stream", content)
			return
		}
		// show in browser
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

	adminID := c.Param("admin_id")
	adminInput := api_admin.AdminIdInput{UnixID: adminID}
	getAdminValueId, err := api_admin.GetAdminId(adminInput)

	if err != nil {
		response := helper.APIResponse(err.Error(), http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// middleware api admin
	currentAdmin := c.MustGet("currentUserAdmin").(api_admin.AdminId)

	if c.Param("admin_id") != getAdminValueId && currentAdmin.UnixAdmin != getAdminValueId {
		response := helper.APIResponse("Your not Admin, cannot Access", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	envVars := []string{
		"ADMIN_ID",
		"DB_USER",
		"DB_PASS",
		"DB_NAME",
		"DB_PORT",
		"INSTANCE_HOST",
		"SERVICE_HOST",
		"SERVICE_PORT",
		"JWT_SECRET",
		"STATUS_ACCOUNT",
	}

	data := make(map[string]interface{})
	for _, key := range envVars {
		data[key] = os.Getenv(key)
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

func (h *userInvestorHandler) ServiceStart(c *gin.Context) {
	// check env open or not
	serviceStatus := "Service is active"

	errService := c.Errors
	if errService != nil {
		response := helper.APIResponse("Service investor is not running", http.StatusInternalServerError, "error", serviceStatus)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response := helper.APIResponse("Service investor is running", http.StatusOK, "success", serviceStatus)
	c.JSON(http.StatusOK, response)
}

func (h *userInvestorHandler) ServiceCheckDB(c *gin.Context) {
	// check env open or not
	// Check database status
	db := database.NewConnectionDB()
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer sqlDB.Close()

	errDB := sqlDB.Ping()
	if errDB != nil {
		response := helper.APIResponse("Database is not running", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	currentUser := c.MustGet("currentUser").(core.User)

	data := gin.H{"user": "Wellcome to Service: " + currentUser.Name, "status": "Database is running"}

	response := helper.APIResponse("Service is running", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)

}

// deactive account
func (h *userInvestorHandler) DeactiveUser(c *gin.Context) {
	var input core.DeactiveUserInput
	// check input from user
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("User Not Found", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// cheack id from get param and fetch data from service admin to check id admin and status account admin
	// var adminInput investor.AdminIdInput
	adminID := c.Param("admin_id")
	adminInput := api_admin.AdminIdInput{UnixID: adminID}
	getAdminValueId, err := api_admin.GetAdminId(adminInput)

	if err != nil {
		response := helper.APIResponse(err.Error(), http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// middleware api admin
	currentAdmin := c.MustGet("currentUserAdmin").(api_admin.AdminId)

	// check id admin
	// adminId := getAdminValueId
	if c.Param("admin_id") == getAdminValueId && currentAdmin.UnixAdmin == getAdminValueId {
		// get id user

		// deactive user
		deactive, err := h.userService.DeactivateAccountUser(input, currentAdmin.UnixAdmin)

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
	var input core.DeactiveUserInput
	// check input from user
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("User Not Found", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// cheack id from get param and fetch data from service admin to check id admin and status account admin
	// var adminInput investor.AdminIdInput
	adminID := c.Param("admin_id")
	adminInput := api_admin.AdminIdInput{UnixID: adminID}
	getAdminValueId, err := api_admin.GetAdminId(adminInput)

	if err != nil {
		response := helper.APIResponse(err.Error(), http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// middleware api admin
	currentAdmin := c.MustGet("currentUserAdmin").(api_admin.AdminId)

	// check id admin
	// adminId := getAdminValueId
	if c.Param("admin_id") == getAdminValueId && currentAdmin.UnixAdmin == getAdminValueId {
		// get id user

		// deactive user
		active, err := h.userService.ActivateAccountUser(input, currentAdmin.UnixAdmin)

		data := gin.H{
			"success_active": active,
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

// delete user by admin
func (h *userInvestorHandler) DeleteUserByAdmin(c *gin.Context) {
	var input core.DeleteUserInput
	// check input from user
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("User Not Found", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// cheack id from get param and fetch data from service admin to check id admin and status account admin
	adminID := c.Param("admin_id")
	adminInput := api_admin.AdminIdInput{UnixID: adminID}
	getAdminValueId, err := api_admin.GetAdminId(adminInput)

	if err != nil {
		response := helper.APIResponse(err.Error(), http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// middleware api admin
	currentAdmin := c.MustGet("currentUserAdmin").(api_admin.AdminId)

	if c.Param("admin_id") == getAdminValueId && currentAdmin.UnixAdmin == getAdminValueId {
		// get id user

		// deactive user
		delete, err := h.userService.DeleteAccountUser(input.UnixID)

		data := gin.H{
			"success_delete": delete,
		}

		if err != nil {
			response := helper.APIResponse("Failed to delete user", http.StatusBadRequest, "error", "Not Found User investor")
			c.JSON(http.StatusBadRequest, response)
			return
		}
		response := helper.APIResponse("User has been delete", http.StatusOK, "success", data)
		c.JSON(http.StatusOK, response)
	} else {
		response := helper.APIResponse("Your not Admin, cannot Access", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusNotFound, response)
		return
	}
}

// get all user by admin
func (h *userInvestorHandler) GetAllUserData(c *gin.Context) {
	// cheack id from get param and fetch data from service admin to check id admin and status account admin
	// var adminInput investor.AdminIdInput
	adminID := c.Param("admin_id")
	adminInput := api_admin.AdminIdInput{UnixID: adminID}
	getAdminValueId, err := api_admin.GetAdminId(adminInput)

	currentAdmin := c.MustGet("currentUserAdmin").(api_admin.AdminId)

	if err != nil {
		response := helper.APIResponse(err.Error(), http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if c.Param("admin_id") == getAdminValueId && currentAdmin.UnixAdmin == getAdminValueId {
		users, err := h.userService.GetAllUsers()
		if err != nil {
			response := helper.APIResponse("Failed to get user", http.StatusBadRequest, "error", nil)
			c.JSON(http.StatusBadRequest, response)
			return
		}
		response := helper.APIResponse("List of user", http.StatusOK, "success", users)
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

	var input core.RegisterUserInput

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

	// save token ke db
	// end save token ke db
	if err != nil {
		if err != nil {
			response := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", nil)
			c.JSON(http.StatusBadRequest, response)
			return
		}
	}

	formatter := core.FormatterUser(newUser, token)

	if formatter.StatusAccount == "active" {
		_, err = h.userService.SaveToken(newUser.UnixID, token)

		if err != nil {
			response := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", nil)
			c.JSON(http.StatusBadRequest, response)
			return
		}

		response := helper.APIResponse("Account has been registered and active", http.StatusOK, "success", formatter)
		c.JSON(http.StatusOK, response)
		return
	}

	data := gin.H{
		"status": "Account has been registered, but you must wait admin to active your account",
	}

	response := helper.APIResponse("Account has been registered but you must wait admin or review to active your account", http.StatusOK, "success", data)

	c.JSON(http.StatusOK, response)
}

func (h *userInvestorHandler) Login(c *gin.Context) {

	var input core.LoginInput

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
	// generate token
	token, err := h.authService.GenerateToken(loggedinUser.UnixID)

	// save toke to database
	_, err = h.userService.SaveToken(loggedinUser.UnixID, token)

	if err != nil {
		response := helper.APIResponse("Login failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	// end save token to database

	if err != nil {
		if err != nil {
			response := helper.APIResponse("Login failed", http.StatusBadRequest, "error", nil)
			c.JSON(http.StatusBadRequest, response)
			return
		}
	}

	// check role acvtive and not send massage your account deactive
	if loggedinUser.StatusAccount == "deactive" {
		errorMessage := gin.H{"errors": "Your account is deactive by admin"}
		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	formatter := core.FormatterUser(loggedinUser, token)

	response := helper.APIResponse("Succesfuly loggedin", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *userInvestorHandler) CheckEmailAvailability(c *gin.Context) {
	var input core.CheckEmailInput

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
	var input core.CheckPhoneInput

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

// get user by middleware
func (h *userInvestorHandler) GetUser(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(core.User)

	// check f account deactive
	if currentUser.StatusAccount == "deactive" {
		errorMessage := gin.H{"errors": "Your account is deactive by admin"}
		response := helper.APIResponse("Get user failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	// if you logout you can't get user
	if currentUser.Token == "" {
		errorMessage := gin.H{"errors": "Your account is logout"}
		response := helper.APIResponse("Get user failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := core.FormatterUser(currentUser, "")

	response := helper.APIResponse("Successfuly get user by middleware", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *userInvestorHandler) UpdateUser(c *gin.Context) {
	var inputID core.GetUserIdInput

	// check id is valid or not
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Update user failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData core.UpdateUserInput

	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Update user failed, input data failure", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(core.User)

	if currentUser.UnixID != inputID.UnixID {
		response := helper.APIResponse("Update user failed, because you are not auth", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	// if you logout you can't get user
	if currentUser.Token == "" {
		errorMessage := gin.H{"errors": "Your account is logout"}
		response := helper.APIResponse("Get user failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updatedUser, err := h.userService.UpdateUserByUnixID(currentUser.UnixID, inputData)
	if err != nil {
		response := helper.APIResponse("Update user failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := core.FormatterUserDetail(currentUser, updatedUser)

	response := helper.APIResponse("User has been updated", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
	return
}

func (h *userInvestorHandler) UpdatePassword(c *gin.Context) {
	var inputID core.GetUserIdInput

	// check id is valid or not
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Update password failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData core.UpdatePasswordInput

	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Update password failed, input data failure", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(core.User)

	if currentUser.UnixID != inputID.UnixID {
		response := helper.APIResponse("Update password failed, because you are not auth", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	// if you logout you can't get user
	if currentUser.Token == "" {
		errorMessage := gin.H{"errors": "Your account is logout"}
		response := helper.APIResponse("Get user failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updatedUser, err := h.userService.UpdatePasswordByUnixID(currentUser.UnixID, inputData)
	if err != nil {
		response := helper.APIResponse("Update password failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := core.FormatterUserDetail(currentUser, updatedUser)

	response := helper.APIResponse("Password has been updated", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
	return
}

// Logout user
func (h *userInvestorHandler) LogoutUser(c *gin.Context) {
	// get data from middleware
	currentUser := c.MustGet("currentUser").(core.User)

	// check if token is empty
	if currentUser.Token == "" {
		response := helper.APIResponse("Logout failed, your logout right now", http.StatusForbidden, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// delete token in database
	_, err := h.userService.DeleteToken(currentUser.UnixID)
	if err != nil {
		response := helper.APIResponse("Logout failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Logout success", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
	return
}
