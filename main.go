package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/lib/pq"
	"github.com/subosito/gotenv"
)

type Person struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Gender      string `json:"gender"`
	CoolLevel   bool   `json:"coollevel"`
	Name        string `json:"name"`
}

var people []Person

var db *sql.DB

func init() {
	gotenv.Load() //loads all variable from env files
}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func logUnFatal(err error, w http.ResponseWriter) {
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
}

func loopOverRows(rows *sql.Rows, w http.ResponseWriter) {
	var person Person
	people = []Person{}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&person.ID, &person.Description, &person.Gender, &person.CoolLevel, &person.Name)
		logUnFatal(err, w)

		people = append(people, person)
	}

	json.NewEncoder(w).Encode(people)
}

func main() {
	pgUrl, err := pq.ParseURL(os.Getenv("ELEPHANTSQL_URL"))
	logFatal(err)

	db, err = sql.Open("postgres", pgUrl)
	logFatal(err)

	err = db.Ping()
	logFatal(err)

	router := mux.NewRouter()

	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})
	//handlers, models, repositories (for queries), pkg(package, this will include utils, and any shared code)

	router.HandleFunc("/people", getPeople).Methods("GET")
	router.HandleFunc("/search", searchPeople).Methods("GET")

	http.ListenAndServe(":8080", handlers.CORS(headers, methods, origins)(router))
}

func getPeople(w http.ResponseWriter, r *http.Request) {
	people = []Person{}

	rows, err := db.Query("select * from cool_people")
	logUnFatal(err, w)

	loopOverRows(rows, w)
}

type searchParams struct {
	Name string
}

func searchPeople(w http.ResponseWriter, r *http.Request) { //limit the query to a have a certain number, make it so the front end can choose how many records they get (using limit and offset)
	params := r.URL.Query()

	q, ok := params["q"]
	if !ok {
		json.NewEncoder(w).Encode(errors.New("no values given"))
		return
	}

	paramsSearch := searchParams{
		Name: q[0],
	}

	query := []string{"select * from cool_people where upper(name) like upper('%", paramsSearch.Name, "%')"}
	sqlQuery := strings.Join(query, "")

	rows, err := db.Query(sqlQuery)
	logUnFatal(err, w)

	loopOverRows(rows, w)
}
