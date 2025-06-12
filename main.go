package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/roudra323/gin-simple-api/api"
	"github.com/roudra323/gin-simple-api/domain"
	"github.com/roudra323/gin-simple-api/repository"
	"github.com/roudra323/gin-simple-api/service"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Database migration running...")
	err = db.AutoMigrate(&domain.Video{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	videoRepo := repository.NewVideoRepository(db)
	videoService := service.NewVideoService(videoRepo)
	videoHandler := api.NewVideohandler(videoService)

	router := gin.Default()

	videoRoutes := router.Group("/videos")
	{
		videoRoutes.GET("", videoHandler.GetVideos)
		videoRoutes.POST("", videoHandler.CreateVideo)
		videoRoutes.GET("/:id", videoHandler.GetVideoByID)
	}

	log.Println("Server starting on port 8080...")
	router.Run(":8080")
}
