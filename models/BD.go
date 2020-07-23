package models

import (
	"log"

	"github.com/frank1995alfredo/api/models/articulos"
	"github.com/frank1995alfredo/api/models/maestroalumno"
	"github.com/jinzhu/gorm"

	//github.com/lib/pq ... libreria para manejar los pq, controla los orm
	_ "github.com/lib/pq"
)

//DB ...
var DB *gorm.DB

//ConectorBD ... permite conectar a la base de datos
func ConectorBD() {
	bd, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=golang password=1234 sslmode=disable")
	if err != nil {
		log.Println(err.Error())
	}

	//bd.AutoMigrate(&Libro{})
	//bd.AutoMigrate(&personas.Persona{})
	bd.AutoMigrate(&maestroalumno.Maestro{}, &maestroalumno.Alumno{})
	bd.Model(&maestroalumno.Alumno{}).AddForeignKey("maes_id", "maestros(maestro_id)", "CASCADE", "CASCADE")
	bd.AutoMigrate(&articulos.Articulo{}, &articulos.Categoria{})
	bd.Model(&articulos.Articulo{}).AddForeignKey("cat_id", "categoria(categoria_id)", "CASCADE", "CASCADE")

	DB = bd
}
