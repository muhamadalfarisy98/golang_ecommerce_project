package main

import (
	"ecommerce/config"
	"ecommerce/docs"
	"ecommerce/routes"
	"ecommerce/utils"
	"log"

	"github.com/joho/godotenv"
)

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @termsOfService http://swagger.io/terms/

func main() {

	// for load godotenv
	// for env
	environment := utils.Getenv("ENVIRONMENT", "development")
	if environment == "development" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	//programmatically set swagger info
	docs.SwaggerInfo.Title = "Swagger ClotheStore API"
	docs.SwaggerInfo.Description = `This is a documentation of Ecommerce endpoint.\n By Muhamad Alfarisy \n
	Initial account : Admin (username = faris123, password = 123)`
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	// database connection
	db := config.ConnectDataBase()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	r := routes.SetupRouter(db)
	r.Run("localhost:8080")
}
