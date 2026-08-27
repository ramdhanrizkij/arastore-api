package model
import (
	"github.com/google/uuid"
	
	"time"
)

type Address struct {
	ID              uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id" format:"uuid" example:"018f7606-a3f7-7c40-8e4b-2d47c6e04c8d"`
	UserID		uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	User 		*User `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Label		string `gorm:"type:varchar(50)" json:"label"`
	RecipientName	string `gorm:"type:varchar(100)" json:"recipient_name"`
	Phone		string `gorm:"type:varchar(20)" json:"address"`
	Address		string `gorm:"type:text" json:"address"`
	City		string `gorm:"type:varchar(100)" json:"city"`
	Province	string `gorm:"type:varchar(100)" json:"province"`
	PostalCode	string `gorm:"type:varchar(100)" json:"postal_code"`
	
	CreatedAt   time.Time `json:"created_at" example:"2026-05-13T10:00:00+07:00"`
	UpdatedAt   time.Time `json:"updated_at" example:"2026-05-13T10:00:00+07:00"`
}

func (Address) TableName() string {
	return "addresses"
}