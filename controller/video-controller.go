package controller

import (
	"gtest/entity"
	"gtest/service"
	"gtest/validators"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) error
	ShowAll(ctx *gin.Context)
}

type controller struct {
	service service.VideoService
}

// Custom Validate
var validate *validator.Validate

func New(service service.VideoService) VideoController {

	validate = validator.New()
	validate.RegisterValidation("is-cool", validators.ValidateCoolTitle)

	return &controller{
		service: service,
	}

}

func (c *controller) FindAll() []entity.Video {

	return c.service.FindAll()

}

func (c *controller) Save(ctx *gin.Context) error {

	var video entity.Video

	// Bind struct json with binding rules declared
	err := ctx.BindJSON(&video)
	if err != nil {
		return err
	}

	// valid struct with validation rules declared
	err = validate.Struct(video)
	if err != nil {
		return err
	}

	c.service.Save(video)
	return nil

}

func (c *controller) ShowAll(ctx *gin.Context) {

	videos := c.FindAll()

	data := gin.H{
		"title":  "Video Page",
		"videos": videos,
	}

	ctx.HTML(http.StatusOK, "index.html", data)
}
