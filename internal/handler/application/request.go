package application

import "fmt"

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

type CreateApplicationRequest struct {
	CandidateID uint   `json:"candidateId"`
	OpeningID   uint   `json:"openingId"`
	Notes       string `json:"notes"`
}

func (r *CreateApplicationRequest) Validate() error {
	if r.CandidateID == 0 && r.OpeningID == 0 {
		return fmt.Errorf("request body is empty or malformed")
	}

	if r.CandidateID == 0 {
		return errParamIsRequired("candidateId", "uint")
	}
	if r.OpeningID == 0 {
		return errParamIsRequired("openingId", "uint")
	}
	return nil
}

type UpdateApplicationRequest struct {
	Status string `json:"status"`
	Notes  string `json:"notes"`
}

func (r *UpdateApplicationRequest) Validate() error {
	if r.Status != "" || r.Notes != "" {
		return nil
	}
	return fmt.Errorf("at least one valid field must be provided")
}