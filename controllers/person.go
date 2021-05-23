package controllers

import (
	"database/sql"
	"encoding/json"
	"wpsite/models"
	personRepository "wpsite/repository/person"
	"wpsite/utils"

	"net/http"
)

type Controller struct{}

var people []models.Person

func (c Controller) GetPeople(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var person models.Person
		var error models.Error

		params := r.URL.Query()

		limit, ok := params["limit"]
		if !ok {
			limit = append(limit, "50")
		}

		offset, ok := params["offset"]
		if !ok {
			offset = append(offset, "0")
		}

		paramsSearch := models.SearchParams{
			Limit:  limit[0],
			Offset: offset[0],
		}

		people = []models.Person{}
		personRepo := personRepository.PersonRepository{}
		people, err := personRepo.GetPeople(db, person, people, paramsSearch)

		if err != nil {
			error.Message = "Server error"
			utils.SendError(w, http.StatusInternalServerError, error) //500
			return
		}

		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, people)
	}
}

func (c Controller) SearchPeople(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var person models.Person
		var error models.Error

		params := r.URL.Query()

		q, ok := params["q"]
		if !ok {
			error.Message = "No values given"
			utils.SendError(w, http.StatusInternalServerError, error) //500
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

		paramsSearch := models.SearchParams{
			Name:   q[0],
			Limit:  limit[0],
			Offset: offset[0],
		}

		people = []models.Person{}
		personRepo := personRepository.PersonRepository{}
		people, err := personRepo.SearchPeople(db, person, people, paramsSearch)

		if err != nil {
			error.Message = "Server error"
			utils.SendError(w, http.StatusInternalServerError, error) //500
			return
		}

		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, people)
	}
}

func (c Controller) AddPeople(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var person models.Person
		var personID int
		var error models.Error

		json.NewDecoder(r.Body).Decode(&person)

		// if person.Gender != "male" && person.Gender != "female" && person.Gender != "other" && person.Gender != "prefer not to say" {
		// 	json.NewEncoder(w).Encode(errors.New("gender option not available yet"))
		// 	return
		// }

		people = []models.Person{}
		personRepo := personRepository.PersonRepository{}
		personID, err := personRepo.AddPeople(db, person)

		if err != nil {
			error.Message = "Server error"
			utils.SendError(w, http.StatusInternalServerError, error) //500
			return
		}

		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, personID)
	}
}

func (c Controller) UpdatePerson(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var error models.Error
		var person models.Person

		json.NewDecoder(r.Body).Decode(&person)

		if person.ID == 1 {
			error.Message = "Omer is a god, I wouldnt dare change his information"
			utils.SendError(w, http.StatusBadRequest, error) //400
			return
		}
		if person.ID == 2 {
			error.Message = "Carly is under omer's protection and can't be changed"
			utils.SendError(w, http.StatusBadRequest, error) //400
			return
		}

		people = []models.Person{}
		personRepo := personRepository.PersonRepository{}
		personID, err := personRepo.UpdatePerson(db, person)

		if err != nil {
			error.Message = "Server error"
			utils.SendError(w, http.StatusInternalServerError, error) //500
			return
		}

		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, personID)
	}
}

func (c Controller) RemovePerson(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var error models.Error

		params := r.URL.Query()

		id, ok := params["id"]
		if !ok {
			error.Message = "No ID given"
			utils.SendError(w, http.StatusBadRequest, error) //400
			return
		}

		for _, v := range id {
			params := models.SearchParams{
				Id: v,
			}

			if params.Id == "1" {
				error.Message = "Omer is a god, I wouldn't dare delete him"
				utils.SendError(w, http.StatusBadRequest, error) //400
				return
			}
			if params.Id == "2" {
				error.Message = "Carly is under Omer's protection"
				utils.SendError(w, http.StatusBadRequest, error) //400
				return
			}

			people = []models.Person{}
			personRepo := personRepository.PersonRepository{}
			rowsDeleted, err := personRepo.RemovePerson(db, params)

			if err != nil {
				error.Message = "Server error"
				utils.SendError(w, http.StatusInternalServerError, error) //500
				return
			}

			if rowsDeleted == 0 {
				error.Message = "Not Found"
				utils.SendError(w, http.StatusNotFound, error) //404
				return
			}

			w.Header().Set("Content-Type", "application/json")
			utils.SendSuccess(w, rowsDeleted)
		}

	}
}
