package infrastructure

import (
	"go-mailer/letters/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type LetterRepository struct {
	db      *gorm.DB
	context *Context
}

func NewLetterRepository(db *gorm.DB, ctx *Context) *LetterRepository {
	return &LetterRepository{
		db:      db,
		context: ctx,
	}
}

func (rep *LetterRepository) AutoMigrate() error {

	return (*rep).db.AutoMigrate(&domain.Letter{})
}

func (rep *LetterRepository) Save(letter *domain.Letter) error {
	err := rep.db.Save(letter).Error
	return err
}

func (rep *LetterRepository) GetNextForSend(limit uint, lockDuration time.Duration) []*domain.Letter {
	var lettersList = make([]*domain.Letter, limit)

	err := rep.db.Transaction(func(tx *gorm.DB) error {
		now := time.Now()
		lock := now.Add(lockDuration)

		err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			Where("status = ? AND (locked_at IS NULL OR locked_at <= ?)", domain.Pending, now).
			Limit(int(limit)).
			Find(&lettersList).Error

		if err != nil {
			return err
		}

		if len(lettersList) == 0 {
			return nil
		}

		ids := make([]domain.LetterId, len(lettersList))

		for i := 0; i < len(lettersList); i++ {
			ids[i] = lettersList[i].ID
		}

		err = tx.Model(&domain.Letter{}).
			Where("id IN ?", ids).
			Update("locked_at", lock).Error

		if err != nil {
			return err
		}

		return nil
	})

	rep.context.HandleError(err)
	return lettersList
}

func (rep *LetterRepository) CreateBatch(letters []*domain.Letter) error {
	return rep.db.Create(&letters).Error
}
