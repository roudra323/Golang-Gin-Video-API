package api

import "github.com/roudra323/gin-simple-api/service"

type VideoHandler struct {
	service service.VideoService
}

func NewVideohandler(s service.VideoService) *VideoHandler {
	return &VideoHandler{service: s}
}
