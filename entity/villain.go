package entity

type Villain struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Universe string `json:"universe"`
	ImageURL string `json:"imageURL"`
}
