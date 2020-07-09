package modelos

import (
	"fmt"
)

//Libro ... modelo libro
type Libro struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Titulo string `json:"titulo"`
	Autor  string `json:"autor"`
}

func nombre() {
	fmt.Print()
}
