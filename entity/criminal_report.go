package entity

type CriminalReport struct {
	ID          int    `json:"id"`
	HeroID      int    `json:"hero_id"`
	VillainID   int    `json:"villain_id"`
	Description string `json:"description"`
	Incident    string `json:"incident"`
}
