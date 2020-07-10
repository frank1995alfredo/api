package models

import (
	"log"

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

	bd.AutoMigrate(&Libro{})

	DB = bd
}
