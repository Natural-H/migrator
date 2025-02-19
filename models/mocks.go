package models

import (
	"github.com/go-faker/faker/v4"
	"github.com/go-faker/faker/v4/pkg/options"
	"gorm.io/gorm"
	"math/rand"
	"time"
)

func MockDB(db *gorm.DB) {
	autores := []*Autores{
		{Nombre: "George Orwell"},
		{Nombre: "J.K. Rowling"},
		{Nombre: "J.R.R. Tolkien"},
		{Nombre: "Jane Austen"},
		{Nombre: "F. Scott Fitzgerald"},
		{Nombre: "Harper Lee"},
		{Nombre: "Gabriel García Márquez"},
		{Nombre: "Leo Tolstoy"},
		{Nombre: "Mark Twain"},
		{Nombre: "Charles Dickens"},
		{Nombre: "Herman Melville"},
		{Nombre: "Mary Shelley"},
		{Nombre: "Aldous Huxley"},
		{Nombre: "Ernest Hemingway"},
		{Nombre: "Virginia Woolf"},
		{Nombre: "Fyodor Dostoevsky"},
		{Nombre: "Miguel de Cervantes"},
		{Nombre: "Homer"},
		{Nombre: "Dante Alighieri"},
		{Nombre: "William Shakespeare"},
	}

	db.Create(&autores)

	// Create some genres
	generos := []*Generos{
		{Nombre: "Science Fiction"},
		{Nombre: "Fantasy"},
		{Nombre: "Dystopian"},
		{Nombre: "Romance"},
		{Nombre: "Classic"},
		{Nombre: "Adventure"},
		{Nombre: "Gothic"},
		{Nombre: "Historical Fiction"},
		{Nombre: "Philosophical"},
		{Nombre: "Epic Poetry"},
	}

	db.Create(&generos)

	// Create some editorials
	editoriales := []*Editoriales{
		{Nombre: "Secker & Warburg"},
		{Nombre: "Bloomsbury"},
		{Nombre: "Allen & Unwin"},
		{Nombre: "T. Egerton, Whitehall"},
		{Nombre: "Charles Schrodinger's Sons"},
		{Nombre: "J.B. Lippincott & Co."},
		{Nombre: "Editorial Sudamericana"},
		{Nombre: "The Russian Messenger"},
		{Nombre: "Chatto & Windus"},
		{Nombre: "Chapman & Hall"},
		{Nombre: "Harper & Brothers"},
		{Nombre: "Lackington, Hughes, Harding, Mavor & Jones"},
		{Nombre: "Chatto & Sus"},
		{Nombre: "Charles Scribner's Sons"},
		{Nombre: "Hogarth Press"},
		{Nombre: "The Mexican Messenger"},
		{Nombre: "Francisco de Robles"},
		{Nombre: "Unknown (Ancient)"},
		{Nombre: "Niccolò di Lorenzo"},
		{Nombre: "First Folio"},
	}

	db.Create(&editoriales)

	// Create some books with their respective authors, genres, and editorials
	libros := []*Libros{
		{
			Titulo:          "1984",
			AnioPublicacion: "1949",
			Editorial:       editoriales[0],
			Autores:         []*Autores{autores[0]},
			Generos:         []*Generos{generos[2], generos[0]},
		},
		{
			Titulo:          "Harry Potter and the Philosopher's Stone",
			AnioPublicacion: "1997",
			Editorial:       editoriales[1],
			Autores:         []*Autores{autores[1]},
			Generos:         []*Generos{generos[1]},
		},
		{
			Titulo:          "The Hobbit",
			AnioPublicacion: "1937",
			Editorial:       editoriales[2],
			Autores:         []*Autores{autores[2]},
			Generos:         []*Generos{generos[1], generos[5]},
		},
		{
			Titulo:          "Pride and Prejudice",
			AnioPublicacion: "1813",
			Editorial:       editoriales[3],
			Autores:         []*Autores{autores[3]},
			Generos:         []*Generos{generos[3], generos[4]},
		},
		{
			Titulo:          "The Great Gatsby",
			AnioPublicacion: "1925",
			Editorial:       editoriales[4],
			Autores:         []*Autores{autores[4]},
			Generos:         []*Generos{generos[4], generos[8]},
		},
		{
			Titulo:          "To Kill a Mockingbird",
			AnioPublicacion: "1960",
			Editorial:       editoriales[5],
			Autores:         []*Autores{autores[5]},
			Generos:         []*Generos{generos[4], generos[8]},
		},
		{
			Titulo:          "One Hundred Years of Solitude",
			AnioPublicacion: "1967",
			Editorial:       editoriales[6],
			Autores:         []*Autores{autores[6]},
			Generos:         []*Generos{generos[8], generos[4]},
		},
		{
			Titulo:          "War and Peace",
			AnioPublicacion: "1869",
			Editorial:       editoriales[7],
			Autores:         []*Autores{autores[7]},
			Generos:         []*Generos{generos[8], generos[4]},
		},
		{
			Titulo:          "The Adventures of Huckleberry Finn",
			AnioPublicacion: "1884",
			Editorial:       editoriales[8],
			Autores:         []*Autores{autores[8]},
			Generos:         []*Generos{generos[5], generos[4]},
		},
		{
			Titulo:          "Great Expectations",
			AnioPublicacion: "1861",
			Editorial:       editoriales[9],
			Autores:         []*Autores{autores[9]},
			Generos:         []*Generos{generos[4], generos[8]},
		},
		{
			Titulo:          "Moby-Dick",
			AnioPublicacion: "1851",
			Editorial:       editoriales[10],
			Autores:         []*Autores{autores[10]},
			Generos:         []*Generos{generos[5], generos[4]},
		},
		{
			Titulo:          "Frankenstein",
			AnioPublicacion: "1818",
			Editorial:       editoriales[11],
			Autores:         []*Autores{autores[11]},
			Generos:         []*Generos{generos[6], generos[0]},
		},
		{
			Titulo:          "Brave New World",
			AnioPublicacion: "1932",
			Editorial:       editoriales[12],
			Autores:         []*Autores{autores[12]},
			Generos:         []*Generos{generos[0], generos[2]},
		},
		{
			Titulo:          "The Old Man and the Sea",
			AnioPublicacion: "1952",
			Editorial:       editoriales[13],
			Autores:         []*Autores{autores[13]},
			Generos:         []*Generos{generos[4], generos[8]},
		},
		{
			Titulo:          "Mrs Dalloway",
			AnioPublicacion: "1925",
			Editorial:       editoriales[14],
			Autores:         []*Autores{autores[14]},
			Generos:         []*Generos{generos[4], generos[8]},
		},
		{
			Titulo:          "Crime and Punishment",
			AnioPublicacion: "1866",
			Editorial:       editoriales[15],
			Autores:         []*Autores{autores[15]},
			Generos:         []*Generos{generos[8], generos[9]},
		},
		{
			Titulo:          "Don Quixote",
			AnioPublicacion: "1605",
			Editorial:       editoriales[16],
			Autores:         []*Autores{autores[16]},
			Generos:         []*Generos{generos[5], generos[4]},
		},
		{
			Titulo:          "The Iliad",
			AnioPublicacion: "800 BC",
			Editorial:       editoriales[17],
			Autores:         []*Autores{autores[17]},
			Generos:         []*Generos{generos[8], generos[9]},
		},
		{
			Titulo:          "The Divine Comedy",
			AnioPublicacion: "1320",
			Editorial:       editoriales[18],
			Autores:         []*Autores{autores[18]},
			Generos:         []*Generos{generos[8], generos[9]},
		},
		{
			Titulo:          "Hamlet",
			AnioPublicacion: "1603",
			Editorial:       editoriales[19],
			Autores:         []*Autores{autores[19]},
			Generos:         []*Generos{generos[4], generos[9]},
		},
	}

	db.Create(&libros)

	const registers = 15000
	const prestamos = 100000
	var users [registers]*Usuarios
	var lends [prestamos]*Prestamos

	for i := 0; i < registers; i++ {
		users[i] = &Usuarios{
			NombreUsuario: faker.Username(options.WithRandomStringLength(20)),
			Nombre:        faker.Name(),
			Email:         faker.Email(options.WithCustomDomain("fake.com")),
			Contrasenia:   faker.Password(),
		}
	}

	db.CreateInBatches(&users, min(registers/10, 30))

	for i := 0; i < prestamos; i++ {
		lendDate := RandomTimeBetweenSixMonths()
		returnDate := lendDate.AddDate(0, 1, 0)

		lends[i] = &Prestamos{
			Usuario:         users[rand.Uint64()%registers],
			Libro:           libros[rand.Uint64()%uint64(len(libros))],
			FechaPrestamo:   &lendDate,
			FechaDevolucion: &returnDate,
		}
	}

	db.CreateInBatches(&lends, min(prestamos/30, 30))
}

func RandomTimeBetweenSixMonths() time.Time {
	// Get the current time
	now := time.Now()

	// Calculate six months ago and six months from now
	sixMonthsAgo := now.AddDate(0, -6, 0)
	sixMonthsLater := now.AddDate(0, 6, 0)

	// Calculate the difference in seconds between the two times
	diff := sixMonthsLater.Unix() - sixMonthsAgo.Unix()

	// Generate a random number of seconds within the range
	randomSeconds := rand.Int63n(diff)

	// Add the random seconds to six months ago to get the random time
	randomTime := sixMonthsAgo.Add(time.Duration(randomSeconds) * time.Second)

	return randomTime
}
