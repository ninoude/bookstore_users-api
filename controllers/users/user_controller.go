package users

import (
	_ "fmt"
	"github.com/gin-gonic/gin"
	"github.com/ninoude/bookstore_users-api/domain/users"
	"github.com/ninoude/bookstore_users-api/services"
	"github.com/ninoude/bookstore_users-api/utils/errors"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		//TODO: return bad request to the caller
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		//TODO: Handle user creation error
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me!")
}
