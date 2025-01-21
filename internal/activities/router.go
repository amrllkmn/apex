package activities

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) *gin.Engine {
	activityService := NewService()
	activitiesHandler := NewHandler(activityService)
	r.GET("v1/activities/:id", activitiesHandler.HandleReadActivity)
	r.POST("v1/activities", activitiesHandler.HandleCreateActivity)
	return r
}
