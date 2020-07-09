package modelos

import "github.com/jinzhu/gorm"

//Libro ... modelo libro
type Libro struct {
	gorm.Model
	ID     uint   `json:"id" gorm:"primary_key"`
	Titulo string `json:"titulo"`
	Autor  string `json:"autor"`
}
