package handler

import (
	"github.com/JofeGonzalezMeLi/goweb/cmd/internal/users"
	"github.com/gin-gonic/gin"
)

type request struct {
	Edad           int     `json:"edad" binding:"required"`
	Nombre         string  `json:"nombre" binding:"required"`
	Apellido       string  `json:"apellido" binding:"required"`
	Email          string  `json:"email" binding:"required"`
	Fecha_creacion string  `json:"fecha_creacion" binding:"required"`
	Altura         float64 `json:"altura" binding:"required"`
	Activo         bool    `json:"activo" binding:"required"`
}

type User struct {
	service users.Service
}

func NewUser(u users.Service) *User {
	return &User{
		service: u,
	}
}

func (u *User) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != "123456" {
			ctx.JSON(401, gin.H{
				"error": "no tiene permisos para realizar la petición solicitada",
			})
			return
		}

		us, err := u.service.GetAll()
		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(200, gin.H{
			"usuarios": us,
		})
	}
}

func (u *User) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != "123456" {
			ctx.JSON(401, gin.H{
				"error": "no tiene permisos para realizar la petición solicitada",
			})
			return
		}
		var req request
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		us, err := u.service.Store(req.Edad, req.Nombre, req.Apellido, req.Email, req.Fecha_creacion, req.Altura, req.Activo)
		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(200, us)
	}
}
