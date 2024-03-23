package handler

import "fmt"

type CreateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func errorParamsIsRequired(name, valueType string) error {
	return fmt.Errorf("param %s is required and its type should be %s", name, valueType)
}

func (r *CreateOpeningRequest) Validate() error {
	if r.Role == "" {
		return errorParamsIsRequired("role", "string")
	}

	if r.Company == "" {
		return errorParamsIsRequired("company", "string")
	}

	if r.Location == "" {
		return errorParamsIsRequired("location", "string")
	}

	if r.Remote == nil {
		return errorParamsIsRequired("remote", "bool")
	}

	if r.Link == "" {
		return errorParamsIsRequired("link", "string")
	}

	if r.Salary <= 0 {
		return errorParamsIsRequired("salary", "int64")
	}

	return nil
}
