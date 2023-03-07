package repositories

type InMemoryVideoRepository struct {
	videoMap map[string]int
}

func CreateInMemoryVideoRepository() *InMemoryVideoRepository {
	return &InMemoryVideoRepository{videoMap: map[string]int{}}
}
func (repo *InMemoryVideoRepository) IncrementViewCountById(videoId string) {
	count, present := repo.videoMap[videoId]
	if present {
		repo.videoMap[videoId] = count + 1
	} else {
		repo.videoMap[videoId] = 1
	}
}

func (repo *InMemoryVideoRepository) GetViewCountById(videoId string) int {
	return repo.videoMap[videoId]
}
