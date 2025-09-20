package domain

import "github.com/gofrs/uuid"

type UUID string

func NewUUID() UUID {
	id, err := uuid.NewV7()

	if err != nil {
		panic(err)
	}

	return UUID(id.String())
}
