package controllers

import (
	"net/http"
	//sgsd

	"github.com/frank1995alfredo/api/models"
	"github.com/gin-gonic/gin"
)

//ObtenerLibros ... controlador para obtener todos los libros
func ObtenerLibros(c *gin.Context) {
	var libros []models.Libro

	models.DB.Order("id").Find(&libros)

	c.Header("Access-Control-Allow-Origin", "*")

	c.JSON(http.StatusOK, gin.H{"data": libros})
}

//CrearLibro ... funcion para inserta un libro nuevo
func CrearLibro(c *gin.Context) {
	var input models.CrearLibroInput

	//para validar los inputs
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//crea el libro en la base de datos
	libro := models.Libro{Titulo: input.Titulo, Autor: input.Autor}
	models.DB.Create(&libro)

	c.Header("Access-Control-Allow-Origin", "*")

	c.JSON(http.StatusOK, gin.H{"data": libro})
}

//BuscarLibro ... funcion para buscar un libro
func BuscarLibro(c *gin.Context) {
	var libro models.Libro

	if err := models.DB.Where("id=?", c.Param("id")).First(&libro).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No encontrado"})
		return
	}

	c.Header("Access-Control-Allow-Origin", "*")

	c.JSON(http.StatusOK, gin.H{"data": libro})
}

//ActualizarLibro ... funcion para actualizar un libro :v
func ActualizarLibro(c *gin.Context) {
	var libro models.Libro

	if err := models.DB.Where("id = ?", c.Param("id")).First(&libro).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Libro no encontrado"})
		return
	}

	//validamos la entrada de los datos
	var input models.ActualizarLibro
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&libro).Updates(input)

	c.Header("Access-Control-Allow-Origin", "*")

	c.JSON(http.StatusOK, gin.H{"data": libro})
}

//EliminarLibro ... funcion que permite eliminar un libro
func EliminarLibro(c *gin.Context) {

	var libro models.Libro
	if err := models.DB.Where("id=?", c.Param("id")).First(&libro).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Libro no encontrado"})
		return
	}

	models.DB.Delete(&libro)

	c.Header("Access-Control-Allow-Origin", "*")

	c.JSON(http.StatusOK, gin.H{"data": true})
}
