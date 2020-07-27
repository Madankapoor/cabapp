package config

import (
	"log"
	"os"

	"github.com/Madankapoor/cabapp/backendapi/cabprofileservice/api"
	"github.com/Madankapoor/cabapp/backendapi/cabprofileservice/db"
	"github.com/Madankapoor/cabapp/backendapi/cabprofileservice/middleware"
	"github.com/gin-contrib/gzip"
	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

// Setup cabprofileservice
func Setup() *gin.Engine {
	//Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file, please create one in the root directory")
	}
	//Start Mysql database
	db.Init()
	//Start Redis on database 1 - it's used to store the JWT but you can use it for anythig else
	//Example: db.GetRedis().Set(KEY, VALUE, at.Sub(now)).Err()
	db.InitRedis("1")
	//Start the default gin server
	r := gin.Default()
	r.Use(db.Inject(db.GetDB()))
	r.Use(middleware.CORSMiddleware())
	r.Use(middleware.RequestIDMiddleware())
	r.Use(gzip.Gzip(gzip.DefaultCompression))

	// API Setup
	api.ApplyRoutes(r)
	if os.Getenv("ENV") == "PRODUCTION" {
		gin.SetMode(gin.ReleaseMode)
	}
	return r
}
