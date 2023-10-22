package db

const (
	initRequest = `
		CREATE TABLE IF NOT EXISTS person (
    	person_uuid UUID PRIMARY KEY,
    	name VARCHAR(255) NOT NULL,
    	surname VARCHAR(255) NOT NULL,
    	patronymic VARCHAR(255),
    	possible_age INT,
    	possible_gender VARCHAR(255),
    	possible_nationality VARCHAR(255)
	);
`

	dropRequest = `
		DROP TABLE IF EXISTS person;
	`

	deletePersonById = `
		DELETE FROM person WHERE person_uuid = $1;
	`

	addPerson = `
		INSERT INTO person (person_uuid, name, surname, patronymic, possible_age, possible_gender, possible_nationality)
		VALUES ($1, $2, $3, $4, $5, $6, $7);
	`

	updatePersonById = `
		UPDATE person
)
