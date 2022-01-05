package models

type Disc struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type DiscResponse struct {
	Discs []Disc `json:"disc"`
}
