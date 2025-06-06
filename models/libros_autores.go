package models

type LibrosAutores struct {
	LibroAutorID uint    `gorm:"primaryKey"`
	LibroID      uint    `gorm:"uniqueIndex:udx_libro_autor_id"`
	Libro        Libros  `gorm:"references:LibroID"`
	AutorID      uint    `gorm:"uniqueIndex:udx_autor_id"`
	Autor        Autores `gorm:"references:AutorID"`
}
