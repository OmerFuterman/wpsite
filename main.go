package main

import (
	"net/http"
	"wpsite/controllers"
	"wpsite/driver"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load() //loads all variable from env files
}

func main() {
	db := driver.ConnectDB()
	controller := controllers.Controller{}

	router := mux.NewRouter()

	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})

	router.HandleFunc("/people", controller.GetPeople(db)).Methods("GET")       // /people?limit=$1&offset=$2 (this will display all people with a max results of limit being offset by offset)
	router.HandleFunc("/search", controller.SearchPeople(db)).Methods("GET")    // /search?q=$1&limit=$2&offset=$3 (this will query the database for results that include q in their name, max results limited by limit, being offset by offset)
	router.HandleFunc("/add", controller.AddPeople(db)).Methods("POST")         // /add (adds new record into database, expects information in the body with the following items: description, gender, coollevel, name)
	router.HandleFunc("/update", controller.UpdatePerson(db)).Methods("PUT")    // /update (updates an existing record in the datase, expects the following items in the body: id (this is the id of the record to be changed), description, gender, coollevel, name)
	router.HandleFunc("/remove", controller.RemovePerson(db)).Methods("DELETE") // /remove?id=$1 (removes a record with id $1)

	http.ListenAndServe(":8080", handlers.CORS(headers, methods, origins)(router))
}
