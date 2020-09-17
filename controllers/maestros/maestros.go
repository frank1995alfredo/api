package maestros

import (
	"net/http"

	database "github.com/frank1995alfredo/api/database"
	maestroalumno "github.com/frank1995alfredo/api/models/maestroalumno"
	"github.com/gin-gonic/gin"
)

//ObtenerMaestros ... funcion para presentar todos los maestros
func ObtenerMaestros(c *gin.Context) {
	var maestros []maestroalumno.Maestro

	database.DB.Order("maestro_id").Find(&maestros)

	c.SecureJSON(http.StatusOK, gin.H{"data": maestros})
}

//CrearMaestro ...  funciona para craar un maestro
func CrearMaestro(c *gin.Context) {

	//CrearMaestroInput ... structura para validar los inputs
	type CrearMaestroInput struct {
		Nombre         string `json:"nombre"`
		Apellido       string `json:"apellido"`
		NumCedula      string `json:"numcedula"`
		NombreAlumno   string `json:"nombreAlumno"`
		ApellidoAlumno string `json:"apellidoAlumno"`
		EdadAlumno     int    `json:"edadAlumno"`
	}

	//input toma los atributos de CrearMaestroInput, sirve para validar los input del frontend
	var input CrearMaestroInput

	//validamos los inputs
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	maestro := maestroalumno.Maestro{Nombre: input.Nombre,
		Apellido:  input.Apellido,
		NumCedula: input.NumCedula,
		Alumnos: []maestroalumno.Alumno{{Nombre: input.NombreAlumno,
			Apellido: input.ApellidoAlumno, Edad: input.EdadAlumno}}}

	tx := database.DB.Begin()
	err := tx.Debug().Save(&maestro).Error
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()
	//fin de la transaccion

	c.SecureJSON(http.StatusOK, gin.H{"data": maestro})
}

//BuscarMaestro ... funcion para buscar el maestro
func BuscarMaestro(c *gin.Context) {
	var maestro maestroalumno.Maestro

	if err := database.DB.Where("id=?", c.Param("id")).First(&maestro).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No existe el maestro"})
		return
	}

	c.SecureJSON(http.StatusOK, gin.H{"data": maestro})
}

//ActualizarMaestro ... funcion para actualiar un maetro
func ActualizarMaestro(c *gin.Context) {

	//ActualizarMaestroInput ... estructura para valida los input al momento de actualizar los campo
	type ActualizarMaestroInput struct {
		Nombre    string `json:"nombre" binding:"required"`
		Apellido  string `json:"apellido" binding:"required"`
		NumCedula string `json:"numcedula" binding:"required"`
	}

	var maestro maestroalumno.Maestro

	if err := database.DB.Where("id=?", c.Param("id")).First(&maestro).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Maestro no econtrado"})
	}

	//validamos la entrada de los datos
	var input ActualizarMaestroInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	//inicio de la transaccion
	tx := database.DB.Begin()
	err := tx.Model(&maestro).Updates(input).Error
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()
	//fin de la transaccion

	c.SecureJSON(http.StatusOK, gin.H{"data": maestro})
}

//EliminarMaestro ... funcion para eliminar un maestro
func EliminarMaestro(c *gin.Context) {
	var maestro maestroalumno.Maestro

	if err := database.DB.Where("id=?", c.Param("id")).First(&maestro).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//inicio de la transaccion
	tx := database.DB.Begin()
	err := tx.Delete(&maestro).Error
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()
	//fin de la transaccion

	c.SecureJSON(http.StatusOK, gin.H{"data": "Maestro eliminado"})
}

//MostrarMaesAlum ... funcion para datos del amestro y el alumno
func MostrarMaesAlum(c *gin.Context) {

	var maesAlum []maestroalumno.MaestroAlumno

	//modificar el query
	database.DB.Table("articulos").Order("articulo_id").Select("articulos.nombre, articulos.precio, articulos.cantidad, articulos.total," +
		"categoria.descripcion").Joins("JOIN categoria ON articulos.cat_id = categoria.categoria_id").Find(&maesAlum)

	c.SecureJSON(http.StatusOK, gin.H{"data": maesAlum})
}
