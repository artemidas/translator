package main

import (
	"github.com/artemidas/translator/controller"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	translator := router.Group("/api/translate")
	{
		translator.GET("/:lang", controller.GenerateTranslation)
		translator.POST("/:lang/create", controller.CreateTranslation)
		translator.PUT("/:lang/update/:id", controller.UpdateTranslation)
		translator.DELETE("/:lang/delete/:id", controller.DeleteTranslation)
	}

	router.Use(static.Serve("/", static.LocalFile("./dist", true)))
	router.Static("/css", "./dist/css")
	router.Static("/img", "./dist/img")
	router.Static("/js", "./dist/js")
	router.StaticFile("/favicon.ico", "./dist/favicon.ico")

	router.Run(":8000")
}
