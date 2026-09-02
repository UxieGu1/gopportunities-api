package user

import "fmt"

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

type RegisterUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"` 
}

func (r *RegisterUserRequest) Validate() error {
	if r.Email == "" && r.Password == "" {
		return fmt.Errorf("request body is empty or malformed")
	}

	if r.Email == "" {
		return errParamIsRequired("email", "string")
	}
	if r.Password == "" {
		return errParamIsRequired("password", "string")
	}
	return nil
}

type UpdateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func (r *UpdateUserRequest) Validate() error {
	if r.Email != "" || r.Password != "" || r.Role != "" {
		return nil
	}
	return fmt.Errorf("at least one valid field must be provided")
}

type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *LoginUserRequest) Validate() error {
	if r.Email == "" {
		return errParamIsRequired("email", "string")
	}
	if r.Password == "" {
		return errParamIsRequired("password", "string")
	}
	return nil
}