package repositories

import (
	"context"
	"database/sql"
	"errors"

	"github.com/arifazola/nontoon/internal/db"
)

type MockVideoJobsRepository struct{
	Err error
}

func (m *MockVideoJobsRepository) AddVideoJobs(context context.Context, id, uploadId string, index int) error {
	if m.Err != nil {
		return errors.New("Failed To Store Jobs")
	}

	return nil
}

func (m *MockVideoJobsRepository) GetLatestUploadedChunk(ctx context.Context, uploadId string) (db.VideoJob, error) {
	var videoJob db.VideoJob
	if m.Err != nil {
		return videoJob, errors.New("Err")
	}

	videoJob.ID = "zzzz"
	videoJob.UploadId = "abc"
	index := sql.NullInt32 {
		Int32: 1,
		Valid: false,
	}
	videoJob.Index = index

	return videoJob, nil

}