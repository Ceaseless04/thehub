package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// function for API routes
func main() {
	
	r := gin.Default();
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello",
		});
	});
	r.Run("localhost:8080");

}
