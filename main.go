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
	u := ""
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", gin.H{
			"url": u,
		})
	})
	router.POST("/", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		fmt.Print(file.Filename)
		fileName := file.Filename
		c.SaveUploadedFile(file, "tmp")
		uploadFile("file-haitatu", fileName, "tmp")
		u, err := generateV4GetObjectSignedURL("file-haitatu", fileName, os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))
		fmt.Print(u)
		if err != nil {
			fmt.Print(err)
		}
		c.HTML(http.StatusOK, "index.tmpl.html", gin.H{
			"url": u,
		})
	})

	router.GET("/hello", func(c *gin.Context) {
		c.Writer.Header().Set("Content-Disposition", "attachment;filename=aaa")
		c.File("static/Cat-03.jpg")
	})
	router.GET("/json", func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Next()
	})

	router.Run(":" + port)
}
