package domains

type Characters struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Game       string `json:"game"`
	Occupation string `json:"occupation"`
	Skill      string `json:"skill"`
}
