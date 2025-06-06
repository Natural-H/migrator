package models

type LibrosGeneros struct {
	LibroGeneroID uint    `gorm:"primaryKey"`
	LibroID       uint    `gorm:"uniqueIndex:udx_libro_libro_id"`
	Libro         Libros  `gorm:"references:LibroID"`
	GeneroID      uint    `gorm:"uniqueIndex:udx_libro_libro_id"`
	Genero        Generos `gorm:"references:GeneroID"`
}
