package babies

import (
	"errors"
	"time"
)


type CreateInput struct {
	FullName string
	BirthDate time.Time 
}

type Baby struct {
	ID string
	fullName string
	birthDate time.Time
	createdAt time.Time
}

func (c *CreateInput) Validate() error {
	if len(c.FullName) != 0 {
		return errors.New("baby must have a name. cannot be empty")
	}
	return nil
}