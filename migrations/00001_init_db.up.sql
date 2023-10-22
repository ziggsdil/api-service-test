BEGIN ;

CREATE TABLE IF NOT EXISTS person (
    person_uuid UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    surname VARCHAR(255) NOT NULL,
    patronymic VARCHAR(255),
    possible_age INT,
    possible_gender VARCHAR(255),
    possible_nationality VARCHAR(255)
);

COMMIT;