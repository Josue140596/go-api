package main

import (
	"go/api/internal/user"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	userService := user.NewService()
	user := user.MakeEndpoints(userService)
	router.GET("/user", gin.HandlerFunc(user.Get))
	router.POST("/user", gin.HandlerFunc(user.Create))

	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
