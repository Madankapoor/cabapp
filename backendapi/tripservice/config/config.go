package config

import (
	"fmt"
	"log"
	"os"

	"github.com/Madankapoor/cabapp/backendapi/tripservice/api"
	"github.com/Madankapoor/cabapp/backendapi/tripservice/db"
	"github.com/Madankapoor/cabapp/backendapi/tripservice/middleware"
	"github.com/dghubble/sling"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// InjectServiceClients is used to inject http service clients.
func InjectServiceClients() gin.HandlerFunc {
	return func(c *gin.Context) {
		locationservice := sling.New().Base(fmt.Sprintf("http://%s", os.Getenv("LOCATIONSERVICE"))).Client(nil)
		c.Set("locationservice", locationservice)
		c.Next()
	}
}

// Setup tripservice
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
	// Service Clients Inject
	r.Use(InjectServiceClients())
	// API Setup
	api.ApplyRoutes(r)
	if os.Getenv("ENV") == "PRODUCTION" {
		gin.SetMode(gin.ReleaseMode)
	}
	return r
}
