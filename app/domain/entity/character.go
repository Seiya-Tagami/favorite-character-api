package entity

import (
	"errors"
	"time"
)

type Character struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"not null"`
	Belonging string    `json:"belonging" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (c *Character) Validate() error {
	if len(c.Name) == 0 {
		return errors.New("name is required")
	}
	if len(c.Belonging) == 0 {
		return errors.New("belonging is required")
	}

	return nil
}