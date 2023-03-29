package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

var (
	router = gin.Default()
)

func StartApplication() {

	mapUrls()

	fmt.Println("start server on Port 8080")
	log.Fatal(router.Run(":8081"))
}
