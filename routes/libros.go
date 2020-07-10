package routes

import (
	"github.com/frank1995alfredo/api/controllers"
	"github.com/frank1995alfredo/api/models"
	"github.com/gin-gonic/gin"
)

//RutasLibros ... funcion para guardar las rutas
func RutasLibros() {
	r := gin.Default()

	models.ConectorBD()

	r.GET("/libros", controllers.ObtenerLibros)
	r.POST("/crearLibro", controllers.CrearLibro)
	r.GET("/buscarLibro/:id", controllers.BuscarLibro)
	r.PATCH("/actualizarLibro/:id", controllers.ActualizarLibro)
	r.DELETE("eliminarLibro/:id", controllers.EliminarLibro)
	r.Run()
}
