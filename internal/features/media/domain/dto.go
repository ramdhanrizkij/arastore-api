package domain

import "github.com/google/uuid"

// Request
type UploadMediaRequest struct {
	IsTemporary bool `json:"is_temporary"` // Flag media sementara
}

// Response
type MediaResponse struct {
	ID        uuid.UUID `json:"id"`
	Bucket    string    `json:"bucket"`
	Key       string    `json:"key"`
	URL       string    `json:"url"`
	FileName  string    `json:"file_name"`
	MimeType  string    `json:"mime_type"`
	FileSize  int64     `json:"file_size"`
	Width     int       `json:"width,omitempty"`
	Height    int       `json:"height,omitempty"`
	CreatedAt string    `json:"created_at"`
}

type MediaDetailResponse struct {
	ID          uuid.UUID `json:"id"`
	Bucket      string    `json:"bucket"`
	Key         string    `json:"key"`
	URL         string    `json:"url"`
	FileName    string    `json:"file_name"`
	MimeType    string    `json:"mime_type"`
	FileSize    int64     `json:"file_size"`
	Width       int       `json:"width,omitempty"`
	Height      int       `json:"height,omitempty"`
	IsTemporary bool      `json:"is_temporary"`
	CreatedAt   string    `json:"created_at"`
	UpdatedAt   string    `json:"updated_at"`
}
