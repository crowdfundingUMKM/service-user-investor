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

	// Rounting start
	api.POST("register_investor", userHandler.RegisterUser)
	// make api for login

	// end Rounting
	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	router.Run(port)

}
