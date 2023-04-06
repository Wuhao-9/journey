package router

import (
	"net/http"
	v1 "race-proj/router/api/v1"
	"race-proj/setting"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.SetMode(setting.Run_mode)

	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	router.Static("/project", "/home/ubuntu/myProj/race-proj/project")
	router.LoadHTMLFiles("project/html/HomePage.html", "project/html/scan.html")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "HomePage.html", nil)
	})

	router.GET("/favicon.ico", func(ctx *gin.Context) {
		ctx.File("project/logo.ico")
	})

	apiv1 := router.Group("/api/v1/")
	{
		apiv1.GET("/dirsearch", v1.Request_DIRSearch)
	}

	return router
}
