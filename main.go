package main

import (
	"gtest/controller"
	"gtest/middlewares"
	"gtest/service"
	"net/http"

	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func main() {

	engine := gin.New()

	engine.Static("css", "./templates/css")
	engine.LoadHTMLGlob("templates/*.html")

	engine.Use(
		middlewares.Logger(),
		gin.Recovery(),
		middlewares.BasicAuth(),
		gindump.Dump(),
	)

	apiRoutes := engine.Group("/api")
	{

		apiRoutes.GET("/videos", func(ctx *gin.Context) {
			ctx.JSON(200, videoController.FindAll())
		})

		apiRoutes.POST("/videos", func(ctx *gin.Context) {
			err := videoController.Save(ctx)

			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Video Input is valid."})
			}

		})

	}

	viewRoutes := engine.Group("/views")
	{

		viewRoutes.GET("/videos", videoController.ShowAll)

	}

	engine.Run(":3000")

}

func HomePage(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "World",
	})

}
