package repositories


//interface VideoRepository having Get and Insert method 

type VideoRepository interface {
	IncrementViewCountById(videoId string)
	GetViewCountById(videoId string) int
}
