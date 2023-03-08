package routes

import (
	"github.com/abhishek5463/assignment-1/controllers"
	"github.com/abhishek5463/assignment-1/repositories"
	"github.com/abhishek5463/assignment-1/services"
	"github.com/gin-gonic/gin"
)

//all the requests will sendfrom this function

func VideoRoutes(ir *gin.Engine) {
	repository := repositories.CreateInMemoryVideoRepository()
	service := services.CreateVideoService(repository)
	controller := controllers.CreateVideoController(*service)
	ir.POST("/increment-view-count/:video_id", controller.IncrementViewCount())
	ir.GET("/view-count/:video_id", controller.GetVideoCount())
}
