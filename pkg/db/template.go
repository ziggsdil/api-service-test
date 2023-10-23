package db

const (
	initRequest = `
		CREATE TABLE IF NOT EXISTS person (
    	person_uuid UUID PRIMARY KEY,
    	name VARCHAR(255) NOT NULL,
    	surname VARCHAR(255) NOT NULL,
    	patronymic VARCHAR(255),
    	age INT,
    	gender VARCHAR(255),
    	nationality VARCHAR(255)
	);
`

	dropRequest = `
		DROP TABLE IF EXISTS person;
	`

	deletePersonById = `
		DELETE FROM person WHERE person_uuid = $1;
	`

	addPerson = `
		INSERT INTO person (person_uuid, name, surname, patronymic, age, gender, nationality)
		VALUES ($1, $2, $3, $4, $5, $6, $7);
	`

	selectPersonById = `
		SELECT person_uuid, name, surname, patronymic, age, gender, nationality
		FROM person
		WHERE person_uuid = $1
	`

	selectPeople = `
		SELECT person_uuid, name, surname, patronymic, age, gender, nationality
		FROM person
	`

	selectPeopleByAge = `
		SELECT person_uuid, name, surname, patronymic, age, gender, nationality
		FROM person
		WHERE age = $1
	`

	selectPeopleByGender = `
		SELECT person_uuid, name, surname, patronymic, age, gender, nationality
		FROM person
		WHERE gender = $1
	`

	selectPeopleByNationality = `
		SELECT person_uuid, name, surname, patronymic, age, gender, nationality
		FROM person
		WHERE nationality = $1
	`

	selectPeopleByName = `
		SELECT person_uuid, name, surname, patronymic, age, gender, nationality
		FROM person
		WHERE name = $1
	`
)
