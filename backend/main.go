package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"

	"github.com/arifazola/nontoon/controllers"
	"github.com/arifazola/nontoon/database"
	"github.com/arifazola/nontoon/internal/db"
	"github.com/arifazola/nontoon/repositories"
	"github.com/arifazola/nontoon/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)
func main(){
	go func() {
		//ffmpeg -i sample.mp4 -codec: copy -start_number 0 -hls_time 1 -hls_list_size 0 -f hls output.m3u8
		fmt.Println("Executing command")
		our, _ := exec.Command("ffmpeg", "-i", "sample.mp4", "-codec:", "copy", "-start_number", "0", "-hls_time", "10", "-hls_list_size", "0", "-f", "hls", "output.m3u8").CombinedOutput()

		fmt.Println("Command finish: ", string(our))
	}()

	router := gin.Default()
	router.Use(cors.Default())

	databaseUrl := "postgres://postgres:test1234@localhost:5432/nontoon?sslmode=disable"

	dbConn, err := database.NewDB(databaseUrl)

	if err != nil {
		log.Fatal(err)
	}

	defer dbConn.Close()

	queries := db.New(dbConn)

	localStorage := repositories.LocalStorage{
		BasePath: "./files",
	}

	videoJobsRepo := repositories.VideoJobsRepository{
		Queries: queries,
	}

	videoService := services.VideoService {
		FileStorage: &localStorage,
		VideoJobs: &videoJobsRepo,
	}

	videoController := controllers.VideoController{
		VideoService: &videoService,
	}

	chunkController := controllers.ChunkController {
		VideoService: &videoService,
	}
	
	router.GET("/videos", controllers.GetAllVideos)

	router.POST("/videos", videoController.UploadVideo)

	router.POST("/videos/chunks", videoController.UploadChunk)

	router.POST("/videos/merge", videoController.CompleteUpload)

	router.GET("/chunks/:uploadId", chunkController.GetLatestUploadedChunk)

	router.StaticFS("/assets", http.Dir("./assets"))

	router.Run("localhost:8080")
}