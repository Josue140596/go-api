package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	Controller func(c *gin.Context)
	Endpoints  struct {
		Create Controller
		Get    Controller
		GetAll Controller
		Update Controller
		Delete Controller
	}
	CreateReq struct {
		FirstName string `form:"name" json:"firstName" binding:"required"`
		LastName  string `form:"lastName" json:"lastName" binding:"required"`
		Email     string `form:"email" json:"email" binding:"required"`
		Password  string `form:"password" json:"password" binding:"required"`
		Phone     string `form:"phone" json:"phone"`
	}
)

func MakeEndpoints() Endpoints {
	return Endpoints{
		Create: makeCreateEndpoint(),
		Get:    makeGetEndpoint(),
		GetAll: makeGetAllEndpoint(),
		Update: makeUpdateEndpoint(),
		Delete: makeDeleteEndpoint(),
	}
}

func makeCreateEndpoint() Controller {
	return func(c *gin.Context) {
		var json CreateReq
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"user": json.FirstName, "ok": true})
	}
}

func makeGetEndpoint() Controller {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"user": "Bryan", "ok": true})
	}
}

func makeGetAllEndpoint() Controller {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "Get all")
		c.JSON(http.StatusOK, gin.H{"user": "Bryan", "ok": true})
	}
}

func makeUpdateEndpoint() Controller {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "Update")
		c.JSON(http.StatusOK, gin.H{"user": "Bryan", "ok": true})
	}
}
func makeDeleteEndpoint() Controller {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "Delete")
		c.JSON(http.StatusOK, gin.H{"user": "Bryan", "ok": true})
	}
}
