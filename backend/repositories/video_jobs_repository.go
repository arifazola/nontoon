package repositories

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/arifazola/nontoon/internal/db"
)

type VideoJobsRepository struct {
	Queries *db.Queries
}

func (repo *VideoJobsRepository) AddVideoJobs(context context.Context, id, uploadId string, index int) error {
	var index32 = sql.NullInt32{
		Int32: int32(index),
		Valid: true,
	}

	fmt.Println("Index Video Jobs", index32)
	fmt.Println("Index Video Jobs", index)

	return repo.Queries.AddVideoJob(context, db.AddVideoJobParams{
		ID: id,
		UploadId: uploadId,
		Index: index32,
	})
}

func (repo *VideoJobsRepository) GetLatestUploadedChunk(ctx context.Context, uploadId string) (db.VideoJob, error) {
	latestChunk, err := repo.Queries.GetLatestUploadedChunk(ctx, uploadId)

	var videoJob db.VideoJob
	if err != nil {
		return videoJob, err
	}

	return latestChunk, nil
}