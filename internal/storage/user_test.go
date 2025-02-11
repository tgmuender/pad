package storage

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestUserBeforeCreate(t *testing.T) {
	// Create an in-memory SQLite database
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect database: %v", err)
	}

	// Auto migrate the User schema
	err = db.AutoMigrate(&User{})
	if err != nil {
		t.Fatalf("failed to migrate database: %v", err)
	}

	// Create a new user
	user := User{
		Issuer:  "http://localhost:9000",
		Subject: "test_subject",
		Email:   "test@example.com",
	}

	// Save the user to the database
	err = db.Create(&user).Error
	if err != nil {
		t.Fatalf("failed to create user: %v", err)
	}

	// Check if the UUID was generated
	assert.NotEqual(t, uuid.Nil, user.Id, "expected UUID to be generated")
}

func TestUserFields(t *testing.T) {
	// Create a new user
	user := User{
		Issuer:  "test_issuer",
		Subject: "test_subject",
		Email:   "test@example.com",
	}

	// Check if the fields are set correctly
	assert.Equal(t, "test_issuer", user.Issuer)
	assert.Equal(t, "test_subject", user.Subject)
	assert.Equal(t, "test@example.com", user.Email)
}
