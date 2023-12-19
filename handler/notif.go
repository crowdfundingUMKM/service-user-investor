package handler

import (
	"net/http"
	api_admin "service-user-investor/api/admin"
	"service-user-investor/auth"
	"service-user-investor/core"
	"service-user-investor/helper"

	"github.com/gin-gonic/gin"
)

type notifHandler struct {
	userService core.Service
	authService auth.Service
}

func NewNotifHandler(userService core.Service, authService auth.Service) *notifHandler {
	return &notifHandler{userService, authService}
}

// func (h *notifHandler) ReportToAdmin(c *gin.Context) {
// 	currentUser := c.MustGet("currentUser").(core.User)

// 	// check f account deactive
// 	if currentUser.StatusAccount == "deactive" {
// 		errorMessage := gin.H{"errors": "Your account is deactive by admin"}
// 		response := helper.APIResponse("Get user failed", http.StatusUnprocessableEntity, "error", errorMessage)
// 		c.JSON(http.StatusUnprocessableEntity, response)
// 		return
// 	}
// 	// if you logout you can't get user
// 	if currentUser.Token == "" {
// 		errorMessage := gin.H{"errors": "Your account is logout"}
// 		response := helper.APIResponse("Get user failed", http.StatusUnprocessableEntity, "error", errorMessage)
// 		c.JSON(http.StatusUnprocessableEntity, response)
// 		return
// 	}

// 	// file Report

// 	f, uploadedFile, err := c.Request.FormFile("file_report")

// 	if err != nil {
// 		// data := gin.H{"is_uploaded": false}
// 		response := helper.APIResponse("Failed to upload file report", http.StatusBadRequest, "error", err)
// 		c.JSON(http.StatusBadRequest, response)
// 		return
// 	}

// 	// initiate cloud storage os.Getenv("GCS_BUCKET")
// 	bucket := fmt.Sprintf("%s", os.Getenv("GCS_BUCKET"))
// 	subfolder := fmt.Sprintf("%s", os.Getenv("GCS_SUBFOLDER"))
// 	// var err error
// 	ctx := appengine.NewContext(c.Request)

// 	storageClient, err = storage.NewClient(ctx, option.WithCredentialsFile("keys.json"))

// 	if err != nil {
// 		// data := gin.H{"is_uploaded": false}
// 		response := helper.APIResponse("Failed to upload file report to GCP", http.StatusBadRequest, "error", err)

// 		c.JSON(http.StatusBadRequest, response)
// 		return
// 	}
// 	defer f.Close()

// 	objectName := fmt.Sprintf("%s/file-report-%s-%s", subfolder, currentUser.UnixID, uploadedFile.Filename)
// 	sw := storageClient.Bucket(bucket).Object(objectName).NewWriter(ctx)

// 	if _, err := io.Copy(sw, f); err != nil {
// 		// data := gin.H{"is_uploaded": false}
// 		response := helper.APIResponse("Failed to upload file report to GCP", http.StatusBadRequest, "error", err)

// 		c.JSON(http.StatusBadRequest, response)
// 		return
// 	}

// 	if err := sw.Close(); err != nil {
// 		// data := gin.H{"is_uploaded": false}
// 		response := helper.APIResponse("Failed to upload file report to GCP", http.StatusBadRequest, "error", err)

// 		c.JSON(http.StatusBadRequest, response)
// 		return
// 	}

// 	u, err := url.Parse("/" + bucket + "/" + sw.Attrs().Name)
// 	if err != nil {
// 		// data := gin.H{"is_uploaded": false}
// 		response := helper.APIResponse("Failed to upload file report to GCP", http.StatusBadRequest, "error", err)

// 		c.JSON(http.StatusBadRequest, response)
// 		return
// 	}
// 	path := u.String()

// 	if path == "" {
// 		path = ""
// 	}

// 	// crete input
// 	var input core.ReportToAdminInput

// 	err = c.ShouldBindJSON(&input)
// 	if err != nil {
// 		errors := helper.FormatValidationError(err)
// 		errorMessage := gin.H{"errors": errors}
// 		response := helper.APIResponse("Invalid report to admin", http.StatusUnprocessableEntity, "error", errorMessage)
// 		c.JSON(http.StatusUnprocessableEntity, response)
// 		return
// 	}
// 	newNotif, err := h.userService.ReportAdmin(currentUser.UnixID, path, input)
// 	if err != nil {
// 		response := helper.APIResponse("Report to admin failed", http.StatusBadRequest, "error", nil)
// 		c.JSON(http.StatusBadRequest, response)
// 		return
// 	}

// 	formatter := core.FormatterNotify(newNotif)

// 	response := helper.APIResponse("Report to admin success", http.StatusOK, "success", formatter)
// 	c.JSON(http.StatusOK, response)

// }
func (h *notifHandler) ReportToAdmin(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(core.User)

	// check if account is deactivated
	if currentUser.StatusAccount == "deactive" {
		errorMessage := gin.H{"errors": "Your account is deactive by admin"}
		response := helper.APIResponse("Get user failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// if the user is logged out, they can't get user data
	if currentUser.Token == "" {
		errorMessage := gin.H{"errors": "Your account is logout"}
		response := helper.APIResponse("Get user failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// create input
	var input core.ReportToAdminInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Invalid report to admin", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newNotif, err := h.userService.ReportAdmin(currentUser.UnixID, input)
	if err != nil {
		response := helper.APIResponse("Report to admin failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := core.FormatterNotify(newNotif)

	response := helper.APIResponse("Report to admin success", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

// GetNotifToAdmin
func (h *notifHandler) GetNotifToAdmin(c *gin.Context) {
	currentAdmin := c.MustGet("currentUserAdmin").(api_admin.AdminId)

	if currentAdmin.UnixAdmin == "" {
		errorMessage := gin.H{"errors": "Your account admin is logout"}
		response := helper.APIResponse("Get all reports", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	users, err := h.userService.GetAllUsers()
	if err != nil {
		response := helper.APIResponse("Failed All report investor", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("List of report investor", http.StatusOK, "success", users)
	c.JSON(http.StatusOK, response)
	return

}
