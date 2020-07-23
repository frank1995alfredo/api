package personas

//Persona ... modelo persona
type Persona struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	Nombre    string `json:"nombre"`
	Apellido  string `json:"apellido"`
	NumCedula string `json:"numcedula"`
}

//CrearPersonaInput ... estructura para validar e insertar a una persona
type CrearPersonaInput struct {
	Nombre    string `json:"nombre" binding:"required"`
	Apellido  string `json:"apellido" binding:"required"`
	NumCedula string `json:"numcedula" binding:"required"`
}

//ActualizarPersonaInput ... estructura para validar y actualizar
type ActualizarPersonaInput struct {
	Nombre    string `json:"nombre"`
	Apellido  string `json:"apellido"`
	NumCedula string `json:"numcedula"`
}
