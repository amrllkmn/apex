package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello")
	})
	r.GET("/hello/:name", HandleHello)
	r.Run() // Run and serve on 8080
}
