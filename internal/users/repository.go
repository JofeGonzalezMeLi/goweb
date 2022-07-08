package users

import (
	"encoding/json"
	"log"
	"os"

	"github.com/JofeGonzalezMeLi/goweb/cmd/internal/domain"
)

var users []domain.User
var lastId int

type Repository interface {
	GetAll() ([]domain.User, error)
	Store(id, edad int, nombre, apellido, email, fecha_creacion string, altura float64, activo bool) (domain.User, error)
	LastId() (int, error)
}

type repository struct{}

func NewRepository() Repository {
	j, _ := os.ReadFile("./usuarios.json")
	if err := json.Unmarshal(j, &users); err != nil {
		log.Println(string(j))
		log.Fatal(err)
	}
	return &repository{}
}

func (r *repository) GetAll() ([]domain.User, error) {
	return users, nil
}

func (r *repository) LastId() (int, error) {
	return lastId, nil
}

func (r *repository) Store(id, edad int, nombre, apellido, email, fecha_creacion string, altura float64, activo bool) (domain.User, error) {
	u := domain.User{Id: id, Edad: edad, Nombre: nombre, Apellido: apellido, Email: email, Fecha_creacion: fecha_creacion, Altura: altura, Activo: activo}
	users = append(users, u)
	lastId = u.Id
	return u, nil
}
