package models

type Editoriales struct {
	EditorialID uint      `gorm:"primaryKey"`
	Nombre      string    `gorm:"not null;unique;size:128;index"`
	Libros      []*Libros `gorm:"foreignKey:EditorialID;references:EditorialID"`
}
