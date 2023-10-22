package handler

import (
	"encoding/json"
	"fmt"
	"github.com/ziggsdil/api-service-test/pkg/db"
	"io"
	"net/http"
)

const (
	ageApi      = "https://api.agify.io/?name=%s"
	genderApi   = "https://api.genderize.io/?name=%s"
	nationalApi = "https://api.nationalize.io/?name=%s"
)

func (h *Handler) Add(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var person db.Person

	err := h.parseBody(r.Body, &person)
	if err != nil {
		h.renderer.RenderError(w, err)
		return
	}

	person.Age, err = GetAge(person.Name)
	if err != nil {
		h.renderer.RenderError(w, err)
		return
	}
	person.Gender, err = GetGender(person.Name)
	if err != nil {
		h.renderer.RenderError(w, err)
		return
	}
	person.Nationality, err = GetNationality(person.Name)
	if err != nil {
		h.renderer.RenderError(w, err)
		return
	}

	err = h.db.AddPerson(ctx, &person)
	if err != nil {
		h.renderer.RenderError(w, err)
		return
	}
	// todo: log and info
	fmt.Printf("Person added: %+v", person)
}

type AgifyResponse struct {
	Count int    `json:"count"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
}

func GetAge(name string) (int, error) {
	var agifyResponse AgifyResponse
	err := fetchData(fmt.Sprintf(ageApi, name), &agifyResponse)
	if err != nil {
		return 0, err
	}
	return agifyResponse.Age, nil
}

type GenderizeResponse struct {
	Count       int     `json:"count"`
	Name        string  `json:"name"`
	Gender      string  `json:"gender"`
	Probability float64 `json:"probability"`
}

func GetGender(name string) (string, error) {
	var genderizeResponse GenderizeResponse
	err := fetchData(fmt.Sprintf(genderApi, name), &genderizeResponse)
	if err != nil {
		return "", err
	}
	return genderizeResponse.Gender, nil
}

type NationalityResponse struct {
	Count   int    `json:"count"`
	Name    string `json:"name"`
	Country []struct {
		CountryName string  `json:"country_id"`
		Probability float64 `json:"probability"`
	} `json:"country"`
}

func GetNationality(name string) (string, error) {
	var nationalityResponse NationalityResponse
	err := fetchData(fmt.Sprintf(nationalApi, name), &nationalityResponse)
	if err != nil {
		return "", err
	}
	return nationalityResponse.Country[0].CountryName, nil // так как ответ уже отсортирован по убыванию
}

func fetchData(apiUrl string, responseData interface{}) error {
	resp, err := http.Get(apiUrl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, responseData)
	if err != nil {
		return err
	}

	return nil
}
