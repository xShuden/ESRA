package main

import (
	"esra/database"
	"esra/handlers" // "handlers" paketini doğru şekilde import ettiğinizden emin olun

	"github.com/gin-gonic/gin"
)

func main() {
	// Veri tabanı bağlantısını başlat
	database.Initialize()

	router := gin.Default()

	// HTML şablonları ve statik dosyaları yükle
	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("templates/*")

	// Route'lar
	router.GET("/", handlers.ShowProjects)
	router.GET("/add", handlers.ShowAddProject)
	router.POST("/add", handlers.AddProject)
	router.GET("/projects", handlers.ShowProjects)
	router.GET("/project/:id", handlers.ShowProjectDetails)

	router.Run(":8080")
}
