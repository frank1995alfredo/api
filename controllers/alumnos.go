package controllers

import (
	"net/http"

	"github.com/frank1995alfredo/api/models"
	"github.com/frank1995alfredo/api/models/maestroalumno"

	"github.com/gin-gonic/gin"
)

//ObtenerAlumnos ... funcion para presentar todos los maestros
func ObtenerAlumnos(c *gin.Context) {
	var alumno []maestroalumno.Alumno

	models.DB.Order("alumno_id").Find(&alumno)

	c.Header("Access-Control-Allow-Origin", "*")

	c.JSON(http.StatusOK, gin.H{"data": alumno})
}

//CrearAlumno ... funcion para insertar a una alumno
func CrearAlumno(c *gin.Context) {
	var input maestroalumno.Alumno

	//validaops los inputs
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//crea la persona en la base de datos
	alumno := maestroalumno.Alumno{Nombre: input.Nombre, Apellido: input.Apellido, Edad: input.Edad, MaesID: input.MaesID}

	models.DB.Create(&alumno)

	c.Header("Access-Control-Allow-Origin", "*")

	c.JSON(http.StatusBadRequest, gin.H{"data": alumno})
}
