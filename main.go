package main

import (
	"go-final-project/docs"
	"go-final-project/routers"
	"log"
	"os"

	"github.com/joho/godotenv"                 // gin
	swagger "github.com/swaggo/files"          // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := routers.SetupRouter()
	port := os.Getenv("port")

	docs.SwaggerInfo.Title = "Todo Swaggo"
	docs.SwaggerInfo.Description = "Todo API Documentations"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.Schemes = []string{"http"}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swagger.Handler))

	router.Run(":" + port)

}
