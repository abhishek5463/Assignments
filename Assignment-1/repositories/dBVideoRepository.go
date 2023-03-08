package repositories

type DBVideoRepository struct {
}

// different type of implementation of videoRepository interface

func (repo *DBVideoRepository) IncrementViewCountById(videoId string) {

}

func (repo *DBVideoRepository) GetViewCountById(videoId string) {

}
