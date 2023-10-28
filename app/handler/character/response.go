package character

import (
	"github.com/Seiya-Tagami/favorite-character-api/domain/entity"
	"time"
)

type response struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Belonging string    `json:"belonging"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ToResponse(character *entity.Character) response {
	return response{
		ID:        character.ID,
		Name:      character.Name,
		Belonging: character.Belonging,
		CreatedAt: character.CreatedAt,
		UpdatedAt: character.UpdatedAt,
	}
}

func ToListResponse(characters *[]entity.Character) []response {
	listResponse := []response{}
	for _, character := range *characters {
		response := ToResponse(&character)
		listResponse = append(listResponse, response)
	}
	return listResponse
}
