package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lilianabarbosa15/APIHeros/handlers"
	"github.com/lilianabarbosa15/APIHeros/repository"
)

func main() {
	// Nueva base de datos:
	bd := repository.NewBaseDatos()
	handler := handlers.NewHandlerSuperheroes(bd) //asociar base de datos con el handler de comentarios

	// Se agrega multiplexor y enrutador:
	mux := mux.NewRouter()
	//asociar handlers al mux
	mux.HandleFunc("/api/superhero", handler.ListarSuperheroes()).Methods("GET")
	mux.HandleFunc("/api/superhero", handler.TraerSuperheroe()).Queries("hero", "{hero}").Methods("GET") //traer un solo superheroe

	// Definición de servidor que esté escuchando:
	log.Fatal(http.ListenAndServe(":8080", mux)) //Localhost y la ruta la toma el mux que se acaba de crear
}
