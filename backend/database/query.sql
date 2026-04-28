-- database/query.sql

-- name: GetAllVideoJobs :many

SELECT id, "uploadId", index
	FROM public."VideoJobs";