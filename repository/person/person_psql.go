package personRepository

import (
	"database/sql"
	"strings"
	"wpsite/models"
)

type PersonRepository struct{}

func (b PersonRepository) GetPeople(db *sql.DB, person models.Person, people []models.Person, paramsSearch models.SearchParams) ([]models.Person, error) {

	rows, err := db.Query("select * from cool_people limit $1 offset $2", paramsSearch.Limit, paramsSearch.Offset)

	if err != nil {
		return []models.Person{}, err
	}

	for rows.Next() {
		err := rows.Scan(&person.ID, &person.Description, &person.Gender, &person.CoolLevel, &person.Name)
		if err != nil {
			return []models.Person{}, err
		}

		people = append(people, person)
	}

	return people, nil
}

func (b PersonRepository) SearchPeople(db *sql.DB, person models.Person, people []models.Person, paramsSearch models.SearchParams) ([]models.Person, error) {
	query := []string{"select * from cool_people where upper(name) like upper('%", paramsSearch.Name, "%') limit ", paramsSearch.Limit, " offset ", paramsSearch.Offset}
	sqlQuery := strings.Join(query, "")

	rows, err := db.Query(sqlQuery)

	if err != nil {
		return []models.Person{}, err
	}

	for rows.Next() {
		err := rows.Scan(&person.ID, &person.Description, &person.Gender, &person.CoolLevel, &person.Name)
		if err != nil {
			return []models.Person{}, err
		}

		people = append(people, person)
	}

	return people, nil
}

func (b PersonRepository) AddPeople(db *sql.DB, person models.Person) (int, error) {
	err := db.QueryRow("insert into cool_people (description, gender, coollevel, name) values($1, $2, $3, $4) RETURNING id;",
		person.Description, person.Gender, person.CoolLevel, person.Name).Scan(&person.ID)
	if err != nil {
		return 0, err
	}

	return person.ID, err
}

func (b PersonRepository) UpdatePerson(db *sql.DB, person models.Person) (int64, error) {
	result, err := db.Exec("update cool_people set description=$1, gender=$2, coollevel=$3, name=$4 where id=$5 RETURNING id;",
		&person.Description, &person.Gender, &person.CoolLevel, &person.Name, &person.ID)
	if err != nil {
		return 0, err
	}

	rowsUpdated, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsUpdated, err
}

func (b PersonRepository) RemovePerson(db *sql.DB, params models.SearchParams) (int64, error) {
	result, err := db.Exec("delete from cool_people where id=$1;", params.Id)
	if err != nil {
		return 0, err
	}

	rowsDeleted, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsDeleted, err
}
