package personas

//Persona ... modelo persona
type Persona struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	Nombre    string `json:"nombre"`
	Apellido  string `json:"apellido"`
	NumCedula string `json:"numcedula"`
}
