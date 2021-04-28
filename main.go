package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/lib/pq"
	"github.com/subosito/gotenv"
)

type Person struct {
	ID          int    `json:"id"`
	LastName    string `json:"lastname"`
	FirstName   string `json:"firstname"`
	Description string `json:"description"`
	Gender      string `json:"gender"`
	CoolLevel   bool   `json:"coollevel"`
}

var people []Person

var db *sql.DB

func init() {
	gotenv.Load() //loads all variable from env files
}

func logFatalDBconnect(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func logFatal(err error, w http.ResponseWriter) {
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
}

func main() {
	pgUrl, err := pq.ParseURL(os.Getenv("ELEPHANTSQL_URL"))
	logFatalDBconnect(err)

	db, err = sql.Open("postgres", pgUrl)
	logFatalDBconnect(err)

	err = db.Ping()
	logFatalDBconnect(err)

	router := mux.NewRouter()

	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})

	router.HandleFunc("/people", getPeople).Methods("GET")

	//log.Fatal(http.ListenAndServe(":8080", router))
	http.ListenAndServe(":8080", handlers.CORS(headers, methods, origins)(router))
}

func getPeople(w http.ResponseWriter, r *http.Request) {
	var person Person
	people = []Person{}

	rows, err := db.Query("select * from cool_people")
	logFatal(err, w)

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&person.ID, &person.LastName, &person.FirstName, &person.Description, &person.Gender, &person.CoolLevel)
		logFatal(err, w)

		people = append(people, person)
	}

	json.NewEncoder(w).Encode(people)
}
