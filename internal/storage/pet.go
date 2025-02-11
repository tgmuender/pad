package storage

import (
	"context"
	"fmt"
	"github.com/cockroachdb/cockroach-go/v2/crdb/crdbgorm"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

var Db *gorm.DB

// PetEntity represents a pet in the database
type PetEntity struct {
	// The unique identifier of the pet
	ID uuid.UUID `gorm:"type:uuid;primary_key;"`
	// The unique identifier of the owner of the pet
	OwnerID uuid.UUID `gorm:"type:uuid;"`
	// The owner of the pet
	Owner User `gorm:"foreignKey:OwnerID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	// The date and time the pet was created
	CreatedAt time.Time
	// The date and time the pet was last updated
	UpdatedAt time.Time
	// The data of the pet
	Data string `gorm:"type:jsonb" json:"data"`
}

// MealEntity represents a meal in the database
type MealEntity struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	PetID     uuid.UUID `gorm:"type:uuid;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Data      string `gorm:"type:jsonb" json:"data"`
}

// InsertPet persists the given pet into the database
func InsertPet(petEntity *PetEntity) error {
	fmt.Println("Inserting into database")

	return crdbgorm.ExecuteTx(context.Background(), Db, nil,
		func(tx *gorm.DB) error {
			return tx.Create(&petEntity).Error
		},
	)
}

// FindByOwner queries the database to find all pets whose owner is set to the given owner.
func FindByOwner(user *User) []PetEntity {
	if user == nil {
		return []PetEntity{}
	}

	var result []PetEntity
	/*	Db.Where(
		&PetEntity{
			Owner: Owner{
				Issuer:  owner.Issuer,
				OwnerId: owner.OwnerId,
			},
		},
	).Find(&result)*/
	return result
}

// FindMealsByPetID queries the database to find all meals whose pet ID is set to the given ID.
func FindMealsByPetID(petId uuid.UUID) []MealEntity {
	if petId == uuid.Nil {
		return []MealEntity{}
	}

	var result []MealEntity
	Db.Where(
		&MealEntity{
			PetID: petId,
		},
	).Find(&result)
	return result
}
