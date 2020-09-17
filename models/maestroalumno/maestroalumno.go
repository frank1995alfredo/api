package maestroalumno

//Maestro ... estructura de maestro
type Maestro struct {
	MaestroID uint     `json:"maestroid" gorm:"primary_key"`
	Nombre    string   `json:"nombre"`
	Apellido  string   `json:"apellido"`
	NumCedula string   `json:"numcedula"`
	Alumnos   []Alumno `json:"alumnos" gorm:"foreignkey:MaesID"`
}

//Alumno ...
type Alumno struct {
	AlumnoID uint   `json:"alumnoid" gorm:"primary_key"`
	Nombre   string `json:"nombreAlumno"`
	Apellido string `json:"apellidoAlumno"`
	Edad     int    `json:"edadAlumno"`
	MaesID   uint   `json:"maesid"`
}

//MaestroAlumno ...
type MaestroAlumno struct {
	Alumno
	Maestro
}
