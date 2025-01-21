package activities

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) *gin.Engine {
	activityService := NewService()
	activitiesHandler := NewHandler(activityService)
	r.POST("v1/activities", activitiesHandler.HandleCreateActivity)
	r.GET("v1/activities", activitiesHandler.HandleGetActivities)
	r.GET("v1/activities/:id", activitiesHandler.HandleReadActivity)
	return r
}
