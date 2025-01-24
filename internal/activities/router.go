package activities

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(r *gin.Engine, db *gorm.DB) *gin.Engine {
	activityRepo := NewRepo(db)
	activityService := NewService(activityRepo)
	activitiesHandler := NewHandler(activityService)
	v1Activities := r.Group("v1/activities")
	{
		v1Activities.POST("/", activitiesHandler.HandleCreateActivity)
		v1Activities.GET("/", activitiesHandler.HandleGetActivities)
		v1Activities.GET("/:id", activitiesHandler.HandleReadActivity)
		v1Activities.PATCH("/:id", activitiesHandler.HandleUpdateActivity)
		v1Activities.DELETE("/:id", activitiesHandler.HandleDeleteActivity)
	}

	return r
}
