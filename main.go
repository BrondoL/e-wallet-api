package main

import (
	"fmt"
	"os"

	"git.garena.com/sea-labs-id/batch-02/aulia-nabil/assignment-05-golang-backend/config"
	"github.com/gin-gonic/gin"
)

func main() {
	config.GetConn()

	router := gin.Default()
	router.Static("/docs/", "./pkg/swaggerui")

	version := os.Getenv("API_VERSION")
	router.Group(fmt.Sprintf("/api/%s", version))

	router.Run(":8000")
}
