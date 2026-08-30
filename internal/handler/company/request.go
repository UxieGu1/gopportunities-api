package company

import "fmt"

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

type CreateCompanyRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Website     string `json:"website"`
	Email       string `json:"email"`
	Location    string `json:"location"`
}

func (r *CreateCompanyRequest) Validate() error {
	if r.Name == "" && r.Email == "" {
		return fmt.Errorf("request body is empty or malformed")
	}

	if r.Name == "" {
		return errParamIsRequired("name", "string")
	}
	if r.Email == "" {
		return errParamIsRequired("email", "string")
	}
	return nil
}

type UpdateCompanyRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Website     string `json:"website"`
	Email       string `json:"email"`
	Location    string `json:"location"`
}

func (r *UpdateCompanyRequest) Validate() error {
	if r.Name != "" || r.Description != "" || r.Website != "" || r.Email != "" || r.Location != "" {
		return nil
	}
	return fmt.Errorf("at least one valid field must be provided")
}