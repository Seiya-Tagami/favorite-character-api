package entity

import (
	"errors"
	"time"
)

type Character struct {
	ID				int			`json:"id" gorm:"primaryKey"`
	Name 			string	`json:"name" gorm:"not null"`
	Belonging string 	`json:"belonging" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewCharacter(name, belonging string) (*Character, error) {
	if name == "" {
		return nil, errors.New("nameを入力してください。")
	}

	if belonging == "" {
		return nil, errors.New("belongingを入力してください。")
	}

	character := &Character{
		Name: name,
		Belonging: belonging,
	}

	return character, nil
}

func (c *Character) Set(name, belonging string) error {
	if name == "" {
		return errors.New("nameを入力してください。")
	}

	if belonging == "" {
		return errors.New("belongingを入力してください。")
	}

	c.Name = name
	c.Belonging = belonging

	return nil
}