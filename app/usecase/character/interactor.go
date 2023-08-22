package character

import (
	"github.com/Seiya-Tagami/favorite-character-api/data/response"
	"github.com/Seiya-Tagami/favorite-character-api/domain/entity"
	characterRepository "github.com/Seiya-Tagami/favorite-character-api/domain/repository/character"
)

type Interactor interface {
	ListCharacters() ([]response.CharacterResponse, error)
	FindCharacterById(id int) (response.CharacterResponse, error)
	CreateCharacter(character entity.Character) (response.CharacterResponse, error)
	UpdateCharacter(character entity.Character, id int) (response.CharacterResponse, error)
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

func (i *interactor) ListCharacters() ([]response.CharacterResponse, error) {
	characters := []entity.Character{}
	err := i.characterRepository.SelectALL(&characters)
	if err != nil {
		return []response.CharacterResponse{}, err
	}

	charactersRes := []response.CharacterResponse{}
	for _, v := range characters {
		t := response.CharacterResponse{
			ID:        v.ID,
			Name:      v.Name,
			Belonging: v.Belonging,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
		charactersRes = append(charactersRes, t)
	}
	println(charactersRes)

	return charactersRes, nil
}

func (i *interactor) FindCharacterById(id int) (response.CharacterResponse, error) {
	character := entity.Character{}
	err := i.characterRepository.SelectById(&character, id)
	if err != nil {
		return response.CharacterResponse{}, err
	}

	characterRes := response.CharacterResponse{
		ID:        character.ID,
		Name:      character.Name,
		Belonging: character.Belonging,
		CreatedAt: character.CreatedAt,
		UpdatedAt: character.UpdatedAt,
	}

	return characterRes, nil
}

func (i *interactor) CreateCharacter(character entity.Character) (response.CharacterResponse, error) {
	err := i.characterRepository.Insert(&character)
	if err != nil {
		return response.CharacterResponse{}, err
	}

	characterRes := response.CharacterResponse{
		ID:        character.ID,
		Name:      character.Name,
		Belonging: character.Belonging,
		CreatedAt: character.CreatedAt,
		UpdatedAt: character.UpdatedAt,
	}

	return characterRes, nil
}

func (i *interactor) UpdateCharacter(character entity.Character, id int) (response.CharacterResponse, error) {
	if err := i.characterRepository.UpdateById(&character, id); err != nil {
		return response.CharacterResponse{}, err
	}

	characterRes := response.CharacterResponse{
		ID:        id,
		Name:      character.Name,
		Belonging: character.Belonging,
		CreatedAt: character.CreatedAt,
		UpdatedAt: character.UpdatedAt,
	}

	return characterRes, nil
}

func (i *interactor) DeleteById(id int) error {
	if err := i.characterRepository.DeleteById(id); err != nil {
		return err
	}

	return nil
}

