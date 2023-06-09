package main

import (
	"fmt"
	"log"
	"os"

	"service-user-investor/auth"
	"service-user-investor/database"
	"service-user-investor/handler"
	"service-user-investor/investor"

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
	// L.InitLog()

	// SETUP REPO
	db := database.NewConnectionDB()
	userInvestorRepository := investor.NewRepository(db)

	// SETUP SERVICE
	userInvestorService := investor.NewService(userInvestorRepository)
	authService := auth.NewService()

	// setup handler
	userHandler := handler.NewUserHandler(userInvestorService, authService)

	// END SETUP

	// RUN SERVICE
	router := gin.Default()
	api := router.Group("api/v1")

	// admin request
	// api.GET("/admin/log_service_toAdmin/:id", userHandler.GetLogtoAdmin)
	api.GET("/admin/service_status/:admin_id", userHandler.ServiceHealth)
	// make endpoint deactive user
	api.POST("/admin/deactive_user/:admin_id", userHandler.DeactiveUser)
	api.POST("/admin/active_user/:admin_id", userHandler.ActiveUser)

	// Rounting start for investor
	api.POST("/register_investor", userHandler.RegisterUser)
	api.POST("/login_investor", userHandler.Login)
	api.POST("/email_check", userHandler.CheckEmailAvailability)
	api.POST("/phone_check", userHandler.CheckPhoneAvailability)

	// errorr email_check

	// end Rounting
	// make env SERVICE_HOST and SERVICE_PORT
	url := fmt.Sprintf("%s:%s", os.Getenv("SERVICE_HOST"), os.Getenv("SERVICE_PORT"))
	router.Run(url)

}
