package models

type LibrosGeneros struct {
	LibroGeneroID uint    `gorm:"primaryKey"`
	LibroID       uint    `gorm:"index:idx_libro_libro_id,unique"`
	Libro         Libros  `gorm:"references:LibroID"`
	GeneroID      uint    `gorm:"index:idx_libro_libro_id,unique"`
	Genero        Generos `gorm:"references:GeneroID"`
}
