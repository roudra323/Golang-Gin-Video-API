package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
	"github.com/roudra323/gin-simple-api/api"
	"github.com/roudra323/gin-simple-api/domain"
	"github.com/roudra323/gin-simple-api/repository"
	"github.com/roudra323/gin-simple-api/service"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Enhanced CORS for Socket.IO compatibility
		origin := c.Request.Header.Get("Origin")
		c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		// Important: Remove Origin header to prevent conflicts
		c.Request.Header.Del("Origin")
		c.Next()
	}
}

func main() {
	// Database setup
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Database migration running...")
	err = db.AutoMigrate(&domain.Video{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Repository and service setup
	videoRepo := repository.NewVideoRepository(db)
	videoService := service.NewVideoService(videoRepo)
	videoHandler := api.NewVideohandler(videoService)

	// Gin router setup
	router := gin.Default()
	router.Use(CORSMiddleware())

	// Socket.IO setup
	socketServer := socketio.NewServer(nil)
	socketHandler := api.NewSocketHandler(videoService)
	socketHandler.RegisterHandlers(socketServer)

	// REST API routes
	videoRoutes := router.Group("/videos")
	{
		videoRoutes.GET("", videoHandler.GetVideos)
		videoRoutes.POST("", videoHandler.CreateVideo)
		videoRoutes.GET("/:id", videoHandler.GetVideoByID)
	}

	// ✅ FIXED: Socket.IO routes - Use direct ServeHTTP instead of gin.WrapH
	router.GET("/socket.io/*any", func(c *gin.Context) {
		socketServer.ServeHTTP(c.Writer, c.Request)
	})

	router.POST("/socket.io/*any", func(c *gin.Context) {
		socketServer.ServeHTTP(c.Writer, c.Request)
	})

	// Start Socket.IO server
	go func() {
		if err := socketServer.Serve(); err != nil {
			log.Fatalf("socketio listen error: %s\n", err)
		}
	}()
	defer socketServer.Close()

	log.Println("🎯 Server starting on port 8080...")
	log.Println("📡 REST API: http://localhost:8080/videos")
	log.Println("🔌 Socket.IO: ws://localhost:8080/socket.io/")
	router.Run(":8080")
}
