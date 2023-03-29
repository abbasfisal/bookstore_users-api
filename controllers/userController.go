package controllers

import (
	"bookstore_users-api/domain/users"
	"bookstore_users-api/services"
	"bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(ctx *gin.Context) {

	var user users.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		restErr := errors.RestErr{
			Message: "invalid json body",
			Status:  http.StatusBadRequest,
			Error:   "bad request",
		}
		ctx.JSON(restErr.Status, restErr)
		return
	}

	result, cErr := services.CreateUser(user)
	if cErr != nil {
		ctx.JSON(cErr.Status, cErr)
		return
	}
	ctx.JSON(http.StatusCreated, result)

}
func GetUser(ctx *gin.Context) {
	ctx.String(http.StatusNotImplemented, "implement Me :)")
}
func SearchUser(ctx *gin.Context) {
	ctx.String(http.StatusNotImplemented, "implement Me :)")
}
