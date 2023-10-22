package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
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
		// todo: log and info
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
	_, err := db.client.ExecContext(ctx, addPerson, person.Id, person.Name, person.Surname, person.Patronymic, person.Age, person.Gender, person.Nationality)
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
	// todo: если меняется имя, то необходимо менять и все остальные данные которые зависят от имени.

	query.WriteString(strings.Join(setClauses, ", "))
	query.WriteString(" WHERE person_uuid = $" + strconv.Itoa(idx))
	values = append(values, person.Id)
	fmt.Println(query.String(), values)

	_, err := db.client.ExecContext(ctx, query.String(), values...)
	return err
}

//func (db *Database) Update(ctx context.Context, person Person) error {
//	var query strings.Builder
//	query.WriteString("UPDATE person SET")
//
//	var values []interface{}
//	var setClauses []string
//
//	if person.Name != "" {
//		setClauses = append(setClauses, "name = ?")
//		values = append(values, person.Name)
//	}
//	if person.Surname != "" {
//		setClauses = append(setClauses, "surname = ?")
//		values = append(values, person.Surname)
//	}
//	if person.Patronymic != "" {
//		setClauses = append(setClauses, "patronymic = ?")
//		values = append(values, person.Patronymic)
//	}
//	if person.Age != 0 {
//		setClauses = append(setClauses, "age = ?")
//		values = append(values, person.Age)
//	}
//	if person.Gender != "" {
//		setClauses = append(setClauses, "gender = ?")
//		values = append(values, person.Gender)
//	}
//	if person.Nationality != "" {
//		setClauses = append(setClauses, "nationality = ?")
//		values = append(values, person.Nationality)
//	}
//
//	query.WriteString(strings.Join(setClauses, ", "))
//	query.WriteString(" WHERE person_uuid = ?")
//	values = append(values, person.Id)
//
//	_, err := db.client.ExecContext(ctx, query.String(), values...)
//	return err
//}

func (db *Database) Init(ctx context.Context) error {
	_, err := db.client.ExecContext(ctx, initRequest)
	return err
}

func (db *Database) Close() error {
	return db.client.Close()
}
