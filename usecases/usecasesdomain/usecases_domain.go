package usecasesdomain

type CharacterDetails struct {
	Name         string `json:"name"`
	Class        string `json:"class"`
	Life         string `json:"life"`
	Strength     string `json:"strength"`
	Skill        string `json:"skill"`
	Intelligence string `json:"intelligence"`
	Attack       string `json:"attack"`
	Velocity     string `json:"velocity"`
}

type CharacterStatus struct {
	Name    string `json:"name"`
	Class   string `json:"class"`
	IsAlive bool   `json:"is_alive"`
}
