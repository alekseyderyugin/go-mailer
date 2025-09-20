package domain

import (
	"go-mailer/shared/domain"
)

type ClientId string

func NewClientId() ClientId {
	return ClientId(domain.NewUUID())
}
