package handler

import (
	"fmt"
	"strconv"

	"github.com/JofeGonzalezMeLi/goweb/cmd/internal/users"
	"github.com/gin-gonic/gin"
)

type request struct {
	Edad           int     `json:"edad"`
	Nombre         string  `json:"nombre"`
	Apellido       string  `json:"apellido"`
	Email          string  `json:"email"`
	Fecha_creacion string  `json:"fecha_creacion"`
	Altura         float64 `json:"altura"`
	Activo         bool    `json:"activo"`
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

func (u *User) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != "123456" {
			ctx.JSON(401, gin.H{
				"error": "no tiene permisos para realizar la petición solicitada",
			})
			return
		}
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(404, gin.H{
				"error": "Invalid ID",
			})
		}
		var req request
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		if req.Edad == 0 {
			ctx.JSON(404, gin.H{
				"error": "El campo Edad es necesario",
			})
		}
		if req.Nombre == "" {
			ctx.JSON(404, gin.H{
				"error": "El campo Nombre es necesario",
			})
		}
		if req.Apellido == "" {
			ctx.JSON(404, gin.H{
				"error": "El campo Apellido es necesario",
			})
		}
		if req.Email == "" {
			ctx.JSON(404, gin.H{
				"error": "El campo Emaill es necesario",
			})
		}
		if req.Fecha_creacion == "" {
			ctx.JSON(404, gin.H{
				"error": "El campo Edad es necesario",
			})
		}
		if req.Altura == 0 {
			ctx.JSON(404, gin.H{
				"error": "El campo Edad es necesario",
			})
		}
		u, err := u.service.Update(id, req.Edad, req.Nombre, req.Apellido, req.Email, req.Fecha_creacion, req.Altura, req.Activo)
		if err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, u)
	}
}

func (u *User) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != "123456" {
			ctx.JSON(401, gin.H{
				"error": "no tiene permisos para realizar la petición solicitada",
			})
			return
		}
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(404, gin.H{
				"error": "Invalid ID",
			})
		}
		err = u.service.Delete(id)
		if err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, gin.H{
			"data": fmt.Sprintf("El Usuario con id %v ha sido borrado", id),
		})
	}
}

func (u *User) UpdateLastNameAndAge() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != "123456" {
			ctx.JSON(401, gin.H{
				"error": "no tiene permisos para realizar la petición solicitada",
			})
			return
		}
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(404, gin.H{
				"error": "Invalid ID",
			})
		}
		var req request
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		if req.Edad == 0 {
			ctx.JSON(404, gin.H{
				"error": "El campo Edad es necesario",
			})
		}
		if req.Apellido == "" {
			ctx.JSON(404, gin.H{
				"error": "El campo Apellido es necesario",
			})
		}
		u, err := u.service.UpdateLastNameAndAge(id, req.Edad, req.Apellido)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, u)
	}
}
