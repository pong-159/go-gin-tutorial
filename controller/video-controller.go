package controller

import (
	"go-gin/tutorial-api/entity"
	"go-gin/tutorial-api/service"
	"go-gin/tutorial-api/validators"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) error
}
type controller struct {
	service service.VideoService
}

func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) error {
	var video entity.Video
	err := ctx.ShouldBindJSON(&video)

	if err != nil {
		return err
	}

	err2 := validate.Struct(video)

	if err2 != nil {
		return err2
	}
	c.service.Save(video)
	return nil
}




func New(service service.VideoService) VideoController {
	validate = validator.New()
	validate.RegisterValidation("is-cool", validators.ValidateCoolTitle)
	return &controller{
		service : service,
	}
}

