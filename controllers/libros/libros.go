package libros

import (
	"net/http"
	//sgsd

	database "github.com/frank1995alfredo/api/database"
	libros "github.com/frank1995alfredo/api/models/libros"
	"github.com/gin-gonic/gin"
)

//ObtenerLibros ... controlador para obtener todos los libros
func ObtenerLibros(c *gin.Context) {
	var libros []libros.Libro

	database.DB.Order("id").Find(&libros)

	c.JSON(http.StatusOK, gin.H{"data": libros})
}

//CrearLibro ... funcion para inserta un libro nuevo
func CrearLibro(c *gin.Context) {
	var input libros.CrearLibroInput

	//para validar los inputs
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//crea el libro en la base de datos
	libro := libros.Libro{Titulo: input.Titulo, Autor: input.Autor}
	database.DB.Create(&libro)

	c.JSON(http.StatusOK, gin.H{"data": libro})
}

//BuscarLibro ... funcion para buscar un libro
func BuscarLibro(c *gin.Context) {
	var libro libros.Libro

	if err := database.DB.Where("id=?", c.Param("id")).First(&libro).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No encontrado"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": libro})
}

//ActualizarLibro ... funcion para actualizar un libro :v
func ActualizarLibro(c *gin.Context) {
	var libro libros.Libro

	if err := database.DB.Where("id = ?", c.Param("id")).First(&libro).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Libro no encontrado"})
		return
	}

	//validamos la entrada de los datos
	var input libros.ActualizarLibro
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	//inicio de la transaccion
	tx := database.DB.Begin()
	err := tx.Model(&libro).Updates(input).Error
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()
	//fin de la transaccion

	c.JSON(http.StatusOK, gin.H{"data": libro})
}

//EliminarLibro ... funcion que permite eliminar un libro
func EliminarLibro(c *gin.Context) {

	var libro libros.Libro
	if err := database.DB.Where("id=?", c.Param("id")).First(&libro).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Libro no encontrado"})
		return
	}

	//inicio de la transaccion
	tx := database.DB.Begin()
	err := tx.Delete(&libro).Error
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()
	//fin de la transaccion

	c.JSON(http.StatusOK, gin.H{"data": true})
}
