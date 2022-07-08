package users

import (
	"encoding/json"
	"fmt"
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
	Update(id, edad int, nombre, apellido, email, fecha_creacion string, altura float64, activo bool) (domain.User, error)
	Delete(id int) error
	UpdateLastNameAndAge(id, edad int, apellido string) (domain.User, error)
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

func (r *repository) Update(id, edad int, nombre, apellido, email, fecha_creacion string, altura float64, activo bool) (domain.User, error) {
	u := domain.User{Id: id, Edad: edad, Nombre: nombre, Apellido: apellido, Email: email, Fecha_creacion: fecha_creacion, Altura: altura, Activo: activo}

	update := false

	for i := range users {
		if users[i].Id == id {
			u.Id = id
			users[i] = u
			update = true
		}
	}
	if !update {
		return domain.User{}, fmt.Errorf("Usuario %v no encontrado", id)
	}
	return u, nil
}

func (r *repository) Delete(id int) error {
	delete := false
	var index int
	for i := range users {
		if users[i].Id == id {
			index = i
			delete = true
		}
	}
	if !delete {
		return fmt.Errorf("Usuario %v no encontrado", id)
	}
	users = append(users[:index], users[index+1:]...)
	return nil
}

func (r *repository) UpdateLastNameAndAge(id, edad int, apellido string) (domain.User, error) {
	update := false
	var u domain.User
	for i := range users {
		if users[i].Id == id {
			users[i].Apellido = apellido
			users[i].Edad = edad
			update = true
			u = users[i]
		}
	}
	if !update {
		return domain.User{}, fmt.Errorf("Usuario %v no encontrado", id)
	}
	return u, nil
}
