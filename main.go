package main 

import (
	// "fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"go-gin/tutorial-api/service"
	"go-gin/tutorial-api/controller"
)

var (
	videoService service.VideoService = service.New()
	videoController controller.VideoController = controller.New(videoService)
)


func main() {
	server := gin.Default()
 
	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, videoController.FindAll(),
	)})

	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, videoController.Save(ctx),
	
	)})

	err := server.Run(":8000")

	if err != nil{
		panic(err)
	}
	
}