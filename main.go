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

	socketio "github.com/googollee/go-socket.io"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

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

	router.Use(CORSMiddleware())

	socketServer := socketio.NewServer(nil)

	socketHandler := api.NewSocketHandler(videoService)
	socketHandler.RegisterHandlers(socketServer)

	videoRoutes := router.Group("/videos")
	{
		videoRoutes.GET("", videoHandler.GetVideos)
		videoRoutes.POST("", videoHandler.CreateVideo)
		videoRoutes.GET("/:id", videoHandler.GetVideoByID)
	}

	// --- NEW: Socket.IO routes ---
	// Gin will forward all requests for /socket.io/ to our socketServer.
	router.GET("/socket.io/*any", gin.WrapH(socketServer))
	router.POST("/socket.io/*any", gin.WrapH(socketServer))

	go func() {
		if err := socketServer.Serve(); err != nil {
			log.Fatalf("socketio listen error: %s\n", err)
		}
	}()
	// --- NEW: Ensure the socket server is closed when the app exits ---
	defer socketServer.Close()

	log.Println("Server starting on port 8080...")
	router.Run(":8080")
}
