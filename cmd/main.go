package main

import (
	"github.com/amrllkmn/apex/internal/activities"
	"github.com/gin-gonic/gin"
)

func main() {
	activityService := activities.NewService()
	activitiesHandler := activities.NewHandler(activityService)
	r := gin.Default()
	r.GET("/hello/:name", activities.HandleHello)
	r.GET("/activities/:id", activitiesHandler.HandleReadActivity)
	r.Run() // Run and serve on 8080
}
