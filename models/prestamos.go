package models

import "time"

type Prestamos struct {
	PrestamoID      uint       `gorm:"primaryKey"`
	UsuarioID       uint       `gorm:"not null;index"`
	Usuario         *Usuarios  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	LibroID         uint       `gorm:"not null;index"`
	Libro           *Libros    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	FechaPrestamo   *time.Time `gorm:"not null"`
	FechaDevolucion *time.Time
	Devuelto        bool `gorm:"not null;default:false"`
}
