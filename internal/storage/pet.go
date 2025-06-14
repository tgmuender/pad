package storage

import (
	"context"
	"errors"
	"fmt"
	"github.com/cockroachdb/cockroach-go/v2/crdb/crdbgorm"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

var Db *gorm.DB

// Pet represents a pet in the database
type Pet struct {
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

// PetOverview represents a pet with additional profile picture metadata
// This struct is used to provide an overview of a pet, including its profile picture.
type PetOverview struct {
	Pet
	ProfilePicture *FileMetadata
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
func InsertPet(petEntity *Pet) error {
	fmt.Println("Inserting into database")

	return crdbgorm.ExecuteTx(context.Background(), Db, nil,
		func(tx *gorm.DB) error {
			return tx.Create(&petEntity).Error
		},
	)
}

// GetPetByName queries the database to find a pet with the given name.
func GetPetByName(name string, ownerId uuid.UUID) (*Pet, error) {
	var pet Pet
	result := Db.Where("data ->> 'name' = ? and owner_id = ?", name, ownerId).First(&pet)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, result.Error
		}
	}
	return &pet, nil
}

// FindByOwner queries the database to find all pets whose owner is set to the given owner.
func FindByOwner(user *User) []PetOverview {
	if user == nil {
		return []PetOverview{}
	}

	var pets []Pet
	Db.Where("owner_id = ?", user.Id).Find(&pets)

	var petOverviews []PetOverview
	for _, pet := range pets {
		var fm *FileMetadata
		dbResult := Db.Model(&FileMetadata{}).
			Where("pet_id = ? and type = ?", pet.ID, ProfilePicture).
			Order("created_at desc").
			First(&fm)
		if dbResult.Error == nil {
			petOverviews = append(petOverviews, PetOverview{
				Pet:            pet,
				ProfilePicture: fm,
			})
		} else {
			// Either the profile picture does not exist or there was an error retrieving it
			petOverviews = append(petOverviews, PetOverview{
				Pet:            pet,
				ProfilePicture: nil,
			})
		}
	}

	return petOverviews
}

// ExistsForOwner checks if a pet with the given ID exists and is owned by the given user.
func ExistsForOwner(petId uuid.UUID, user *User) bool {
	if user == nil {
		return false
	}

	var count int64
	Db.Model(&Pet{}).Where("id = ? and owner_id = ?", petId, user.Id).Count(&count)
	return count > 0
}

func DeletePet(petId uuid.UUID) error {
	return crdbgorm.ExecuteTx(context.Background(), Db, nil,
		func(tx *gorm.DB) error {
			return tx.Delete(&Pet{ID: petId}).Error
		},
	)
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
