package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"

	config "github.com/Madankapoor/cabapp/backendapi/passengerprofileservice/config"
	"github.com/Madankapoor/cabapp/backendapi/passengerprofileservice/db"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
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

/**
* TestCreatePassenger
* Test passenger create
*
* Must return response code 200
 */
func TestCreatePassenger(t *testing.T) {
	testRouter := SetupRouter()
	type InputPassenger struct {
		Email    string `json:"email"`
		Name     string `json:"name"`
		MobileNo string `json:"mobileno"`
	}
	var passenger InputPassenger
	passenger.Name = "MadanKapoor"
	passenger.Email = fmt.Sprintf("madankapoor%d@gmail.com", rand.Intn(100))
	passenger.MobileNo = fmt.Sprintf("877%d2946", rand.Intn(100))

	data, _ := json.Marshal(passenger)
	req, err := http.NewRequest("POST", "/api/v1.0/passenger/", bytes.NewBufferString(string(data)))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		fmt.Println(err)
	}
	resp := httptest.NewRecorder()

	testRouter.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
}

/**
* TestDeletePassenger
* Test deleting an passenger
*
* Must return response code 204
 */
func TestDeletePassenger(t *testing.T) {
	testRouter := SetupRouter()
	url := fmt.Sprintf("/api/v1.0/passenger/%d", 1)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		fmt.Println(err)
	}
	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusNoContent, resp.Code)
}

/**
* TestCleanUp
* Deletes the passengers
*
* Must pass
 */
func TestCleanUp(t *testing.T) {
	var err error
	_, err = db.GetDB().DB().Exec("TRUNCATE TABLE passengers")
	if err != nil {
		t.Error(err)
	}
}
