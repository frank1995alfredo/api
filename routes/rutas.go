package routes

import (
	"github.com/frank1995alfredo/api/controllers"
	"github.com/frank1995alfredo/api/controllers/articulo"

	_ "github.com/frank1995alfredo/api/controllers"          //fsdf
	_ "github.com/frank1995alfredo/api/controllers/articulo" // sfdsd

	//sdfsdf
	"github.com/frank1995alfredo/api/models"
	"github.com/gin-gonic/gin"
)

//Rutas ... funcion para guardar las rutas
func Rutas() {
	r := gin.Default()

	defer models.DB.Close()

	libro := r.Group("/libros")
	{

		libro.GET("/obtenerLibros", controllers.ObtenerLibros)
		libro.POST("/crearLibro", controllers.CrearLibro)
		libro.GET("/buscarLibro/:id", controllers.BuscarLibro)
		libro.PATCH("/actualizarLibro/:id", controllers.ActualizarLibro)
		libro.DELETE("/eliminarLibro/:id", controllers.EliminarLibro)
	}
	persona := r.Group("/personas")
	{
		persona.GET("/obtenerPersona", controllers.ObtenerPersona)
		persona.POST("/crearPersona", controllers.CrearPersona)
		persona.GET("/buscarPersona/:id", controllers.BuscarPersona)
		persona.PATCH("/actualizarPersona/:id", controllers.ActualizarPersona)
		persona.DELETE("/eliminarPersona/:id", controllers.EliminarPersona)
	}

	alumnos := r.Group("/alumnos")
	{
		alumnos.GET("/obtenerAlumno", controllers.ObtenerAlumnos)
		alumnos.POST("/crearAlumno", controllers.CrearAlumno)
	}

	maestros := r.Group("/maestros")
	{
		maestros.GET("/obtenerMaestro", controllers.ObtenerMaestros)
		maestros.POST("/crearMaestro", controllers.CrearMaestro)
		maestros.GET("/buscarMaestro/:id", controllers.BuscarMaestro)
		maestros.PATCH("/actualizarMaestro/:id", controllers.ActualizarMaestro)
	}

	categorias := r.Group("/categorias")
	{
		categorias.GET("/obtenerCategoria", articulo.ObtenerCategoria)
		categorias.POST("/crearCategoria", articulo.CrearCategoria)
	}

	articulos := r.Group("/articulos")
	{
		articulos.GET("/obtenerArticulo", articulo.ObtenerArticulos)
		articulos.GET("/listArticulo", articulo.PresentarArticuloCategoria)
		articulos.POST("/crearArticulo", articulo.CrearArticulo)
		articulos.GET("/buscarArticulo/:id", articulo.BuscarArticulo)
		articulos.PATCH("/actualizarArticulo/:id", articulo.ActualizarArticulo)
	}
	//RutasLibros()
	models.ConectorBD()
	r.Run()
}
