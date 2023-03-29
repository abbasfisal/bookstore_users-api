package controllers

import (
	"bookstore_users-api/domain/users"
	"bookstore_users-api/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(ctx *gin.Context) {

	var user users.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message":         "error while unmarshaling data",
			"error":           err,
			"error to string": err.Error(),
		})
		return
	}

	result, cErr := services.CreateUser(user)
	if cErr != nil {
		//TODO json error
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": cErr.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "user created successfully",
		"data":    result,
	})
	
}
func GetUser(ctx *gin.Context) {
	ctx.String(http.StatusNotImplemented, "implement Me :)")
}
func SearchUser(ctx *gin.Context) {
	ctx.String(http.StatusNotImplemented, "implement Me :)")
}
