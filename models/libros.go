package models

//Libro ... modelo libro
type Libro struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Titulo string `json:"titulo"`
	Autor  string `json:"autor"`
}
