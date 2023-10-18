package entity

type Hero struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Universe string `json:"universe"`
	Skill    string `json:"skill"`
	ImageURL string `json:"imageURL"`
}
