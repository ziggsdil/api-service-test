package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
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

// todo: возможно нужно обновлять отдельные части
//func (db *Database) Update(ctx context.Context, id string) error {
//	_, err := db.client.ExecContext(ctx, updatePersonById, id)
//	return nil
//}

func (db *Database) Init(ctx context.Context) error {
	_, err := db.client.ExecContext(ctx, initRequest)
	return err
}

func (db *Database) Close() error {
	return db.client.Close()
}
