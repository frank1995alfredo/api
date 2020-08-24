package articulos

import (
	"net/http"

	database "github.com/frank1995alfredo/api/database"
	articulos "github.com/frank1995alfredo/api/models/articulos"
	"github.com/gin-gonic/gin"
)

/*************ARTICULOS****************/

//ObtenerArticulos ... funcion para obtener todos los articulos
func ObtenerArticulos(c *gin.Context) {
	var articulos []articulos.Articulo

	database.DB.Order("articulo_id").Find(&articulos)

	c.Header("Access-Control-Allow-Origin", "*")

	c.JSON(http.StatusOK, gin.H{"data": articulos})
}

//CrearArticulo ... funcion para crear un articulos
func CrearArticulo(c *gin.Context) {

	var input articulos.CrearArticuloInput

	//validamoos los inputs
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//operacion matematica para calcular el total
	cantidad := input.Cantidad
	precio := input.Precio
	total := float64(cantidad) * precio

	//crea el articulo en la base de datos
	articulo := articulos.Articulo{Nombre: input.Nombre, Precio: input.Precio, Cantidad: input.Cantidad, Total: total, CatID: input.CatID}

	//inicio de la transaccion
	tx := database.DB.Begin()
	err := tx.Create(&articulo).Error //si no hay un error, se guarda el articulo
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()
	//fin de la transaccion

	//c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	c.JSON(http.StatusOK, gin.H{"data": articulo})

}

// BuscarArticulo ... funcion para buscar a un articulo
func BuscarArticulo(c *gin.Context) {
	var articulo articulos.Articulo

	if err := database.DB.Where("id=?", c.Param("id")).First(&articulo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No existe ese articulo"})
		return
	}

	c.Header("Access-Control-Allow-Origin", "*")

	c.JSON(http.StatusOK, gin.H{"data": articulo})
}

//ActualizarArticulo ... funcion para actualizar persona
func ActualizarArticulo(c *gin.Context) {
	var articulo articulos.Articulo

	if err := database.DB.Where("id=?", c.Param("id")).First(&articulo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Articulo no encontrado"})
	}

	c.Header("Access-Control-Allow-Origin", "*")
	//validamos la entrada de los datos
	var input articulos.ActualiarArticuloInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	cantidad := input.Cantidad
	precio := input.Precio
	total := float64(cantidad) * precio

	art := articulos.Articulo{Nombre: input.Nombre, Precio: input.Precio, Cantidad: input.Cantidad, Total: total}

	//inicio de la transaccion
	tx := database.DB.Begin()
	err := tx.Model(&articulo).Update(art).Error
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()
	//fin de la transaccion

	c.Header("Access-Control-Allow-Origin", "*")

	c.JSON(http.StatusOK, gin.H{"data": articulo})
}

//EliminarArticulo ... funcion que permite eliminar un libro
func EliminarArticulo(c *gin.Context) {
	var articulo articulos.Articulo

	if err := database.DB.Where("id=?", c.Param("id")).First(&articulo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Articulo no encontrado"})
		return
	}

	//inicio de la transaccion
	tx := database.DB.Begin()
	err := tx.Delete(&articulo).Error
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()
	//fin de la transaccion

	c.Header("Access-Control-Allow-Origin", "*")

	c.JSON(http.StatusOK, gin.H{"data": "Persona eliminada"})
}

/**************CATEGORIA**************/

//ObtenerCategoria ... funcion para obtener todos las categorias
func ObtenerCategoria(c *gin.Context) {
	var categoria []articulos.Categoria

	database.DB.Order("categoria_id").Select("descripcion").Find(&categoria)

	c.Header("Access-Control-Allow-Origin", "*")

	c.JSON(http.StatusOK, gin.H{"data": categoria})
}

//CrearCategoria  ... funcion para crear las categorias
func CrearCategoria(c *gin.Context) {
	var input articulos.CrearCategoriaInput

	//validaops los inputs
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//crea la persona en la base de datos
	categoria := articulos.Categoria{Descripcion: input.Descripcion}

	//inicio de la transaccion
	tx := database.DB.Begin()
	err := tx.Create(&categoria).Error
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()
	//fin de la transaccion

	c.Header("Access-Control-Allow-Origin", "*")

	c.JSON(http.StatusOK, gin.H{"data": categoria})
}

/***************OBTENER ARTICULO Y CATEGORIA*******************/

//PresentarArticuloCategoria ...
func PresentarArticuloCategoria(c *gin.Context) {

	var catArticulo []articulos.ArticuloCategoria

	database.DB.Table("articulos").Order("articulo_id").Select("articulos.nombre, articulos.precio, articulos.cantidad, articulos.total," +
		"categoria.descripcion").Joins("JOIN categoria ON articulos.cat_id = categoria.categoria_id").Find(&catArticulo)

	c.Header("Access-Control-Allow-Origin", "*")

	c.SecureJSON(http.StatusOK, gin.H{"data": catArticulo})

}
