package domain

type User struct {
	Id             int     `json:"id"`
	Edad           int     `json:"edad" binding:"required"`
	Nombre         string  `json:"nombre" binding:"required"`
	Apellido       string  `json:"apellido" binding:"required"`
	Email          string  `json:"email" binding:"required"`
	Fecha_creacion string  `json:"fecha_creacion" binding:"required"`
	Altura         float64 `json:"altura" binding:"required"`
	Activo         bool    `json:"activo" binding:"required"`
}