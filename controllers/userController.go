package controllers

import (
	"bookstore_users-api/domain/users"
	"bookstore_users-api/services"
	"bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateUser(ctx *gin.Context) {

	var user users.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
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
	userId, userErr := strconv.ParseInt(ctx.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("user id must be  a Number  ")
		ctx.JSON(err.Status, err)
		return
	}

	user, getErr := services.GetUser(userId)
	if getErr != nil {
		ctx.JSON(getErr.Status, getErr)
		return
	}
	ctx.JSON(http.StatusOK, user)
	//ctx.String(http.StatusNotImplemented, "implement Me :)")
}
func SearchUser(ctx *gin.Context) {
	ctx.String(http.StatusNotImplemented, "implement Me :)")
}
