package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/roudra323/gin-simple-api/domain"
	"github.com/roudra323/gin-simple-api/service"
)

type VideoHandler struct {
	service service.VideoService
}

func NewVideohandler(s service.VideoService) *VideoHandler {
	return &VideoHandler{service: s}
}

func (h *VideoHandler) GetVideos(c *gin.Context) {

	videos, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve videos",
		})
		return
	}

	c.JSON(http.StatusOK, videos)
}

func (h *VideoHandler) CreateVideo(c *gin.Context) {
	var video domain.Video

	if err := c.ShouldBindBodyWithJSON(&video); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	newVideo, err := h.service.Create(video)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create video",
		})
		return
	}

	c.JSON(http.StatusCreated, newVideo)
}

func (h *VideoHandler) GetVideoByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID format",
		})
		return
	}

	video, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Video not found"})
		return
	}

	c.JSON(http.StatusOK, video)
}
