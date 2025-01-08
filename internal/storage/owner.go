package storage

// An Owner Stores information about the authenticated party
type Owner struct {
	Issuer  string `gorm:"column:auth_issuer;type:varchar(255);not null"`
	OwnerId string `gorm:"column:auth_subject;type:varchar(255);not null"`
	Email   string `gorm:"column:auth_email;type:varchar(255);not null"`
}
