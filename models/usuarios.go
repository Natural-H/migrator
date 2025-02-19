package models

type Usuarios struct {
	UsuarioID     uint         `gorm:"primaryKey"`
	NombreUsuario string       `gorm:"not null;uniqueIndex;size:128"`
	Nombre        string       `gorm:"not null;size:128"`
	Email         string       `gorm:"unique;not null;size:128"`
	Contrasenia   string       `gorm:"not null;size:128"`
	Prestamos     []*Prestamos `gorm:"foreignKey:UsuarioID"`
}
