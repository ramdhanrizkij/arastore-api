package domain

import "context"

type MediaRepository interface {
	Create(ctx context.Context, media *Media) error
}

type MediaService interface {
	Upload(ctx context.Context, req *UploadMediaRequest, file []byte, fileName, mimeType string) (*MediaResponse, error)
}
