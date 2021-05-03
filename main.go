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

	router.HandleFunc("/people", getPeople).Methods("GET")       // /people?limit=$1&offset=$2 (this will display all people with a max results of limit being offset by offset)
	router.HandleFunc("/search", searchPeople).Methods("GET")    // /search?q=$1&limit=$2&offset=$3 (this will query the database for results that include q in their name, max results limited by limit, being offset by offset)
	router.HandleFunc("/add", addPeople).Methods("POST")         // /add (adds new record into databse, expects information in the body with the following items: description, gender, coollevel, name)
	router.HandleFunc("/update", updatePerson).Methods("PUT")    // /update (updates an existing record in the datase, expects the following items in the body: id (this is the id of the record to be changed), description, gender, coollevel, name)
	router.HandleFunc("/remove", removePerson).Methods("DELETE") // /remove?id=$1 (removes a record with id $1)

	http.ListenAndServe(":8080", handlers.CORS(headers, methods, origins)(router))
}

type searchParams struct {
	Id     string
	Name   string
	Limit  string
	Offset string
}

func getPeople(w http.ResponseWriter, r *http.Request) {
	people = []Person{}
	params := r.URL.Query()

	limit, ok := params["limit"]
	if !ok {
		limit = append(limit, "50")
	}

	offset, ok := params["offset"]
	if !ok {
		offset = append(offset, "0")
	}

	paramsSearch := searchParams{
		Limit:  limit[0],
		Offset: offset[0],
	}

	rows, err := db.Query("select * from cool_people limit $1 offset $2", paramsSearch.Limit, paramsSearch.Offset)
	logUnFatal(err, w)

	loopOverRows(rows, w)
}

func searchPeople(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	q, ok := params["q"]
	if !ok {
		json.NewEncoder(w).Encode(errors.New("no values given"))
		return
	}

	limit, ok := params["limit"]
	if !ok {
		limit = append(limit, "50")
	}

	offset, ok := params["offset"]
	if !ok {
		offset = append(offset, "0")
	}

	paramsSearch := searchParams{
		Name:   q[0],
		Limit:  limit[0],
		Offset: offset[0],
	}

	query := []string{"select * from cool_people where upper(name) like upper('%", paramsSearch.Name, "%') limit ", paramsSearch.Limit, " offset ", paramsSearch.Offset}
	sqlQuery := strings.Join(query, "")

	rows, err := db.Query(sqlQuery)

	logUnFatal(err, w)

	loopOverRows(rows, w)
}

func addPeople(w http.ResponseWriter, r *http.Request) {
	var person Person
	var personID int

	json.NewDecoder(r.Body).Decode(&person)

	if person.Gender != "male" && person.Gender != "female" && person.Gender != "other" && person.Gender != "prefer not to say" {
		json.NewEncoder(w).Encode(errors.New("gender option not available yet"))
		return
	}

	err := db.QueryRow("insert into cool_people (description, gender, coollevel, name) values($1, $2, $3, $4) RETURNING id;",
		person.Description, person.Gender, person.CoolLevel, person.Name).Scan(&personID)
	logUnFatal(err, w)

	json.NewEncoder(w).Encode(personID)
}

func updatePerson(w http.ResponseWriter, r *http.Request) {
	var person Person
	json.NewDecoder(r.Body).Decode(&person)

	if person.ID == 1 {
		json.NewEncoder(w).Encode(errors.New("omer is a god, i wouldnt dare change his information"))
		return
	}
	if person.ID == 2 {
		json.NewEncoder(w).Encode(errors.New("carly is under omers protection and cant be changed"))
		return
	}

	result, err := db.Exec("update cool_people set description=$1, gender=$2, coollevel=$3, name=$4 where id=$5 RETURNING id;",
		&person.Description, &person.Gender, &person.CoolLevel, &person.Name, &person.ID)
	logUnFatal(err, w)

	rowsUpdated, err := result.RowsAffected()
	logUnFatal(err, w)

	json.NewEncoder(w).Encode(rowsUpdated)
}

func removePerson(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	id, ok := params["id"]
	if !ok {
		json.NewEncoder(w).Encode(errors.New("no id given"))
		return
	}

	paramsSearch := searchParams{
		Id: id[0],
	}
	if paramsSearch.Id == "1" {
		json.NewEncoder(w).Encode(errors.New("omer is a god, i wouldn't dare delete him"))
		return
	}
	if paramsSearch.Id == "2" {
		json.NewEncoder(w).Encode(errors.New("carly is under omers protection"))
		return
	}

	result, err := db.Exec("delete from cool_people where id=$1;", paramsSearch.Id)
	logFatal(err)

	rowsDeleted, err := result.RowsAffected()
	logFatal(err)

	json.NewEncoder(w).Encode(rowsDeleted)
}
