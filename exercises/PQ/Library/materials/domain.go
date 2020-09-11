package materials

import "time"

type MaterialType int

const (
	BookType MaterialType = iota
	NewspaperType
	MagazineType
)

// Domain
type Material struct {
	UniqueCode     string       `json:"uniqueCode"`
	Name           string       `json:"name"`
	DateOfEmission time.Time    `json:"dateOfEmission"`
	NumberOfPages  uint         `json:"numberOfPages"`
	MaterialType   MaterialType `json:"materialType"`
}

type Book struct {
	Material
	AuthorName string `json:"authorName"`
	Genre      string `json:"genre"`
}

type Newspaper struct {
	Material
	Url string `json:"url"`
}

type Magazine struct {
	Material
	Sections []Section `json:"sections"`
	Url      string    `json:"url"`
}

type Section struct {
	Code    string `json:"code"`
	Content string `json:"content"`
}

// DTOs
type DTOMaterial struct {
	UniqueCode     string
	Name           string
	DateOfEmission time.Time
	NumberOfPages  uint
	MaterialType   MaterialType
}

type DTOBook struct {
	DTOMaterial
	AuthorName string
	Genre      string
}

type DTONewspaper struct {
	DTOMaterial
	Url string
}

type DTOMagazine struct {
	DTOMaterial
	Sections []DTOSection
	Url      string
}

type DTOSection struct {
	Code    string
	Content string
}

// MaterialTyped interface

type MaterialTyped interface {
	GetMaterialType() MaterialType
}

func (m Material) GetMaterialType() MaterialType {
	return m.MaterialType
}

func (b Book) GetMaterialType() MaterialType {
	return b.MaterialType
}

func (m Magazine) GetMaterialType() MaterialType {
	return m.MaterialType
}

func (n Newspaper) GetMaterialType() MaterialType {
	return n.MaterialType
}

func (m DTOMaterial) GetMaterialType() MaterialType {
	return m.MaterialType
}

func (b DTOBook) GetMaterialType() MaterialType {
	return b.MaterialType
}

func (m DTOMagazine) GetMaterialType() MaterialType {
	return m.MaterialType
}

func (n DTONewspaper) GetMaterialType() MaterialType {
	return n.MaterialType
}
