package models

type ChunkMetadata struct {
	UploadID, Filename      string
	ChunkIndex, TotalChunks int
}