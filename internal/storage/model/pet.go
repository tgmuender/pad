package storage

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

var Db *gorm.DB

type PetEntity struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	Owner     Owner     `gorm:"embedded"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Data      string `gorm:"type:jsonb" json:"data"`
}

// An Owner Stores information about the authenticated party
type Owner struct {
	Issuer  string `gorm:"column:auth_issuer;type:varchar(255);not null"`
	OwnerId string `gorm:"column:auth_subject;type:varchar(255);not null"`
	Email   string `gorm:"column:auth_email;type:varchar(255);not null"`
}

func InsertPet(petEntity *PetEntity) error {
	fmt.Println("Inserting into database")

	if err := Db.Create(&petEntity).Error; err != nil {
		return err
	}
	return nil
}
