package main

import (
	"github.com/amrllkmn/apex/internal/activities"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r = activities.SetupRouter(r)
	r.Run() // Run and serve on 8080
}
