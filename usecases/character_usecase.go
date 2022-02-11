package usecases

import (
	"log"

	validator "github.com/bytedance/go-tagexpr/v2/validator"
	"github.com/victoramsantos/game-contest/domain"
	"github.com/victoramsantos/game-contest/usecases/usecasesdomain"
)

type characterUsecase struct {
	repository domain.CharacterRepository
}

func NewCharacterUsecase(repository domain.CharacterRepository) domain.CharacterUsecase {
	return &characterUsecase{
		repository: repository,
	}
}

func (this *characterUsecase) NewCharacter(name string, className string) (*domain.Character, error) {
	class, err := this.repository.GetClassByName(className)
	if err != nil {
		log.Println("No class found for", className)
		return nil, err
	}

	character := domain.Character{
		Name:    name,
		Class:   class,
		IsAlive: true,
	}

	err = validator.Validate(character)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	err = this.repository.CreateCharacter(character)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &character, nil
}

func (this *characterUsecase) GetCharacterDetails(name string) (*usecasesdomain.CharacterDetails, error) {
	character, err := this.repository.GetCharacterByName(name)
	if err != nil {
		log.Println("No class found for", name)
		return nil, err
	}

	return character.ToDetails(), nil
}

func (this *characterUsecase) ListCharacters() ([]usecasesdomain.CharacterStatus, error) {
	characters, err := this.repository.ListCharacters()
	if err != nil {
		log.Println("No class found for")
		return nil, err
	}
	return fromDomainListToStatusList(characters), nil
}

func fromDomainListToStatusList(characters []domain.Character) []usecasesdomain.CharacterStatus {
	list := make([]usecasesdomain.CharacterStatus, 0)
	for _, character := range characters {
		list = append(list, *character.ToStatus())
	}
	return list
}
