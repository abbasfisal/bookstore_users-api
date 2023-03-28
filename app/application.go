package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {

	mapUrls()

	fmt.Println("start server on Port 8080")
	router.Run(":8080")
}
