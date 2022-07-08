package main

import (
	"github.com/JofeGonzalezMeLi/goweb/cmd/cmd/server/handler"
	"github.com/JofeGonzalezMeLi/goweb/cmd/internal/users"
	"github.com/gin-gonic/gin"
)

func main() {
	repository := users.NewRepository()
	service := users.NewService(repository)
	u := handler.NewUser(service)

	router := gin.Default()
	ug := router.Group("/users")
	ug.POST("/", u.Store())
	ug.GET("/", u.GetAll())
	router.Run()
}
