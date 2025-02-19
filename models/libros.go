package models

type Libros struct {
	LibroID         uint   `gorm:"primaryKey"`
	Titulo          string `gorm:"not null;index;size:128"`
	AnioPublicacion string `gorm:"not null;index"`
	EditorialID     uint   `gorm:"not null;index"`
	Editorial       *Editoriales
	Autores         []*Autores `gorm:"many2many:libros_autores;joinForeignKey:LibroID;joinReferences:AutorID"`
	Generos         []*Generos `gorm:"many2many:libros_generos;joinForeignKey:LibroID;joinReferences:GeneroID"`
}
