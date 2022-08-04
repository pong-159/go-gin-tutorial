package main

import (
	// "fmt"
	"go-gin/tutorial-api/controller"
	"go-gin/tutorial-api/middleware"
	"go-gin/tutorial-api/service"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	videoService service.VideoService = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func setUpLogOutput(){
	f, _ := os.Create("gin.log")

	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

}


func main() {
	
	server := gin.New()

	setUpLogOutput()
	server.Use(gin.Recovery(), middleware.Logger(), gindump.Dump())
 
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