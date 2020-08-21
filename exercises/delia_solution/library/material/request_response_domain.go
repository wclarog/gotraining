package material

import "time"

type MaterialRequest struct {
	Name          string    `json:"name"`
	EmissionDate  time.Time `json:"emissionDate"`
	NumberOfPages int       `json:"numberOfPages"`
}

type BookRequest struct {
	MaterialRequest
	Author string `json:"author"`
	Genre  string `json:"genre"`
}

type MaterialByCodeRequest struct {
	Id string
}

type NoContentResponse struct {
	Err error `json:"err,omitempty"`
}

type UpdateBookRequest struct {
	Code string      `json:"book"`
	Book BookRequest `json:"book"`
}

type BookResponse struct {
	Book Book  `json:"book,omitempty"`
	Err  error `json:"err,omitempty"`
}

type AddBookRequest struct {
	Book BookRequest `json:"book"`
}

type GetRequest struct{}
