package controllerdomain

type CharacterRequest struct {
	CharacterName string `json:"character_name"`
	ClassName     string `json:"class_name"`
}
