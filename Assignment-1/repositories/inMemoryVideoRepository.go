package repositories

import "sync"

//implementation of Get and Insert methods of interface videoRepository for in-memory

type InMemoryVideoRepository struct {
	videoMap map[string]int
}

func CreateInMemoryVideoRepository() *InMemoryVideoRepository {
	return &InMemoryVideoRepository{videoMap: map[string]int{}}
}

var mut sync.Mutex

func (repo *InMemoryVideoRepository) IncrementViewCountById(videoId string) {
	count, present := repo.videoMap[videoId]
	if present {
		mut.Lock()
		repo.videoMap[videoId] = count + 1
		mut.Unlock()
	} else {
		mut.Lock()
		repo.videoMap[videoId] = 1
		mut.Unlock()
	}
}

func (repo *InMemoryVideoRepository) GetViewCountById(videoId string) int {
	mut.Lock()
	count := repo.videoMap[videoId]
	mut.Unlock()
	return count
}
