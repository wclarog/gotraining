package material

import (
	"time"
)

const (
	BookType TypeMaterial = iota
	MagazineType
	NewsPaperType
)

type TypeMaterial int

type GenericMaterial interface {
	GetTypeMaterial() TypeMaterial
}

type Material struct {
	UniqueCode    string       `json:"unique_code"`
	Name          string       `json:"name"`
	EmissionDate  time.Time    `json:"emissionDate"`
	NumberOfPages int          `json:"numberOfPages"`
	Type          TypeMaterial `json:"type"`
}

type MaterialDTO struct {
	UniqueCode    string
	Name          string
	EmissionDate  time.Time
	NumberOfPages int
	Type          TypeMaterial
}

type Book struct {
	Material
	Author string `json:"author"`
	Genre  string `json:"genre"`
}

type BookDTO struct {
	MaterialDTO
	Author string
	Genre  string
}

type Magazine struct {
	Material
	Url      string    `json:"url"`
	Sections []Section `json:"sections"`
}

type MagazineDTO struct {
	MaterialDTO
	Url      string
	Sections []Section
}

type NewsPaper struct {
	Material
	Url string `json:"url"`
}

type NewsPaperDTO struct {
	MaterialDTO
	Url string `json:"url"`
}

type Section struct {
	Name string
	Code string
}

func (m MaterialDTO) GetTypeMaterial() TypeMaterial {
	return m.Type
}

func (b BookDTO) GetTypeMaterial() TypeMaterial {
	return b.Type
}

func (m MagazineDTO) GetTypeMaterial() TypeMaterial {
	return m.Type
}

func (n NewsPaperDTO) GetTypeMaterial() TypeMaterial {
	return n.Type
}
