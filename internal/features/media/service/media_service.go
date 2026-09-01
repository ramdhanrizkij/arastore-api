package service

import (
	"context"
	"fmt"
	"path"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/ramdhanrizkij/arastore-api/internal/core/storage"
	"github.com/ramdhanrizkij/arastore-api/internal/features/media/domain"
	"go.uber.org/zap"
)

type mediaService struct {
	log     *zap.Logger
	repo    domain.MediaRepository
	storage storage.Provider
	bucket  string
}

func NewMediaService(
	log *zap.Logger,
	repo domain.MediaRepository,
	storage storage.Provider,
	bucket string,
) domain.MediaService {
	return &mediaService{
		log:     log,
		repo:    repo,
		storage: storage,
		bucket:  bucket,
	}
}

// Upload implements [domain.MediaService].
func (s *mediaService) Upload(ctx context.Context, req *domain.UploadMediaRequest, file []byte, fileName string, mimeType string) (*domain.MediaResponse, error) {
	objectKey := s.generateObjectKey(fileName)

	storage, err := s.storage.Put(ctx, &storage.PutObjectRequest{
		Bucket:      s.bucket,
		Key:         objectKey,
		Reader:      strings.NewReader(string(file)),
		Size:        int64(len(file)),
		ContentType: mimeType,
	})

	if err != nil {
		s.log.Error("failed to upload file to storage",
			zap.Error(err),
			zap.String("bucket", s.bucket),
			zap.String("key", objectKey),
		)
		return nil, fmt.Errorf("failed to upload file: %w", err)
	}

	media := &domain.Media{
		ID:          uuid.New(),
		BucketName:  s.bucket,
		ObjectKey:   objectKey,
		FileName:    fileName,
		MimeType:    mimeType,
		FileSize:    int(storage.Size),
		IsTemporary: req.IsTemporary,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := s.repo.Create(ctx, media); err != nil {
		s.log.Error("failed to save media metadata",
			zap.Error(err),
			zap.String("object_key", objectKey),
		)

		_ = s.storage.Delete(ctx, s.bucket, objectKey)
		return nil, fmt.Errorf("failed to save media metadata: %w", err)
	}

	return &domain.MediaResponse{
		ID:        media.ID,
		Bucket:    s.bucket,
		Key:       objectKey,
		URL:       storage.URL,
		FileName:  fileName,
		MimeType:  mimeType,
		FileSize:  storage.Size,
		CreatedAt: media.CreatedAt.Format(time.RFC3339),
	}, nil
}

func (s *mediaService) generateObjectKey(fileName string) string {
	now := time.Now()
	ext := path.Ext(fileName)
	name := strings.TrimSuffix(fileName, ext)

	// Generate unique name
	uniqueID := uuid.New().String()[:8]

	return fmt.Sprintf("%s/%d/%02d/%s-%s%s",
		"uploads",
		now.Year(),
		now.Month(),
		name,
		uniqueID,
		ext,
	)
}
