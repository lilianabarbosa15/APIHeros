package repository

import (
	models "github.com/lilianabarbosa15/APIHeros/models"
)

// Persistencia en memoria (estructura de datos que se puede almacenar de los comentarios que se pueden leer):
type BaseDatos struct {
	Memoria map[string]models.Superheroe
}

// Función para crear un nuevo repositorio:
// con esto se devuelve base de datos inicializada
func NewBaseDatos() *BaseDatos {
	return &BaseDatos{
		Memoria: map[string]models.Superheroe{
			"Wolverine": {
				Nombre: "Wolverine",
				Biography: struct {
					FullName string "json:\"fullName\""
				}{"Bruce Wayne"},
				Powerstats: struct {
					Intelligence int "json:\"intelligence\""
					Strength     int "json:\"strength\""
					Speed        int "json:\"speed\""
					Durability   int "json:\"durability\""
					Power        int "json:\"power\""
					Combat       int "json:\"combat\""
				}{63, 32, 50, 100, 89, 100},
				Images: struct {
					Xs string "json:\"xs\""
					Sm string "json:\"sm\""
					Md string "json:\"md\""
					Lg string "json:\"lg\""
				}{"https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/xs/717-wolverine.jpg",
					"https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/sm/717-wolverine.jpg",
					"https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/md/717-wolverine.jpg",
					"https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/lg/717-wolverine.jpg"},
			},
			"Spider-Man": {
				Nombre: "Spider-Man",
				Biography: struct {
					FullName string "json:\"fullName\""
				}{"Peter Parker"},
				Powerstats: struct {
					Intelligence int "json:\"intelligence\""
					Strength     int "json:\"strength\""
					Speed        int "json:\"speed\""
					Durability   int "json:\"durability\""
					Power        int "json:\"power\""
					Combat       int "json:\"combat\""
				}{90, 55, 67, 75, 74, 85},
				Images: struct {
					Xs string "json:\"xs\""
					Sm string "json:\"sm\""
					Md string "json:\"md\""
					Lg string "json:\"lg\""
				}{"https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/xs/620-spider-man.jpg",
					"https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/sm/620-spider-man.jpg",
					"https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/md/620-spider-man.jpg",
					"https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/lg/620-spider-man.jpg"},
			},
			"Iron Man": {
				Nombre: "Iron Man",
				Biography: struct {
					FullName string "json:\"fullName\""
				}{"Tony Stark"},
				Powerstats: struct {
					Intelligence int "json:\"intelligence\""
					Strength     int "json:\"strength\""
					Speed        int "json:\"speed\""
					Durability   int "json:\"durability\""
					Power        int "json:\"power\""
					Combat       int "json:\"combat\""
				}{100, 85, 58, 85, 100, 64},
				Images: struct {
					Xs string "json:\"xs\""
					Sm string "json:\"sm\""
					Md string "json:\"md\""
					Lg string "json:\"lg\""
				}{"https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/xs/346-iron-man.jpg",
					"https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/sm/346-iron-man.jpg",
					"https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/md/346-iron-man.jpg",
					"https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/lg/346-iron-man.jpg"},
			},
			"Black Widow": {
				Nombre: "Black Widow",
				Biography: struct {
					FullName string "json:\"fullName\""
				}{"Natasha Romanoff"},
				Powerstats: struct {
					Intelligence int "json:\"intelligence\""
					Strength     int "json:\"strength\""
					Speed        int "json:\"speed\""
					Durability   int "json:\"durability\""
					Power        int "json:\"power\""
					Combat       int "json:\"combat\""
				}{75, 13, 33, 30, 36, 100},
				Images: struct {
					Xs string "json:\"xs\""
					Sm string "json:\"sm\""
					Md string "json:\"md\""
					Lg string "json:\"lg\""
				}{"https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/xs/107-black-widow.jpg",
					"https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/sm/107-black-widow.jpg",
					"https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/md/107-black-widow.jpg",
					"https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/lg/107-black-widow.jpg"},
			},
			"Thor": {
				Nombre: "Thor",
				Biography: struct {
					FullName string "json:\"fullName\""
				}{"Thor Odinson"},
				Powerstats: struct {
					Intelligence int "json:\"intelligence\""
					Strength     int "json:\"strength\""
					Speed        int "json:\"speed\""
					Durability   int "json:\"durability\""
					Power        int "json:\"power\""
					Combat       int "json:\"combat\""
				}{69, 100, 83, 100, 100, 100},
				Images: struct {
					Xs string "json:\"xs\""
					Sm string "json:\"sm\""
					Md string "json:\"md\""
					Lg string "json:\"lg\""
				}{"https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/xs/659-thor.jpg",
					"https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/sm/659-thor.jpg",
					"https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/md/659-thor.jpg",
					"https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/lg/659-thor.jpg"},
			},
		},
	}
}
