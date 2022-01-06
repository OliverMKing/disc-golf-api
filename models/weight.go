package models

type Measurement struct {
	Value float32 `json:"value"`
	Unit  Unit    `json:"unit"`
}

// Unit enum
type Unit string

const (
	Gram       Unit = "g"
	Centimeter Unit = "cm"
	Kilogram   Unit = "kg"
	Percent    Unit = "%"
)
