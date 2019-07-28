package lawsuit

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
	"github.com/pkg/errors"
)

type Lawsuit struct {
	ID            uint           `gorm:"primary_key"`
	OrderID       string         `gorm:"column:order_id"`
	LawsuitNumber string         `gorm:"column:lawsuit_number"`
	Processed     bool           `gorm:"column:processed"`
	Criticized    bool           `gorm:"column:criticized"`
	CreatedAt     time.Time      `gorm:"column:created_at"`
	UpdatedAt     time.Time      `gorm:"column:updated_at"`
	Output        postgres.Jsonb `gorm:"column:output"`
}

func (Lawsuit) TableName() string {
	return "lawsuit"
}

func (l *Lawsuit) Ok() error {
	if l.LawsuitNumber == "" {
		return errors.New("LawsuitNumber is invalid")
	}

	return nil
}

func GetByOrderIDAndLawsuitNumber(db *gorm.DB, orderID string, lawsuitNumber string) (*Lawsuit, error) {
	l := new(Lawsuit)
	if err := db.First(l, "order_id = ? and lawsuit_number = ?", orderID, lawsuitNumber).Error; err != nil {
		return nil, err
	}
	return l, nil
}

func New(lawsuitNumber string) *Lawsuit {
	l := new(Lawsuit)
	l.LawsuitNumber = lawsuitNumber
	l.OrderID = uuid.New().String()

	return l
}

func Create(db *gorm.DB, lawsuit *Lawsuit) error {
	lawsuit.CreatedAt = time.Now()

	err := lawsuit.Ok()
	if err != nil {
		return err
	}

	err = db.Create(lawsuit).Error
	if err != nil {
		return err
	}

	return nil
}

func Update(db *gorm.DB, lawsuit *Lawsuit) error {
	lawsuit.UpdatedAt = time.Now()

	err := lawsuit.Ok()
	if err != nil {
		return err
	}

	fmt.Println(lawsuit)

	err = db.Save(lawsuit).Error
	if err != nil {
		return err
	}

	return nil
}
