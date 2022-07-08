package main

import (
	"github.com/JofeGonzalezMeLi/goweb/cmd/cmd/server/handler"
	"github.com/JofeGonzalezMeLi/goweb/cmd/internal/users"
	"github.com/JofeGonzalezMeLi/goweb/cmd/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	_ = godotenv.Load()

	db := store.NewStore("./usuarios.json")

	repository := users.NewRepository(db)
	service := users.NewService(repository)
	u := handler.NewUser(service)

	router := gin.Default()
	ug := router.Group("/users")
	ug.POST("/", u.Store())
	ug.GET("/", u.GetAll())
	ug.PUT("/:id", u.Update())
	ug.DELETE("/:id", u.Delete())
	ug.PATCH("/:id", u.UpdateLastNameAndAge())
	router.Run()
}
