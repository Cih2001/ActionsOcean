package model

import (
	"testing"
)

func TestGetAllEmployees(t *testing.T) {
	employees, err := GetAllEmployees()
	if err != nil {
		t.Fatal(err)
	}

	if len(employees) == 0 {
		t.Fatalf("empty responce from server.\n")
	}
}

func TestEmployeeFind(t *testing.T) {
	employee := EmployeeModel{}
	err := employee.Find("a0d5e87a-af04-473d-b1f5-3105bbf986c8")
	if err != nil && err != ErrDoesNotExist {
		t.Fatal(err)
	}
}

func TestEmployeesFind(t *testing.T) {
	employeeIDs := []string{
		"a0d5e87a-af04-473d-b1f5-3105bbf986c8",
		"ae18aa6f-49f0-443f-b7ff-20aae3729045",
		"e1f1dd6d-af29-430c-b74f-c346eb66ef7a",
		"a0d5e87a-af04-473d-b1f5-3105bbf986c8",
		"ae18aa6f-49f0-443f-b7ff-20aae3729045",
		"e1f1dd6d-af29-430c-b74f-c346eb66ef7a",
		"a0d5e87a-af04-473d-b1f5-3105bbf986c8",
		"ae18aa6f-49f0-443f-b7ff-20aae3729045",
		"e1f1dd6d-af29-430c-b74f-c346eb66ef7a",
	}

	result, err := FindEmployees(employeeIDs)
	if err != nil {
		t.Fatal(err)
	}
	if err == nil && len(result) != len(employeeIDs) {
		t.Fatalf("some ids are not retrieved")
	}

	employeeIDs = []string{
		"ae18aa6f-49f0-443f-b7ff-20aae3729040",
		"ae18aa6f-49f0-443f-b7ff-20aae3729045",
		"e1f1dd6d-af29-430c-b74f-c346eb66ef7a",
		"a0d5e87a-af04-473d-b1f5-3105bbf986c8",
		"ae18aa6f-49f0-443f-b7ff-20aae3729045",
		"e1f1dd6d-af29-430c-b74f-c346eb66ef7a",
		"a0d5e87a-af04-473d-b1f5-3105bbf986c8",
		"ae18aa6f-49f0-443f-b7ff-20aae3729045",
		"e1f1dd6d-af29-430c-b74f-c346eb66ef7a",
	}
	_, err = FindEmployees(employeeIDs)
	if err == nil {
		t.Fatal("expected an error")
	}
}
