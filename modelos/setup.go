package modelos

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//DB ...
var DB *gorm.DB

//ConectorBD ... permite conectar a la base de datos
func ConectorBD() {
	bd, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=golang password=1234 sslmode=disable")
	if err != nil {
		panic("ERROR al conectar con la base de datos.")
	}
	bd.AutoMigrate(&Libros{})

	DB = bd
}
