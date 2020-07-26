package main

import (
	"os"

	"github.com/Madankapoor/cabapp/backendapi/passengerprofileservice/config"
)

func main() {
	r := config.Setup()
	port := os.Getenv("PORT")
	r.Run(":" + port)
}
