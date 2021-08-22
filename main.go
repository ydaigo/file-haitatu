package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")
	router.Static("/css", "css")
	router.Static("/js", "js")
	u, err := generateV4GetObjectSignedURL("file-haitatu", "test.sh", os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(u)
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", "hello")
	})
	router.POST("/", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		fmt.Print(file.Filename)
		fileName := file.Filename
		c.SaveUploadedFile(file, "tmp")
		uploadFile("file-haitatu", fileName, "tmp")
		c.HTML(http.StatusOK, "index.tmpl.html", fileName)
	})

	router.Run(":" + port)
}
