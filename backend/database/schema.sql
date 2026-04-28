-- database/schema.sql
CREATE TABLE IF NOT EXISTS public."VideoJobs"
(
    id text COLLATE pg_catalog."default" NOT NULL,
    "uploadId" text COLLATE pg_catalog."default" NOT NULL,
    index integer,
    CONSTRAINT "VideoJobs_pkey" PRIMARY KEY (id)
)