package models

import "time"

type Prestamos struct {
	PrestamoID      uint `gorm:"primaryKey"`
	UsuarioID       uint `gorm:"not null;index"`
	Usuario         *Usuarios
	LibroID         uint `gorm:"not null;index"`
	Libro           *Libros
	FechaPrestamo   *time.Time `gorm:"not null"`
	FechaDevolucion *time.Time
}
