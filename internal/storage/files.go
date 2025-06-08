package storage

import (
	"context"
	"github.com/cockroachdb/cockroach-go/v2/crdb/crdbgorm"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type FileMetadata struct {
	// Unique identifier for the file
	ID int64 `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	// The object key for this file in the storage system (S3)
	ObjectKey string `json:"object_key"`
	// The unique identifier of the owner of the pet
	OwnerID uuid.UUID `gorm:"type:uuid;"`
	// The unique identifier of the pet this file is associated with
	PetID uuid.UUID `gorm:"type:uuid;"`
	// The name of the file
	Name string `json:"name"`
	// The date and time this item was created
	CreatedAt time.Time `json:"created_at"`
}

func InsertFileMetadata(context context.Context, fileMetadata *FileMetadata) error {
	return crdbgorm.ExecuteTx(context, Db, nil,
		func(tx *gorm.DB) error {
			return tx.Create(&fileMetadata).Error
		},
	)
}
