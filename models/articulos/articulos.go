package articulos

//Articulo ...
type Articulo struct {
	ArticuloID int     `json:"idarticulo" gorm:"primary_key"`
	CatID      int     `json:"catid"`
	Nombre     string  `json:"nombre" gorm:"size:50"`
	Precio     float64 `json:"precio"`
	Cantidad   int     `json:"cantidad"`
	Total      float64 `json:"total"`
}

//Categoria ...
type Categoria struct {
	CategoriaID int        `json:"idcategoria" gorm:"primary_key"`
	Descripcion string     `json:"descripcion" gorm:"size:50"`
	Articulos   []Articulo `json:"articulos" gorm:"foreignkey:CatID"`
}

/*ArticuloCategoria ... es una estructura que permite heredar las dos estructuras que estan relacionadas,
permite hacer las consultas a las base de datos*/
type ArticuloCategoria struct {
	Articulo
	Categoria
}
