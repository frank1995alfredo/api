package controllers

import (
	"net/http"

	"github.com/frank1995alfredo/api/models"
	"github.com/gin-gonic/gin"
)

//ObtenerLibros ... controlador para obtener todos los libros
func ObtenerLibros(c *gin.Context) {
	var libros []models.Libro
	models.DB.Find(&libros)

	c.JSON(http.StatusOK, gin.H{"data": libros})
}
