package db

import "github.com/google/uuid"

type Person struct {
	Id          uuid.UUID `json:"id" postgres:"person_uuid"`
	Name        string    `json:"name" postgres:"name"`
	Surname     string    `json:"surname" postgres:"surname"`
	Patronymic  string    `json:"patronymic,omitempty" postgres:"patronymic"`
	Age         int       `json:"age,omitempty" postgres:"age"`
	Gender      string    `json:"gender,omitempty" postgres:"gender"`
	Nationality string    `json:"nationality,omitempty" postgres:"nationality"`
}
