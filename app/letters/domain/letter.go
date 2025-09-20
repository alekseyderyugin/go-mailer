package domain

import (
	"gorm.io/gorm"
	"time"
)

type LetterStatus string
type HtmlMessage string
type PlainMessage string

const (
	Pending    LetterStatus = "pending"
	Processing LetterStatus = "processing"
	Sent       LetterStatus = "sent"
	SendFailed LetterStatus = "send_failed"
)

type Letter struct {
	ID        LetterId `gorm:"type:char(36);primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	From         string
	FromName     string
	To           []Address `gorm:"type:json;serializer:json"`
	Cc           []Address `gorm:"type:json;serializer:json"`
	Bcc          []Address `gorm:"type:json;serializer:json"`
	Subject      string
	HtmlMessage  HtmlMessage  `gorm:"type:text"`
	PlainMessage PlainMessage `gorm:"type:text"`
	Status       LetterStatus `gorm:"type:text"`
	ClientID     ClientId
	LockedAt     *time.Time
}

func NewLetter(
	id LetterId,
	from string,
	to []Address,
	subject string,
	htmlMessage HtmlMessage,
	plainMessage PlainMessage,
	clientId ClientId,
) *Letter {
	entity := Letter{
		ID:           id,
		From:         from,
		To:           to,
		Cc:           []Address{},
		Bcc:          []Address{},
		Subject:      subject,
		HtmlMessage:  htmlMessage,
		PlainMessage: plainMessage,
		Status:       Pending,
		ClientID:     clientId,
	}

	now := time.Now()
	entity.CreatedAt = now
	entity.UpdatedAt = now

	return &entity
}

func (letter *Letter) GetFrom() string {

	return letter.From
}

func (letter *Letter) GetTo() []string {
	var result = make([]string, len(letter.To))
	for i := 0; i < len(letter.To); i++ {
		result[i] = letter.To[i].Email
	}
	return result
}

func (letter *Letter) GetHtmlMessage() string {
	return string(letter.HtmlMessage)
}
