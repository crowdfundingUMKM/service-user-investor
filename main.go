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

	// setup log
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

	// admin request
	api.GET("/admin/log_service_toAdmin/:admin_id", userHandler.GetLogtoAdmin)
	api.GET("/admin/service_status/:admin_id", userHandler.ServiceHealth)
	// make endpoint deactive user
	api.POST("/admin/deactive_user/:admin_id", userHandler.DeactiveUser)
	api.POST("/admin/active_user/:admin_id", userHandler.ActiveUser)

	// make endpoint get all user by admin

	// make update password by admin

	// make delete user by admin

	// make

	// Rounting start for investor
	api.POST("/register_investor", userHandler.RegisterUser)
	api.POST("/login_investor", userHandler.Login)
	api.POST("/email_check", userHandler.CheckEmailAvailability)
	api.POST("/phone_check", userHandler.CheckPhoneAvailability)

	//make get user by auth

	//make update profile user by unix_id

	//make update password user by unix_id

	//make create image profile user by unix_id

	//make update image profile user by unix_id

	//make service health for investor

	// end Rounting
	// make env SERVICE_HOST and SERVICE_PORT
	url := fmt.Sprintf("%s:%s", os.Getenv("SERVICE_HOST"), os.Getenv("SERVICE_PORT"))
	router.Run(url)

}
