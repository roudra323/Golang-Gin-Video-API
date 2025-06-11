package repository

import (
	"github.com/roudra323/gin-simple-api/domain"
	"gorm.io/gorm"
)

type VideoRepository interface {
	Save(video domain.Video) (domain.Video, error)
	FindAll() ([]domain.Video, error)
	FindByID(id uint64) (domain.Video, error)
	Update(video domain.Video) (domain.Video, error)
	Delete(video domain.Video) error
}

type videoRepository struct {
	db *gorm.DB
}

func (r *videoRepository) Save(video domain.Video) (domain.Video, error) {
	err := r.db.Create(&video).Error
	return video, err
}

func (r *videoRepository) FindAll() ([]domain.Video, error) {
	var videos []domain.Video
	err := r.db.Find(&videos).Error
	return videos, err
}

func (r *videoRepository) FindByID(id uint64) (domain.Video, error) {
	var video domain.Video
	err := r.db.First(&video, id).Error
	return video, err
}

func (r *videoRepository) Update(video domain.Video) (domain.Video, error) {
	err := r.db.Save(&video).Error
	return video, err
}

func (r *videoRepository) Delete(video domain.Video) error {
	err := r.db.Delete(&video).Error
	return err
}
