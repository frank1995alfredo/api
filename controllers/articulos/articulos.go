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
	var artic []articulos.Articulo

	database.DB.Order("articulo_id").Find(&artic)

	c.Header("Access-Control-Allow-Origin", "*")

	c.SecureJSON(http.StatusOK, gin.H{"data": artic})
}

//CrearArticulo ... funcion para crear un articulos
func CrearArticulo(c *gin.Context) {

	//CrearArticuloInput ... estructura con los campos que se van a validar para insertar
	type CrearArticuloInput struct {
		CategoriaID int     `json:"idcategoria" binding:"required"`
		Nombre      string  `json:"nombre" binding:"required"`
		Precio      float64 `json:"precio" binding:"required"`
		Cantidad    int     `json:"cantidad" binding:"required"`
	}

	var input CrearArticuloInput
	//var catego = &articulos.Categoria{} //esto me ayuda a realizar una busqueda

	//validamoos los inputs
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//operacion matematica para calcular el total
	cantidad := input.Cantidad
	precio := input.Precio
	total := float64(cantidad) * precio

	articulo := articulos.Articulo{Nombre: input.Nombre, Precio: input.Precio,
		Cantidad: input.Cantidad, Total: total, CatID: input.CategoriaID}

	/*categoria := &articulos.Categoria{CategoriaID: input.CategoriaID,
	Articulos: []articulos.Articulo{{Nombre: input.Nombre,
		Precio:   input.Precio,
		Cantidad: input.Cantidad, Total: total, CatID: input.CategoriaID}}}*/

	//inicio de la transaction
	tx := database.DB.Begin()
	err := tx.Debug().Create(&articulo).Error //si no hay un error, se guarda el articulo
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()
	//fin de la transaction

	c.SecureJSON(http.StatusOK, gin.H{"data": articulo})
}

// BuscarArticulo ... funcion para buscar a un articulo
func BuscarArticulo(c *gin.Context) {
	var articulo articulos.Articulo

	if err := database.DB.Where("articulo_id=?", c.Param("id")).First(&articulo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No existe ese articulo"})
		return
	}

	c.SecureJSON(http.StatusOK, gin.H{"data": articulo})
}

//ActualizarArticulo ... funcion para actualizar persona
func ActualizarArticulo(c *gin.Context) {
	//ActualiarArticuloInput ... estrucura con los campos que se van a validar para actualizar
	type ActualiarArticuloInput struct {
		Nombre   string  `json:"nombre"`
		Precio   float64 `json:"precio"`
		CatID    int     `json:"catid"`
		Cantidad int     `json:"cantidad"`
	}

	var articulo articulos.Articulo
	var input ActualiarArticuloInput

	//validamos la entrada de los datos
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Where("articulo_id=?", c.Param("id")).First(&articulo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Articulo no encontrado"})
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

	c.SecureJSON(http.StatusOK, gin.H{"data": articulo})
}

//EliminarArticulo ... funcion que permite eliminar un libro
func EliminarArticulo(c *gin.Context) {
	var articulo articulos.Articulo

	if err := database.DB.Where("articulo_id=?", c.Param("id")).First(&articulo).Error; err != nil {
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

	c.SecureJSON(http.StatusOK, gin.H{"data": "Articulo eliminado"})
}

/**************CATEGORIA**************/

//ObtenerCategoria ... funcion para obtener todos las categorias
func ObtenerCategoria(c *gin.Context) {
	var categoria []articulos.Categoria

	database.DB.Order("categoria_id").Find(&categoria)

	c.SecureJSON(http.StatusOK, gin.H{"data": categoria})
}

//CrearCategoria  ... funcion para crear las categorias
func CrearCategoria(c *gin.Context) {
	//CrearCategoriaInput ... estrucura con los campos que se van a validar para insertar
	type CrearCategoriaInput struct {
		Descripcion string `json:"descripcion" binding:"required"`
	}

	var input CrearCategoriaInput

	//validaops los inputs
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Debe ingresar una descripcion"})
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

	c.SecureJSON(http.StatusOK, gin.H{"data": categoria})
}

// ActualizarCategoria ... metodo para actulizar una categoria
func ActualizarCategoria(c *gin.Context) {
	//CrearCategoriaInput ... estrucura con los campos que se van a validar para insertar
	type CrearCategoriaInput struct {
		Descripcion string `json:"descripcion" binding:"required"`
	}

	var categoria articulos.Categoria
	var input CrearCategoriaInput

	if err := database.DB.Where("id=?", c.Param("categoria_id")).First(&categoria).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Categoria no encontrado"})
	}

	cat := articulos.Categoria{Descripcion: input.Descripcion}

	//inicio de la transaccion
	tx := database.DB.Begin()
	err := tx.Model(&categoria).Update(cat).Error
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()
	//fin de la transaccion

	c.SecureJSON(http.StatusOK, gin.H{"data": categoria})

}

/***************OBTENER ARTICULO Y CATEGORIA*******************/

//PresentarArticuloCategoria ...
func PresentarArticuloCategoria(c *gin.Context) {

	//var catArticulo []articulos.ArticuloCategoria
	var categoria []articulos.Categoria

	/*database.DB.Table("articulos").Order("articulo_id").Select("articulos.articulo_id," +
	"articulos.nombre, articulos.precio, articulos.cantidad, articulos.total," +
	"categoria.descripcion").Joins("JOIN categoria ON articulos.cat_id = categoria.categoria_id").Find(&catArticulo) */
	database.DB.Debug().Order("categoria_id").Preload("Articulos").Find(&categoria)
	c.SecureJSON(http.StatusOK, gin.H{"data": categoria})

}
