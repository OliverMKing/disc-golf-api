package models

type Disc struct {
	Name        string      `json:"name" gorm:"primaryKey"`
	Distributor string      `json:"distributor" gorm:"primaryKey"`
	MaxWeight   Measurement `json:"maxWeight" gorm:"embedded;embeddedPrefix:max_weight_"`
	Diameter    Measurement `json:"diameter" gorm:"embedded;embeddedPrefix:diameter_"`
	Height      Measurement `json:"height" gorm:"embedded;embeddedPrefix:height_"`
	RimDepth    Measurement `json:"rimDepth" gorm:"embedded;embeddedPrefix:rim_depth_"`
	Speed       int         `json:"speed"`
	Glide       int         `json:"glide"`
	Turn        int         `json:"turn"`
	Fade        int         `json:"fade"`
	Stability   string      `json:"stability"`
	PrimaryUse  string      `json:"primaryUse"`
	// PlasticGrades []string    `json:"plasticGrades"`
	Link string `json:"link"`
}

type DiscsResponse struct {
	Discs []Disc `json:"discs"`
}
