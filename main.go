package main

import (
	"log"
	"go-gin-jwt-authorization-example/configs"
	_ "go-gin-jwt-authorization-example/docs"
	"go-gin-jwt-authorization-example/routes"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found: %s", err)
	}
}

// @title Online Course API
// @version 1.0.0
// @description Online Course API

// @host localhost:4000
// @BasePath /api/v1
// @query.collection.format multi

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationurl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationurl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information

// @x-extension-openapi {"example": "value on a json format"}

func main() {
	// Load env variables
	configs.InitEnv()

	// Init database
	if err := configs.InitPersistence(); err != nil {
		log.Printf("Database connection error: %s", err)
	}
	defer configs.DB.Close()

	// Init redis
	configs.InitRedis()

	r := routes.SetupRouter()
	r.Run()
}
