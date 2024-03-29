package lawsuit

import (
	"regexp"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
	"github.com/pkg/errors"
)

const (
	LawsuitTableName = "lawsuit"
)

var (
	LawsuitNumberRegex = regexp.MustCompile(`^\d{7}\-\d{2}\.\d{4}\.\d{1}\.\d{2}\.\d{4}$`)
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
	return LawsuitTableName
}

func IsLawsuitNumber(lawsuitNumber string) bool {
	if !LawsuitNumberRegex.MatchString(lawsuitNumber) {
		return false
	}

	return true
}

func (l *Lawsuit) Ok() error {
	if !IsLawsuitNumber(l.LawsuitNumber) {
		return errors.New("lawsuit-number is invalid")
	}

	return nil
}

func GetByOrderID(db *gorm.DB, orderID string) (*Lawsuit, error) {
	l := new(Lawsuit)

	if err := db.First(l, "order_id = ?", orderID).Error; err != nil {
		return nil, err
	}

	return l, nil
}

func GetByOrderIDAndLawsuitNumber(db *gorm.DB, orderID string, lawsuitNumber string) (*Lawsuit, error) {
	l := new(Lawsuit)

	if err := db.First(l, "order_id = ? and lawsuit_number = ?", orderID, lawsuitNumber).Error; err != nil {
		return nil, err
	}

	return l, nil
}

func New(lawsuitNumber string) *Lawsuit {
	return &Lawsuit{LawsuitNumber: lawsuitNumber, OrderID: uuid.New().String()}
}

func Create(db *gorm.DB, lawsuit *Lawsuit) error {
	lawsuit.CreatedAt = time.Now()

	if err := lawsuit.Ok(); err != nil {
		return err
	}

	if err := db.Create(lawsuit).Error; err != nil {
		return err
	}

	return nil
}

func Update(db *gorm.DB, lawsuit *Lawsuit) error {
	lawsuit.UpdatedAt = time.Now()

	if err := lawsuit.Ok(); err != nil {
		return err
	}

	if err := db.Save(lawsuit).Error; err != nil {
		return err
	}

	return nil
}
