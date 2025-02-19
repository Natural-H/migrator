package models

type Generos struct {
	GeneroID uint      `gorm:"primaryKey"`
	Nombre   string    `gorm:"not null;unique;size:128;index"`
	Libros   []*Libros `gorm:"many2many:libros_generos;joinForeignKey:GeneroID;joinReferences:LibroID"`
}
