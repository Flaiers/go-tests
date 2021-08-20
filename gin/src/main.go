package main

import "github.com/gin-gonic/gin"

func get(c *gin.Context) {
	c.String(200, "Hello, World!")
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())
	router.GET("/", get)
	router.Run("0.0.0.0:8000")
}
