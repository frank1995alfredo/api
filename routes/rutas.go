package routes

import (
	alumnos "github.com/frank1995alfredo/api/controllers/alumnos"
	articulos "github.com/frank1995alfredo/api/controllers/articulos"
	libros "github.com/frank1995alfredo/api/controllers/libros"
	maestros "github.com/frank1995alfredo/api/controllers/maestros"
	personas "github.com/frank1995alfredo/api/controllers/personas"

	//sdfsdf
	database "github.com/frank1995alfredo/api/database"

	config "github.com/frank1995alfredo/api/config"
	"github.com/gin-gonic/gin"
)

//Rutas ... funcion para guardar las rutas
func Rutas() {
	r := gin.Default()
	r.Use(config.CORS)

	//RutasLibros()
	database.ConectorBD()
	defer database.DB.Close()

	libro := r.Group("/libros")
	{

		libro.GET("/obtenerLibros", libros.ObtenerLibros)
		libro.POST("/crearLibro", libros.CrearLibro)
		libro.GET("/buscarLibro/:id", libros.BuscarLibro)
		libro.PATCH("/actualizarLibro/:id", libros.ActualizarLibro)
		libro.DELETE("/eliminarLibro/:id", libros.EliminarLibro)
	}
	persona := r.Group("/personas")
	{
		persona.GET("/obtenerPersona", personas.ObtenerPersona)
		persona.POST("/crearPersona", personas.CrearPersona)
		persona.GET("/buscarPersona/:id", personas.BuscarPersona)
		persona.PATCH("/actualizarPersona/:id", personas.ActualizarPersona)
		persona.DELETE("/eliminarPersona/:id", personas.EliminarPersona)
	}

	alumno := r.Group("/alumnos")
	{
		alumno.GET("/obtenerAlumno", alumnos.ObtenerAlumnos)
		alumno.POST("/crearAlumno", alumnos.CrearAlumno)
	}

	maestro := r.Group("/maestros")
	{
		maestro.GET("/obtenerMaestro", maestros.ObtenerMaestros)
		maestro.POST("/crearMaestro", maestros.CrearMaestro)
		maestro.GET("/buscarMaestro/:id", maestros.BuscarMaestro)
		maestro.PATCH("/actualizarMaestro/:id", maestros.ActualizarMaestro)
	}

	categorias := r.Group("/categorias")
	{
		categorias.GET("/obtenerCategoria", articulos.ObtenerCategoria)
		categorias.POST("/crearCategoria", articulos.CrearCategoria)
	}

	articulo := r.Group("/articulos")
	{
		articulo.GET("/obtenerArticulo", articulos.ObtenerArticulos)
		articulo.GET("/listArticuloCat", articulos.PresentarArticuloCategoria)
		articulo.POST("/crearArticulo", articulos.CrearArticulo)
		articulo.POST("/eliminarArticulo/:id", articulos.EliminarArticulo)
		articulo.GET("/buscarArticulo/:id", articulos.BuscarArticulo)
		articulo.PATCH("/actualizarArticulo/:id", articulos.ActualizarArticulo)
	}
	r.Run()
}
