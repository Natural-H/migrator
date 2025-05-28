package models

type Libros struct {
	LibroID         uint         `gorm:"primaryKey"`
	Titulo          string       `gorm:"not null;index;size:128;unique"`
	AnioPublicacion string       `gorm:"not null;index"`
	EditorialID     *uint        `gorm:"index"`
	Editorial       *Editoriales `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	PortadaURL      string       `gorm:"size:500;default:'https://covers.openlibrary.org/b/olid/OL51694024M-M.jpg'"`
	Cantidad        uint         `gorm:"default:100"`
	Autores         []*Autores   `gorm:"many2many:libros_autores;joinForeignKey:LibroID;joinReferences:AutorID"`
	Generos         []*Generos   `gorm:"many2many:libros_generos;joinForeignKey:LibroID;joinReferences:GeneroID"`
}
