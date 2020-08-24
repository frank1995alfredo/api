package personas

import (
	"net/http"

	database "github.com/frank1995alfredo/api/database"
	personas "github.com/frank1995alfredo/api/models/personas"
	"github.com/gin-gonic/gin"
)

//ObtenerPersona ...
func ObtenerPersona(c *gin.Context) {
	var personas []personas.Persona

	database.DB.Order("id").Find(&personas)

	c.JSON(http.StatusOK, gin.H{"data": personas})

}

//CrearPersona ... funcion para inserar a una persona
func CrearPersona(c *gin.Context) {
	var input personas.CrearPersonaInput
	var per []personas.Persona

	//validaops los inputs
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Where("num_cedula=?", input.NumCedula).First(&per).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ya existe una cedula con ese numero"})
		return
	}

	//crea la persona en la base de datos
	persona := personas.Persona{Nombre: input.Nombre, Apellido: input.Apellido, NumCedula: input.NumCedula}

	tx := database.DB.Begin()
	err := tx.Create(&persona).Error //si no hay un error, se guarda el articulo
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()

	c.JSON(http.StatusBadRequest, gin.H{"data": persona})
}

//BuscarPersona ... funcion para buscar a una persona
func BuscarPersona(c *gin.Context) {
	var persona personas.Persona

	if err := database.DB.Where("id=?", c.Param("id")).First(&persona).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No existe esa persona con ese id"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": persona})
}

//ActualizarPersona ... funcion para actualizar persona
func ActualizarPersona(c *gin.Context) {
	var persona personas.Persona

	if err := database.DB.Where("id=?", c.Param("id")).First(&persona).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Persona no econtrada"})
	}

	//validamos la entrada de los datos
	var input personas.ActualizarPersonaInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	tx := database.DB.Begin()
	err := tx.Model(&persona).Updates(input).Error //si no hay un error, se guarda el articulo
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()

	c.JSON(http.StatusOK, gin.H{"data": persona})
}

//EliminarPersona ... funcion que permite eliminar un libro
func EliminarPersona(c *gin.Context) {
	var persona personas.Persona

	if err := database.DB.Where("id=?", c.Param("id")).First(&persona).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Persona no encontrada"})
		return
	}

	tx := database.DB.Begin()
	err := tx.Delete(&persona).Error //si no hay un error, se guarda el articulo
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()

	c.JSON(http.StatusOK, gin.H{"data": "Persona eliminada"})
}
