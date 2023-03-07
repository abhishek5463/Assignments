package controllers

import (
	"net/http"

	"github.com/abhishek5463/assignment-1/services"
	"github.com/gin-gonic/gin"
)

type VideoController struct {
	vidSer services.VideoService
}

func CreateVideoController(vidSer services.VideoService) *VideoController {
	return &VideoController{
		vidSer: vidSer,
	}
}

func (c *VideoController) IncrementViewCount() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		vars := ctx.Params
		videoId, _ := vars.Get("video_id")
		c.vidSer.IncrementVideoViewCount(videoId)
	}
}
func (c *VideoController) GetVideoCount() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		vars := ctx.Params
		videoId, _ := vars.Get("video_id")

		ctx.JSON(http.StatusOK, gin.H{
			"count": c.vidSer.GetVideoViewCount(videoId),
		})

	}
}
