package storage

import (
	"context"
	"fmt"
	"github.com/cockroachdb/cockroach-go/v2/crdb/crdbgorm"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// User represents an application user in the database, which can be an owner of a pet or a person which has access
// to shared pets.
type User struct {
	// The unique identifier of thh user
	Id uuid.UUID `gorm:"type:uuid;primary_key;"`
	// The issuer/idp of the user
	Issuer string `gorm:"column:auth_issuer;type:varchar(255);not null"`
	// The subject of the user in their authentication system/oidc provider
	Subject string `gorm:"column:auth_subject;type:varchar(255);not null"`
	// The email address of the user
	Email string `gorm:"column:auth_email;type:varchar(255);not null"`
	// The name of the user
	Name string `gorm:"column:name;type:varchar(255);null"`
	// CreatedAt is the time the user was created
	CreatedAt time.Time
	// UpdatedAt is the time the user was last updated
	UpdatedAt time.Time
}

// BeforeCreate is a hook that is called before the user is created in the database. It generates a new UUID for the user.
func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.Id = uuid.New()
	return nil
}

// InsertUser inserts a new user into the database.
func InsertUser(user *User) error {
	fmt.Println("Inserting into database")

	return crdbgorm.ExecuteTx(context.Background(), Db, nil,
		func(tx *gorm.DB) error {
			return tx.Create(&user).Error
		},
	)
}

// FindUserByEmail queries the database to find a user with the given email address.
func FindUserByEmail(email string) (*User, error) {
	var user User
	result := Db.Where("auth_email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
