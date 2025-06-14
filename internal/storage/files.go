package storage

import (
	"context"
	"github.com/cockroachdb/cockroach-go/v2/crdb/crdbgorm"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// FileType represents the type of file being stored.
type FileType int

const (
	// Other represents any other file type.
	Other FileType = iota + 1
	// ProfilePicture represents a profile picture file type.
	ProfilePicture
)

// FileMetadata represents metadata for a file stored in the system.
type FileMetadata struct {
	// Unique identifier for the file
	ID int64 `gorm:"primaryKey;autoIncrement;not null"`
	// The object key for this file in the storage system (S3)
	ObjectKey string `gorm:"not null;type:varchar;"`
	// The identifier of the user who uploaded this file
	UploaderID uuid.UUID `gorm:"not null;type:uuid;"`
	// The unique identifier of the pet this file is associated with
	PetID uuid.UUID `gorm:"not null;type:uuid;"`
	// The name of the file
	Name string `gorm:"not null;type:varchar;"`
	// The type of the file (e.g., profile picture)
	Type FileType `gorm:"not null;type:int;default:1"`
	// The date and time this item was created
	CreatedAt time.Time
}

// InsertFileMetadata inserts a new file metadata record into the database.
func InsertFileMetadata(context context.Context, fileMetadata *FileMetadata) error {
	return crdbgorm.ExecuteTx(context, Db, nil,
		func(tx *gorm.DB) error {
			return tx.Create(&fileMetadata).Error
		},
	)
}
