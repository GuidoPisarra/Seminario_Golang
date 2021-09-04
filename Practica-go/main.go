package main

//gorilla mux es un enrutadorDefinir middleware en las rutas, es decir, aplicar funciones
//que se ejecutan antes de cada petición HTTP y que permiten detener la ejecución o loguear determinadas cosas.
// Definición de rutas con verbos HTTP. Lectura de parámetros GET
import (
	"encoding/json"
	"log"
	"net/http"

	//"log"
	"github.com/gorilla/mux"
)

//Defino Persona
type Person struct {
	// le digo que puede recibir en formato json y que no tenga en cuenta a los vacios
	ID        string `json:"id,omitempty"`
	FirstName string `json:"firstname,omitempty"`
	LastName  string `json:"lastname,omitempty"`
	// con el * le digo que apunte al tipo de dato Address
	Address *Address `json:"address,omitempty"`
}
type Address struct {
	State  string `json:"state,omitempty"`
	City   string `json:"city,omitempty"`
	Street string `json:"street,omitempty"`
	Nro    string `json:"nro,omitempty"`
}

//mock de prueba
var people []Person

// la w es de write es decir que puede "escribir " una respuesta
// response writer es el que devuelve el arreglo de datos
func GetPeopleEndPoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(people)
}
func GetPerson(w http.ResponseWriter, req *http.Request) {

}
func CreatePersonEndPoint(w http.ResponseWriter, req *http.Request) {

}

func DeletePersonEndPoint(w http.ResponseWriter, req *http.Request) {

}
func main() {
	//routes
	router := mux.NewRouter()

	people = append(people, Person{ID: "1", FirstName: "Juan", LastName: "Gomez", Address: &Address{State: "Buenos Aires", City: "Azul", Street: "San Martin", Nro: "888"}})
	people = append(people, Person{ID: "2", FirstName: "Pepe", LastName: "Grillo", Address: &Address{State: "Buenos Aires", City: "Tres Arroyos", Street: "Colon", Nro: "333"}})

	//endpoints
	router.HandleFunc("/people", GetPeopleEndPoint).Methods("GET")
	router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePersonEndPoint).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePersonEndPoint).Methods("DELETE")

	http.ListenAndServe(":3000", router)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
