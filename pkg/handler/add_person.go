package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"

	"github.com/gookit/slog"

	"github.com/ziggsdil/api-service-test/pkg/db"
)

const (
	ageAPI      = "https://api.agify.io/?name=%s"
	genderAPI   = "https://api.genderize.io/?name=%s"
	nationalAPI = "https://api.nationalize.io/?name=%s"
)

func (h *Handler) Add(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	wg := &sync.WaitGroup{}

	var person db.Person

	err := h.parseBody(r.Body, &person)
	if err != nil {
		h.renderer.RenderError(w, err)
		return
	}

	var ageErr, genderErr, nationalityErr error
	errs := make([]error, 0, 3)
	wg.Add(1)
	go func() {
		defer wg.Done()
		slog.Infof("Starting goroutine api service age for %s", person.Name)
		person.Age, ageErr = GetAge(person.Name)
		slog.Infof("Goroutine is finished Person age: %d", person.Age)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		slog.Infof("Starting goroutine api service gender for %s", person.Name)
		person.Gender, genderErr = GetGender(person.Name)
		slog.Infof("Goroutine is finished Person gender: %s", person.Gender)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		slog.Infof("Starting goroutine api service nationality for %s", person.Name)
		person.Nationality, nationalityErr = GetNationality(person.Name)
		slog.Infof("Goroutine is finished Person nationality: %s", person.Nationality)
	}()
	errs = append(errs, ageErr, genderErr, nationalityErr)
	if h.handleError(w, errs...) {
		return
	}

	wg.Wait()

	err = h.db.AddPerson(ctx, &person)
	if err != nil {
		h.renderer.RenderError(w, err)
		return
	}
	slog.Infof("Person added: %+v", person)
	h.renderer.RenderOK(w)
}

func (h *Handler) handleError(w http.ResponseWriter, errs ...error) bool {
	slog.Errorf("Failed to add person: %v", errs[0].Error())
	for _, err := range errs {
		h.renderer.RenderError(w, err)
		return true
	}
	return false
}

type AgifyResponse struct {
	Count int    `json:"count"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
}

func GetAge(name string) (int, error) {
	var agifyResponse AgifyResponse
	err := fetchData(fmt.Sprintf(ageAPI, name), &agifyResponse)
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
	err := fetchData(fmt.Sprintf(genderAPI, name), &genderizeResponse)
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
	err := fetchData(fmt.Sprintf(nationalAPI, name), &nationalityResponse)
	if err != nil {
		return "", err
	}
	return nationalityResponse.Country[0].CountryName, nil // так как ответ уже отсортирован по убыванию
}

func fetchData(apiURL string, responseData interface{}) error {
	//nolint: gosec
	resp, err := http.Get(apiURL)
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
