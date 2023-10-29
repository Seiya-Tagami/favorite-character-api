package character

import (
	"time"

	"github.com/Seiya-Tagami/favorite-character-api/domain/entity"
)

type Response struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Belonging string    `json:"belonging"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ToResponse(character *entity.Character) Response {
	return Response{
		ID:        character.ID,
		Name:      character.Name,
		Belonging: character.Belonging,
		CreatedAt: character.CreatedAt,
		UpdatedAt: character.UpdatedAt,
	}
}

func ToListResponse(characters *[]entity.Character) []Response {
	var listResponse []Response
	for _, character := range *characters {
		response := ToResponse(&character)
		listResponse = append(listResponse, response)
	}
	return listResponse
}
