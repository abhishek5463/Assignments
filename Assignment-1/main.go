package main

import (
	"github.com/abhishek5463/assignment-1/routes"
	"github.com/gin-gonic/gin"
)

const PORT = "8081"

// main function

func main() {
	router := gin.New()
	routes.VideoRoutes(router)

	router.Run(":" + PORT)
}
