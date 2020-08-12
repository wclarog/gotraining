package library

import "time"

type Material struct {
	UniqueCode     string    `json:"unique_code"`
	Name           string    `json:"name"`
	DateOfEmission time.Time `json:"date_of_emission"`
	NumberOfPages  int       `json:"number_of_pages"`
}

type DTOMaterial struct {
	UniqueCode     string    `json:"unique_code"`
	Name           string    `json:"name"`
	DateOfEmission time.Time `json:"date_of_emission"`
	NumberOfPages  int       `json:"number_of_pages"`
}

type Book struct {
	Material
	AuthorName string `json:"author_name"`
	Genre      string `json:"genre"`
}

type DTOBook struct {
	DTOMaterial
	AuthorName string `json:"author_name"`
	Genre      string `json:"genre"`
}

type Newspaper struct {
	Material
	Url string `json:"url"`
}

type DTONewspaper struct {
	DTOMaterial
	Url string `json:"url"`
}

type Magazine struct {
	Material
	Sections []string `json:"sections"`
	Url      string   `json:"url"`
}

type DTOMagazine struct {
	DTOMaterial
	Sections []string `json:"sections"`
	Url      string   `json:"url"`
}
