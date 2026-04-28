package interfaces

import (
	"context"

	"github.com/arifazola/nontoon/internal/db"
)

type VideoJobs interface {
	AddVideoJobs(context context.Context, id, uploadId string, index int) error
	GetLatestUploadedChunk(ctx context.Context, uploadId string) (db.VideoJob, error)
}