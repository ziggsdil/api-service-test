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

	//updatePerson = `
	//	UPDATE person
	//	SET name = $1, surname = $2, age = $3
	//	WHERE person_uuid = $4;
	//`
	updatePerson = `
	UPDATE person SET Surname = $1, Age = $2 WHERE person_uuid = $3
`
)
