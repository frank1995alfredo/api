package main

import (
	"github.com/frank1995alfredo/api/controllers"
	"github.com/frank1995alfredo/api/models"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	models.ConectorBD()

	r.GET("/libros", controllers.ObtenerLibros)

	r.Run()
}
