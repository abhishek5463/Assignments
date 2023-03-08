package services

import (
	"testing"

	"github.com/abhishek5463/assignment-1/repositories"
)

// unit tests for  IncrementVideoCount function
func TestVideoService_IncrementVideoViewCount(t *testing.T) {
	type args struct {
		videoId string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{videoId: "1"},
		},
		{
			name: "test2",
			args: args{videoId: "1"},
		},
		{
			name: "test3",
			args: args{videoId: "2"},
		},
		{
			name: "test4",
			args: args{videoId: "2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			videoService := CreateVideoService(repositories.CreateInMemoryVideoRepository())
			videoService.IncrementVideoViewCount(tt.args.videoId)
		})
	}
}

//unit tests for  GetVideoCount function

func TestVideoService_GetVideoViewCount(t *testing.T) {
	type args struct {
		videoId string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{videoId: "1"},
			want: 2,
		},
		{
			name: "test2",
			args: args{videoId: "1"},
			want: 2,
		},
		{
			name: "test3",
			args: args{videoId: "2"},
			want: 2,
		},
		{
			name: "test4",
			args: args{videoId: "2"},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			videoService := CreateVideoService(repositories.CreateInMemoryVideoRepository())
			videoService.IncrementVideoViewCount(tt.args.videoId)
			videoService.IncrementVideoViewCount(tt.args.videoId)
			if got := videoService.GetVideoViewCount(tt.args.videoId); got != tt.want {
				t.Errorf("VideoService.GetVideoViewCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
