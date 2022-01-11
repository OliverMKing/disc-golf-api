package models

type Disc struct {
	Id            int         `json:"id"`
	Name          string      `json:"name"`
	Distributor   string      `json:"distributor"`
	MaxWeight     Measurement `json:"maxWeight"`
	Diameter      Measurement `json:"diameter"`
	Height        Measurement `json:"height"`
	RimDepth      Measurement `json:"rimDepth"`
	Speed         int         `json:"speed"`
	Glide         int         `json:"glide"`
	Turn          int         `json:"turn"`
	Fade          int         `json:"fade"`
	Stability     string      `json:"stability"`
	PrimaryUse    string      `json:"primaryUse"`
	PlasticGrades []string    `json:"plasticGrades"`
	Link          string      `json:"link"`
}

type DiscsResponse struct {
	Discs []Disc `json:"discs"`
}
