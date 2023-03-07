package repositories

type VideoRepository interface {
	IncrementViewCountById(videoId string)
	GetViewCountById(videoId string) int
}
