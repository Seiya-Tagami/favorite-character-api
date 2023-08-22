package response

import "time"


type CharacterResponse struct {
	ID				int			`json:"id"`
	Name 			string	`json:"name"`
	Belonging string 	`json:"belonging"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}