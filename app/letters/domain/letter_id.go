package domain

import (
	"go-mailer/shared/domain"
)

type LetterId string

func NewLetterID() LetterId {
	return LetterId(domain.NewUUID())
}
