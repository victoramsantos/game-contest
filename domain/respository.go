package domain

type CharacterRepository interface {
	GetClassByName(className string) (*Class, error)
	GetCharacterByName(name string) (*Character, error)
	CreateCharacter(character Character) error
	ListCharacters() ([]Character, error)
}
