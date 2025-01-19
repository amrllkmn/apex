package main

import (
	"github.com/gin-gonic/gin"
	"github.com/amrllkmn/apex/internal/activities"
)

func main() {
	r := gin.Default()
	r.GET("/hello/:name", activities.HandleHello)
	r.Run() // Run and serve on 8080
}
