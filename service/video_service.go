package service

import (
	"github.com/roudra323/gin-simple-api/domain"
	"github.com/roudra323/gin-simple-api/repository"
)

type VideoService interface {
	Create(video domain.Video) (domain.Video, error)
	GetAll() ([]domain.Video, error)
	GetByID(id uint64) (domain.Video, error)
	Update(video domain.Video) (domain.Video, error)
	Delete(id uint64) error
}

type videoService struct {
	repo repository.VideoRepository
}

func NewVideoService(r repository.VideoRepository) VideoService {
	return &videoService{repo: r}
}

// Create contains the business logic for creating a new video.
func (s *videoService) Create(video domain.Video) (domain.Video, error) {
	// some business logic
	return s.repo.Save(video)
}

// GetAll contains the business logic for retrieving all videos.
func (s *videoService) GetAll() ([]domain.Video, error) {
	// some business logic
	return s.repo.FindAll()
}

// GetByID contains the business logic for retrieving a single video.
func (s *videoService) GetByID(id uint64) (domain.Video, error) {
	// some business logic
	return s.repo.FindByID(id)
}

// Update contains the business logic for updating a video.
func (s *videoService) Update(video domain.Video) (domain.Video, error) {
	// You might add logic here to check if the user performing the update
	// is the original author of the video.
	return s.repo.Update(video)
}

// Delete contains the business logic for deleting a video.
func (s *videoService) Delete(id uint64) error {
	video, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(video)
}
