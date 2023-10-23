package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/gookit/slog"
	"reflect"
	"strconv"
	"strings"
)

type Database struct {
	client *sql.DB
}

func NewDatabase(cfg Config) (*Database, error) {
	connInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database,
	)

	client, err := sql.Open("postgres", connInfo)
	if err != nil {
		return nil, err
	}
	return &Database{
		client: client,
	}, nil
}

func (db *Database) Delete(ctx context.Context, id string) error {
	_, err := db.client.ExecContext(ctx, deletePersonById, id)
	return err
}

func (db *Database) AddPerson(ctx context.Context, person *Person) error {
	person.Id = uuid.New()
	_, err := db.client.ExecContext(ctx, addPerson,
		person.Id, person.Name, person.Surname, person.Patronymic, person.Age, person.Gender, person.Nationality)
	return err
}

func (db *Database) UserById(ctx context.Context, id string) (*Person, error) {
	var person Person
	err := db.client.QueryRowContext(ctx, selectPersonById, id).Scan(
		&person.Id, &person.Name, &person.Surname, &person.Patronymic, &person.Age, &person.Gender, &person.Nationality,
	)
	if err != nil {
		return nil, err
	}
	return &person, nil
}

func (db *Database) Update(ctx context.Context, person Person) error {
	var query strings.Builder
	query.WriteString("UPDATE person SET")
	query.WriteString(" ")

	var values []interface{}
	var setClauses []string

	v := reflect.ValueOf(person)
	t := v.Type()

	idx := 1
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		if !value.IsZero() && field.Name != "Id" {
			setClauses = append(setClauses, field.Name+" = $"+strconv.Itoa(idx))
			idx++
			values = append(values, value.Interface())
		}
	}

	query.WriteString(strings.Join(setClauses, ", "))
	query.WriteString(" WHERE person_uuid = $" + strconv.Itoa(idx))
	values = append(values, person.Id)
	slog.Infof("Request for update person with id: %s with values: %+v", person.Id, values)

	_, err := db.client.ExecContext(ctx, query.String(), values...)
	return err
}

func (db *Database) Users(ctx context.Context) ([]*Person, error) {
	var people []*Person
	rows, err := db.client.QueryContext(ctx, selectPeople)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var person Person
		err = rows.Scan(
			&person.Id, &person.Name, &person.Surname, &person.Patronymic, &person.Age, &person.Gender, &person.Nationality,
		)
		if err != nil {
			return nil, err
		}
		people = append(people, &person)
	}
	return people, nil
}

func (db *Database) Init(ctx context.Context) error {
	_, err := db.client.ExecContext(ctx, initRequest)
	return err
}

func (db *Database) queryUsers(ctx context.Context, query string, arg string) ([]*Person, error) {
	var people []*Person
	rows, err := db.client.QueryContext(ctx, query, arg)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var person Person
		err = rows.Scan(
			&person.Id, &person.Name, &person.Surname, &person.Patronymic, &person.Age, &person.Gender, &person.Nationality,
		)
		if err != nil {
			return nil, err
		}
		people = append(people, &person)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return people, nil
}

func (db *Database) UsersByAge(ctx context.Context, age string) ([]*Person, error) {
	return db.queryUsers(ctx, selectPeopleByAge, age)
}

func (db *Database) UsersByGender(ctx context.Context, gender string) ([]*Person, error) {
	return db.queryUsers(ctx, selectPeopleByGender, gender)
}

func (db *Database) UsersByNationality(ctx context.Context, nationality string) ([]*Person, error) {
	return db.queryUsers(ctx, selectPeopleByNationality, nationality)
}

func (db *Database) UsersByName(ctx context.Context, surname string) ([]*Person, error) {
	return db.queryUsers(ctx, selectPeopleByName, surname)
}
