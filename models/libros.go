package models

//Libro ... modelo libro
type Libro struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Titulo string `json:"titulo"`
	Autor  string `json:"autor"`
}

//CrearLibroInput ... estructura para validar e insertar libro
type CrearLibroInput struct {
	Titulo string `json:"titulo" binding:"required"`
	Autor  string `json:"autor" binding:"required"`
}

//ActualizarLibro ... estructura para validar y actulizar el libro
type ActualizarLibro struct {
	Titulo string `json:"titulo"`
	Autor  string `json:"autor"`
}
