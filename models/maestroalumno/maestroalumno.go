package maestroalumno

//Maestro ... estructura de maestro
type Maestro struct {
	MaestroID int      `json:"idmaestro" gorm:"primary_key"`
	Nombre    string   `json:"nombreMaestro" gorm:"size:50"`
	Apellido  string   `json:"apellidoMaestro" gorm:"size:50"`
	NumCedula string   `json:"numcedula" gorm:"size:11"`
	Alumnos   []Alumno `gorm:"Foreingkey:MaesID"`
}

//Alumno ... estructura de alumno, estructura que se relaciona con maestro
type Alumno struct {
	AlumnoID int    `json:"idalumno" gorm:"primary_key"`
	Nombre   string `json:"nombreAlumno" gorm:"size:50"`
	Apellido string `json:"apellidoAlumno" gorm:"size:50"`
	Edad     int    `json:"edad" gorm:"not null"`
	MaesID   int    `json:"maesid" gorm:"not null"`
}

/*MaestroAlumno ... estructura que hereda de Maestro y Alumno
no se hace migracion de esta estructura*/
type MaestroAlumno struct {
	Maestro
	Alumno
}

//CrearMaestroInput ... structura para valida los inpu
type CrearMaestroInput struct {
	Nombre    string `json:"nombre" binding:"required"`
	Apellido  string `json:"apellido" binding:"required"`
	NumCedula string `json:"numcedula" binding:"required"`
}

//ActualizarMaestroInput ... estructura para valida los input al momento de actualizar los campo
type ActualizarMaestroInput struct {
	Nombre    string `json:"nombre" binding:"required"`
	Apellido  string `json:"apellido" binding:"required"`
	NumCedula string `json:"numcedula" binding:"required"`
}

//CrearAlumnoInput ... estructura para validar los inputs
type CrearAlumnoInput struct {
	Nombre   string `json:"nombre" binding:"required"`
	Apellido string `json:"apellido" binding:"required"`
	Edad     int    `json:"edad" binding:"required"`
	MaesID   int    `json:"maesid" binding:"required"`
}
