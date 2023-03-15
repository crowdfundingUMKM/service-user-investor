package main

import (
	"fmt"
	"log"
	"os"

	"service-user-investor/auth"
	"service-user-investor/handler"
	"service-user-investor/investor"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// initial connected db
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var (
		dbUser    = os.Getenv("DB_USER")       // e.g. 'my-db-user'
		dbPwd     = os.Getenv("DB_PASS")       // e.g. 'my-db-password'
		dbName    = os.Getenv("DB_NAME")       // e.g. 'my-database'
		dbPort    = os.Getenv("DB_PORT")       // e.g. '3306'
		dbTCPHost = os.Getenv("INSTANCE_HOST") // e.g. '127.0.0.1' ('172.17.0.1' if deployed to GAE Flex)
	)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPwd, dbTCPHost, dbPort, dbName)

	// end connected

	// open connection
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	fmt.Println(db)
	if err != nil {
		log.Fatal(err.Error())
	}

	// SETUP REPO

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
	api.POST("users_investor", userHandler.RegisterUser)

	// end Rounting
	router.Run()

}
