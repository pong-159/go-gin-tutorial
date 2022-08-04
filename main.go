package main

import (
	// "fmt"
	"go-gin/tutorial-api/controller"
	"go-gin/tutorial-api/middleware"
	"go-gin/tutorial-api/repository"
	"go-gin/tutorial-api/service"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	// gindump "github.com/tpkeeper/gin-dump"
)

var (
	videoRepository repository.VideoRepository = repository.NewVideoRepository()
	videoService service.VideoService = service.New(videoRepository)
	videoController controller.VideoController = controller.New(videoService)
)

func setUpLogOutput(){
	f, _ := os.Create("gin.log")

	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

}


func main() {

	defer videoRepository.CloseDB()
	
	server := gin.New() 

	setUpLogOutput()
	// server.Use(gin.Recovery(), middleware.Logger(), gindump.Dump())
	server.Use(gin.Recovery(), middleware.Logger())
 
	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, videoController.FindAll(),
	)})

	server.POST("/videos", func(ctx *gin.Context) {
		err := videoController.Save(ctx)
		if err != nil{
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error" : err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK,"Video was Saved")
	
	})

	server.PUT("/videos/:id", func(ctx *gin.Context) {
		err := videoController.Update(ctx)
		if err != nil{
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error" : err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK,"Video was Updated")
	
	})

	server.DELETE("/videos/:id", func(ctx *gin.Context) {
		err := videoController.Delete(ctx)
		if err != nil{
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error" : err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK,"Video was Deleted")
	
	})

	err := server.Run(":8000")

	if err != nil{
		panic(err)
	}
	
}