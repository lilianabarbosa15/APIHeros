package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	models "github.com/lilianabarbosa15/APIHeros/models"
	repository "github.com/lilianabarbosa15/APIHeros/repository"
)

// Otra estructura a la que se asocia métodos:
type HandlerSuperheroes struct {
	BD *repository.BaseDatos
}

// Función que crea Handlers
func NewHandlerSuperheroes(bd *repository.BaseDatos) *HandlerSuperheroes {
	return &HandlerSuperheroes{BD: bd} //puntero de Handler de comentarios
}

// Función para traer todos los comentarios en la base de datos: //GET
func (hc *HandlerSuperheroes) ListarSuperheroes() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { //función anonima que recibe ciertos parametros de entrada
		superheroes := []models.Superheroe{} //slice vacío en un principio
		//names := []string{}
		for _, superheroe := range hc.BD.Memoria { //name, superheroe	//se recorre el comentario de la base de datos de la memoria
			//names = append(names, name)
			superheroes = append(superheroes, superheroe)
		}
		//jsonComent, err := json.Marshal(names)
		jsonComent, err := json.Marshal(superheroes)
		if err != nil {
			http.Error(w, "fallo al codificar en json", http.StatusInternalServerError) //Error 5xx
		}
		w.Header().Set("Content-Type", "application/json") //comentario convertido a json
		w.WriteHeader(http.StatusOK)
		w.Write(jsonComent) //.Write(), recibe arreglo de bytes
	})
}

// Función para traer todos un solo superheroe de la base de datos
func (hc *HandlerSuperheroes) TraerSuperheroe() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r) //ident := r.PathValue("nombre") //w.Write([]byte(ident))
		hero := vars["hero"]
		//superheroes := []models.Superheroe{}
		for _, superheroe := range hc.BD.Memoria { //name, superheroe	//se recorre cada fila de la base de datos de la memoria
			//fmt.Println(hero)
			if hero == superheroe.Nombre {
				//superheroes = append(superheroes, superheroe)
				jsonComent, err := json.Marshal(superheroe) //superheroes
				if err != nil {
					http.Error(w, "fallo al codificar en json", http.StatusInternalServerError) //Error 5xx
				}
				w.Header().Set("Content-Type", "application/json") //comentario convertido a json
				w.WriteHeader(http.StatusOK)
				w.Write(jsonComent) //.Write(), recibe arreglo de bytes
			}
		}
		http.Error(w, "recurso no encontrado", http.StatusNotFound) //Error 4xx
	})
}
