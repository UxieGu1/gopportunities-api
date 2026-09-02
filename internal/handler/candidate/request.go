package candidate

import "fmt"

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

type CreateCandidateRequest struct {
	UserID    uint   `json:"userId"`
	Name      string `json:"name"`
	Linkedin  string `json:"linkedin"`
	ResumeURL string `json:"resumeUrl"`
	Skills    string `json:"skills"`
}

func (r *CreateCandidateRequest) Validate() error {
	if r.Name == "" && r.Linkedin == "" && r.ResumeURL == "" && r.Skills == "" {
		return fmt.Errorf("request body is empty or malformed")
	}

	if r.Name == "" {
		return errParamIsRequired("name", "string")
	}
	return nil
}

type UpdateCandidateRequest struct {
	Name      string `json:"name"`
	Linkedin  string `json:"linkedin"`
	ResumeURL string `json:"resumeUrl"`
	Skills    string `json:"skills"`
}

func (r *UpdateCandidateRequest) Validate() error {
	if r.Name != "" || r.Linkedin != "" || r.ResumeURL != "" || r.Skills != "" {
		return nil
	}
	return fmt.Errorf("at least one valid field must be provided")
}
