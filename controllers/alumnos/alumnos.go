package alumnos

import (
	"net/http"

	database "github.com/frank1995alfredo/api/database"
	maestroalumno "github.com/frank1995alfredo/api/models/maestroalumno"

	"github.com/gin-gonic/gin"
)

//ObtenerAlumnos ... funcion para presentar todos los maestros
func ObtenerAlumnos(c *gin.Context) {
	var alumno []maestroalumno.Alumno

	database.DB.Order("alumno_id").Find(&alumno)

	c.SecureJSON(http.StatusOK, gin.H{"data": alumno})
}

//CrearAlumno ... funcion para insertar a una alumno
func CrearAlumno(c *gin.Context) {

	//CrearAlumnoInput ... estructura para validar los inputs
	type CrearAlumnoInput struct {
		Nombre   string `json:"nombre" binding:"required"`
		Apellido string `json:"apellido" binding:"required"`
		Edad     int    `json:"edad" binding:"required"`
		MaesID   uint   `json:"maesid" binding:"required"`
	}
	var input CrearAlumnoInput

	//validaops los inputs
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//crea la persona en la base de datos
	alumno := maestroalumno.Alumno{Nombre: input.Nombre, Apellido: input.Apellido, Edad: input.Edad, MaesID: input.MaesID}

	database.DB.Create(&alumno)

	tx := database.DB.Begin()
	err := tx.Create(&alumno).Error //si no hay un error, se guarda el articulo
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()

	c.SecureJSON(http.StatusBadRequest, gin.H{"data": alumno})
}
