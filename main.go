package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	r.LoadHTMLGlob("templates/*")
	r.Static("static", "./static")

	// Handle HTTP requests
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "test.gohtml", gin.H{})
	})

	r.POST("/request", func(c *gin.Context) {
		message := c.PostForm("request")
		c.HTML(http.StatusOK, "request.gohtml", gin.H{
			"Message": message,
		})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	println("stared server")
	r.Run() // listen and serve on 0.0.0.0:8080
	println("bye bye")
}
