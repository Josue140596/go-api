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
		c.String(http.StatusOK, "create")
		c.JSON(http.StatusOK, gin.H{"user": "Bryan", "ok": true})
	}
}

func makeGetEndpoint() Controller {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "Get")
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
