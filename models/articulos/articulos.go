package articulos

//Articulo ...
type Articulo struct {
	ArticuloID int     `json:"idarticulo" gorm:"primary_key"`
	CatID      int     `json:"catid" gorm:"not null"`
	Nombre     string  `json:"nombre" gorm:"size:50"`
	Precio     float64 `json:"precio"`
	Cantidad   int     `json:"cantidad"`
	Total      float64 `json:"total"`
}

//Categoria ...
type Categoria struct {
	CategoriaID int        `json:"idcategoria" gorm:"primary_key"`
	Descripcion string     `json:"descripcion" gorm:"size:50"`
	Articulos   []Articulo `json:"articulos" gorm:"Foreingkey:CatID"`
}

/*ArticuloCategoria ... es una estructura que permite heredar las dos estructuras que estan relacionadas,
permite hacer las consultas a las base de datos*/
type ArticuloCategoria struct {
	Articulo
	Categoria
}

//CrearArticuloInput ... estrucura con los campos que se van a validar para insertar
type CrearArticuloInput struct {
	Nombre   string  `json:"nombre" binding:"required"`
	CatID    int     `json:"catid,string" binding:"required"`
	Precio   float64 `json:"precio,string" binding:"required"`
	Cantidad int     `json:"cantidad,string" binding:"required"`
}

//ActualiarArticuloInput ... estrucura con los campos que se van a validar para actualizar
type ActualiarArticuloInput struct {
	Nombre   string  `json:"nombre" binding:"required"`
	Precio   float64 `json:"precio,string" binding:"required"`
	Cantidad int     `json:"cantidad,string" binding:"required"`
}

//CrearCategoriaInput ... estrucura con los campos que se van a validar para insertar
type CrearCategoriaInput struct {
	Descripcion string `json:"descripcion" binding:"required"`
}

//ActualizarCategoriaInput ... estrucura con los campos que se van a validar para insertar
type ActualizarCategoriaInput struct {
	Descripcion string `json:"descripcion" binding:"required"`
}
