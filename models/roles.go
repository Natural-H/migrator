package models

type Roles struct {
	RolID uint   `gorm:"primaryKey"`
	Name  string `gorm:"not null;uniqueIndex;size:128"`
}

var Rols = []Roles{
	{Name: "Usuario"},
	{Name: "Administrador"},
}
