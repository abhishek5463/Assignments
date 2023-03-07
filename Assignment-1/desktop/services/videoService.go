package services

import "github.com/abhishek5463/assignment-1/repositories"

type VideoService struct {
	vidRepo repositories.VideoRepository
}

func CreateVideoService(vidRepo repositories.VideoRepository) *VideoService {
	return &VideoService{
		vidRepo: vidRepo,
	}
}

func (vs *VideoService) IncrementVideoViewCount(videoId string) {
	vs.vidRepo.IncrementViewCountById(videoId)
}

func (vs *VideoService) GetVideoViewCount(videoId string) int {
	return vs.vidRepo.GetViewCountById(videoId)
}
