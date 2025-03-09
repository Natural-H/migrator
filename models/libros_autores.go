package models

type LibrosAutores struct {
	LibroAutorID uint    `gorm:"primaryKey"`
	LibroID      uint    `gorm:"index:idx_libro_autor_id"`
	Libro        Libros  `gorm:"references:LibroID"`
	AutorID      uint    `gorm:"index:idx_autor_id"`
	Autor        Autores `gorm:"references:AutorID"`
}
