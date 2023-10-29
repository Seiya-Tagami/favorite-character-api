package character

import (
	"github.com/Seiya-Tagami/favorite-character-management-api/domain/entity"
	characterRepository "github.com/Seiya-Tagami/favorite-character-management-api/domain/repository/character"
)

type Interactor interface {
	ListCharacters() ([]entity.Character, error)
	FindCharacterById(id int) (entity.Character, error)
	CreateCharacter(character entity.Character) (entity.Character, error)
	UpdateCharacter(character entity.Character, id int) (entity.Character, error)
	DeleteById(id int) error
}

type interactor struct {
	characterRepository characterRepository.Repository
}

func New(
	characterRepository characterRepository.Repository,
) Interactor {
	return &interactor{
		characterRepository,
	}
}

func (i *interactor) ListCharacters() ([]entity.Character, error) {
	characters := []entity.Character{}
	err := i.characterRepository.SelectALL(&characters)
	if err != nil {
		return []entity.Character{}, err
	}
	return characters, nil
}

func (i *interactor) FindCharacterById(id int) (entity.Character, error) {
	character := entity.Character{}
	err := i.characterRepository.SelectById(&character, id)
	if err != nil {
		return entity.Character{}, err
	}

	return character, nil
}

func (i *interactor) CreateCharacter(character entity.Character) (entity.Character, error) {
	err := i.characterRepository.Insert(&character)
	if err != nil {
		return entity.Character{}, err
	}

	return character, nil
}

func (i *interactor) UpdateCharacter(character entity.Character, id int) (entity.Character, error) {
	if err := i.characterRepository.UpdateById(&character, id); err != nil {
		return entity.Character{}, err
	}

	return character, nil
}

func (i *interactor) DeleteById(id int) error {
	if err := i.characterRepository.DeleteById(id); err != nil {
		return err
	}

	return nil
}
