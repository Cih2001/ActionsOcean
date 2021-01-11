package model

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

const (
	EmployeeServerAddr = "https://employees-api.vercel.app/api/"
)

var (
	ErrDoesNotExist = errors.New("does not exist")
	ErrServerError  = errors.New("internal server error")
	ErrEmptyID      = errors.New("empty id")
)

type EmployeeModel struct {
	ID         string `json:"id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	Department string `json:"department"`
	Role       string `json:"role"`
}

// GetAllEmployees returns a list of all employees.
func GetAllEmployees() (result []EmployeeModel, err error) {
	resp, err := http.Get(EmployeeServerAddr + "employees")
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	type ResponceModel struct {
		Data []EmployeeModel `json:"data"`
	}
	respModel := ResponceModel{}

	err = json.Unmarshal(body, &respModel)
	result = respModel.Data

	return
}

// Find finds an employee by its ID.
func (model *EmployeeModel) Find(id string) error {
	if id == "" {
		return ErrEmptyID
	}

	resp, err := http.Get(EmployeeServerAddr + "employees/" + id)
	if err != nil {
		return err
	}

	switch resp.StatusCode {
	case 200:
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		err = json.Unmarshal(body, model)
	case 404:
		err = ErrDoesNotExist
	case 500:
		fallthrough
	default:
		err = ErrServerError
	}

	return err
}

// FindEmployees gets a slice of ids and return a slice of respective EmployeeModels
func FindEmployees(ids []string) (result []EmployeeModel, err error) {
	// FindEmployees can easily be implemented with the use of GetAllEmployees function and
	// probably with a better performance. However, As I want to demonstrate how one can use
	// go routines to send out multiple requests concurrently, I have decided not to use
	// GetAllEmployees and send a single request for each individual.

	if len(ids) == 0 {
		err = ErrEmptyID
		return
	}

	maxRequestCount := 10 // TODO: must be in a config file
	type Response struct {
		ID       string
		Employee EmployeeModel
		Err      error
	}
	channel := make(chan Response, maxRequestCount)

	for _, id := range ids {
		go func(id string) {
			employee := EmployeeModel{}
			err := employee.Find(id)
			channel <- Response{id, employee, err}
		}(id)
	}

	for i := 0; i < len(ids); i++ {
		response := <-channel
		if response.Err != nil {
			err = errors.New("error finding employee " + response.ID)
			break
		}
		result = append(result, response.Employee)
	}

	return
}
