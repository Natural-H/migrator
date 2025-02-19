package models

type Autores struct {
	AutorID uint      `gorm:"primaryKey"`
	Nombre  string    `gorm:"not null;index;size:128;unique"`
	Libros  []*Libros `gorm:"many2many:libros_autores;joinForeignKey:AutorID;joinReferences:LibroID"`
}
