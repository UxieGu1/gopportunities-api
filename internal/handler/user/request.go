package user

import (
	"fmt"
	"net/mail"
	"strings"
)

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

type RegisterUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func (r *RegisterUserRequest) Validate() error {
	r.Email = strings.TrimSpace(strings.ToLower(r.Email))
	r.Role = strings.ToUpper(strings.TrimSpace(r.Role))

	if r.Email == "" && r.Password == "" {
		return fmt.Errorf("request body is empty or malformed")
	}

	if r.Email == "" {
		return errParamIsRequired("email", "string")
	}
	if r.Password == "" {
		return errParamIsRequired("password", "string")
	}
	if _, err := mail.ParseAddress(r.Email); err != nil {
		return fmt.Errorf("email must be valid")
	}
	if len(r.Password) < 8 {
		return fmt.Errorf("password must contain at least 8 characters")
	}
	if r.Role != "" && r.Role != "COMPANY" && r.Role != "CANDIDATE" {
		return fmt.Errorf("role must be COMPANY or CANDIDATE")
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
	r.Email = strings.TrimSpace(strings.ToLower(r.Email))
	if r.Email == "" {
		return errParamIsRequired("email", "string")
	}
	if r.Password == "" {
		return errParamIsRequired("password", "string")
	}
	if _, err := mail.ParseAddress(r.Email); err != nil {
		return fmt.Errorf("email must be valid")
	}
	return nil
}
