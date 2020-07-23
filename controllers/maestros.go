package controllers

import (
	"net/http"

	"github.com/frank1995alfredo/api/models"
	"github.com/frank1995alfredo/api/models/maestroalumno"
	"github.com/gin-gonic/gin"
)

//ObtenerMaestros ... funcion para presentar todos los maestros
func ObtenerMaestros(c *gin.Context) {
	var maestros []maestroalumno.Maestro

	models.DB.Order("maestro_id").Find(&maestros)

	c.Header("Access-Control-Allow-Origin", "*")

	c.JSON(http.StatusOK, gin.H{"data": maestros})
}

//CrearMaestro ...  funciona para craar un maestro
func CrearMaestro(c *gin.Context) {
	var input maestroalumno.CrearMaestroInput
	var per []maestroalumno.Maestro

	//validaops los inputs
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.DB.Where("num_cedula=?", input.NumCedula).First(&per).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Numero de cedula ya esta usado"})
		return
	}

	//crea crea al maestro en la base de datos
	maestro := maestroalumno.Maestro{Nombre: input.Nombre, Apellido: input.Apellido, NumCedula: input.NumCedula}

	models.DB.Create(&maestro)

	c.Header("Access-Control-Allow-Origin", "*")

	c.JSON(http.StatusOK, gin.H{"data": maestro})

}

//BuscarMaestro ... funcion para buscar el maestro
func BuscarMaestro(c *gin.Context) {
	var maestro maestroalumno.Maestro

	if err := models.DB.Where("id=?", c.Param("id")).First(&maestro).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No existe el maestro"})
		return
	}

	c.Header("Access-Control-Allow-Origin", "*")

	c.JSON(http.StatusOK, gin.H{"data": maestro})
}

//ActualizarMaestro ... funcion para actualiar un maetro
func ActualizarMaestro(c *gin.Context) {
	var maestro maestroalumno.Maestro

	if err := models.DB.Where("id=?", c.Param("id")).First(&maestro).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Maestro no econtrado"})
	}

	//validamos la entrada de los datos
	var input maestroalumno.ActualizarMaestroInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&maestro).Updates(input)

	c.Header("Access-Control-Allow-Origin", "*")

	c.JSON(http.StatusOK, gin.H{"data": maestro})
}

//EliminarMaestro ... funcion para eliminar un maestro
func EliminarMaestro(c *gin.Context) {
	var maestro maestroalumno.Maestro

	if err := models.DB.Where("id=?", c.Param("id")).First(&maestro).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Delete(&maestro)

	c.Header("Access-Control-Allow-Origin", "*")

	c.JSON(http.StatusOK, gin.H{"data": "Maestro eliminado"})
}

//MostrarMaesAlum ... funcion para datos del amestro y el alumno
func MostrarMaesAlum(c *gin.Context) {

	var maesAlum []maestroalumno.MaestroAlumno

	models.DB.Table("articulos").Order("articulo_id").Select("articulos.nombre, articulos.precio, articulos.cantidad, articulos.total," +
		"categoria.descripcion").Joins("JOIN categoria ON articulos.cat_id = categoria.categoria_id").Find(&maesAlum)

	c.Header("Access-Control-Allow-Origin", "*")

	c.JSON(http.StatusOK, gin.H{"data": maesAlum})
}
