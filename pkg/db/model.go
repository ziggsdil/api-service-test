package db

import "github.com/google/uuid"

type Person struct {
	Id          uuid.UUID `json:"id" postgres:"person_uuid"`
	Name        string    `json:"name" postgres:"name"`
	Surname     string    `json:"surname" postgres:"surname"`
	Patronymic  string    `json:"patronymic,omitempty" postgres:"patronymic"`
	Age         int       `json:"age,omitempty" postgres:"possible_age"`
	Gender      string    `json:"gender,omitempty" postgres:"possible_gender"`
	Nationality string    `json:"nationality,omitempty" postgres:"possible_nationality"`
}
