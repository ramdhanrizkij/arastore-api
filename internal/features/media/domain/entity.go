package domain

import (
	"time"

	"github.com/google/uuid"
)

type Media struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id" format:"uuid"`
	BucketName  string    `gorm:"type:varchar(100);not null" json:"bucket_name"`
	ObjectKey   string    `gorm:"type:varchar(500) not null" json:"object_key"`
	FileName    string    `gorm:"varchar(256)" json:"file_name"`
	MimeType    string    `gorm:"type:varchar(100)" json:"mime_type"`
	FileSize    int       `gorm:"type:bigint" json:"file_size"`
	Width       int       `gorm:"type:int" json:"width"`
	Height      int       `gorm:"type:int" json:"height"`
	IsTemporary bool      `gorm:"type:boolean;not null;default:false" json:"is_temporary"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (Media) TableName() string {
	return "media"
}
