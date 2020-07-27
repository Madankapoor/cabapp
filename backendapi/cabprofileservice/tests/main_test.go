package tests

import (
	"log"
	"testing"

	config "github.com/Madankapoor/cabapp/backendapi/cabprofileservice/config"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	//gin.SetMode(gin.TestMode)
	r := config.Setup()
	return r
}

func main() {
	r := SetupRouter()
	r.Run()
}

/**
* TestIntDB
* It tests the connection to the database and init the db for this test
*
* Must pass
 */
func TestIntDB(t *testing.T) {
	//Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file, please create one in the root directory")
	}
	db.Init()
	db.InitRedis("1")
}
