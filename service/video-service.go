package service

import (
	"go-gin/tutorial-api/entity"
	"go-gin/tutorial-api/repository"
)

type VideoService interface {
	Save(entity.Video) entity.Video
	Update(entity.Video) 
	Delete(entity.Video) 
	FindAll() []entity.Video
}

type videoService struct{
	videoRepository repository.VideoRepository
}

func New(repo repository.VideoRepository) *videoService{
	return &videoService{
		videoRepository : repo,
	}
}

func (service *videoService) Save(video entity.Video) entity.Video {
	service.videoRepository.Save(video)
	return video
}

func (service *videoService) FindAll() []entity.Video{
	return service.videoRepository.FindAll()
}

func (service *videoService) Update(video entity.Video)  {
	service.videoRepository.Update(video)

}

func (service *videoService) Delete(video entity.Video) {
	service.videoRepository.Delete(video)

}