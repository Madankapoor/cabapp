package main

import (
	"os"

	config "github.com/Madankapoor/cabapp/backendapi/cabprofileservice/config"
)

func main() {
	r := config.Setup()
	port := os.Getenv("PORT")
	r.Run(":" + port)
}
