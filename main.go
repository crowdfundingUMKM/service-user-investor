package main

import (
	"fmt"
	"log"
	"os"

	"service-user-investor/auth"
	"service-user-investor/config"
	"service-user-investor/core"
	"service-user-investor/database"
	"service-user-investor/handler"
	"service-user-investor/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// setup log if Production Mode
	// config.InitLog()

	// SETUP REPO
	db := database.NewConnectionDB()
	userInvestorRepository := core.NewRepository(db)

	// SETUP SERVICE
	userInvestorService := core.NewService(userInvestorRepository)
	authService := auth.NewService()

	// setup handler
	userHandler := handler.NewUserHandler(userInvestorService, authService)

	// END SETUP

	// RUN SERVICE
	router := gin.Default()

	// setup cors
	corsConfig := config.InitCors()
	router.Use(cors.New(corsConfig))

	// group api
	api := router.Group("api/v1")

	// admin request -> service user admin
	api.GET("/admin/log_service_toAdmin/:admin_id", middleware.AuthApiAdminMiddleware(authService, userInvestorService), userHandler.GetLogtoAdmin)
	api.GET("/admin/service_status/:admin_id", middleware.AuthApiAdminMiddleware(authService, userInvestorService), userHandler.ServiceHealth)
	api.POST("/admin/deactive_user/:admin_id", middleware.AuthApiAdminMiddleware(authService, userInvestorService), userHandler.DeactiveUser)
	api.POST("/admin/active_user/:admin_id", middleware.AuthApiAdminMiddleware(authService, userInvestorService), userHandler.ActiveUser)

	// make endpoint get all user by admin
	api.GET("/admin/get_all_user/:admin_id", middleware.AuthApiAdminMiddleware(authService, userInvestorService), userHandler.GetAllUserData)

	// make endpoint update user by admin

	// make update password by admin

	// make delete user by admin
	api.POST("/admin/delete_user/:admin_id", middleware.AuthApiAdminMiddleware(authService, userInvestorService), userHandler.DeleteUserByAdmin)

	// make endpoint delete user SoftDelete and change on login and ferify if delete status account

	// update /admin/ if change data must with time and push array data user id

	// Rounting start for investor
	// starting endpoint
	//make service health for investor
	api.GET("/service_start", userHandler.ServiceStart)
	api.GET("/service_check", middleware.AuthMiddleware(authService, userInvestorService), userHandler.ServiceCheckDB)
	api.POST("/register_investor", userHandler.RegisterUser)
	api.POST("/login_investor", userHandler.Login)
	api.POST("/email_check", userHandler.CheckEmailAvailability)
	api.POST("/phone_check", userHandler.CheckPhoneAvailability)

	//make get user by auth
	api.GET("/get_user", middleware.AuthMiddleware(authService, userInvestorService), userHandler.GetUser)

	//make update profile user by unix_id
	api.PUT("/update_profile/:unix_id", middleware.AuthMiddleware(authService, userInvestorService), userHandler.UpdateUser)

	//make update password user by unix_id
	api.PUT("/update_password/:unix_id", middleware.AuthMiddleware(authService, userInvestorService), userHandler.UpdatePassword)

	//make create image profile user by unix_id

	//make update image profile user by unix_id

	//make delete image profile user by unix_id

	// make logout user by unix_id
	api.POST("/logout_investor", middleware.AuthMiddleware(authService, userInvestorService), userHandler.LogoutUser)

	// end Rounting
	// make env SERVICE_HOST and SERVICE_PORT
	url := fmt.Sprintf("%s:%s", os.Getenv("SERVICE_HOST"), os.Getenv("SERVICE_PORT"))
	router.Run(url)

}
