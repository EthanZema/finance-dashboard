package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func main() {
	engine := gin.Default()
	engine.SetTrustedProxies([]string{"127.0.0.1"})
	engine.GET("/", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
	})
	engine.Run(":8080")
}